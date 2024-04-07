package lsp

type DidOpenTextNotification struct {
	Notifications
	Params DidOpenTextNotificationParams `json:"params"`
}

type DidOpenTextNotificationParams struct {
	TextDocument TextDocumentItem `json:"textDocument"`
}
