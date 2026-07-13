package main

import (
	"fmt"
	"io"
	"net/http"
)

func Cumprimenta(escritor io.Writer, nome string) {
	fmt.Fprintf(escritor, "Olá, %s", nome)
}

func HandleMeuCumprimento(w http.ResponseWriter, r *http.Request) {
	Cumprimenta(w, "mundo")
}

func main() {
	err := http.ListenAndServe(":5000", http.HandlerFunc(HandleMeuCumprimento))

	if err != nil {
		fmt.Println(err)
	}
}
