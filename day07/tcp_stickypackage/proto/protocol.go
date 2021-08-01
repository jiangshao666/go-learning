package proto

import (
	"bufio"
	"bytes"
	"encoding/binary"
)

// 按照自定的格式 [len]+[body] 序列化
func Encode(msg string) ([]byte, error) {
	var length = int32(len(msg))
	var pkg = new(bytes.Buffer)
	err := binary.Write(pkg, binary.LittleEndian, length)
	if err != nil {
		return nil, err
	}
	err = binary.Write(pkg, binary.LittleEndian, []byte(msg))
	if err != nil {
		return nil, err
	}
	return pkg.Bytes(), nil
}

// 按照自定的格式 [len]+[body] 反序列化
func Decode(reader *bufio.Reader) (string, error) {
	lengthBytes, _ := reader.Peek(4)
	lenBuffer := bytes.NewBuffer(lengthBytes)
	var length int32
	err := binary.Read(lenBuffer, binary.LittleEndian, &length)
	if err != nil {
		return "", err
	}
	if int32(reader.Buffered()) < length+4 {
		return "", err
	}

	pack := make([]byte, length+4)
	_, err = reader.Read(pack)
	if err != nil {
		return "", err
	}
	return string(pack[4:]), nil
}
