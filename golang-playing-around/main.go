package main

import (
	"bytes"
	"encoding/binary"
	"io"
)

func main() {
	firstBuf := make([]byte, 16)
	binary.BigEndian.PutUint64(firstBuf, 1488)
	binary.BigEndian.PutUint64(firstBuf, 1696)
	var buf bytes.Buffer
	valReader := bytes.NewReader(firstBuf)

	bufValCopy := new(bytes.New)
	io.TeeReader(valReader, bufValCopy)

}
