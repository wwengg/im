/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
<<<<<<< HEAD
=======
	"github.com/wwengg/im/global"
	"github.com/wwengg/im/internal/gate"

>>>>>>> dev
	"github.com/spf13/cobra"
	"github.com/wwengg/im/global"
	"github.com/wwengg/im/internal/gate"
)

// gateCmd represents the gate command
var gateCmd = &cobra.Command{
	Use:   "gate",
	Short: "A brief description of your command",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("gate called")
		global.InitSlog()

		global.InitSRPC()
<<<<<<< HEAD

		global.InitDB()
		
		gate.InitGateTcp()

		gate.InitRpcxService("GateRPCX", global.CONFIG.RPC, global.CONFIG.RpcService)
=======

		global.InitDB()

		gate.InitRpcxService(global.CONFIG.Zinx.Name, global.CONFIG.RPC, global.CONFIG.RpcService)

		gate.InitGateTcp()
>>>>>>> dev

	},
}

func init() {
	rootCmd.AddCommand(gateCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// gateCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// gateCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
