/*
Copyright 2021 Upbound Inc.
*/

package controller

import (
	ctrl "sigs.k8s.io/controller-runtime"

	"github.com/upbound/upjet/pkg/controller"

	storage "github.com/flerwin/provider-jet-vultr/internal/controller/object/storage"
	providerconfig "github.com/flerwin/provider-jet-vultr/internal/controller/providerconfig"
	instance "github.com/flerwin/provider-jet-vultr/internal/controller/vultr/instance"
)

// Setup creates all controllers with the supplied logger and adds them to
// the supplied manager.
func Setup(mgr ctrl.Manager, o controller.Options) error {
	for _, setup := range []func(ctrl.Manager, controller.Options) error{
		storage.Setup,
		providerconfig.Setup,
		instance.Setup,
	} {
		if err := setup(mgr, o); err != nil {
			return err
		}
	}
	return nil
}
