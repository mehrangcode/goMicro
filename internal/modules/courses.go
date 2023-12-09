package modules

import (
	"net/http"

	"mehrang.ir/school/utils"
)

func List(w http.ResponseWriter, r *http.Request) {
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
