/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"github.com/wwengg/im/global"
	"github.com/wwengg/im/internal/logic"
)

// logicCmd represents the logic command
var logicCmd = &cobra.Command{
	Use:   "logic",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("logic called")
		global.InitSlog()

		global.InitSRPC()

		global.InitEtcdV3()

		global.InitDB()

		global.InitRedisBase()

		logic.ServeForLogic()

		logic.InitRpcxService("Logic", global.CONFIG.RPC, global.CONFIG.RpcService)
	},
}

func init() {
	rootCmd.AddCommand(logicCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// logicCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// logicCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
