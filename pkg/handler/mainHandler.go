package handler

import "net/http"

func (h *Handler) Home(w http.ResponseWriter, r *http.Request) {
	//w.Write([]byte("Asd"))
	h.render(w, "home.tmpl", nil, 200)
}
