package main

import (
	"errors"
	"fmt"
	"os"

	"github.com/spf13/cobra"

	"goesp/ping"
	"goesp/server"
)

var (
	ip string
)

func main() {
	rootCommand := &cobra.Command{
		Use:   "goesp",
		Short: "use goesp to connect to a esp module and perform operations",
		Long: `goesp acts as a gateway between the esp module and your machine.
	It lets you expose an api endpoint where you can send data from the esp module and 
	store it in a database.
	`,
	}

	pingCommand := cobra.Command{
		Use:     "ping",
		Short:   "ping lets you ping an esp module",
		Example: "ping --ip 192.168.58.50",
		Run:     run,
	}

	serveCommand := cobra.Command{
		Use:     "serve",
		Short:   "serve lets you listen to the data sent by esp module",
		Example: "serve --port 3000",
		Run:     serve,
	}

	rootCommand.PersistentFlags().StringVar(&ip, "ip", "", "ping allows you to communicate with esp module")

	rootCommand.AddCommand(&pingCommand)
	rootCommand.AddCommand(&serveCommand)

	if err := rootCommand.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

}

func serve(cmd *cobra.Command, _ []string) {
	server.Serve()
}

func run(cmd *cobra.Command, _ []string) {
	if ip != "" {
		rtt, err := ping.Ping(ip)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		fmt.Printf("ping returned from ip %s in %v\n", ip, rtt.Seconds())
	} else {
		cmd.Help()
		fmt.Println(errors.New("please specify ip address using -ip flag"))
		os.Exit(1)
	}

}
