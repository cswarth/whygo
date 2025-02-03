//go:generate pwd
//go:generate protoc --plugin=/usr/local/go/bin/protoc-gen-go --go_out=.  --proto_path=../internal/proto/v0/  v0_messages.proto
//go:generate protoc --plugin=/usr/local/go/bin/protoc-gen-go --go_out=.  --proto_path=../internal/proto/v0/  v0_constraints.proto
//go:generate protoc --plugin=/usr/local/go/bin/protoc-gen-go --go_out=.  --proto_path=../internal/proto/v0/  v0_summaries.proto
//go:generate ls cmd/generated/whylogs

package main

import (
	view "awesomeProject/cmd/internal/view"
	"fmt"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {
	fmt.Println("hello world")

	view := view.ProfileView{}
	view.ReadDelimited()

}
