package main

// cobra cli that create and get all users

import (
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "app",
	Short: "app is a CLI for the user management",
}

var createUserCmd = &cobra.Command{
	Use:   "create",
	Short: "Create a user",
	Long:  `Create a user with the given name and email`,
	Run: func(cmd *cobra.Command, args []string) {
		//TODO implement me
		panic("implement me")
	},
}

var getAllUsersCmd = &cobra.Command{
	Use:   "get",
	Short: "Get all users",
	Long:  `Get all users`,
	Run: func(cmd *cobra.Command, args []string) {
		//TODO implement me
		panic("implement me")
	},
}

func main() {
	if err := rootCmd.Execute(); err != nil {
		panic(err)
	}

}

func init() {
	rootCmd.AddCommand(createUserCmd)
	rootCmd.AddCommand(getAllUsersCmd)
}
