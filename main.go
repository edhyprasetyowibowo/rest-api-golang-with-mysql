package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"github.com/repo/rest_api/models"
	"github.com/repo/rest_api/users"
	"github.com/repo/rest_api/utils"
)

func main() {

	http.HandleFunc("/users", GetUsers)
	http.HandleFunc("/users/create", PostUsers)
	http.HandleFunc("/users/update", UpdateUsers)
	http.HandleFunc("/users/delete", DeleteUsers)

	err := http.ListenAndServe(":7000", nil)

	if err != nil {
		log.Fatal(err)
	}
}

//GetUsers
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

// PostUsers
func PostUsers(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		if r.Header.Get("Content-Type") != "application/json" {
			http.Error(w, "Gunakan content type application / json", http.StatusBadRequest)
			return
		}

		ctx, cancel := context.WithCancel(context.Background())
		defer cancel()

		var user models.Users

		if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
			utils.ResponseJSON(w, err, http.StatusBadRequest)
			return
		}

		if err := users.Insert(ctx, user); err != nil {
			utils.ResponseJSON(w, err, http.StatusInternalServerError)
			return
		}

		res := map[string]string{
			"status": "Succesfully",
		}

		utils.ResponseJSON(w, res, http.StatusCreated)
		return
	}

	http.Error(w, "Tidak di ijinkan", http.StatusMethodNotAllowed)
	return
}

// UpdateUsers
func UpdateUsers(w http.ResponseWriter, r *http.Request) {
    if r.Method == "PUT" {
 
        if r.Header.Get("Content-Type") != "application/json" {
            http.Error(w, "Gunakan content type application / json", http.StatusBadRequest)
            return
        }
 
        ctx, cancel := context.WithCancel(context.Background())
        defer cancel()
 
        var user models.Users
 
        if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
            utils.ResponseJSON(w, err, http.StatusBadRequest)
            return
        }
 
        fmt.Println(user)
 
        if err := users.Update(ctx, user); err != nil {
            utils.ResponseJSON(w, err, http.StatusInternalServerError)
            return
        }
 
        res := map[string]string{
            "status": "Succesfully",
        }
 
        utils.ResponseJSON(w, res, http.StatusCreated)
        return
    }
 
    http.Error(w, "Tidak di ijinkan", http.StatusMethodNotAllowed)
    return
}


// DeleteUsers
func DeleteUsers(w http.ResponseWriter, r *http.Request) {
 
    if r.Method == "DELETE" {
 
        ctx, cancel := context.WithCancel(context.Background())
        defer cancel()
 
        var user models.Users
 
        id := r.URL.Query().Get("id")
 
        if id == "" {
            utils.ResponseJSON(w, "id tidak boleh kosong", http.StatusBadRequest)
            return
        }
        user.ID, _ = strconv.Atoi(id)
 
        if err := users.Delete(ctx, user); err != nil {
 
            kesalahan := map[string]string{
                "error": fmt.Sprintf("%v", err),
            }
 
            utils.ResponseJSON(w, kesalahan, http.StatusInternalServerError)
            return
        }
 
        res := map[string]string{
            "status": "Succesfully",
        }
 
        utils.ResponseJSON(w, res, http.StatusOK)
        return
    }
 
    http.Error(w, "Tidak di ijinkan", http.StatusMethodNotAllowed)
    return
}
