package handlers

import (
	"io"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"strings"
	"time"

	"github.com/Yandex-Practicum/go1fl-sprint6-final/internal/service"
)

func HtmlHandler(w http.ResponseWriter, r *http.Request) {
	filePath, err := filepath.Abs(filepath.Join("..", "index.html"))
	if err != nil {
		http.Error(w, "error in file location path", http.StatusInternalServerError)
		return
	}
	data, err := os.ReadFile(filePath)
	if err != nil {
		http.Error(w, "I can't read Html-file", http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	w.WriteHeader(http.StatusOK)
	w.Write(data)
}

func UploadHandler(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseMultipartForm(10 << 20); err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	file, header, err := r.FormFile("myFile")
	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	defer file.Close()

	data, err := io.ReadAll(file)
	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	converted := service.Conversation(string(data))

	ext := filepath.Ext(header.Filename)
	filename := strings.ReplaceAll(time.Now().UTC().String(), " ", "_")
	filename = strings.ReplaceAll(filename, ":", "-") + ext
	localFile, err := os.OpenFile(filepath.Join("..", filename), os.O_CREATE|os.O_WRONLY, 0755)
	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	defer localFile.Close()

	if _, err = localFile.Write([]byte(converted)); err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "text/html; charset=utf-8")

	if _, err := w.Write([]byte(converted)); err != nil {
		log.Printf("Error sending response: %v", err)
	}
}
