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
	w.Header().Add("Content-Type", "text/html")
	http.ServeFile(w, r, "./index.html")
}

func UploadHandler(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseMultipartForm(10 << 20); err != nil {
		http.Error(w, "Internal Server Error1", http.StatusInternalServerError)
		return
	}

	file, header, err := r.FormFile("myFile")
	if err != nil {
		http.Error(w, "Internal Server Error2", http.StatusInternalServerError)
		return
	}

	defer file.Close()

	data, err := io.ReadAll(file)
	if err != nil {
		http.Error(w, "Internal Server Error3", http.StatusInternalServerError)
		return
	}

	converted := service.Conversation(string(data))

	ext := filepath.Ext(header.Filename)
	filename := strings.ReplaceAll(time.Now().UTC().String(), " ", "_")
	filename = strings.ReplaceAll(filename, ":", "-") + ext
	localFile, err := os.OpenFile(filepath.Join("..", filename), os.O_CREATE|os.O_WRONLY, 0755)
	if err != nil {
		http.Error(w, "Internal Server Error4", http.StatusInternalServerError)
		return
	}

	defer localFile.Close()

	if _, err = localFile.Write([]byte(converted)); err != nil {
		http.Error(w, "Internal Server Error5", http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "text/plain; charset=utf-8")

	if _, err := w.Write([]byte(converted)); err != nil {
		log.Printf("Error sending response: %v", err)
	}
}
