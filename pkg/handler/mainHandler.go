package handler

import "net/http"

func (h *Handler) Home(w http.ResponseWriter, r *http.Request) {
	h.render(w, "home.tmpl", nil, 200)
}
