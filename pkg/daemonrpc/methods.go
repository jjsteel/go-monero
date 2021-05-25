package daemonrpc

import (
	"context"
	"fmt"
)

const (
	MethodGetBlockCount      = "get_block_count"
	MethodGetBlockTemplate   = "get_block_template"
	MethodGetCoinbaseTxSum   = "get_coinbase_tx_sum"
	MethodGetConnections     = "get_connections"
	MethodGetFeeEstimate     = "get_fee_estimate"
	MethodGetInfo            = "get_info"
	MethodGetLastBlockHeader = "get_last_block_header"
	MethodOnGetBlockHash     = "on_get_block_hash"
	MethodSyncInfo           = "sync_info"

	EndpointGetTransactionPool      = "/get_transaction_pool"
	EndpointGetTransactionPoolStats = "/get_transaction_pool_stats"
)

type GetBlockCountResult struct {
	Count  uint64 `json:"count"`
	Status string `json:"status"`
}

func (c *Client) GetBlockCount(ctx context.Context) (*GetBlockCountResult, error) {
	var resp = &GetBlockCountResult{}

	if err := c.JsonRPC(ctx, MethodGetBlockCount, nil, resp); err != nil {
		return nil, fmt.Errorf("get: %w", err)
	}

	return resp, nil
}

func (c *Client) OnGetBlockHash(ctx context.Context, height uint64) (string, error) {
	var (
		resp   = ""
		params = []uint64{height}
	)

	if err := c.JsonRPC(ctx, MethodOnGetBlockHash, params, &resp); err != nil {
		return "", fmt.Errorf("get: %w", err)
	}

	return resp, nil
}

type GetBlockTemplateResult struct {
	BlockhashingBlob  string `json:"blockhashing_blob"`
	BlocktemplateBlob string `json:"blocktemplate_blob"`
	Difficulty        int64  `json:"difficulty"`
	ExpectedReward    int64  `json:"expected_reward"`
	Height            int    `json:"height"`
	PrevHash          string `json:"prev_hash"`
	ReservedOffset    int    `json:"reserved_offset"`
	Status            string `json:"status"`
	Untrusted         bool   `json:"untrusted"`
}

func (c *Client) GetBlockTemplate(ctx context.Context, walletAddress string, reserveSize uint) (*GetBlockTemplateResult, error) {
	var (
		resp   = &GetBlockTemplateResult{}
		params = map[string]interface{}{
			"wallet_address": walletAddress,
			"reserve_size":   reserveSize,
		}
	)

	if err := c.JsonRPC(ctx, MethodGetBlockTemplate, params, resp); err != nil {
		return nil, fmt.Errorf("get: %w", err)
	}

	return resp, nil
}

type GetConnectionsResult struct {
	Connections []struct {
		Address         string `json:"address"`
		AvgDownload     int    `json:"avg_download"`
		AvgUpload       int    `json:"avg_upload"`
		ConnectionID    string `json:"connection_id"`
		CurrentDownload int    `json:"current_download"`
		CurrentUpload   int    `json:"current_upload"`
		Height          int    `json:"height"`
		Host            string `json:"host"`
		Incoming        bool   `json:"incoming"`
		IP              string `json:"ip"`
		LiveTime        int    `json:"live_time"`
		LocalIP         bool   `json:"local_ip"`
		Localhost       bool   `json:"localhost"`
		PeerID          string `json:"peer_id"`
		Port            string `json:"port"`
		RecvCount       int    `json:"recv_count"`
		RecvIdleTime    int    `json:"recv_idle_time"`
		SendCount       int    `json:"send_count"`
		SendIdleTime    int    `json:"send_idle_time"`
		State           string `json:"state"`
		SupportFlags    int    `json:"support_flags"`
	} `json:"connections"`
	Status string `json:"status"`
}

func (c *Client) GetConnections(ctx context.Context) (*GetConnectionsResult, error) {
	var resp = &GetConnectionsResult{}

	if err := c.JsonRPC(ctx, MethodGetConnections, nil, resp); err != nil {
		return nil, fmt.Errorf("get: %w", err)
	}

	return resp, nil
}

