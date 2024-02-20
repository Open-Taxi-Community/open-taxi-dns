package dnsapi

import (
	"fmt"
	"golang.org/x/net/http2"
	"net/http"
	"os"

	"github.com/Open-Taxi-Community/open-taxi-dns/internal/dnsapi"
	"github.com/Open-Taxi-Community/open-taxi-dns/pkg/dnsapi/v1/dnsapiconnect"
	"github.com/spf13/cobra"
	"golang.org/x/net/http2/h2c"
)

var (
	projectId string
	httpPort  int
)

func init() {

	globalFlags := rootCmd.PersistentFlags()
	globalFlags.StringVar(&projectId, "project-id", os.Getenv("PROJECT_ID"), "Project ID of firestore project")
	globalFlags.IntVar(&httpPort, "http-port", 8000, "Port to run the HTTP server on")

	rootCmd.AddCommand()
	rootCmd.AddCommand(run)
	rootCmd.AddCommand(runLocal)
}

var rootCmd = &cobra.Command{
	Use:   "dnsapi",
	Short: "Root command for the DNS API server",
	PersistentPostRunE: func(cmd *cobra.Command, args []string) error {

		return nil
	},
}

var run = &cobra.Command{
	Use:   "run",
	Short: "command to start the DNS API server",
	RunE: func(cmd *cobra.Command, args []string) error {
		runServer("", httpPort)
		return nil
	},
}

var runLocal = &cobra.Command{
	Use:   "run-local",
	Short: "command to start the DNS API server locally",
	RunE: func(cmd *cobra.Command, args []string) error {
		runServer("localhost", httpPort)
		return nil
	},
}

func runServer(address string, port int) error {
	mux := http.NewServeMux()
	mux.Handle(dnsapiconnect.NewDnsServiceHandler(dnsapi.NewDnsApiServer(), nil))

	fmt.Println("Starting DNS API server on", fmt.Sprintf("%s:%d", address, port))
	return http.ListenAndServe(
		fmt.Sprintf("%s:%d", address, port),
		// Use h2c so we can use HTTP/2 without TLS
		h2c.NewHandler(mux, &http2.Server{}),
	)
}
