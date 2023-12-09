package modules

import (
	"log"
	"net/http"

	"mehrang.ir/school/pkg/database"
	"mehrang.ir/school/utils"
)

func List(w http.ResponseWriter, r *http.Request) {
	raw, err := database.DB.Query("SELECT * FROM users")
	if err != nil {
		utils.WriteJson(w, http.StatusOK, struct{ Message string }{Message: err.Error()})
		return
	} else {
		log.Println(raw)
	}
	defer raw.Close()
	payload := []struct {
		Title   string
		Teacher string
		Content string
	}{
		{
			Title:   "Go Crash Course",
			Teacher: "Mehran Ganji",
			Content: "The Content of Course",
		},
		{
			Title:   "Go Crash Course",
			Teacher: "Mehran Ganji",
			Content: "The Content of Course",
		},
		{
			Title:   "Go Crash Course",
			Teacher: "Mehran Ganji",
			Content: "The Content of Course",
		},
	}
	utils.WriteJson(w, http.StatusOK, payload)
}
