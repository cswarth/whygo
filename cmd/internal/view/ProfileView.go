package view

import (
	"bufio"
	"errors"
	"google.golang.org/protobuf/encoding/protowire"
	"io"
	"os"
)

type ProfileView struct {
}

func (view *ProfileView) ReadDelimited() {
	//     size = _read_varint(stream, offset=offset)
	//    if size == 0:
	//        return proto_class_name()
	//    buf = stream.read(size)
	//    msg = proto_class_name()
	//    delimited_read_failure = False
	file, err := os.Open("cmd/internal/view/testdata/TEST_PROFILE.bin")
	if err != nil {
		panic(err)
	}
	defer file.Close()
	bytes := make([]byte, 4)
	bufr := bufio.NewReader(file)
	_, err = bufr.Read(bytes)
	if err != nil {
		panic(err)
	}
	size, err := readVarint(bufr)
	if err != nil {
		panic(err)
	}
	buf := make([]byte, size)
	_, err = bufr.Read(buf)
	if err != nil {
		panic(err)
	}
	//msg = proto_class_name()
	//delimited_read_failure = False

}

// readVarint reads a varint from the given stream at the specified offset.
func readVarint(stream io.Reader) (int, error) {

	var buf []byte
	temp := make([]byte, 1)

	// Read the first byte
	_, err := stream.Read(temp)
	if err != nil {
		return 0, err
	}
	buf = append(buf, temp[0])

	// Continue reading while the MSB is 1
	for temp[0]&0x80 != 0 {
		_, err := stream.Read(temp)
		if err == io.EOF {
			return 0, errors.New("unexpected EOF")
		}
		if err != nil {
			return 0, err
		}
		buf = append(buf, temp[0])
	}

	// Decode varint
	varint, _ := protowire.ConsumeVarint(buf)
	return int(varint), nil
}