type GetInfoResult struct {
	AltBlocksCount           int    `json:"alt_blocks_count"`
	BlockSizeLimit           int    `json:"block_size_limit"`
	BlockSizeMedian          int    `json:"block_size_median"`
	BootstrapDaemonAddress   string `json:"bootstrap_daemon_address"`
	BusySyncing              bool   `json:"busy_syncing"`
	CumulativeDifficulty     int64  `json:"cumulative_difficulty"`
	Difficulty               int64  `json:"difficulty"`
	FreeSpace                int64  `json:"free_space"`
	GreyPeerlistSize         int    `json:"grey_peerlist_size"`
	Height                   int    `json:"height"`
	HeightWithoutBootstrap   int    `json:"height_without_bootstrap"`
	IncomingConnectionsCount int    `json:"incoming_connections_count"`
	Mainnet                  bool   `json:"mainnet"`
	Offline                  bool   `json:"offline"`
	OutgoingConnectionsCount int    `json:"outgoing_connections_count"`
	RPCConnectionsCount      int    `json:"rpc_connections_count"`
	Stagenet                 bool   `json:"stagenet"`
	StartTime                int    `json:"start_time"`
	Status                   string `json:"status"`
	Synchronized             bool   `json:"synchronized"`
	Target                   int    `json:"target"`
	TargetHeight             int    `json:"target_height"`
	Testnet                  bool   `json:"testnet"`
	TopBlockHash             string `json:"top_block_hash"`
	TxCount                  int    `json:"tx_count"`
	TxPoolSize               int    `json:"tx_pool_size"`
	Untrusted                bool   `json:"untrusted"`
	WasBootstrapEverUsed     bool   `json:"was_bootstrap_ever_used"`
	WhitePeerlistSize        int    `json:"white_peerlist_size"`
}

func (c *Client) GetInfo(ctx context.Context) (*GetInfoResult, error) {
	var resp = &GetInfoResult{}

	if err := c.JsonRPC(ctx, MethodGetInfo, nil, resp); err != nil {
		return nil, fmt.Errorf("get: %w", err)
	}

	return resp, nil
}

type GetLastBlockHeaderResult struct {
	BlockHeader struct {
		BlockSize                 int    `json:"block_size"`
		BlockWeight               int    `json:"block_weight"`
		CumulativeDifficulty      int64  `json:"cumulative_difficulty"`
		CumulativeDifficultyTop64 int    `json:"cumulative_difficulty_top64"`
		Depth                     int    `json:"depth"`
		Difficulty                int64  `json:"difficulty"`
		DifficultyTop64           int    `json:"difficulty_top64"`
		Hash                      string `json:"hash"`
		Height                    int    `json:"height"`
		LongTermWeight            int    `json:"long_term_weight"`
		MajorVersion              int    `json:"major_version"`
		MinerTxHash               string `json:"miner_tx_hash"`
		MinorVersion              int    `json:"minor_version"`
		Nonce                     int    `json:"nonce"`
		NumTxes                   int    `json:"num_txes"`
		OrphanStatus              bool   `json:"orphan_status"`
		PowHash                   string `json:"pow_hash"`
		PrevHash                  string `json:"prev_hash"`
		Reward                    int64  `json:"reward"`
		Timestamp                 int    `json:"timestamp"`
		WideCumulativeDifficulty  string `json:"wide_cumulative_difficulty"`
		WideDifficulty            string `json:"wide_difficulty"`
	} `json:"block_header"`
}

func (c *Client) GetLastBlockHeader(ctx context.Context) (*GetLastBlockHeaderResult, error) {
	var resp = &GetLastBlockHeaderResult{}

	if err := c.JsonRPC(ctx, MethodGetLastBlockHeader, nil, resp); err != nil {
		return nil, fmt.Errorf("get: %w", err)
	}

	return resp, nil
}

type GetCoinbaseTxSumResult struct {
	Credits             int    `json:"credits"`
	EmissionAmount      int64  `json:"emission_amount"`
	EmissionAmountTop64 int    `json:"emission_amount_top64"`
	FeeAmount           int    `json:"fee_amount"`
	FeeAmountTop64      int    `json:"fee_amount_top64"`
	Status              string `json:"status"`
	TopHash             string `json:"top_hash"`
	Untrusted           bool   `json:"untrusted"`
	WideEmissionAmount  string `json:"wide_emission_amount"`
	WideFeeAmount       string `json:"wide_fee_amount"`
}

func (c *Client) GetCoinbaseTxSum(ctx context.Context, height, count uint64) (*GetCoinbaseTxSumResult, error) {
	var (
		resp   = &GetCoinbaseTxSumResult{}
		params = map[string]uint64{
			"height": height,
			"count":  count,
		}
	)

	if err := c.JsonRPC(ctx, MethodGetCoinbaseTxSum, params, resp); err != nil {
		return nil, fmt.Errorf("get: %w", err)
	}

	return resp, nil
}

type GetFeeEstimateResult struct {
	Credits          int    `json:"credits"`
	Fee              int    `json:"fee"`
	QuantizationMask int    `json:"quantization_mask"`
	Status           string `json:"status"`
	TopHash          string `json:"top_hash"`
	Untrusted        bool   `json:"untrusted"`
}

func (c *Client) GetFeeEstimate(ctx context.Context, graceBlocks uint64) (*GetFeeEstimateResult, error) {
	var (
		resp   = new(GetFeeEstimateResult)
		params = map[string]uint64{
			"grace_blocks": graceBlocks,
		}
	)

	if err := c.JsonRPC(ctx, MethodGetFeeEstimate, params, resp); err != nil {
		return nil, fmt.Errorf("get: %w", err)
	}

	return resp, nil
}

