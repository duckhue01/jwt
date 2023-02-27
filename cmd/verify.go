/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// verifyCmd represents the verify command
var verifyCmd = &cobra.Command{
	Use:   "verify",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println(cmd.Flag("is-expire").Value.String())
		fmt.Println(cmd.Flag("contain").Value.String())
		fmt.Println(cmd.Flag("is-valid").Value.String())

	},
}

func init() {
	rootCmd.AddCommand(verifyCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// verifyCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	verifyCmd.Flags().Bool("is-expire", false, "Verify if token is expire or not")
	verifyCmd.Flags().StringSlice("contain", make([]string, 0), "Verify if Token contains specified claims")
	verifyCmd.Flags().Bool("is-valid", false, "Verify if token is a valid token or not. It means that token is issued from trusted party and doesn't change")

}
