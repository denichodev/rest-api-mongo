package cmd

import (
	"fmt"
	"net/http"

	"github.com/denichodev/rest-api-mongo/handler"
	"github.com/spf13/cobra"
)

// serveCmd represents the serve command
var serveCmd = &cobra.Command{
	Use:   "serve",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: run,
}

func run(*cobra.Command, []string) {
	router := handler.CreateRouter()

	fmt.Println("App running on port 3030")
	http.ListenAndServe(":3030", router)
}

func init() {
	rootCmd.AddCommand(serveCmd)
}
