package cmd

import (
	"errors"
	"fmt"

	"github.com/Nevermore-FMS/poesitory/cli/poesitory/comm"
	"github.com/spf13/cobra"
)

var identifier string
var path string
var downloadUrl string

var pullCmd = &cobra.Command{
	Use:   "pull [identifier]",
	Short: "Pulls a plugin from Poesitory",
	Long: `Pulls a plugin from Poesitory
	- The identifier parameter comes in the following form:

		pluginname[#channel][@version]

		If channel is omitted, it will default to STABLE
		If version is omitted, it will default to the latest version for the given channel
	`,

	Args: cobra.ExactArgs(1),
	PreRunE: func(cmd *cobra.Command, args []string) error {
		identifier = args[0]
		pluginVersion, err := comm.GetPluginVersion(identifier)
		if err != nil {
			return err
		}
		if pluginVersion == nil || pluginVersion.Plugin == nil {
			return errors.New("unable to locate plugin version on poesitory")
		}
		if pluginVersion.DownloadUrl == nil {
			return errors.New("unable to obtain download url")
		}
		downloadUrl = string(*pluginVersion.DownloadUrl)
		if len(path) == 0 {
			path = string(pluginVersion.Plugin.Name)
		}
		return nil
	},
	RunE: func(cmd *cobra.Command, args []string) error {
		stream, err := comm.DownloadUrl(downloadUrl)
		if err != nil {
			return err
		}
		defer stream.Close()
		err = comm.ExtractTarGz(stream, path)
		if err != nil {
			return err
		}
		fmt.Printf("Pulled plugin %s", identifier)
		return nil
	},
}

func init() {
	pullCmd.Flags().StringVar(&path, "path", "", "path to place the plugin")
	rootCmd.AddCommand(pullCmd)
}
