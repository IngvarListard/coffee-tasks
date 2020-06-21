package prot1

import (
	"bytes"
	"encoding/binary"
	"io"
	"log"
	"math"
	"time"
)

type ByteSize int

const (
	NilCode uint = 1 << iota
	BoolCode
	Int32Code
	Int64Code
	Float64Code
	StringCode
	TimeCode

	OneByte    ByteSize = 1
	TwoBytes   ByteSize = 2
	FourBytes  ByteSize = 4
	EightBytes ByteSize = 8

	protocolVersion uint = 1
)

type Table struct {
	Name      string
	Headers   []string
	Types     []uint
	Data      []interface{}
	rowsCount uint
	reader    io.Reader
}

type Record struct {
	Type    uint
	Payload []byte
}

// MarshalBinary binary serialization
func (t *Table) MarshalBinary() ([]byte, error) {
	buf := new(bytes.Buffer)

	// запись версии протокола
	err := writeUint(buf, OneByte, protocolVersion)
	if err != nil {
		log.Fatal("ошибка записи версии", err)
	}

	// запись имени таблицы
	name := "test table name"
	err = writeUint(buf, OneByte, uint(len(name)))
	_, err = buf.Write([]byte(name))
	if err != nil {
		log.Fatal("ошибка записи имени таблицы")
	}

	// запись количества колонок в таблице
	err = writeUint(buf, OneByte, uint(len(t.Headers)))

	// запись имен столбцов
	for _, header := range t.Headers {
		err = writeUint(buf, OneByte, uint(len(header)))
		buf.Write([]byte(header))
	}

	// запись количества строк таблицы
	err = writeUint(buf, TwoBytes, uint(len(t.Data)))
	if err != nil {
		return nil, err
	}

	// запись таблицы
	for _, row := range t.Data {
		for i, col := range row.([]interface{}) {
			switch t.Types[i] {
			case Int32Code:
				b := make([]byte, FourBytes)
				binary.BigEndian.PutUint32(b, uint32(col.(int32)))
				err = writeUint(buf, OneByte, Int32Code)
				err = binary.Write(buf, binary.BigEndian, b)
				if err != nil {
					return nil, err
				}
			case Int64Code:
				b := make([]byte, EightBytes)
				binary.BigEndian.PutUint64(b, uint64(col.(int64)))
				err = writeUint(buf, OneByte, Int64Code)
				err = binary.Write(buf, binary.BigEndian, b)
				if err != nil {
					return nil, err
				}
			case Float64Code:
				b := make([]byte, EightBytes)
				binary.BigEndian.PutUint64(b, uint64(math.Float64bits(col.(float64))))

				err = writeUint(buf, OneByte, Float64Code)
				err = binary.Write(buf, binary.BigEndian, b)
				if err != nil {
					return nil, err
				}
			case StringCode:
				err = writeUint(buf, OneByte, StringCode)
				s := col.(string)
				err = writeUint(buf, FourBytes, uint(len(s)))
				err = binary.Write(buf, binary.BigEndian, []byte(col.(string)))
				if err != nil {
					return nil, err
				}
			case NilCode:
				err = writeUint(buf, OneByte, NilCode)
				if err != nil {
					return nil, err
				}
			case BoolCode:
				err = writeUint(buf, OneByte, BoolCode)
				err = binary.Write(buf, binary.BigEndian, col)
				if err != nil {
					return nil, err
				}
			case TimeCode:
				err = writeUint(buf, OneByte, TimeCode)
				val := col.(time.Time)
				timeBinary, err := val.MarshalBinary()
				err = binary.Write(buf, binary.BigEndian, timeBinary)
				if err != nil {
					return nil, err
				}
			}
		}
	}

	return buf.Bytes(), nil
}

func writeUint(w io.Writer, b ByteSize, i uint) error {
	var num interface{}
	switch b {
	case OneByte:
		num = uint8(i)
	case TwoBytes:
		num = uint16(i)
	case FourBytes:
		num = uint32(i)
	case EightBytes:
		num = uint64(i)
	}
	return binary.Write(w, binary.BigEndian, num)
}
