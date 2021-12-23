package daemon

import (
	"fmt"
	"time"

	"github.com/dustin/go-humanize"
	"github.com/spf13/cobra"

	"github.com/jjsteel/go-monero/cmd/monero/display"
	"github.com/jjsteel/go-monero/cmd/monero/options"
	"github.com/jjsteel/go-monero/pkg/rpc/daemon"
)

type getInfoCommand struct {
	JSON bool
}

func (c *getInfoCommand) Cmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "get-info",
		Short: "general information about the node and the network",
		RunE:  c.RunE,
	}

	cmd.Flags().BoolVar(&c.JSON, "json",
		false, "whether or not to output the result as json")

	return cmd
}

func (c *getInfoCommand) RunE(_ *cobra.Command, _ []string) error {
	ctx, cancel := options.RootOpts.Context()
	defer cancel()

	client, err := options.RootOpts.Client()
	if err != nil {
		return fmt.Errorf("client: %w", err)
	}

	resp, err := client.GetInfo(ctx)
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
func (c *getInfoCommand) pretty(v *daemon.GetInfoResult) {
	table := display.NewTable()

	table.AddRow("Adjusted Time:", time.Unix(int64(v.AdjustedTime), 0))
	table.AddRow("Alternative Blocks:", v.AltBlocksCount)
	table.AddRow("Block Size Limit:", humanize.Bytes(v.BlockSizeLimit))
	table.AddRow("Block Size Median:", humanize.Bytes(v.BlockSizeMedian))
	table.AddRow("Block Size Limit:", humanize.Bytes(v.BlockWeightLimit))
	table.AddRow("Block Size Median:", humanize.Bytes(v.BlockWeightMedian))
	table.AddRow("Bootstrap Daemon Address:", v.BootstrapDaemonAddress)
	table.AddRow("Busy Syncing:", v.BusySyncing)
	table.AddRow("Cumulative Difficulty:", v.CumulativeDifficulty)
	table.AddRow("DatabaseSize:", humanize.Bytes(v.DatabaseSize))
	table.AddRow("Difficulty:", v.Difficulty)
	table.AddRow("Free Space:", humanize.Bytes(v.FreeSpace))
	table.AddRow("Grey Peer List Size:", v.GreyPeerlistSize)
	table.AddRow("Height:", v.Height)
	table.AddRow("Height Without Bootstrap:", v.HeightWithoutBootstrap)
	table.AddRow("Incoming Connections:", v.IncomingConnectionsCount)
	table.AddRow("Mainnet:", v.Mainnet)
	table.AddRow("Nettype:", v.Nettype)
	table.AddRow("Offline:", v.Offline)
	table.AddRow("Outgoing Connections:", v.OutgoingConnectionsCount)
	table.AddRow("RPC Connections:", v.RPCConnectionsCount)
	table.AddRow("Stagenet:", v.Stagenet)
	table.AddRow("Start Time:", time.Unix(int64(v.StartTime), 0))
	table.AddRow("Synchronized:", v.Synchronized)
	table.AddRow("Target:", v.Target)
	table.AddRow("Target Height:", v.TargetHeight)
	table.AddRow("Testnet:", v.Testnet)
	table.AddRow("TopBlockHash:", v.TopBlockHash)
	table.AddRow("Transactions:", v.TxCount)
	table.AddRow("Transaction Pool Size:", v.TxPoolSize)
	table.AddRow("Update Available:", v.UpdateAvailable)
	table.AddRow("Version:", v.Version)
	table.AddRow("Was Bootstrap Ever Used:", v.WasBootstrapEverUsed)
	table.AddRow("White Peer List:", v.WhitePeerlistSize)
	table.AddRow("WideCumulativeDifficulty:", v.WideCumulativeDifficulty)
	table.AddRow("WideDifficulty:", v.WideDifficulty)

	fmt.Println(table)
}

func init() {
	RootCommand.AddCommand((&getInfoCommand{}).Cmd())
}
