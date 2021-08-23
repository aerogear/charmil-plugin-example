package update

import (
	"context"
	"errors"
	"fmt"
	"os"

	"github.com/aerogear/charmil-plugin-example/pkg/connection"
	"github.com/aerogear/charmil-plugin-example/pkg/dump"
	"github.com/aerogear/charmil/core/utils/localize"

	"github.com/aerogear/charmil-plugin-example/pkg/cmd/artifact/util"
	"github.com/aerogear/charmil-plugin-example/pkg/cmd/flag"
	flagutil "github.com/aerogear/charmil-plugin-example/pkg/cmdutil/flags"

	"github.com/aerogear/charmil/core/utils/iostreams"

	"github.com/aerogear/charmil/core/utils/logging"

	"github.com/spf13/cobra"

	"github.com/aerogear/charmil-plugin-example/pkg/cmd/factory"
	"github.com/aerogear/charmil-plugin-example/pkg/config"
)

type Options struct {
	artifact string
	group    string

	file string

	registryID   string
	outputFormat string

	IO         *iostreams.IOStreams
	CfgHandler *config.CfgHandler
	Connection factory.ConnectionFunc
	Logger     func() (logging.Logger, error)
	localizer  localize.Localizer
}

// NewUpdateCommand creates a new command for updating binary content of registry artifacts.
func NewUpdateCommand(f *factory.Factory) *cobra.Command {
	opts := &Options{
		IO:         f.IOStreams,
		CfgHandler: f.CfgHandler,
		Connection: f.Connection,
		Logger:     f.Logger,
		localizer:  f.Localizer,
	}

	cmd := &cobra.Command{
		Use:   "update",
		Short: "Update artifact",
		Long: `
Update artifact from file or directly standard input

Artifacts can be typically in JSON format for most of the supported types, but may be in another format for a few (for example, PROTOBUF).
The type of the content should be compatible with the current artifact's type.

When successful, this creates a new version of the artifact, making it the most recent (and therefore official) version of the artifact.

An artifact is update using the content provided in the body of the request.  
This content is updated under a unique artifactId provided by user.
		`,
		Example: `
## update artifact from group and artifact-id
rhoas service-registry artifact update --artifact-id=my-artifact --group my-group my-artifact.json 

## update artifact from group and artifact-id
rhoas service-registry artifact update --artifact-id=my-artifact --group my-group my-artifact.json 
`,
		Args: cobra.RangeArgs(0, 1),
		RunE: func(cmd *cobra.Command, args []string) error {
			validOutputFormats := flagutil.ValidOutputFormats
			if opts.outputFormat != "" && !flagutil.IsValidInput(opts.outputFormat, validOutputFormats...) {
				return flag.InvalidValueError("output", opts.outputFormat, validOutputFormats...)
			}

			if opts.artifact == "" {
				return errors.New("artifact is required. Please specify artifact by using --artifact-id flag")
			}

			if len(args) > 0 {
				opts.file = args[0]
			}

			if opts.registryID != "" {
				return runUpdate(opts)
			}

			if !opts.CfgHandler.Cfg.HasServiceRegistry() {
				return errors.New("no service Registry selected. Use 'rhoas service-registry use' use to select your registry")
			}

			opts.registryID = fmt.Sprint(opts.CfgHandler.Cfg.InstanceID)
			return runUpdate(opts)
		},
	}

	cmd.Flags().StringVarP(&opts.outputFormat, "output", "o", "json", opts.localizer.LocalizeByID("registry.cmd.flag.output.description"))
	cmd.Flags().StringVarP(&opts.file, "file", "f", "", "File location of the artifact")

	cmd.Flags().StringVarP(&opts.artifact, "artifact-id", "a", "", "Id of the artifact")
	cmd.Flags().StringVarP(&opts.group, "group", "g", util.DefaultArtifactGroup, "Artifact group")
	cmd.Flags().StringVar(&opts.registryID, "instance-id", "", "Id of the registry to be used. By default uses currently selected registry")

	flagutil.EnableOutputFlagCompletion(cmd)

	return cmd
}

func runUpdate(opts *Options) error {

	logger, err := opts.Logger()
	if err != nil {
		return err
	}

	conn, err := opts.Connection(connection.DefaultConfigRequireMasAuth)
	if err != nil {
		return err
	}

	dataAPI, _, err := conn.API().ServiceRegistryInstance(opts.registryID)
	if err != nil {
		return err
	}

	if opts.group == util.DefaultArtifactGroup {
		logger.Info("Group was not specified. Using", util.DefaultArtifactGroup, "artifacts group.")
		opts.group = util.DefaultArtifactGroup
	}

	var specifiedFile *os.File
	if opts.file != "" {
		logger.Info("Opening file: " + opts.file)
		specifiedFile, err = os.Open(opts.file)
		if err != nil {
			return err
		}
	} else {
		logger.Info("Reading file content from stdin")
		specifiedFile, err = util.CreateFileFromStdin()
		if err != nil {
			return err
		}
	}

	ctx := context.Background()
	request := dataAPI.ArtifactsApi.UpdateArtifact(ctx, opts.group, opts.artifact)
	request = request.Body(specifiedFile)
	metadata, _, err := request.Execute()
	if err != nil {
		return err
	}
	logger.Info("Artifact updated")

	dump.PrintDataInFormat(opts.outputFormat, metadata, opts.IO.Out)

	return nil
}