type SyncInfoResult struct {
	Credits               int    `json:"credits"`
	Height                int    `json:"height"`
	NextNeededPruningSeed int    `json:"next_needed_pruning_seed"`
	Overview              string `json:"overview"`
	Status                string `json:"status"`
	TargetHeight          int    `json:"target_height"`
	TopHash               string `json:"top_hash"`
	Untrusted             bool   `json:"untrusted"`
	Peers                 []struct {
		Info struct {
			Address           string `json:"address"`
			AddressType       int    `json:"address_type"`
			AvgDownload       int    `json:"avg_download"`
			AvgUpload         int    `json:"avg_upload"`
			ConnectionID      string `json:"connection_id"`
			CurrentDownload   int    `json:"current_download"`
			CurrentUpload     int    `json:"current_upload"`
			Height            int    `json:"height"`
			Host              string `json:"host"`
			Incoming          bool   `json:"incoming"`
			IP                string `json:"ip"`
			LiveTime          int    `json:"live_time"`
			LocalIP           bool   `json:"local_ip"`
			Localhost         bool   `json:"localhost"`
			PeerID            string `json:"peer_id"`
			Port              string `json:"port"`
			PruningSeed       int    `json:"pruning_seed"`
			RecvCount         int    `json:"recv_count"`
			RecvIdleTime      int    `json:"recv_idle_time"`
			RPCCreditsPerHash int    `json:"rpc_credits_per_hash"`
			RPCPort           int    `json:"rpc_port"`
			SendCount         int    `json:"send_count"`
			SendIdleTime      int    `json:"send_idle_time"`
			State             string `json:"state"`
			SupportFlags      int    `json:"support_flags"`
		} `json:"info"`
	} `json:"peers"`
}

func (c *Client) SyncInfo(ctx context.Context) (*SyncInfoResult, error) {
	var resp = new(SyncInfoResult)

	if err := c.JsonRPC(ctx, MethodSyncInfo, nil, resp); err != nil {
		return nil, fmt.Errorf("get: %w", err)
	}

	return resp, nil
}

type GetTransactionPoolResult struct {
	Credits        int `json:"credits"`
	SpentKeyImages []struct {
		IDHash    string   `json:"id_hash"`
		TxsHashes []string `json:"txs_hashes"`
	} `json:"spent_key_images"`
	Status       string `json:"status"`
	TopHash      string `json:"top_hash"`
	Transactions []struct {
		BlobSize           int    `json:"blob_size"`
		DoNotRelay         bool   `json:"do_not_relay"`
		DoubleSpendSeen    bool   `json:"double_spend_seen"`
		Fee                int    `json:"fee"`
		IDHash             string `json:"id_hash"`
		KeptByBlock        bool   `json:"kept_by_block"`
		LastFailedHeight   int    `json:"last_failed_height"`
		LastFailedIDHash   string `json:"last_failed_id_hash"`
		LastRelayedTime    int    `json:"last_relayed_time"`
		MaxUsedBlockHeight int    `json:"max_used_block_height"`
		MaxUsedBlockIDHash string `json:"max_used_block_id_hash"`
		ReceiveTime        int    `json:"receive_time"`
		Relayed            bool   `json:"relayed"`
		TxBlob             string `json:"tx_blob"`
		TxJSON             string `json:"tx_json"`
		Weight             int    `json:"weight"`
	} `json:"transactions"`
	Untrusted bool `json:"untrusted"`
}

func (c *Client) GetTransactionPool(ctx context.Context) (*GetTransactionPoolResult, error) {
	var (
		resp = new(GetTransactionPoolResult)
	)

	if err := c.Other(ctx, EndpointGetTransactionPool, nil, resp); err != nil {
		return nil, fmt.Errorf("other: %w", err)
	}

	return resp, nil
}

type GetTransactionPoolStatsResult struct {
	Credits   int `json:"credits"`
	PoolStats struct {
		BytesMax   int `json:"bytes_max"`
		BytesMed   int `json:"bytes_med"`
		BytesMin   int `json:"bytes_min"`
		BytesTotal int `json:"bytes_total"`
		FeeTotal   int `json:"fee_total"`
		Histo      []struct {
			Bytes int `json:"bytes"`
			Txs   int `json:"txs"`
		} `json:"histo"`
		Histo98Pc       int `json:"histo_98pc"`
		Num10M          int `json:"num_10m"`
		NumDoubleSpends int `json:"num_double_spends"`
		NumFailing      int `json:"num_failing"`
		NumNotRelayed   int `json:"num_not_relayed"`
		Oldest          int `json:"oldest"`
		TxsTotal        int `json:"txs_total"`
	} `json:"pool_stats"`
	Status    string `json:"status"`
	TopHash   string `json:"top_hash"`
	Untrusted bool   `json:"untrusted"`
}

func (c *Client) GetTransactionPoolStats(ctx context.Context) (*GetTransactionPoolStatsResult, error) {
	var (
		resp = new(GetTransactionPoolStatsResult)
	)

	if err := c.Other(ctx, EndpointGetTransactionPoolStats, nil, resp); err != nil {
		return nil, fmt.Errorf("other: %w", err)
	}

	return resp, nil
}
