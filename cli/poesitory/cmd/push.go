package cmd

import (
	"errors"
	"fmt"
	"os"

	"github.com/Nevermore-FMS/poesitory/cli/poesitory/auth"
	"github.com/Nevermore-FMS/poesitory/cli/poesitory/comm"
	"github.com/Nevermore-FMS/poesitory/cli/poesitory/local"
	"github.com/spf13/cobra"
)

var basePath string
var channel string
var version string
var plugin *comm.Plugin

var pushCmd = &cobra.Command{
	Use:   "push",
	Short: "Pushes a plugin to Poesitory",
	Args:  cobra.NoArgs,
	PreRunE: func(cmd *cobra.Command, args []string) (err error) {
		if _, err := os.Stat(basePath); err != nil {
			return err.(*os.PathError).Err
		}
		nevermoreJson, err := local.ReadNevermoreJson(basePath)
		if err != nil {
			return err
		}
		packageJson, err := local.ReadPackageJson(basePath)
		if err != nil {
			return err
		}
		if packageJson.Name != nevermoreJson.Name {
			return errors.New("names in package.json and nevermore.json must match")
		}
		version = packageJson.Version
		comm.AuthorizationHeader = auth.DeferAuthentication(webAuth, userToken, uploadToken)
		plugin, err = comm.GetPlugin(nevermoreJson.Name)
		if err != nil {
			return err
		}
		return nil
	},
	RunE: func(cmd *cobra.Command, args []string) error {
		uploadUrl, err := comm.UploadPluginVersion(plugin.Id.(string), version, channel)
		if err != nil {
			return err
		}
		tar, err := comm.CreateTarGz(basePath)
		if err != nil {
			return err
		}
		err = comm.UploadUrl(uploadUrl, tar)
		if err != nil {
			return err
		}
		fmt.Printf("Pushed plugin %s", plugin.Name)
		return nil
	},
}

func init() {
	pushCmd.Flags().StringVar(&channel, "channel", "STABLE", "the channel to push the plugin to (defaults to STABLE)")
	pushCmd.Flags().StringVar(&basePath, "path", ".", "the path to the plugin (defaults to ./)")
	rootCmd.AddCommand(pushCmd)
}
