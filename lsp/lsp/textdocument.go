package lsp

type TextDocumentItem struct {
	URI        string `json:"uri"`
	LanguageID string `json:"languageId"`
	Version    int    `json:"version"`
	Text       string `json:"text"`
}

type TextDocumentIdentifer struct {
	URI string `json:"uri"`
}

type VersionTextDocumentIdentifier struct {
	TextDocumentIdentifer
	Version int `json:"version"`
}

type TextDocumentPosistionParams struct {
	TextDocument TextDocumentIdentifer `json:"textDocument"`
	Position     Position              `json:"position"`
}

type Position struct {
	Line      int `json:"line"`
	Character int `json:"character"`
}
