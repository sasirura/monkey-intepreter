package lsp

type DefinitionRequest struct {
	Request
	Params DefinitionParams `json:"params"`
}

type DefinitionParams struct {
	TextDocumentPosistionParams
}

type DefinitionResponse struct {
	Response
	Result Location `json:"result"`
}
