package kubernetes

import "github.com/upbound/upjet/pkg/config"

// Configure configures the kubernetes package.
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("vultr_kubernetes", func(r *config.Resource) {
		r.ExternalName = config.IdentifierFromProvider
	})
}
