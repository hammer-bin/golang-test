package main

import (
	"bytes"
	"github.com/stretchr/testify/assert"
	"io"
	"mime/multipart"
	"net/http"
	"net/http/httptest"
	"os"
	"path/filepath"
	"testing"
)

func TestUploadTest(t *testing.T) {
	assert := assert.New(t)
	path := "D:/workspaceGolang/golang-test/turkerweb/WEB4-1/hammer.png"
	file, _ := os.Open(path)
	defer file.Close()

	err := os.RemoveAll("uploads/")
	/*if err != nil {
		return
	}*/
	assert.NoError(err)

	buf := &bytes.Buffer{}
	writer := multipart.NewWriter(buf)
	multi, err := writer.CreateFormFile("upload_file", filepath.Base(path))
	assert.NoError(err)
	io.Copy(multi, file)
	writer.Close()

	res := httptest.NewRecorder()
	req := httptest.NewRequest("POST", "/uploads", buf)
	req.Header.Set("Content-type", writer.FormDataContentType())

	uploadHandler(res, req)
	assert.Equal(http.StatusOK, res.Code)

	uploadFilePath := "uploads/" + filepath.Base(path)
	_, err = os.Stat(uploadFilePath)
	assert.NoError(err)

	uploadFile, _ := os.Open(uploadFilePath)
	originFile, _ := os.Open(path)
	defer uploadFile.Close()
	defer originFile.Close()

	uploadData := []byte{}
	originData := []byte{}
	uploadFile.Read(uploadData)
	originFile.Read(originData)

	assert.Equal(originData, uploadData)
}
