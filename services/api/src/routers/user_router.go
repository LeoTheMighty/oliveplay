package routers

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/go-chi/chi/v5"

	"github.com/leothemighty/oliveplay/utils/services/user"
)

func RegisterUserRoutes(r chi.Router) {
	r.Get("/user", func(w http.ResponseWriter, req *http.Request) {
		params, err := user.FetchUserParams(req)
		if err != nil {
			fmt.Println("Invalid Request Parameters: ", params, "Err: ", err)
			http.Error(w, "Invalid Request Parameters", http.StatusBadRequest)
		}

		response, err := user.FetchUser(context.Background(), params)
		if err != nil {
			fmt.Println("Error Fetching User: ", err)
			http.Error(w, "Unable to fetch user", http.StatusInternalServerError)
		}

		w.WriteHeader(http.StatusAccepted)
		json.NewEncoder(w).Encode(response)
	})
}
