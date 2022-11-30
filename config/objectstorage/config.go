package objectstorage

import "github.com/upbound/upjet/pkg/config"

func Configure(p *config.Provider) {
	p.AddResourceConfigurator("vultr_object_storage", func(r *config.Resource) {
		r.ExternalName = config.IdentifierFromProvider
	})
}
