package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

// upload handler 만들기
func uploadHandler(w http.ResponseWriter, r *http.Request) {
	// request에 실려온 전송된 파일을 읽음
	uploadFile, header, err := r.FormFile("upload_file")
	// error 발생했을 때는 status 바꿔주고 return
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprint(w, err)
		return
	}
	defer uploadFile.Close()

	// 폴더 없으면 폴더 만들어주기
	// 파일모드, 읽기권한 : 0777(8진수형태:read, write, execute 모두 다됨)
	dirname := "./uploads"
	os.MkdirAll(dirname, 0777)
	filepath := fmt.Sprintf("%s/%s", dirname, header.Filename)
	// 비어있는 파일 만들어주기
	file, err := os.Create(filepath)
	// 파일은 항상 닫아줘야함 (handle은 OS 자원이어서 안닫아주면 문제가 생길 수 있음)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprint(w, err)
		return
	}
	defer file.Close()

	io.Copy(file, uploadFile)

	w.WriteHeader(http.StatusOK)
	fmt.Fprint(w, filepath)
}

func main() {

	http.HandleFunc("/uploads", uploadHandler)

	http.Handle("/", http.FileServer(http.Dir("public")))

	http.ListenAndServe(":3000", nil)
}
