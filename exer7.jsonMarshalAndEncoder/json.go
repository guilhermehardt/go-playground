package main

import (
	"encoding/json"
	"fmt"
	"os"
)

type backup struct {
	FileName    string
	Description string
	SizeInBytes int
	Owner       string
}

type unmarshalBackup struct {
	FileName    string
	Description string
	SizeInBytes int
	Owner       string
}

type decoderBackup struct {
	FileName string
}

func main() {
	backup1 := backup{"sites-enable.tar.gz", "Backup Server Nginx", 200, "ubuntu"}

	fmt.Println("### Marshal")
	backup1JSON := testingMarshal(backup1)
	fmt.Println("### Unmarshal")
	testingUnmarshal(backup1JSON)
	fmt.Println("### Encoder")
	testingEncoder(backup1)
	fmt.Println("### Decoder")
	testingDecoder()

}

func testingMarshal(b backup) []byte {

	dJSON, err := json.Marshal(b)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(string(dJSON))
	return dJSON
}

func testingUnmarshal(b []byte) {

	var dStruct unmarshalBackup
	err := json.Unmarshal(b, &dStruct)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(dStruct)
}

func testingEncoder(b backup) {

	encoder := json.NewEncoder(os.Stdout)
	encoder.Encode(b)
}

func testingDecoder() {
	var deco decoderBackup
	decoder := json.NewDecoder(os.Stdin)
	decoder.Decode(&deco)

	fmt.Println(deco.FileName)
}
