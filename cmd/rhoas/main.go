package main

import (
	"context"
	"fmt"
	"os"

	"github.com/aerogear/charmil-plugin-example/pkg/doc"
	"github.com/aerogear/charmil-plugin-example/pkg/localesettings"
	"github.com/aerogear/charmil/core/utils/localize"
	"golang.org/x/text/language"

	"github.com/aerogear/charmil-plugin-example/internal/build"

	"github.com/aerogear/charmil-plugin-example/internal/config"

	"github.com/aerogear/charmil-plugin-example/pkg/cmd/debug"
	"github.com/aerogear/charmil-plugin-example/pkg/cmd/factory"
	"github.com/aerogear/charmil-plugin-example/pkg/cmd/root"
	"github.com/spf13/cobra"
)

var generateDocs = os.Getenv("GENERATE_DOCS") == "true"
var cmdFactory *factory.Factory
var buildVersion string

func main() {
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

	buildVersion = build.Version
	cmdFactory = factory.New(build.Version, localizer)
	logger, err := cmdFactory.Logger()
	if err != nil {
		fmt.Println(cmdFactory.IOStreams.ErrOut, err)
		os.Exit(1)
	}

	err = initConfig(cmdFactory)
	if err != nil {
		logger.Errorf(localizer.LocalizeByID("main.config.error", localize.NewEntry("Error", err)))
		os.Exit(1)
	}

	RootCmd := RootCmd()

	RootCmd.InitDefaultHelpCmd()

	if generateDocs {
		generateDocumentation(RootCmd)
		os.Exit(0)
	}

	err = RootCmd.Execute()
	if err == nil {
		if debug.Enabled() {
			build.CheckForUpdate(context.Background(), logger, localizer)
		}
		return
	}

	if err != nil {
		logger.Error(wrapErrorf(err, localizer))
		build.CheckForUpdate(context.Background(), logger, localizer)
		os.Exit(1)
	}
}

func RootCmd() *cobra.Command {
	return root.NewRootCommand(cmdFactory, buildVersion)
}

/**
* Generates documentation files
 */
func generateDocumentation(rootCommand *cobra.Command) {
	fmt.Fprint(os.Stderr, "Generating docs.\n\n")
	filePrepender := func(filename string) string {
		return ""
	}

	rootCommand.DisableAutoGenTag = true

	linkHandler := func(s string) string { return s }

	err := doc.GenAsciidocTreeCustom(rootCommand, "./docs/commands", filePrepender, linkHandler)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
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

func wrapErrorf(err error, localizer localize.Localizer) error {
	return fmt.Errorf("Error: %w. %v", err, localizer.LocalizeByID("common.log.error.verboseModeHint"))
}
