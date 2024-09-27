/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"github.com/RejwankabirHamim/Prometheus/api"

	"github.com/spf13/cobra"
)

var Port string

// startCmd represents the start command
var startCmd = &cobra.Command{
	Use:   "start",
	Short: "Start the server",
	Long:  "Start the server",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("start called")
		api.RunServer(Port)
	},
}

func init() {
	startCmd.PersistentFlags().StringVarP(&Port, "port", "p", "8181", "Port to listen on")
	rootCmd.AddCommand(startCmd)
}
