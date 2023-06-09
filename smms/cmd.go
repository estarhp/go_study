package main

import (
	"bytes"
	"context"
	"fmt"
	"github.com/carlmjohnson/requests"
	"io"
	"log"
	"mime/multipart"
	"net/http"
	"net/url"
	"os"
)

func main() {
	image, err := os.Open("v2-1c15aad66d7889fe522661a90f623dc4_1440w.jpg")
	if err != nil {
		println(err)
		return
	}
	defer func(image *os.File) {
		err := image.Close()
		if err != nil {
			println(err)
		}
	}(image)
	//设置参数

	var body = new(bytes.Buffer)
	writer := multipart.NewWriter(body)
	part, err := writer.CreateFormFile("smfile", "v2-1c15aad66d7889fe522661a90f623dc4_1440w.jpg")

	if err != nil {
		fmt.Println("创建表单数据失败：", err)
		return
	}

	_, err = io.Copy(part, image)

	if err != nil {
		fmt.Println("添加文件内容失败：", err)
		return
	}

	err = writer.Close()

	if err != nil {
		fmt.Println("关闭表单数据失败：", err)
		return
	}

	transport := http.Transport{
		Proxy: func(req *http.Request) (*url.URL, error) {
			// 设置代理服务器地址和端口号
			proxyUrl, err := url.Parse("http://127.0.0.1:1080")
			if err != nil {
				return nil, err
			}
			return proxyUrl, nil
		},
	}
	var s string
	err = requests.
		URL("https://sm.ms/api/v2/upload").
		BodyReader(body).
		Header("Authorization", "YUXTOHcLUFW7QSyzVRqpYYuXAw8iShY1").
		Transport(&transport).
		ContentType(writer.FormDataContentType()).
		ToString(&s).
		Fetch(context.Background())
	log.Println(s)

	if err != nil {
		log.Println(err)
	}

}
