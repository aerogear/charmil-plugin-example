package describe

import (
	"context"
	"encoding/json"
	"errors"
	"io"

	flagutil "github.com/aerogear/charmil-plugin-example/pkg/cmdutil/flags"
	"github.com/aerogear/charmil-plugin-example/pkg/connection"
	"github.com/aerogear/charmil-plugin-example/pkg/serviceregistry"
	"github.com/aerogear/charmil/core/utils/iostreams"
	"github.com/aerogear/charmil/core/utils/localize"

	"github.com/aerogear/charmil-plugin-example/pkg/cmd/flag"

	"github.com/aerogear/charmil-plugin-example/pkg/cmd/factory"
	"github.com/aerogear/charmil-plugin-example/pkg/config"
	"github.com/aerogear/charmil-plugin-example/pkg/dump"
	srsmgmtv1 "github.com/redhat-developer/app-services-sdk-go/registrymgmt/apiv1/client"
	"github.com/spf13/cobra"
	"gopkg.in/yaml.v2"
)

type Options struct {
	id           string
	name         string
	outputFormat string

	IO         *iostreams.IOStreams
	CfgHandler *config.CfgHandler
	Connection factory.ConnectionFunc
	localizer  localize.Localizer
}

// NewDescribeCommand describes a service instance, either by passing an `--id flag`
// or by using the service instance set in the config, if any
func NewDescribeCommand(f *factory.Factory) *cobra.Command {
	opts := &Options{
		CfgHandler: f.CfgHandler,
		Connection: f.Connection,
		IO:         f.IOStreams,
		localizer:  f.Localizer,
	}

	cmd := &cobra.Command{
		Use:     "describe",
		Short:   f.Localizer.LocalizeByID("registry.cmd.describe.shortDescription"),
		Long:    f.Localizer.LocalizeByID("registry.cmd.describe.longDescription"),
		Example: f.Localizer.LocalizeByID("registry.cmd.describe.example"),
		Args:    cobra.RangeArgs(0, 1),
		RunE: func(cmd *cobra.Command, args []string) error {
			validOutputFormats := flagutil.ValidOutputFormats
			if opts.outputFormat != "" && !flagutil.IsValidInput(opts.outputFormat, validOutputFormats...) {
				return flag.InvalidValueError("output", opts.outputFormat, validOutputFormats...)
			}

			if len(args) > 0 {
				opts.name = args[0]
			}

			if opts.name != "" && opts.id != "" {
				return errors.New(opts.localizer.LocalizeByID("service.error.idAndNameCannotBeUsed"))
			}

			if opts.id != "" || opts.name != "" {
				return runDescribe(opts)
			}

			var registryConfig *config.Config
			if opts.CfgHandler.Cfg == registryConfig || opts.CfgHandler.Cfg.InstanceID == "" {
				return errors.New(opts.localizer.LocalizeByID("registry.common.error.noServiceSelected"))
			}

			opts.id = opts.CfgHandler.Cfg.InstanceID

			return runDescribe(opts)
		},
	}

	cmd.Flags().StringVarP(&opts.outputFormat, "output", "o", "json", opts.localizer.LocalizeByID("registry.cmd.flag.output.description"))
	cmd.Flags().StringVar(&opts.id, "id", "", opts.localizer.LocalizeByID("registry.common.flag.id"))

	flagutil.EnableOutputFlagCompletion(cmd)

	return cmd
}

func runDescribe(opts *Options) error {
	connection, err := opts.Connection(connection.DefaultConfigSkipMasAuth)
	if err != nil {
		return err
	}

	api := connection.API()

	var registry *srsmgmtv1.RegistryRest
	ctx := context.Background()
	if opts.name != "" {
		registry, _, err = serviceregistry.GetServiceRegistryByName(ctx, api.ServiceRegistryMgmt(), opts.name)
		if err != nil {
			return err
		}
	} else {
		registry, _, err = serviceregistry.GetServiceRegistryByID(ctx, api.ServiceRegistryMgmt(), opts.id)
		if err != nil {
			return err
		}
	}

	return printService(opts.IO.Out, registry, opts.outputFormat)
}

func printService(w io.Writer, registry interface{}, outputFormat string) error {
	switch outputFormat {
	case dump.YAMLFormat, dump.YMLFormat:
		data, err := yaml.Marshal(registry)
		if err != nil {
			return err
		}
		return dump.YAML(w, data)
	default:
		data, err := json.Marshal(registry)
		if err != nil {
			return err
		}
		return dump.JSON(w, data)
	}
}
