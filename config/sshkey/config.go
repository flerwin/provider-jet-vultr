package sshkey

import "github.com/upbound/upjet/pkg/config"

// Configure configures the sshkey package.
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("vultr_ssh_key", func(r *config.Resource) {
		r.ExternalName = config.IdentifierFromProvider
	})
}
