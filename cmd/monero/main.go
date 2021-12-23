package main

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"

	"github.com/jjsteel/go-monero/cmd/monero/commands/address"
	"github.com/jjsteel/go-monero/cmd/monero/commands/daemon"
	"github.com/jjsteel/go-monero/cmd/monero/commands/p2p"
	"github.com/jjsteel/go-monero/cmd/monero/commands/wallet"
)

var (
	version = "dev"
	commit  = "dev"
)

var rootCmd = &cobra.Command{
	Use:   "monero",
	Short: "Daemon, Wallet, and p2p command line monero CLI",
}

// nolint:forbidigo
var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "print the version of this cli",
	Run: func(_ *cobra.Command, _ []string) {
		fmt.Println(version, commit)
	},
}

func init() {
	rootCmd.AddCommand(daemon.RootCommand)
	rootCmd.AddCommand(wallet.RootCommand)
	rootCmd.AddCommand(p2p.RootCommand)
	rootCmd.AddCommand(address.RootCommand)
	rootCmd.AddCommand(versionCmd)
}

func main() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
