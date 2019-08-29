package handle

import (
	"crypto/md5"
	"encoding/hex"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"path/filepath"
	"time"

	"github.com/futanaichas/static-service/src/module"
)

// NowDate ...
func NowDate() string {
	return time.Now().Format("2006-01-02")
}

func getFormValue(r *http.Request) (string, string) {
	return r.FormValue("type"), r.FormValue("info")
}

func hanleFileName(data []byte, fileType string, fileInfo string) (hashstr string, fileDate string) {
	hash := md5.Sum(data)
	hashstr = hex.EncodeToString(hash[:])
	fileDate = NowDate()
	return
}

// SaveFile ...
func SaveFile(fileName string, rootPath string, time string, data []byte) error {
	os.Mkdir(filepath.Join(rootPath, time), os.ModePerm)

	file, err := os.Create(filepath.Join(rootPath, time, fileName))
	if err != nil {
		fmt.Println(filepath.Join(rootPath, time, fileName))
		return errors.New("create file fail")
	}
	defer file.Close()

	_, err = file.Write(data)
	if err != nil {
		fmt.Println("write file fail")
		return errors.New("write file fail")
	}

	return nil
}

// Upload 上传处理
func Upload() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// 如果不是post请求直接返回错误
		if r.Method != "POST" {
			w.Write(module.FailRes())
			return
		}
		// 获取上传的文件类型
		fileType, fileInfo := getFormValue(r)
		if fileType == "" || fileInfo == "" {
			w.Write(module.FailRes())
			return
		}

		// 读取上传的文件
		file, _, err := r.FormFile("file")
		if err != nil {
			w.Write(module.FailRes())
			return
		}
		defer file.Close()

		// 将文件读取到应用层
		fileByte, err := ioutil.ReadAll(file)
		if err != nil {
			w.Write(module.FailRes())
			return
		}

		//重命名文件名称
		hash, fileDate := hanleFileName(fileByte, fileType, fileInfo)

		// 储存文件
		rootPath, _ := filepath.Abs("./public/static")
		err = SaveFile(fmt.Sprintf("%s%s.%s", hash, fileInfo, fileType), rootPath, fileDate, fileByte)
		if err != nil {
			w.Write(module.FailRes())
			return
		}

		w.Write(module.OkRes(hash, fmt.Sprintf("/static/%s/%s%s.%s", fileDate, hash, fileInfo, fileType)))
	}
}
