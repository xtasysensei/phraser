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

// createCmd represents the create command
var createCmd = &cobra.Command{
	Use:   "create",
	Short: "Creates a wallet in your store",
	Long:  `This creates an encrypted json file in your store to store the phrases`,
	Run:   createWallet,
}

func init() {
	rootCmd.AddCommand(createCmd)

	createCmd.Flags().StringP("store", "s", "", "name of the store to access")
	createCmd.Flags().StringP("wallet", "w", "", "name of the wallet to be created")
	createCmd.Flags().IntP("amount", "a", 0, "amount of phrases to be inputed")
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

	fmt.Printf("Preparing Wallet %s for input...\n", wallet)
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

	// NOTE:using json.Marshal
	//jsonData, err := json.Marshal(w)
	//cobra.CheckErr(err)
	//err = ioutil.WriteFile(filePath, jsonData, 0644)
	//cobra.CheckErr(err)

	encoder := json.NewEncoder(f)
	encoder.Encode(w)

	fmt.Printf("Phrases successfully added to Wallet %s\n", wallet)

}
