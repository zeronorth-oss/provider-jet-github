package teammembership

import (
	"github.com/crossplane/terrajet/pkg/config"
)

// Configure configures individual resources by adding custom ResourceConfigurators.
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("github_team_membership", func(r *config.Resource) {

		// we need to override the default group that terrajet generated for
		// this resource, which would be "github"
		r.ShortGroup = "team"
		r.Kind = "TeamMembership"
		r.ExternalName = config.IdentifierFromProvider // Should be identifier from team_id

		r.References["team_id"] = config.Reference{
			Type: "github.com/joakimhew/provider-jet-github/apis/team/v1alpha1.Team",
		}
	})
}
