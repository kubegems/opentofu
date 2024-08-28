// Copyright (c) The OpenTofu Authors
// SPDX-License-Identifier: MPL-2.0
// Copyright (c) 2023 HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package tofu

import (
	"context"
	"sync"
	"time"

	"github.com/zclconf/go-cty/cty"

	"github.com/kubegems/opentofu/pkg/addrs"
	"github.com/kubegems/opentofu/pkg/checks"
	"github.com/kubegems/opentofu/pkg/configs"
	"github.com/kubegems/opentofu/pkg/encryption"
	"github.com/kubegems/opentofu/pkg/instances"
	"github.com/kubegems/opentofu/pkg/plans"
	"github.com/kubegems/opentofu/pkg/providers"
	"github.com/kubegems/opentofu/pkg/provisioners"
	"github.com/kubegems/opentofu/pkg/refactoring"
	"github.com/kubegems/opentofu/pkg/states"
	"github.com/kubegems/opentofu/pkg/tfdiags"
)

// ContextGraphWalker is the GraphWalker implementation used with the
// Context struct to walk and evaluate the graph.
type ContextGraphWalker struct {
	NullGraphWalker

	// Configurable values
	Context            *Context
	State              *states.SyncState       // Used for safe concurrent access to state
	RefreshState       *states.SyncState       // Used for safe concurrent access to state
	PrevRunState       *states.SyncState       // Used for safe concurrent access to state
	Changes            *plans.ChangesSync      // Used for safe concurrent writes to changes
	Checks             *checks.State           // Used for safe concurrent writes of checkable objects and their check results
	InstanceExpander   *instances.Expander     // Tracks our gradual expansion of module and resource instances
	ImportResolver     *ImportResolver         // Tracks import targets as they are being resolved
	MoveResults        refactoring.MoveResults // Read-only record of earlier processing of move statements
	Operation          walkOperation
	StopContext        context.Context
	RootVariableValues InputValues
	Config             *configs.Config
	PlanTimestamp      time.Time
	Encryption         encryption.Encryption

	// This is an output. Do not set this, nor read it while a graph walk
	// is in progress.
	NonFatalDiagnostics tfdiags.Diagnostics

	once        sync.Once
	contextLock sync.Mutex
	contexts    map[string]*BuiltinEvalContext

	variableValuesLock sync.Mutex
	variableValues     map[string]map[string]cty.Value

	providerLock  sync.Mutex
	providerCache map[string]providers.Interface

	provisionerLock  sync.Mutex
	provisionerCache map[string]provisioners.Interface
}

func (w *ContextGraphWalker) EnterPath(path addrs.ModuleInstance) EvalContext {
	w.contextLock.Lock()
	defer w.contextLock.Unlock()

	// If we already have a context for this path cached, use that
	key := path.String()
	if ctx, ok := w.contexts[key]; ok {
		return ctx
	}

	ctx := w.EvalContext().WithPath(path)
	w.contexts[key] = ctx.(*BuiltinEvalContext)
	return ctx
}

func (w *ContextGraphWalker) EvalContext() EvalContext {
	w.once.Do(w.init)

	// Our evaluator shares some locks with the main context and the walker
	// so that we can safely run multiple evaluations at once across
	// different modules.
	evaluator := &Evaluator{
		Meta:               w.Context.meta,
		Config:             w.Config,
		Operation:          w.Operation,
		State:              w.State,
		Changes:            w.Changes,
		Plugins:            w.Context.plugins,
		VariableValues:     w.variableValues,
		VariableValuesLock: &w.variableValuesLock,
		PlanTimestamp:      w.PlanTimestamp,
	}

	ctx := &BuiltinEvalContext{
		StopContext:           w.StopContext,
		Hooks:                 w.Context.hooks,
		InputValue:            w.Context.uiInput,
		InstanceExpanderValue: w.InstanceExpander,
		Plugins:               w.Context.plugins,
		MoveResultsValue:      w.MoveResults,
		ImportResolverValue:   w.ImportResolver,
		ProviderCache:         w.providerCache,
		ProviderInputConfig:   w.Context.providerInputConfig,
		ProviderLock:          &w.providerLock,
		ProvisionerCache:      w.provisionerCache,
		ProvisionerLock:       &w.provisionerLock,
		ChangesValue:          w.Changes,
		ChecksValue:           w.Checks,
		StateValue:            w.State,
		RefreshStateValue:     w.RefreshState,
		PrevRunStateValue:     w.PrevRunState,
		Evaluator:             evaluator,
		VariableValues:        w.variableValues,
		VariableValuesLock:    &w.variableValuesLock,
		Encryption:            w.Encryption,
	}

	return ctx
}

func (w *ContextGraphWalker) init() {
	w.contexts = make(map[string]*BuiltinEvalContext)
	w.providerCache = make(map[string]providers.Interface)
	w.provisionerCache = make(map[string]provisioners.Interface)
	w.variableValues = make(map[string]map[string]cty.Value)

	// Populate root module variable values. Other modules will be populated
	// during the graph walk.
	w.variableValues[""] = make(map[string]cty.Value)
	for k, iv := range w.RootVariableValues {
		w.variableValues[""][k] = iv.Value
	}
}

func (w *ContextGraphWalker) Execute(ctx EvalContext, n GraphNodeExecutable) tfdiags.Diagnostics {
	// Acquire a lock on the semaphore
	w.Context.parallelSem.Acquire()
	defer w.Context.parallelSem.Release()

	return n.Execute(ctx, w.Operation)
}
