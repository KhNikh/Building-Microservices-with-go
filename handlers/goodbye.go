package handlers

import (
	"fmt"
	"io"
	"log"
	"net/http"
)

type Goodbye struct {
	l *log.Logger
}

func NewBye(l *log.Logger) *Goodbye {
	return &Goodbye{l}
}
func (g *Goodbye) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	g.l.Println("Good Bye")
	item, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(rw, "Something went wrong", http.StatusBadRequest)
		return
	}
	fmt.Fprintf(rw, "Goodbye %s\n", item)
}
