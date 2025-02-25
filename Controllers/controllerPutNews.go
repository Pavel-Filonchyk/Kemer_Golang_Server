package Controllers

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"os"
    "fmt"
)

type NewsUpdate struct {
	NewsRu string `json:"newsRu"`
	NewsTr string `json:"newsTr"`
	NewsEn string `json:"newsEn"`
}

func UpdateNews(w http.ResponseWriter, r *http.Request) {

	if r.Method != http.MethodPut {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	body, err := ioutil.ReadAll(r.Body)
    fmt.Println(r)
	if err != nil {
		http.Error(w, "Error reading request body", http.StatusBadRequest)
		return
	}
	defer r.Body.Close()


	var news []NewsUpdate
	err = json.Unmarshal(body, &news)
	if err != nil {
		http.Error(w, "Error decoding JSON", http.StatusBadRequest)
		return
	}

	file, err := os.OpenFile("news.json", os.O_RDWR|os.O_TRUNC|os.O_CREATE, 0644)
	if err != nil {
		http.Error(w, "Error opening file", http.StatusInternalServerError)
		return
	}
	defer file.Close()

	updatedNewsJSON, err := json.Marshal(news)
	if err != nil {
		http.Error(w, "Error encoding JSON", http.StatusInternalServerError)
		return
	}

	_, err = file.Write(updatedNewsJSON)
	if err != nil {
		http.Error(w, "Error writing to file", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte("News updated successfully"))
}