// REST API exposed via the serve command.
package registry

import (
	"github.com/aerogear/charmil-plugin-example/pkg/cmd/factory"
	"github.com/aerogear/charmil-plugin-example/pkg/cmd/registry/create"
	"github.com/aerogear/charmil-plugin-example/pkg/cmd/registry/delete"
	"github.com/aerogear/charmil-plugin-example/pkg/cmd/registry/describe"
	"github.com/aerogear/charmil-plugin-example/pkg/cmd/registry/list"
	"github.com/aerogear/charmil-plugin-example/pkg/cmd/registry/use"
	"github.com/aerogear/charmil-plugin-example/pkg/profile"
	"github.com/spf13/cobra"
)

func NewServiceRegistryCommand(f *factory.Factory) *cobra.Command {
	cmd := &cobra.Command{
		Use:         "service-registry",
		Annotations: profile.DevPreviewAnnotation(),
		Short:       profile.ApplyDevPreviewLabel(f.Localizer.MustLocalize("registry.cmd.shortDescription")),
		Long:        f.Localizer.MustLocalize("registry.cmd.longDescription"),
		Example:     f.Localizer.MustLocalize("registry.cmd.example"),
		Args:        cobra.MinimumNArgs(1),
	}

	// add sub-commands
	cmd.AddCommand(
		create.NewCreateCommand(f),
		describe.NewDescribeCommand(f),
		delete.NewDeleteCommand(f),
		list.NewListCommand(f),
		use.NewUseCommand(f),
	)

	return cmd
}
