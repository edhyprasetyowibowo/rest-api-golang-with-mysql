package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"github.com/repo/rest_api/users"
	"github.com/repo/rest_api/utils"
)

func main() {

	http.HandleFunc("/users", GetUsers)
	err := http.ListenAndServe(":7000", nil)

	if err != nil {
		log.Fatal(err)
	}
}


func GetUsers(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		ctx, cancel := context.WithCancel(context.Background())

		defer cancel()

		users, err := users.GetAll(ctx)

		if err != nil {
			fmt.Println(err)
		}

		utils.ResponseJSON(w, users, http.StatusOK)
		return
	}

	http.Error(w, "Tidak di ijinkan", http.StatusNotFound)
	return
}
