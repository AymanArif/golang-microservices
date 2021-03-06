package handlers

import (
	"log"
	"net/http"
)

// Goodbye https
type Goodbye struct {
	l *log.Logger
}

// NewGoodbye https://
func NewGoodbye(l *log.Logger) *Goodbye {
	return &Goodbye{l}
}

func (g *Goodbye) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	g.l.Println("Decompsoited with go-mods: Goodbye. ")
	rw.Write([]byte("Decompsoited with go-mods: Goodbye\n"))
}
