package message

import (
	"sync"

	"github.com/cocktail828/go-tools/messagepb"
	"google.golang.org/protobuf/proto"
)

type Parsed struct {
	Data interface{}
	Meta sync.Map
}

type Message struct {
	*messagepb.Message
	parsed   *Parsed
	reserved sync.Map
}

func New(body []byte) (*Message, error) {
	rawmsg := messagepb.Message{}
	if err := proto.Unmarshal(body, &rawmsg); err != nil {
		return nil, err
	}
	m := &Message{Message: &rawmsg}
	if f, ok := parserRegistry[rawmsg.GetSub()]; ok {
		p, err := f(&rawmsg)
		if err != nil {
			return nil, err
		}
		m.parsed = p
	}
	return m, nil
}

func (m *Message) Store(key string, val interface{}) {
	m.reserved.Store(key, val)
}

func (m *Message) Delete(key string) {
	m.reserved.Delete(key)
}

func (m *Message) Load(key string) (interface{}, bool) {
	return m.reserved.Load(key)
}

func (m *Message) Parsed() *Parsed {
	return m.parsed
}
