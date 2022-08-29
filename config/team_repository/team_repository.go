package teamrepository

import (
	"github.com/crossplane/crossplane-runtime/pkg/fieldpath"
	"github.com/crossplane/crossplane-runtime/pkg/reference"
	"github.com/crossplane/crossplane-runtime/pkg/resource"
	"github.com/crossplane/terrajet/pkg/config"
)

const (
	// SelfPackagePath is the golang path for this package.
	SelfPackagePath       = "github.com/joakimhew/provider-jet-github/config/team_repository"
	PathTeamSlugExtractor = SelfPackagePath + ".TeamSlugExtractor()"
)

// Configure configures individual resources by adding custom ResourceConfigurators.
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("github_team_repository", func(r *config.Resource) {

		// we need to override the default group that terrajet generated for
		// this resource, which would be "github"
		r.ShortGroup = "team"
		r.Kind = "TeamRepository"
		r.ExternalName = config.IdentifierFromProvider

		r.References["team_id"] = config.Reference{
			Type: "github.com/joakimhew/provider-jet-github/apis/team/v1alpha1.Team",
		}

		r.References["repository"] = config.Reference{
			Type: "github.com/joakimhew/provider-jet-github/apis/repository/v1alpha1.Repository",
		}
	})
}

func TeamSlugExtractor() reference.ExtractValueFn {
	return func(mg resource.Managed) string {
		paved, err := fieldpath.PaveObject(mg)
		if err != nil {
			return ""
		}
		r, err := paved.GetString("status.atProvider.slug")
		if err != nil {
			return ""
		}
		return r
	}
}
