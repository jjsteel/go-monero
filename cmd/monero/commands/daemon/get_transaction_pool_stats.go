package daemon

import (
	"fmt"
	"sort"
	"time"

	"github.com/dustin/go-humanize"
	"github.com/spf13/cobra"

	"github.com/jjsteel/go-monero/cmd/monero/display"
	"github.com/jjsteel/go-monero/cmd/monero/options"
	"github.com/jjsteel/go-monero/pkg/rpc/daemon"
)

type getTransactionPoolStatsCommand struct {
	JSON bool
}

func (c *getTransactionPoolStatsCommand) Cmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "get-transaction-pool-stats",
		Short: "statistics about the transaction pool",
		RunE:  c.RunE,
	}

	cmd.Flags().BoolVar(
		&c.JSON,
		"json",
		false,
		"whether or not to output the result as json",
	)

	return cmd
}

func (c *getTransactionPoolStatsCommand) RunE(_ *cobra.Command, _ []string) error {
	ctx, cancel := options.RootOpts.Context()
	defer cancel()

	client, err := options.RootOpts.Client()
	if err != nil {
		return fmt.Errorf("client: %w", err)
	}

	resp, err := client.GetTransactionPoolStats(ctx)
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
func (c *getTransactionPoolStatsCommand) pretty(v *daemon.GetTransactionPoolStatsResult) {
	table := display.NewTable()

	table.AddRow("Bytes Max:", humanize.Bytes(v.PoolStats.BytesMax))
	table.AddRow("Bytes Med:", humanize.Bytes(v.PoolStats.BytesMed))
	table.AddRow("Bytes Min:", humanize.Bytes(v.PoolStats.BytesMin))
	table.AddRow("Bytes Total:", humanize.Bytes(v.PoolStats.BytesTotal))
	table.AddRow("Fee Total:", display.PreciseXMR(v.PoolStats.FeeTotal))
	table.AddRow("Histogram 98pct:", v.PoolStats.Histo98Pc)
	table.AddRow("Txns in Pool for Longer than 10m:", v.PoolStats.Num10M)
	table.AddRow("Double Spends:", v.PoolStats.NumDoubleSpends)
	table.AddRow("Failing Transactions:", v.PoolStats.NumFailing)
	table.AddRow("Not Relayed:", v.PoolStats.NumNotRelayed)
	table.AddRow("Oldest:", humanize.Time(time.Unix(v.PoolStats.Oldest, 0)))
	table.AddRow("Txns Total:", v.PoolStats.TxsTotal)

	table.AddRow("")
	table.AddRow("BYTES", "TXNS")

	sort.Slice(v.PoolStats.Histo, func(i, j int) bool {
		return v.PoolStats.Histo[i].Bytes < v.PoolStats.Histo[j].Bytes
	})
	for _, h := range v.PoolStats.Histo {
		if h.Bytes == 0 {
			continue
		}
		table.AddRow(h.Bytes, h.Txs)
	}

	fmt.Println(table)
}

func init() {
	RootCommand.AddCommand((&getTransactionPoolStatsCommand{}).Cmd())
}
