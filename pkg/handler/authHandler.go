package handler

import (
	"net/http"
)

func (h *Handler) SignUp(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("ASdasd"))
}

func (h *Handler) SignIn(w http.ResponseWriter, r *http.Request) {
	h.render(w, "signIn.tmpl", r)
}
