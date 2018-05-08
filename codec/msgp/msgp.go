// Package msgp implements encoding and decoding of MessagePack, relying on
// serialization code generated by msgp, the code generation library for
// MessagePack at github.com/tinylib/msgp.
package msgp

import (
	"encoding/base64"
	"fmt"

	"github.com/tinylib/msgp/msgp"
	"github.com/zippoxer/bow/codec"
)

//go:generate msgp -unexported
//msgp:ignore Codec

// Id is the msgp equivalent of bow.Id. It must be used instead.
type Id []byte

func (id Id) Marshal(in []byte) ([]byte, error) {
	return id, nil
}

func (id *Id) Unmarshal(b []byte) error {
	*id = b
	return nil
}

func (id Id) String() string {
	return base64.RawURLEncoding.EncodeToString(id)
}

type Codec struct{}

func (c Codec) Marshal(v interface{}, in []byte) (out []byte, err error) {
	m, ok := v.(msgp.Marshaler)
	if !ok {
		return nil, fmt.Errorf("type %T doesn't implement msgp.Marshaler. "+
			"Did you forget to 'go generate'?", v)
	}
	return m.MarshalMsg(in[0:])
}

func (c Codec) Unmarshal(data []byte, v interface{}) error {
	m, ok := v.(msgp.Unmarshaler)
	if !ok {
		return fmt.Errorf("type %T doesn't implement msgp.Marshaler. "+
			"Did you forget to 'go generate'?", v)
	}
	_, err := m.UnmarshalMsg(data)
	return err
}

func (c Codec) Format() codec.Format {
	return codec.MessagePack
}
