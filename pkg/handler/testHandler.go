package handler

import "net/http"

func (h *Handler) CreateTest(w http.ResponseWriter, r *http.Request) {
	h.render(w, "createTest.tmpl", nil, 200)
}
