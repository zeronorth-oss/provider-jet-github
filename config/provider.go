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

package config

import (
	// Note(turkenh): we are importing this to embed provider schema document
	_ "embed"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/joakimhew/provider-jet-github/config/membership"
	"github.com/joakimhew/provider-jet-github/config/repository"
	"github.com/joakimhew/provider-jet-github/config/team"
	teammembership "github.com/joakimhew/provider-jet-github/config/team_membership"

	tjconfig "github.com/crossplane/terrajet/pkg/config"
	teammembers "github.com/joakimhew/provider-jet-github/config/team_members"
	teamrepository "github.com/joakimhew/provider-jet-github/config/team_repository"
)

const (
	resourcePrefix = "github"
	modulePath     = "github.com/joakimhew/provider-jet-github"
)

//go:embed schema.json
var providerSchema string

// GetProvider returns provider configuration
func GetProvider() *tjconfig.Provider {
	defaultResourceFn := func(name string, terraformResource *schema.Resource, opts ...tjconfig.ResourceOption) *tjconfig.Resource {
		r := tjconfig.DefaultResource(name, terraformResource)
		// Add any provider-specific defaulting here. For example:
		//   r.ExternalName = tjconfig.IdentifierFromProvider
		return r
	}

	pc := tjconfig.NewProviderWithSchema([]byte(providerSchema), resourcePrefix, modulePath,
		tjconfig.WithDefaultResourceFn(defaultResourceFn),
		tjconfig.WithIncludeList([]string{
			"github_membership$",
			"github_team$",
			"github_repository$",
			"github_team_repository$",
			"github_team_members$",
			"github_team_membership$",
		}))

	for _, configure := range []func(provider *tjconfig.Provider){
		membership.Configure,
		team.Configure,
		repository.Configure,
		teamrepository.Configure,
		teammembers.Configure,
		teammembership.Configure,
	} {
		configure(pc)
	}

	pc.ConfigureResources()
	return pc
}
