package handlers

import (
	"fmt"
	"log"
	"net/http"
)

type Goodbye struct {
	logger *log.Logger
}

func NewGoodbye(l *log.Logger) *Goodbye {
	return &Goodbye{l}
}

func (h *Goodbye) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	h.logger.Println("Handle Goodbye request")
	fmt.Fprint(rw, "Goodbye")
}
