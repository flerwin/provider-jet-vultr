package vpc

import "github.com/upbound/upjet/pkg/config"

// Configure configures the vpc package.
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("vultr_vpc", func(r *config.Resource) {
		r.ExternalName = config.IdentifierFromProvider
	})
}
