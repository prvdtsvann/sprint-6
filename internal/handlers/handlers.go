package handlers

import (
	"io"
	"net/http"
	"os"
	"path/filepath"
	"time"

	"github.com/Yandex-Practicum/go1fl-sprint6-final/internal/service"
)

func RootHandler(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "./index.html")
}

func UploadHandler(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseMultipartForm(10 << 20); err != nil {
		http.Error(w, "parse form error: "+err.Error(), http.StatusInternalServerError)
		return
	}

	file, header, err := r.FormFile("myFile")
	if err != nil {
		http.Error(w, "failed receiving error: "+err.Error(), http.StatusInternalServerError)
		return
	}
	defer file.Close()

	fileData, err := io.ReadAll(file)
	if err != nil {
		http.Error(w, "failed reading file: "+err.Error(), http.StatusInternalServerError)
		return
	}

	convertedData := service.Convert(string(fileData))

	localFile := time.Now().UTC().Format("20060102_150405") + filepath.Ext(header.Filename)

	err = os.WriteFile(localFile, []byte(convertedData), 0755)
	if err != nil {
		http.Error(w, "failed writing file: "+err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "text/html")
	w.WriteHeader(http.StatusOK)
	_, err = w.Write([]byte("Complited. Result: \n" + convertedData))
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

}
