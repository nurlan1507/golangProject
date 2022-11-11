package handler

import (
	"encoding/json"
	"fmt"
	"net/http"
	"testApp/pkg/helpers"
	"testApp/pkg/models"
)

type testForm struct {
	Validation *helpers.Validation
}

func (h *Handler) CreateTest(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		h.Loggers.ErrorLogger.Println(err)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	newTest := &models.TestModel{}
	err = json.NewDecoder(r.Body).Decode(&newTest)
	testForm := &testForm{Validation: helpers.NewValidation()}
	if err != nil {
		h.Loggers.ErrorLogger.Println(err)
		testForm.Validation.Errors["badRequest"] = "invalid request, try to reload the page"
		resp, _ := json.Marshal(testForm.Validation.Errors)
		w.WriteHeader(400)
		w.Write(resp)
		return
	}
	//newTest.StartAt = time.Now().Add(time.Hour * 3)
	//validation
	testForm.Validation.Check(helpers.NotEmpty(newTest.Title), "title", "title can not be empty")
	testForm.Validation.Check(helpers.NotEmpty(newTest.Description), "description", "description can not be empty")
	//testForm.Validation.Check(helpers.NotEmpty(newTest.GroupId), "group", "invited group can not be empty")
	//testForm.Validation.Check(helpers.NotEmptyTime(newTest.StartAt), "startAt", "start date can not be empty")
	//testForm.Validation.Check(helpers.TimeIsValid(newTest.StartAt), "startAt", "start date can not be equal or less then current time")
	if !testForm.Validation.Valid() {
		w.WriteHeader(400)
		fmt.Println(testForm.Validation.Errors)
		resp, _ := json.Marshal(testForm.Validation.Errors)
		w.Write(resp)
		return
	}

	fmt.Printf("%+v", newTest)
	result, err := h.TestService.CreateTest(newTest)
	if err != nil {
		h.Loggers.ErrorLogger.Println(err)
		w.WriteHeader(400)
	}

	mars, err := json.Marshal(result)
	if err != nil {
		return
	}
	w.Header().Add("SAlam", "salam")
	_, err = w.Write(mars)
	if err != nil {
		h.Loggers.ErrorLogger.Println(err)
		w.WriteHeader(500)
		return
	}
}
