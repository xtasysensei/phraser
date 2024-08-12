/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/spf13/cobra"
	"github.com/xtasysensei/phraser/cmd/utils"
)

// createCmd represents the create command
var createCmd = &cobra.Command{
	Use:   "create",
	Short: "Creates a wallet in your store",
	Long:  `This creates an encrypted json file in your store to store the phrases`,
	Run:   createWallet,
}

func init() {
	rootCmd.AddCommand(createCmd)

}

type WalletPayload struct {
	WalletName      string   `json:"walletname"`
	NumberOfPhrases int      `json:"number_of_phrases"`
	WalletPhrases   []string `json:"walletphrases"`
}

func createWallet(cmd *cobra.Command, args []string) {
	store, err := cmd.Flags().GetString("store")
	cobra.CheckErr(err)

	wallet, err := cmd.Flags().GetString("wallet")
	cobra.CheckErr(err)

	numberOfPhrases, err := cmd.Flags().GetInt("amount")
	cobra.CheckErr(err)

	store = store + ".store"
	if _, err := os.Open(store); os.IsNotExist(err) {
		cobra.CheckErr(err)
	}
	var passphrase string
	fmt.Printf("Enter encrypt/decrypt passphrase> ")
	fmt.Scan(&passphrase)

	fmt.Println("Initializing Wallet...")
	fileWallet := store + "." + wallet + ".json"
	filePath := store + "/" + fileWallet

	if _, err := os.Stat(filePath); err == nil {
		cobra.CheckErr(err)
	} else if !os.IsNotExist(err) {
		cobra.CheckErr(err)
	}

	f, err := os.OpenFile(filePath, os.O_CREATE|os.O_EXCL|os.O_WRONLY, 0644)
	cobra.CheckErr(err)
	fmt.Printf("Wallet %s successfully created in %s\n", wallet, store)
	defer f.Close()

	fmt.Println("---------------------------------")
	fmt.Printf("Preparing Wallet %s for input...\n", wallet)
	fmt.Println("---------------------------------")
	phrases := make([]string, numberOfPhrases)

	var phrase string
	for i := range phrases {
		fmt.Printf("Enter phrase %d> ", i+1)
		fmt.Scan(&phrase)
		phrases[i] = phrase
	}

	w := WalletPayload{
		WalletName:      wallet,
		NumberOfPhrases: numberOfPhrases,
		WalletPhrases:   phrases,
	}

	jsonData, err := json.Marshal(w)
	encryptedData := utils.EncryptFile(passphrase, jsonData)
	cobra.CheckErr(err)
	err = os.WriteFile(filePath, encryptedData, 0644)
	cobra.CheckErr(err)

	fmt.Printf("Phrases successfully added to Wallet %s\n", wallet)

}
