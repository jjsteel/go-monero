package daemon

import (
	"fmt"

	"github.com/spf13/cobra"

	"github.com/jjsteel/go-monero/cmd/monero/display"
	"github.com/jjsteel/go-monero/cmd/monero/options"
	"github.com/jjsteel/go-monero/pkg/rpc/daemon"
)

type getBlockCountCommand struct {
	JSON bool
}

func (c *getBlockCountCommand) Cmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "get-block-count",
		Short: "look up how many blocks are in the longest chain known to the node",
		RunE:  c.RunE,
	}

	cmd.Flags().BoolVar(&c.JSON, "json",
		false, "whether or not to output the result as json")

	return cmd
}

func (c *getBlockCountCommand) RunE(_ *cobra.Command, _ []string) error {
	ctx, cancel := options.RootOpts.Context()
	defer cancel()

	client, err := options.RootOpts.Client()
	if err != nil {
		return fmt.Errorf("client: %w", err)
	}

	resp, err := client.GetBlockCount(ctx)
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
func (c *getBlockCountCommand) pretty(v *daemon.GetBlockCountResult) {
	table := display.NewTable()

	table.AddRow("Count:", v.Count)

	fmt.Println(table)
}

func init() {
	RootCommand.AddCommand((&getBlockCountCommand{}).Cmd())
}
