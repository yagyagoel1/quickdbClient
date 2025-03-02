package utils

import (
	"bufio"
	"fmt"
	"io"
	"strconv"
)

const (
	STRING  = "+"
	ERROR   = "-"
	INTEGER = ":"
	BULK    = "$"
	ARRAY   = "*"
)

type Value struct {
	Typ   string
	Str   string
	Num   int
	Bulk  string
	Array []Value
}
type Resp struct {
	reader *bufio.Reader
}

type Writer struct {
	writer io.Writer
}

func NewWriter(w io.Writer) *Writer {
	return &Writer{writer: w}
}

func NewResp(rd io.Reader) *Resp {
	return &Resp{reader: bufio.NewReader(rd)}
}

func (r *Resp) readLine() (line []byte, n int, err error) {
	for {
		b, err := r.reader.ReadByte()
		if err != nil {
			return nil, 0, err
		}

		n += 1
		line = append(line, b)
		f, et := strconv.ParseInt(string(b), 10, 64)
		fmt.Println("l", string(line), f, et)
		if len(line) >= 2 && line[len(line)-2] == '\r' {
			break
		}

	}
	return line[:len(line)-2], n, nil
}

func (r *Resp) readInteger() (x int, n int, err error) {
	line, n, err := r.readLine()
	if err != nil {
		return 0, 0, err
	}
	fmt.Println("line", string(line), n)
	i64, err := strconv.ParseInt(string(line), 10, 64)
	if err != nil {
		return 0, n, err
	}
	return int(i64), n, nil
}

func (r *Resp) Read() (Value, error) {
	_type, err := r.reader.ReadByte()
	if err != nil {
		return Value{}, err
	}
	fmt.Println("type", string(_type))
	switch string(_type) {
	case ARRAY:
		return r.readArray()
	case BULK:
		return r.readBulk()
	case STRING:
		return r.readString()
	default:
		fmt.Printf("Unknown type %v", string(_type))
		return Value{}, nil
	}

}
func (r *Resp) readArray() (Value, error) {
	v := Value{}
	v.Typ = "array"
	length, _, err := r.readInteger()
	if err != nil {
		return v, err
	}

	v.Array = make([]Value, length)
	for i := 0; i < length; i++ {
		val, err := r.Read()
		if err != nil {
			return v, err
		}
		v.Array[i] = val
	}
	return v, nil

}

func (r *Resp) readBulk() (Value, error) {
	v := Value{}

	v.Typ = "bulk"

	len, _, err := r.readInteger()
	if err != nil {
		return v, err
	}
	if len == -1 {
		v.Bulk = "null"
		return v, nil
	}
	bulk := make([]byte, len)

	r.reader.Read(bulk)

	v.Bulk = string(bulk)
	// Read the trailing CRLF
	r.readLine()

	return v, nil
}

func (r *Resp) readString() (Value, error) {
	v := Value{}

	v.Typ = "string"
	line, _, err := r.readLine()
	if err != nil {
		return v, err
	}
	v.Str = string(line)

	return v, nil
}

func (v Value) Marshal() []byte {
	switch v.Typ {
	case "array":
		return v.marshalArray()
	case "bulk":
		return v.marshalBulk()
	case "string":
		return v.marshalString()
	case "null":
		return v.marshalNull()
	case "error":
		return v.marshalError()
	default:
		return []byte{}
	}

}

func (v Value) marshalString() []byte {
	var bytes []byte
	bytes = append(bytes, STRING...)
	bytes = append(bytes, v.Str...)
	bytes = append(bytes, '\r', '\n')
	return bytes

}

func (v Value) marshalBulk() []byte {
	var bytes []byte
	bytes = append(bytes, BULK...)
	bytes = append(bytes, strconv.Itoa(len(v.Bulk))...)
	bytes = append(bytes, '\r', '\n')
	bytes = append(bytes, v.Bulk...)
	bytes = append(bytes, '\r', '\n')
	return bytes
}
func (v Value) marshalArray() []byte {
	len := len(v.Array)
	var bytes []byte
	bytes = append(bytes, ARRAY...)
	bytes = append(bytes, strconv.Itoa(len)...)
	bytes = append(bytes, '\r', '\n')
	for i := 0; i < len; i++ {
		bytes = append(bytes, v.Array[i].Marshal()...)
	}
	return bytes
}

func (v Value) marshalError() []byte {
	var bytes []byte
	bytes = append(bytes, ERROR...)
	bytes = append(bytes, v.Str...)
	bytes = append(bytes, '\r', '\n')
	return bytes
}

func (v Value) marshalNull() []byte {
	return []byte("$-1\r\n")
}

func (w *Writer) Write(v Value) error {
	var bytes = v.Marshal()
	_, err := w.writer.Write(bytes)
	if err != nil {
		return err
	}
	return nil
}
