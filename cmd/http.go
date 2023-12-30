/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"github.com/wwengg/im/global"
	"github.com/wwengg/im/internal/httpgate"
	"github.com/wwengg/simple/server"
)

var port int

// httpCmd represents the http command
var httpCmd = &cobra.Command{
	Use:   "http",
	Short: "http 网关",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("http called")
		//initHttpPort()

		global.InitSlog()

		global.InitSRPC()

		global.InitDB()

		httpgate.InitGinHttpGate()
		srv := server.NewGateway(&global.CONFIG.Gateway, httpgate.GinEngine)

		srv.Start()

	},
}

func init() {
	rootCmd.AddCommand(httpCmd)

	cobra.OnInitialize(initHttpPort)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	httpCmd.PersistentFlags().IntVar(&port, "port", 0, "http port")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// httpCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

func initHttpPort() {
	if port != 0 {
		global.CONFIG.Gateway.Addr = port
	}
}
