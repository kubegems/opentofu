package providers

import "github.com/kubegems/opentofu/pkg/addrs"

func NewMockSchemaCache() *schemaCache {
	return &schemaCache{
		m: make(map[addrs.Provider]ProviderSchema),
	}
}
