package handlers

import (
	"io"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"time"

	"github.com/Yandex-Practicum/go1fl-sprint6-final/internal/service"
)

func RootHandler(res http.ResponseWriter, req *http.Request) {
	file, err := os.Open("../index.html")
	if err != nil {
		log.Println("ошибка при получении файла", err)
		http.Error(res, "ошибка при чтении файла", http.StatusInternalServerError)
		return
	}

	defer file.Close()

	res.Header().Set("Content-Type", "text/html")

	_, err = io.Copy(res, file)

	if err != nil {
		log.Println("ошибка при отправке файла", err)
		http.Error(res, "ошибка при отправке файла", http.StatusInternalServerError)
		return
	}
}

func UploadHandler(res http.ResponseWriter, req *http.Request) {
	if req.Method != http.MethodPost {
		http.Error(res, "разрешен только метод Post", http.StatusMethodNotAllowed)
		return
	}

	if err := req.ParseMultipartForm(10 << 20); err != nil {
		log.Println("ошибка парсинга", err)
		http.Error(res, "ошибка парсинга", http.StatusInternalServerError)
		return
	}

	file, handler, err := req.FormFile("myFile")
	if err != nil {
		log.Println("ошибка при получении файла", err)
		http.Error(res, "ошибка при получении файла", http.StatusInternalServerError)
		return
	}

	defer file.Close()

	text, err := io.ReadAll(file)

	if err != nil {
		log.Println("ошибка чтения файла", err)
		http.Error(res, "ошибка чтения файла", http.StatusInternalServerError)
		return
	}

	convertString, err := service.ConvertText(string(text))
	if err != nil {
		log.Println("ошибка конвертации файла", err)
		http.Error(res, "ошибка конвертации файла", http.StatusInternalServerError)
		return
	}

	ext := filepath.Ext(handler.Filename)
	newFileName := time.Now().UTC().Format("2006-01-02_15-04-05") + ext

	dir := "results"
	err = os.MkdirAll(dir, 0775)
	if err != nil {
		log.Println("ошибка при создании директории", err)
		http.Error(res, "ошибка при создании директории", http.StatusInternalServerError)
		return
	}

	newFilePath := filepath.Join(dir, newFileName)

	newFile, err := os.Create(newFilePath)
	if err != nil {
		log.Println("ошибка при создании файла", err)
		http.Error(res, "ошибка при создании файла", http.StatusInternalServerError)
		return
	}

	defer newFile.Close()

	_, err = newFile.WriteString(convertString)
	if err != nil {
		log.Println("ошибка записи в файл", err)
		http.Error(res, "ошибка записи в файл", http.StatusInternalServerError)
		return
	}

	res.Header().Set("Content-Type", "text/plain; charset=utf-8")
	res.Write([]byte(convertString))
}
