package main

import (
	_ "github.com/rl404/mal-db/api"
	"github.com/spf13/cobra"
)

// @title MAL-DB API
// @description MyAnimeList database dump and API.
// @contact.name Axel
// @contact.url https://github.com/rl404
// @contact.email axel.rl.404@gmail.com
// @license.name MIT
// @license.url https://opensource.org/licenses/MIT
// @BasePath /
// @schemes http https
func main() {
	cmd := cobra.Command{
		Use:   "mal",
		Short: "MAL DB API",
	}

	cmd.AddCommand(&cobra.Command{
		Use:   "install",
		Short: "Prepare database",
		Run: func(*cobra.Command, []string) {
			install()
		},
	})

	cmd.AddCommand(&cobra.Command{
		Use:   "server",
		Short: "Run API server",
		Run: func(*cobra.Command, []string) {
			server()
		},
	})

	cmd.AddCommand(&cobra.Command{
		Use:   "worker",
		Short: "Run worker",
		Run: func(*cobra.Command, []string) {
			worker()
		},
	})

	toolCmd := cobra.Command{
		Use:   "tools",
		Short: "Run tools",
	}
	cmd.AddCommand(&toolCmd)

	toolCmd.AddCommand(&cobra.Command{
		Use:   "filler",
		Short: "Run filler tools",
		Run: func(*cobra.Command, []string) {
			filler()
		},
	})

	toolCmd.AddCommand(&cobra.Command{
		Use:   "updater",
		Short: "Run updater tools",
		Run: func(*cobra.Command, []string) {
			updater()
		},
	})

	if err := cmd.Execute(); err != nil {
		panic(err)
	}
}
