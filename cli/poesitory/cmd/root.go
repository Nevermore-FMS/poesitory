package cmd

import (
	"errors"
	"log"

	"github.com/spf13/cobra"
	"github.com/spf13/cobra/doc"
)

//flags
var (
	userToken   string
	uploadToken string
	webAuth     bool
)

var rootCmd = &cobra.Command{
	Use:   "poesitory",
	Short: "Poesitory CLI allows you to push and pull Nevermore Plugins to or from Poesitory.",
	CompletionOptions: cobra.CompletionOptions{
		DisableDefaultCmd: true,
	},
	Version:           "0.1.1",
	DisableAutoGenTag: true,
	PersistentPreRunE: func(cmd *cobra.Command, args []string) error {
		var authenticationMethods = 0
		if len(userToken) > 0 {
			authenticationMethods++
		}
		if len(uploadToken) > 0 {
			authenticationMethods++
		}
		if webAuth {
			authenticationMethods++
		}
		if authenticationMethods > 1 {
			return errors.New("only one authentication method can be provided")
		}
		return nil
	},
}

func init() {
	rootCmd.PersistentFlags().StringVarP(&userToken, "user", "u", "", "token for user authentication")
	rootCmd.PersistentFlags().StringVarP(&uploadToken, "upload", "p", "", "token for upload authentication")
	rootCmd.PersistentFlags().BoolVarP(&webAuth, "web", "w", false, "use web authentication")
}

func Execute() {
	cobra.CheckErr(rootCmd.Execute())
}

func GenDocs() {
	err := doc.GenMarkdownTree(rootCmd, "./docs")
	if err != nil {
		log.Fatal(err)
	}
}
