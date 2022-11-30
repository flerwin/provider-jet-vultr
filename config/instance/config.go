package instance

import "github.com/upbound/upjet/pkg/config"

// Configure configures the instance package.
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("vultr_instance", func(r *config.Resource) {
		r.ExternalName = config.IdentifierFromProvider
	})
}
