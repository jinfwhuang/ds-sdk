package main

import (
	"github.com/jinfwhuang/ds-toolkit/experimental/dsn/arweave"
	"github.com/jinfwhuang/ds-toolkit/experimental/dsn/ceramic"
	estuary "github.com/jinfwhuang/ds-toolkit/experimental/dsn/filecoin-estuary"
	"github.com/jinfwhuang/ds-toolkit/experimental/dsn/storj"
)

var (
	data = []byte("123")
)

func main() {

}

func arweaveExample() {
	id, err := arweave.Write(data)
	if err != nil {
		println(err.Error())
	}
	println(id)

	// This is an example id that should be retrieved from Write function.
	// Arweave takes couple of minutes until the id is accepted from the chain,
	// that is why we use another random id that is already in the chain here.
	ret, err := arweave.Read("IUYBL-mW7OpG7Em_kwIpucrg43Br64nGbeMM01yja4w")
	if err != nil {
		println(err.Error())
	}
	println(string(ret))
}

func estuaryExample() {
	resp, err := estuary.Write(data)
	if err != nil {
		println(err.Error())
	}
	println(resp)
}

func storjExample() {
	key := "key_1"

	err := storj.Write(key, data)
	if err != nil {
		println(err.Error())
	}

	retr, err := storj.Read(key)
	if err != nil {
		println(err.Error())
	}
	println(string(retr))
}

func ceramicExample() {
	streamId, err := ceramic.Write(string(data))
	if err != nil {
		println(err.Error())
	}
	println(streamId)

	// This is an example id of a Stream that should be obtained from the above call.
	// However, the writing via HTTP does not seem to be maintained and currently does not work.
	// The example ID is an ID of a stream generated via the CLI.
	retr, err := ceramic.Read("kjzl6cwe1jw147z6im9aagzi92icf72n7evjt5kdhevd97ygkaxjqbicyxz618i")
	println(retr)
	if err != nil {
		println(err.Error())
	}
}
