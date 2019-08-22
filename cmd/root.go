package cmd

import (
	"fmt"
	"os"

	"github.com/piotrpersona/docku/app"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "docku",
	Short: "Upload images to docker registry",
	Long: `Provide docker images config file containing source images
and destination registry.`,
	Args: cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		app.Run(args[0])
	},
}

// Execute will start docku application.
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
