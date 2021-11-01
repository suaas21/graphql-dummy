package main

import (
	"log"

	"github.com/spf13/cobra"
)

var cfgPath string

// rootCmd is the root of all sub commands in the binary
// it doesn't have a Run method as it executes other sub commands
var rootCmd = &cobra.Command{
	Use:     "graphql",
	Short:   "graphql-dummy is a http server to serve public facing api",
	Version: "1.0",
}

func init() {
	// Here all other sub commands should be registered to the rootCmd
	rootCmd.AddCommand(srvCmd)
}

func main() {
	if err := rootCmd.Execute(); err != nil {
		log.Fatal(err)
	}
}
