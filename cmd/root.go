package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "prometheus",
	Short: "Prometheus demo server",
	Long:  `A Prometheus demo server`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("root called")

	},
}

func Execute() {
	fmt.Println("printed")
	err := rootCmd.Execute()
	if err != nil {
		fmt.Println("error")
		os.Exit(1)
	}
}
func init() {

	fmt.Println("never")

}
