package message

import "github.com/cocktail828/go-tools/messagepb"

type parser func(msg *messagepb.Message) (*Parsed, error)

var parserRegistry = map[string]parser{}

func RegisterParser(sub string, p parser) (overwrite bool) {
	if p == nil {
		return false
	}
	_, ok := parserRegistry[sub]
	parserRegistry[sub] = p
	return ok
}
