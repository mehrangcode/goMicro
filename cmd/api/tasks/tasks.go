package tasks

import (
	"log"
	"net/http"

	"mehrangcode.ir/todoapp/cmd/initializers"
	"mehrangcode.ir/todoapp/cmd/models"
	"mehrangcode.ir/todoapp/cmd/utils"
)

func Create(w http.ResponseWriter, r *http.Request) {
	var body models.Task = models.Task{}
	err := utils.ReadJson(w, r, &body)
	if err != nil {
		utils.WriteJson(w, http.StatusOK, err)
		return
	}
	task := models.Task{
		Title:  body.Title,
		Status: body.Status,
		Owner:  body.Owner,
	}
	result := initializers.DB.Create(&task)
	if result.Error != nil {
		_ = utils.WriteJson(w, http.StatusBadRequest,
			struct {
				Error   bool
				Message string
			}{Error: true, Message: "somthing goeas wrong" + result.Error.Error()})
	}
	payload := struct{ ID uint }{ID: task.ID}
	err = utils.WriteJson(w, http.StatusOK, payload)
	log.Print(err)
}
func List(w http.ResponseWriter, r *http.Request) {
	var tasks []models.Task
	initializers.DB.Find(&tasks)

	_ = utils.WriteJson(w, http.StatusOK, tasks)
}
