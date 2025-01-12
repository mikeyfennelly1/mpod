package main

import (
	"fmt"
	"github.com/spf13/cobra"
)

const MPOD_VERSION = "0.0.1"

func main() {
	var rootCmd = &cobra.Command{
		Use:   "mpod",
		Short: "A CLI tool for creating and managing Linux containers",
	}

	var versionCmd = &cobra.Command{
		Use:   "version",
		Short: "Print the version of mpod",
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Printf("mpod %s\n", MPOD_VERSION)
		},
	}

	rootCmd.AddCommand(versionCmd)

	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
	}
}
