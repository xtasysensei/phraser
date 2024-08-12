/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

// getCmd represents the get command
var getCmd = &cobra.Command{
	Use:   "get",
	Short: "Retrieves data stored in a wallet",
	Long: `This retrieves the wallet phrases stored in the encrypted json files
`,
	Run: retrieveWallet,
}

func init() {
	rootCmd.AddCommand(getCmd)

	createCmd.Flags().StringP("store", "s", "", "name of the store to access")
	createCmd.Flags().StringP("wallet", "w", "", "name of the wallet to be created")
}

type GetWalletPayload struct {
	WalletName    string   `json"walletname"`
	WalletPhrases []string `json"walletphrases"`
}

func retrieveWallet(cmd *cobra.Command, args []string) {
	store, err := cmd.Flags().GetString("store")
	cobra.CheckErr(err)

	wallet, err := cmd.Flags().GetString("wallet")
	cobra.CheckErr(err)

	store = store + ".store"
	if _, err := os.Open(store); os.IsNotExist(err) {
		cobra.CheckErr(err)
	}

	fmt.Println("Retrieving Wallet...")
	fileWallet := store + "." + wallet + ".json"
	filePath := store + "/" + fileWallet

	if _, err := os.Stat(filePath); err == nil {
		cobra.CheckErr(err)
	} else if !os.IsNotExist(err) {
		cobra.CheckErr(err)
	}

	f, _ := os.Open(filePath)
	defer f.Close()

	decoder := json.NewDecoder(f)

	var getWallet GetWalletPayload
	if err := decoder.Decode(&getWallet); err != nil {
		cobra.CheckErr(err)
	}

	fmt.Printf(" Phrases for %s\n", getWallet.WalletName)
	fmt.Println("---------------------------------")
	for i, val := range getWallet.WalletPhrases {
		fmt.Printf("%d. %s\n", i+1, val)
	}

}
