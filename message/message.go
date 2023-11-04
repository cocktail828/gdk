package message

import (
	"github.com/cocktail828/gdk/v1/message/messagepb"
	"github.com/pkg/errors"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/types/known/anypb"
)

type Message struct {
	*messagepb.Message
	Parsed interface{}
}

func New(body []byte) (*Message, error) {
	rawmsg := messagepb.Message{}
	if err := proto.Unmarshal(body, &rawmsg); err != nil {
		return nil, err
	}
	return &Message{Message: &rawmsg}, nil
}

func (m *Message) Unmarshal(f func(*anypb.Any) (interface{}, error)) error {
	if f == nil {
		return errors.Errorf("invalid unmarshaller for sub:%v", m.Sub)
	}

	v, err := f(m.GetReserved())
	if err == nil {
		m.Parsed = v
	}
	return err
}
