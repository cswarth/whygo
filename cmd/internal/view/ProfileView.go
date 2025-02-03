package view

import (
	"bufio"
	"os"
)

type ProfileView struct {
}

func (view *ProfileView) ReadDelimited() {
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

}
