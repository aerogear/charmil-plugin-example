package whoami

import (
	"fmt"

	"github.com/aerogear/charmil-plugin-example/pkg/auth/token"
	"github.com/aerogear/charmil-plugin-example/pkg/cmd/factory"
	"github.com/aerogear/charmil-plugin-example/pkg/config"
	"github.com/aerogear/charmil-plugin-example/pkg/connection"
	"github.com/aerogear/charmil/core/utils/iostreams"
	"github.com/aerogear/charmil/core/utils/localize"

	"github.com/spf13/cobra"

	"github.com/aerogear/charmil/core/utils/logging"
)

type Options struct {
	CfgHandler *config.CfgHandler
	Connection factory.ConnectionFunc
	IO         *iostreams.IOStreams
	Logger     func() (logging.Logger, error)
	localizer  localize.Localizer
}

func NewWhoAmICmd(f *factory.Factory) *cobra.Command {
	opts := &Options{
		CfgHandler: f.CfgHandler,
		Connection: f.Connection,
		IO:         f.IOStreams,
		Logger:     f.Logger,
		localizer:  f.Localizer,
	}

	cmd := &cobra.Command{
		Use:     f.Localizer.LocalizeByID("whoami.cmd.use"),
		Short:   f.Localizer.LocalizeByID("whoami.cmd.shortDescription"),
		Long:    f.Localizer.LocalizeByID("whoami.cmd.longDescription"),
		Example: f.Localizer.LocalizeByID("whoami.cmd.example"),
		Args:    cobra.NoArgs,
		RunE: func(cmd *cobra.Command, _ []string) error {
			return runCmd(opts)
		},
	}

	return cmd
}

func runCmd(opts *Options) (err error) {

	logger, err := opts.Logger()
	if err != nil {
		return err
	}

	_, err = opts.Connection(connection.DefaultConfigSkipMasAuth)
	if err != nil {
		return err
	}

	accessTkn, _ := token.Parse(opts.CfgHandler.Cfg.AccessToken)

	tknClaims, _ := token.MapClaims(accessTkn)

	userName, ok := tknClaims["preferred_username"]

	if ok {
		fmt.Fprintln(opts.IO.Out, userName)
	} else {
		logger.Info(opts.localizer.LocalizeByID("whoami.log.info.tokenHasNoUsername"))
	}

	return nil
}
