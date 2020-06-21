package prot1

import (
	"bytes"
	"encoding/binary"
	"errors"
	"fmt"
	"time"
)

func (t *Table) UnmarshalBinary(buf []byte) error {
	t.reader = bytes.NewReader(buf)

	err := t.readHeader()
	if err != nil {
		return fmt.Errorf("UnmarhalBinary: error reading header: %v", err)
	}
	e := t.readTable()
	return e
}

func (t *Table) readHeader() error {
	// чтение версии протокола
	versionBuf := make([]byte, OneByte)
	_, err := t.reader.Read(versionBuf[:])
	version, err := readUint(versionBuf, OneByte)
	if err != nil {
		return err
	}

	if version != protocolVersion {
		return errors.New("wrong protocol version")
	}

	//чтение имени таблицы
	nameLenBuf := make([]byte, OneByte)
	_, err = t.reader.Read(nameLenBuf)
	nameLen, err := readUint(nameLenBuf, OneByte)
	nameBuf := make([]byte, nameLen)
	_, err = t.reader.Read(nameBuf)
	if err != nil {
		return err
	}
	t.Name = string(nameBuf)

	// чтение количества столбцов в таблице
	colNumberBuf := make([]byte, OneByte)
	_, err = t.reader.Read(colNumberBuf)
	colNumber, err := readUint(colNumberBuf, OneByte)
	if err != nil {
		return err
	}

	// чтение имен столбцов
	var colNames []string
	for i := 0; i < int(colNumber); i++ {
		colNameLenBuf := make([]byte, OneByte)
		_, err = t.reader.Read(colNameLenBuf)
		nameLen, err := readUint(colNameLenBuf, OneByte)
		nameBuf := make([]byte, nameLen)
		_, err = t.reader.Read(nameBuf)
		if err != nil {
			return err
		}
		colNames = append(colNames, string(nameBuf))
	}
	t.Headers = colNames

	// чтение количества строк таблицы
	rowsCountBuf := make([]byte, TwoBytes)
	_, err = t.reader.Read(rowsCountBuf)
	rowsCount, err := readUint(rowsCountBuf, TwoBytes)
	t.rowsCount = rowsCount
	return err
}

func (t *Table) readTable() error {
	var rows []interface{}
	for i := 0; i < int(t.rowsCount); i++ {
		var col []interface{}
		for i := 0; i < len(t.Headers); i++ {
			var val interface{}
			typeBuf := make([]byte, OneByte)
			_, err := t.reader.Read(typeBuf)
			typ, err := readUint(typeBuf, OneByte)
			if err != nil {
				return err
			}
			switch typ {
			case NilCode:
				val = nil
				fmt.Println("nil")
			case BoolCode:
				buf := make([]byte, OneByte)
				_, err := t.reader.Read(buf)
				v, err := readUint(buf, OneByte)
				if err != nil {
					return err
				}
				if v == 1 {
					val = true
				} else {
					val = false
				}
				fmt.Println("bool")
			case Int32Code:
				var v int32
				bufInt32 := make([]byte, FourBytes)
				_, err := t.reader.Read(bufInt32)
				err = binary.Read(bytes.NewBuffer(bufInt32), binary.BigEndian, &v)
				if err != nil {
					return err
				}
				val = v

				fmt.Println("int32", val)
			case Int64Code:
				var v int64
				bufInt64 := make([]byte, EightBytes)
				_, err := t.reader.Read(bufInt64)
				err = binary.Read(bytes.NewBuffer(bufInt64), binary.BigEndian, &v)
				if err != nil {
					return err
				}
				val = v
				fmt.Println("int64", val)
			case Float64Code:
				var v float64
				bufFloat64 := make([]byte, EightBytes)
				_, err := t.reader.Read(bufFloat64)
				err = binary.Read(bytes.NewBuffer(bufFloat64), binary.BigEndian, &v)
				if err != nil {
					return err
				}
				val = v
				fmt.Println("float64", val)
			case StringCode:
				stringLenBuf := make([]byte, FourBytes)
				_, err := t.reader.Read(stringLenBuf)
				stringLen, err := readUint(stringLenBuf, FourBytes)
				bufString := make([]byte, stringLen)
				_, err = t.reader.Read(bufString)
				if err != nil {
					return err
				}
				val = string(bufString)
			case TimeCode:
				timeBuf := make([]byte, 15)
				_, err = t.reader.Read(timeBuf)
				t := new(time.Time)
				err = t.UnmarshalBinary(timeBuf)
				if err != nil {
					return err
				}
				val = t
			}
			col = append(col, val)
		}
		rows = append(rows, col)
	}
	t.Data = rows
	return nil
}

func readUint(buf []byte, b ByteSize) (uint, error) {

	reader := bytes.NewReader(buf)

	switch b {
	case OneByte:
		var i uint8
		return uint(i), binary.Read(reader, binary.BigEndian, &i)
	case TwoBytes:
		var i uint16
		return uint(i), binary.Read(reader, binary.BigEndian, &i)
	case FourBytes:
		var i uint32
		return uint(i), binary.Read(reader, binary.BigEndian, &i)
	case EightBytes:
		var i uint64
		return uint(i), binary.Read(reader, binary.BigEndian, &i)
	}

	return 0, errors.New("unexpected byte size")
}
