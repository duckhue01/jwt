/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"

	"github.com/duckhue01/jwt-cli/jwt"
	"github.com/spf13/cobra"
)

const (
	REQUIRE_TOKEN_ERROR = "This command require a JWT token"
)

// decodeCmd represents the decode command
var decodeCmd = &cobra.Command{
	Use:   "decode",
	Short: "Decode a JSON web token",
	Long:  "Decode a JSON web token",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println(len(args))
		if len(args) == 1 {
			res, err := jwt.Decode(args[0])
			if err != nil {

			}
			fmt.Println("Header:")
			fmt.Println("------------")
			fmt.Println(res[0])
			fmt.Println("Claims:")
			fmt.Println("------------")
			fmt.Println(res[1])

		} else {
			fmt.Println(REQUIRE_TOKEN_ERROR)
		}

	},
}

func init() {
	decodeCmd.AddCommand()

	rootCmd.AddCommand(decodeCmd)

	// decodeCmd.Flags().String("toggle", "t", false, "Help message for toggle")

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// decodeCmd.PersistentFlags().String("foo", "", "A help for foo")
}
