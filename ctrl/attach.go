package ctrl

import (
	"ImApplication.go/util"
	"fmt"
	"io"
	"math/rand"
	"net/http"
	"os"
	"strings"
	"time"
)

func init() {
	os.MkdirAll("./mnt", os.ModePerm)
}
func Upload(w http.ResponseWriter, r *http.Request) {
	//上传到本地
	UploadLocal(w, r)
}

// 1.存储位置 ./mnt 自动船舰
// 2.url格式 /mnt/xxxx.png  需要确保网络能访问/mnt/
func UploadLocal(writer http.ResponseWriter,
	request *http.Request) {
	//todo 获得上传的源文件s
	srcfile, head, err := request.FormFile("file")
	if err != nil {
		util.ResponseFail(writer, err.Error())
	}

	//todo 创建一个新文件d
	suffix := ".png"
	//如果前端文件名称包含后缀 xx.xx.png
	ofilename := head.Filename
	tmp := strings.Split(ofilename, ".")
	if len(tmp) > 1 {
		suffix = "." + tmp[len(tmp)-1]
	}
	//如果前端指定filetype
	//formdata.append("filetype",".png")
	filetype := request.FormValue("filetype")
	if len(filetype) > 0 {
		suffix = filetype
	}
	//time.Now().Unix()
	filename := fmt.Sprintf("%d%04d%s",
		time.Now().Unix(), rand.Int31(),
		suffix)
	dstfile, err := os.Create("./mnt/" + filename)
	if err != nil {
		util.ResponseFail(writer, err.Error())
		return
	}

	//todo 将源文件内容copy到新文件
	_, err = io.Copy(dstfile, srcfile)
	if err != nil {
		util.ResponseFail(writer, err.Error())
		return
	}
	//todo 将新文件路径转换成url地址

	url := "/mnt/" + filename
	//todo 响应到前端
	util.ResponseOk(writer, "", url)
}
