// REST API exposed via the serve command.
package registry

import (
	"context"
	"fmt"
	"net/http"
	"os"

	"github.com/aerogear/charmil-plugin-example/pkg/cmd/factory"
	"github.com/aerogear/charmil-plugin-example/pkg/cmd/registry/create"
	"github.com/aerogear/charmil-plugin-example/pkg/cmd/registry/delete"
	"github.com/aerogear/charmil-plugin-example/pkg/cmd/registry/describe"
	"github.com/aerogear/charmil-plugin-example/pkg/cmd/registry/list"
	"github.com/aerogear/charmil-plugin-example/pkg/cmd/registry/use"
	"github.com/aerogear/charmil-plugin-example/pkg/connection"
	"github.com/aerogear/charmil-plugin-example/pkg/httputil"
	"github.com/aerogear/charmil-plugin-example/pkg/localesettings"
	"github.com/aerogear/charmil-plugin-example/pkg/profile"
	"github.com/aerogear/charmil/core/utils/localize"
	"github.com/spf13/cobra"
	"golang.org/x/text/language"

	"github.com/aerogear/charmil/core/utils/logging"
)

func NewServiceRegistryCommand(f *factory.Factory, pluginBuilder *connection.Builder) *cobra.Command {
	cfgFile, err := f.Config.Load()
	if err != nil {
		fmt.Println(f.IOStreams.ErrOut, err)
		os.Exit(1)
	}

	locConfig := &localize.Config{
		Language: &language.English,
		Files:    localesettings.DefaultLocales,
		Format:   "toml",
	}

	localizer, err := localize.New(locConfig)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}

	f.Localizer = localizer

	f.Connection = func(connectionCfg *connection.Config) (connection.Connection, error) {

		var logger logging.Logger

		transportWrapper := func(a http.RoundTripper) http.RoundTripper {
			return &httputil.LoggingRoundTripper{
				Proxied: a,
				Logger:  logger,
			}
		}

		pluginBuilder.WithTransportWrapper(transportWrapper)

		pluginBuilder.WithConnectionConfig(connectionCfg)

		conn, err := pluginBuilder.Build()
		if err != nil {
			return nil, err
		}

		err = conn.RefreshTokens(context.TODO())
		if err != nil {
			return nil, err
		}

		return conn, nil
	}

	cmd := &cobra.Command{
		Use:         "service-registry",
		Annotations: profile.DevPreviewAnnotation(),
		Short:       profile.ApplyDevPreviewLabel(f.Localizer.LocalizeByID("registry.cmd.shortDescription")),
		Long:        f.Localizer.LocalizeByID("registry.cmd.longDescription"),
		Example:     f.Localizer.LocalizeByID("registry.cmd.example"),
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

	if err := f.Config.Save(cfgFile); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}

	return cmd
}
