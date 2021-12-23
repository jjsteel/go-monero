package daemon

import (
	"fmt"

	"github.com/spf13/cobra"

	"github.com/jjsteel/go-monero/cmd/monero/display"
	"github.com/jjsteel/go-monero/cmd/monero/options"
	"github.com/jjsteel/go-monero/pkg/rpc/daemon"
)

type relayTxCommand struct {
	Txns []string `long:"txn" required:"true" description:"transaction to relay"`

	JSON bool
}

func (c *relayTxCommand) Cmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "relay-tx",
		Short: "relay a list of transaction ids",
		RunE:  c.RunE,
	}

	cmd.Flags().BoolVar(&c.JSON, "json",
		false, "whether or not to output the result as json")
	cmd.Flags().StringArrayVar(&c.Txns, "txn",
		[]string{}, "id of a transaction to relay")
	_ = cmd.MarkFlagRequired("txn")

	return cmd
}

func (c *relayTxCommand) RunE(_ *cobra.Command, _ []string) error {
	ctx, cancel := options.RootOpts.Context()
	defer cancel()

	client, err := options.RootOpts.Client()
	if err != nil {
		return fmt.Errorf("client: %w", err)
	}

	resp, err := client.RelayTx(ctx, c.Txns)
	if err != nil {
		return fmt.Errorf("relay tx: %w", err)
	}

	if c.JSON {
		return display.JSON(resp)
	}

	c.pretty(resp)
	return nil
}

// nolint:forbidigo
func (c *relayTxCommand) pretty(v *daemon.RelayTxResult) {
	fmt.Println(v.Status)
}

func init() {
	RootCommand.AddCommand((&relayTxCommand{}).Cmd())
}
