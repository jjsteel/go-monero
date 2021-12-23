package daemon

import (
	"fmt"

	"github.com/spf13/cobra"

	"github.com/jjsteel/go-monero/cmd/monero/display"
	"github.com/jjsteel/go-monero/cmd/monero/options"
	"github.com/jjsteel/go-monero/pkg/rpc/daemon"
)

type getHeightCommand struct {
	JSON bool
}

func (c *getHeightCommand) Cmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "get-height",
		Short: "node's current chain height",
		Long: `Retrieves the current chain height (most recent block height + 1)
including the hash of the most recent block.`,
		RunE: c.RunE,
	}

	cmd.Flags().BoolVar(&c.JSON, "json",
		false, "whether or not to output the result as json")

	return cmd
}

func (c *getHeightCommand) RunE(_ *cobra.Command, _ []string) error {
	ctx, cancel := options.RootOpts.Context()
	defer cancel()

	client, err := options.RootOpts.Client()
	if err != nil {
		return fmt.Errorf("client: %w", err)
	}

	resp, err := client.GetHeight(ctx)
	if err != nil {
		return fmt.Errorf("get block count: %w", err)
	}
	if c.JSON {
		return display.JSON(resp)
	}

	c.pretty(resp)
	return nil
}

// nolint:forbidigo
func (c *getHeightCommand) pretty(v *daemon.GetHeightResult) {
	table := display.NewTable()

	table.AddRow("Hash:", v.Hash)
	table.AddRow("Height:", v.Height)

	fmt.Println(table)
}

func init() {
	RootCommand.AddCommand((&getHeightCommand{}).Cmd())
}
