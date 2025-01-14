package core

import (
	"encoding/binary"
	"errors"
	"io"
)

func readBytes(buf []byte, cursor *int, size int) ([]byte, error) {
	if cursor == nil {
		return nil, errors.New("cursor is nil")
	}
	if *cursor+size > len(buf) {
		return nil, errors.New("cursor out of range")
	}
	end := *cursor + size
	result := buf[*cursor:end]
	*cursor += int(size)
	return result, nil
}

func readByte(buf []byte, cursor *int) (byte, error) {
	if cursor == nil {
		return 0, errors.New("cursor is nil")
	}
	if *cursor >= len(buf) {
		return 0, errors.New("cursor out of range")
	}
	b := buf[*cursor]
	*cursor++
	return b, nil
}

func readZipListLength(buf []byte, cursor *int) int {
	start := *cursor + 8
	end := start + 2
	// zip list buf: [0, 4] -> zlbytes, [4:8] -> zltail, [8:10] -> zllen
	size := int(binary.LittleEndian.Uint16(buf[start:end]))
	*cursor += 10
	return size
}

func (dec *Decoder) readByte() (byte, error) {
	b, err := dec.input.ReadByte()
	if err != nil {
		return 0, err
	}
	dec.readCount++
	return b, nil
}

func (dec *Decoder) readFull(buf []byte) error {
	n, err := io.ReadFull(dec.input, buf)
	if err != nil {
		return err
	}
	dec.readCount += n
	return nil
}
