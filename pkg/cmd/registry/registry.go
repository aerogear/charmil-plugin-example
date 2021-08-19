// REST API exposed via the serve command.
package registry

import (
	"context"
	"fmt"
	"os"

	"github.com/aerogear/charmil-plugin-example/internal/build"
	"github.com/aerogear/charmil-plugin-example/internal/config"
	"github.com/aerogear/charmil-plugin-example/pkg/cmd/factory"
	"github.com/aerogear/charmil-plugin-example/pkg/cmd/registry/create"
	"github.com/aerogear/charmil-plugin-example/pkg/cmd/registry/delete"
	"github.com/aerogear/charmil-plugin-example/pkg/cmd/registry/describe"
	"github.com/aerogear/charmil-plugin-example/pkg/cmd/registry/list"
	"github.com/aerogear/charmil-plugin-example/pkg/cmd/registry/use"
	"github.com/aerogear/charmil-plugin-example/pkg/connection"
	"github.com/aerogear/charmil-plugin-example/pkg/localesettings"
	"github.com/aerogear/charmil-plugin-example/pkg/profile"
	"github.com/aerogear/charmil/core/utils/localize"
	"github.com/spf13/cobra"
	"golang.org/x/text/language"
)

func NewServiceRegistryCommand(f *factory.Factory, pluginBuilder *connection.Builder) *cobra.Command {
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

	cmdFactory := factory.New(build.Version, localizer)
	f.Localizer = cmdFactory.Localizer

	err = initConfig(cmdFactory)
	if err != nil {
		fmt.Print(localizer.LocalizeByID("main.config.error", localize.NewEntry("Error", err)))
		os.Exit(1)
	}

	cmdFactory.Connection = func(cfg *connection.Config) (connection.Connection, error) {
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
		Short:       profile.ApplyDevPreviewLabel(cmdFactory.Localizer.LocalizeByID("registry.cmd.shortDescription")),
		Long:        cmdFactory.Localizer.LocalizeByID("registry.cmd.longDescription"),
		Example:     cmdFactory.Localizer.LocalizeByID("registry.cmd.example"),
		Args:        cobra.MinimumNArgs(1),
	}

	// add sub-commands
	cmd.AddCommand(
		create.NewCreateCommand(cmdFactory),
		describe.NewDescribeCommand(cmdFactory),
		delete.NewDeleteCommand(cmdFactory),
		list.NewListCommand(cmdFactory),
		use.NewUseCommand(cmdFactory),
	)

	return cmd
}

func initConfig(f *factory.Factory) error {
	if !config.HasCustomLocation() {
		rhoasCfgDir, err := config.DefaultDir()
		if err != nil {
			return err
		}

		// create rhoas config directory
		if _, err = os.Stat(rhoasCfgDir); os.IsNotExist(err) {
			err = os.MkdirAll(rhoasCfgDir, 0o700)
			if err != nil {
				return err
			}
		}
	}

	cfgFile, err := f.Config.Load()

	if cfgFile != nil {
		return err
	}

	if !os.IsNotExist(err) {
		return err
	}

	cfgFile = &config.Config{}
	if err := f.Config.Save(cfgFile); err != nil {
		return err
	}
	return nil
}
