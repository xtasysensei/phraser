/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"os"
	"strings"

	"github.com/spf13/cobra"
)

// initCmd represents the init command
var initCmd = &cobra.Command{
	Use:   "init",
	Short: "Initializes a store",
	Long: `This creates a directory that will be
and usage of using your command.`,

	Run: initStore,
}

func init() {
	rootCmd.AddCommand(initCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// initCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// initCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

func initStore(cmd *cobra.Command, args []string) {

	if len(args) == 0 {
		cobra.CheckErr(fmt.Errorf("please secify a name for the new store "))
	}

	if len(args) > 1 {
		cobra.CheckErr(fmt.Errorf("Invalid arguments"))
	}

	dir := fmt.Sprintf(strings.Join(args, " "))
	storeName := dir + ".store"

	err := os.Mkdir(storeName, 0750)
	if err != nil && !os.IsExist(err) {
		cobra.CheckErr(err)
	}

	fmt.Println("Initializing Store...")
	fmt.Printf("Store %s successfully initialized", dir)

}
