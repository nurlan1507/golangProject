package templateData

import "testApp/pkg/models"

type TemplateData struct {
	AuthData *models.UserModel
	Form     any
}

func NewTemplateData(AuthData *models.UserModel, Form any) *TemplateData {
	return &TemplateData{AuthData: AuthData, Form: Form}
}
