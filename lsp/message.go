package lsp

type Request struct {
	RPC    string `json:"jsonrpc"`
	ID     int    `json:"id"`
	Method string `json:"method"`

	// params
}

type Response struct {
	RPC string `json:"jsonrpc"`
	ID  *int   `json:"id,omitempty"`

	// result
	// error
}

type NotificationMessage struct {
	RPC    string `json:"jsonrpc"`
	Method string `json:"method"`
}
