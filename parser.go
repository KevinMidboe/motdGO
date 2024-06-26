package motdGO

import "errors"

// Parser stores some output specific stuff
type Parser struct {
	// Name used for switching in colors.go
	Name string
	// Parser prefix
	Prefix string
	// Parser suffix
	Suffix string
	// Newline representation
	NewLine string
	// Things to be replaced (f.e. \n to <br>)
	Replaces map[string]string
}

var parsers map[string]Parser = map[string]Parser{

	// Default terminal parser
	"terminal": {"terminal", "", "", "\n", nil},
	// Parser for HTML code
	"html": {"html", "<code>", "</code>", "<br>", map[string]string{" ": "&nbsp;"}},
	// Parser for motd file
	"motd": {"motd", "#!/bin/sh\ncat << 'EOF'\n", "EOF", "\n", nil},
}

// GetParser returns a parser by its key
func GetParser(key string) (*Parser, error) {
	parser, ok := parsers[key]
	if !ok {
		return nil, errors.New("Invalid parser key: " + key)
	}
	return &parser, nil
}
