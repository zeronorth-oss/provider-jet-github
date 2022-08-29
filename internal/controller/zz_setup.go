/*
Copyright 2021 The Crossplane Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package controller

import (
	ctrl "sigs.k8s.io/controller-runtime"

	"github.com/crossplane/terrajet/pkg/controller"

	membership "github.com/joakimhew/provider-jet-github/internal/controller/membership/membership"
	providerconfig "github.com/joakimhew/provider-jet-github/internal/controller/providerconfig"
	repository "github.com/joakimhew/provider-jet-github/internal/controller/repository/repository"
	team "github.com/joakimhew/provider-jet-github/internal/controller/team/team"
	teammembers "github.com/joakimhew/provider-jet-github/internal/controller/team/teammembers"
	teammembership "github.com/joakimhew/provider-jet-github/internal/controller/team/teammembership"
	teamrepository "github.com/joakimhew/provider-jet-github/internal/controller/team/teamrepository"
)

// Setup creates all controllers with the supplied logger and adds them to
// the supplied manager.
func Setup(mgr ctrl.Manager, o controller.Options) error {
	for _, setup := range []func(ctrl.Manager, controller.Options) error{
		membership.Setup,
		providerconfig.Setup,
		repository.Setup,
		team.Setup,
		teammembers.Setup,
		teammembership.Setup,
		teamrepository.Setup,
	} {
		if err := setup(mgr, o); err != nil {
			return err
		}
	}
	return nil
}
