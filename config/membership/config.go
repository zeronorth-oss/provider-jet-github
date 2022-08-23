package membership

import (
	"github.com/crossplane/terrajet/pkg/config"
)

// Configure configures individual resources by adding custom ResourceConfigurators.
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("github_membership", func(r *config.Resource) {

		// we need to override the default group that terrajet generated for
		// this resource, which would be "github"
		r.ShortGroup = "membership"
		// r.ExternalName = config.ExternalName{
		// 	SetIdentifierArgumentFn: func(base map[string]interface{}, name string) {
		// 		base["username"] = name
		// 	},
		// 	OmittedFields: []string{"username"},
		// 	GetExternalNameFn: func(base map[string]interface{}) (string, error) {
		// 		return base["id"].(string), nil
		// 	},
		// 	GetIDFn: func(ctx context.Context, externalName string, parameters map[string]interface{}, providerConfig map[string]interface{}) (string, error) {
		// 		return fmt.Sprintf("%s:%s", providerConfig["owner"], externalName), nil
		// 	},
		// 	DisableNameInitializer: true,
		// }
		r.ExternalName = config.IdentifierFromProvider
		// r.ExternalName.GetIDFn = func(ctx context.Context, externalName string, parameters map[string]interface{}, providerConfig map[string]interface{}) (string, error) {
		// 	return fmt.Sprintf("%s:%s", providerConfig["owner"], externalName), nil
		// }
	})
}
