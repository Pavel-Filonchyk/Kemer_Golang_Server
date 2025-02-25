package Controllers

import (
    "encoding/json"
    "net/http"
    "os"
)

type News struct {
    NewsRu string `json:"newsRu"`
    NewsTr string `json:"newsTr"`
    NewsEn string `json:"newsEn"`
}

func GetAllNews(w http.ResponseWriter, r *http.Request) {
    file, err := os.Open("news.json")
    if err != nil {
        http.Error(w, "Unable to open news file: "+err.Error(), http.StatusBadRequest)
        return
    }
    defer file.Close()

    var newsList []News
    err = json.NewDecoder(file).Decode(&newsList)
    if err != nil {
        http.Error(w, "Error decoding JSON: "+err.Error(), http.StatusInternalServerError)
        return
    }

    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(newsList)
}