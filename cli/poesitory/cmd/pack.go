package cmd

import (
	"errors"
	"fmt"
	"io"
	"os"
	"os/exec"

	"github.com/Nevermore-FMS/poesitory/cli/poesitory/comm"
	"github.com/Nevermore-FMS/poesitory/cli/poesitory/local"
	"github.com/spf13/cobra"
)

//var basePath string          - shared with push.go
//var noNpm bool               - shared with push.go
//var pathsToInclude []string  - shared with push.go
var packOutput string
var pluginName string

var packCmd = &cobra.Command{
	Use:   "pack",
	Short: "Creates a tarball of a plugin",
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
		pluginName = packageJson.Name
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
		pathsToInclude = []string{"package.json", "nevermore.json", "README.md"}
		if len(nevermoreJson.FrontendPath) > 0 {
			pathsToInclude = append(pathsToInclude, nevermoreJson.FrontendPath)
		}
		if len(nevermoreJson.PluginPath) > 0 {
			pathsToInclude = append(pathsToInclude, nevermoreJson.PluginPath)
		}
		return nil
	},
	RunE: func(cmd *cobra.Command, args []string) error {
		tar, err := comm.CreateTarGz(basePath, pathsToInclude)
		if err != nil {
			return err
		}
		fileOutPath := packOutput
		if len(fileOutPath) == 0 {
			fileOutPath = fmt.Sprintf("%s.tar.gz", pluginName)
		}
		fo, err := os.Create(fileOutPath)
		if err != nil {
			return err
		}
		defer fo.Close()
		_, err = io.Copy(fo, tar)
		if err != nil {
			return err
		}
		fmt.Printf("Packed plugin %s to %s", pluginName, fileOutPath)
		return nil
	},
}

func init() {
	packCmd.Flags().StringVar(&basePath, "path", ".", "the path to the plugin")
	packCmd.Flags().BoolVar(&noNpm, "no-npm", false, "disables running \"npm run build\" before packing the plugin")
	packCmd.Flags().StringVar(&packOutput, "output", "", "sets the output tarball file (default is ./plugin-name.tar.gz)")
	rootCmd.AddCommand(packCmd)
}
