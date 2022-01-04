package levin

import (
	"encoding/binary"
	"fmt"
)

const (
	TypeInt64 byte = 0x1
	TypeInt32 byte = 0x2
	TypeInt16 byte = 0x3
	TypeInt8  byte = 0x4

	TypeUint64 byte = 0x5
	TypeUint32 byte = 0x6
	TypeUint16 byte = 0x7
	TypeUint8  byte = 0x8

	TypeDouble byte = 0x9

	TypeString byte = 0xa
	TypeBool   byte = 0xb
	TypeObject byte = 0xc
	TypeArray  byte = 0xd

	FlagArray byte = 0x80
)

type Byte byte

func (v Byte) Bytes() []byte {
	return []byte{
		TypeUint8,
		byte(v),
	}
}

type Uint32 uint32

func (v Uint32) Bytes() []byte {
	b := []byte{
		TypeUint32,
		0x00, 0x00, 0x00, 0x00,
	}
	binary.LittleEndian.PutUint32(b[1:], uint32(v))
	return b
}

type Uint64 uint64

func (v Uint64) Bytes() []byte {
	b := []byte{
		TypeUint64,
		0x00, 0x00, 0x00, 0x00,
		0x00, 0x00, 0x00, 0x00,
	}

	binary.LittleEndian.PutUint64(b[1:], uint64(v))

	return b
}

type String string

func (v String) Bytes() []byte {
	b := []byte{TypeString}

	varInB, err := VarIn(len(v))
	if err != nil {
		panic(fmt.Errorf("varin '%d': %w", len(v), err))
	}

	return append(b, append(varInB, []byte(v)...)...)
}
