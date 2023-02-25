/*
Copyright 2021 Upbound Inc.
*/

package config

import (
	// Note(turkenh): we are importing this to embed provider schema document
	_ "embed"
	"fmt"
	"github.com/flerwin/provider-jet-vultr/config/instance"
	"github.com/flerwin/provider-jet-vultr/config/kubernetes"
	"github.com/flerwin/provider-jet-vultr/config/objectstorage"
	"github.com/flerwin/provider-jet-vultr/config/sshkey"
	"github.com/flerwin/provider-jet-vultr/config/vpc"

	ujconfig "github.com/upbound/upjet/pkg/config"
)

const (
	resourcePrefix = "vultr"
	modulePath     = "github.com/flerwin/provider-jet-vultr"
)

//go:embed schema.json
var providerSchema string

//go:embed provider-metadata.yaml
var providerMetadata string

// GetProvider returns provider configuration
func GetProvider() *ujconfig.Provider {
	pc := ujconfig.NewProvider([]byte(providerSchema), resourcePrefix, modulePath, []byte(providerMetadata),
		ujconfig.WithIncludeList(ExternalNameConfigured()),
		ujconfig.WithDefaultResourceOptions(
			ExternalNameConfigurations(),
		))

	pc.RootGroup = fmt.Sprintf("%s.flerwin.io", resourcePrefix)

	for _, configure := range []func(provider *ujconfig.Provider){
		// add custom config functions
		objectstorage.Configure,
		instance.Configure,
		kubernetes.Configure,
		vpc.Configure,
		sshkey.Configure,
	} {
		configure(pc)
	}

	pc.ConfigureResources()
	return pc
}
