package tests

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"math/rand"
	"path/filepath"
	"testing"
	"time"

	"github.com/futanaichas/static-service/src/handle"
)

func TestSaveFile(t *testing.T) {
	root, _ := filepath.Abs("./")
	var data []byte
	rand.Seed(time.Now().Unix())
	for i := 0; i < rand.Intn(100); i++ {
		data = append(data, 'a')
	}
	hash := md5.Sum(data)
	fmt.Println(filepath.Join(root, hex.EncodeToString(hash[:])+".txt"))
	err := handle.SaveFile(hex.EncodeToString(hash[:])+".txt", root, handle.NowDate(), data)
	if err != nil {
		t.Fail()
	}
}
