package config

import (
	"errors"
	"strconv"

	"github.com/aerogear/charmil-plugin-example/pkg/localize"
	"github.com/aerogear/charmil-plugin-example/pkg/profile"

	"github.com/aerogear/charmil-plugin-example/internal/config"
	"github.com/aerogear/charmil-plugin-example/pkg/iostreams"
	"github.com/aerogear/charmil-plugin-example/pkg/logging"

	"github.com/aerogear/charmil-plugin-example/pkg/cmd/factory"
	"github.com/spf13/cobra"
)

type Options struct {
	IO        *iostreams.IOStreams
	Config    config.IConfig
	Logger    func() (logging.Logger, error)
	localizer localize.Localizer
}

func NewConfigCommand(f *factory.Factory) *cobra.Command {
	opts := &Options{
		IO:        f.IOStreams,
		Config:    f.Config,
		Logger:    f.Logger,
		localizer: f.Localizer,
	}

	cmd := &cobra.Command{
		Use:     "config",
		Short:   opts.localizer.MustLocalize("config.cmd.shortDescription"),
		Long:    opts.localizer.MustLocalize("config.cmd.longDescription"),
		Example: opts.localizer.MustLocalize("config.cmd.example"),
	}

	devPreview := &cobra.Command{
		Use:       "dev-preview",
		Short:     opts.localizer.MustLocalize("devpreview.cmd.shortDescription"),
		Long:      opts.localizer.MustLocalize("devpreview.cmd.longDescription"),
		Example:   opts.localizer.MustLocalize("devpreview.cmd.example"),
		ValidArgs: []string{"true", "false"},
		Args:      cobra.ExactValidArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			devPreview, err := strconv.ParseBool(args[0])
			if err != nil {
				return errors.New(opts.localizer.MustLocalize("devpreview.error.enablement"))
			}
			_, err = profile.EnableDevPreview(f, devPreview)
			return err
		},
	}
	cmd.AddCommand(devPreview)
	return cmd
}
