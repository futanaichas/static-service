package handle

import (
	"io/ioutil"
	"net/http"

	"github.com/futanaichas/static-service/src/module"
)

func getFormValue(r *http.Request) (string, string) {
	return r.FormValue("type"), r.FormValue("info")
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
		uploadFile := &File{
			Type: fileType,
			Info: fileInfo,
			Data: fileByte,
		}
		uploadFile.SetRoot("./public/static")

		err = uploadFile.SaveFile()
		if err != nil {
			w.Write(module.FailRes())
			return
		}

		w.Write(module.OkRes(uploadFile.Hash, uploadFile.URL()))
	}
}
