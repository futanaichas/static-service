package tests

import (
	"math/rand"
	"testing"
	"time"

	"github.com/futanaichas/static-service/src/handle"
)

func TestSaveFile(t *testing.T) {
	var data []byte
	rand.Seed(time.Now().Unix())
	for i := 0; i < rand.Intn(100); i++ {
		data = append(data, 'a')
	}

	uploadFile := &handle.File{
		Type: "txt",
		Info: "@test",
		Data: data,
	}
	uploadFile.SetRoot("./")

	err := uploadFile.SaveFile()
	if err != nil {
		t.Fail()
	}
}
