package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/go-chi/chi/v5"
)

/*
Ejercicio 2 - Manipulando el body
Vamos a crear un endpoint llamado /greetings. Con una pequeña estructura con nombre y
apellido que al pegarle deberá responder en texto “Hello + nombre + apellido”
El endpoint deberá ser de método POST
Se deberá usar el package JSON para resolver el ejercicio
La respuesta deberá seguir esta estructura: “Hello Andrea Rivas”
La estructura deberá ser como esta:
{
    “firstName”: “Andrea”,
    “lastName”: “Rivas”
}
*/

type Person struct {
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
}

func main() {
	rt := chi.NewRouter()

	rt.Post("/greetings", func(w http.ResponseWriter, r *http.Request) {
		var person Person
		err := json.NewDecoder(r.Body).Decode(&person)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		message := fmt.Sprintf("Hello %s %s", person.FirstName, person.LastName)
		w.Write([]byte(message))
		w.Header().Add("Content-Type", "text/plain")
		w.WriteHeader(http.StatusOK)
	})

	if err := http.ListenAndServe(":8080", rt); err != nil {
		panic(err)
	}
}
