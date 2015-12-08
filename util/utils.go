package util

import (
	"bytes"
	"encoding/binary"
	"errors"
)

func DecodeValue(page []byte, offset int, a interface{}, order binary.ByteOrder) error {

	size := binary.Size(a)
	if size == -1 {
		return errors.New("Invalid provided type")
	}

	buff := page[offset : offset+size]

	err := binary.Read(bytes.NewBuffer(buff), order, a)

	return err
}

func DecodeBytes(src []byte, dst *[]byte, offset, length int64) {
	*dst = src[offset : offset+length]
}
