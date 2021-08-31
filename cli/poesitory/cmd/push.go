package cmd

import (
	"errors"
	"fmt"
	"os"
	"os/exec"

	"github.com/Nevermore-FMS/poesitory/cli/poesitory/auth"
	"github.com/Nevermore-FMS/poesitory/cli/poesitory/comm"
	"github.com/Nevermore-FMS/poesitory/cli/poesitory/local"
	"github.com/spf13/cobra"
)

var basePath string
var channel string
var version string
var noNpm bool
var plugin *comm.Plugin
var pathsToInclude []string

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
		if len(nevermoreJson.FrontendPath) == 0 && len(nevermoreJson.PluginPath) == 0 {
			return errors.New("nevermore.json must include at least one of \"frontendPath\" or \"pluginPath\"")
		}
		version = packageJson.Version
		comm.AuthorizationHeader = auth.DeferAuthentication(webAuth, userToken, uploadToken)
		plugin, err = comm.GetPlugin(nevermoreJson.Name)
		if err != nil {
			return err
		}
		if !noNpm {
			fmt.Println("Running \"npm run build\"")
			cmd := exec.Command("npm", "run", "build")
			cmd.Dir = basePath
			cmd.Stdout = os.Stdout
			cmd.Stderr = os.Stderr
			err = cmd.Run()
			if err != nil {
				return err
			}
		}
		pathsToInclude = []string{"package.json", "nevermore.json"}
		if len(nevermoreJson.FrontendPath) > 0 {
			pathsToInclude = append(pathsToInclude, nevermoreJson.FrontendPath)
		}
		if len(nevermoreJson.PluginPath) > 0 {
			pathsToInclude = append(pathsToInclude, nevermoreJson.PluginPath)
		}
		return nil
	},
	RunE: func(cmd *cobra.Command, args []string) error {
		uploadUrl, err := comm.UploadPluginVersion(plugin.Id.(string), version, channel)
		if err != nil {
			return err
		}
		tar, err := comm.CreateTarGz(basePath, pathsToInclude)
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
	pushCmd.Flags().StringVar(&channel, "channel", "STABLE", "the channel to push the plugin to")
	pushCmd.Flags().StringVar(&basePath, "path", ".", "the path to the plugin")
	pushCmd.Flags().BoolVar(&noNpm, "no-npm", false, "disables running \"npm run build\" before pushing the plugin")
	rootCmd.AddCommand(pushCmd)
}
