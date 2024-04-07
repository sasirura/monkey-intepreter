package analysis

type State struct {
	// Map of file names to contents
	Documents map[string]string
}

func NewState() State {
	return State{Documents: map[string]string{}}
}

func (s *State) OpenDocument(document, text string) {
	s.Documents[document] = text
}

func (s *State) UpdateDocument(uri, text string) {
	s.Documents[uri] = text
}
