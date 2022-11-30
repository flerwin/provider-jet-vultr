package objectstorage

import "github.com/upbound/upjet/pkg/config"

// Configure configures the objectstorage package.
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("vultr_object_storage", func(r *config.Resource) {
		r.ExternalName = config.IdentifierFromProvider
	})
}
