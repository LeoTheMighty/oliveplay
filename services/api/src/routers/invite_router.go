package routers

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/leothemighty/oliveplay/utils/services/invite"
)

// RegisterPingRoutes attaches routes for the Ping worker.
func RegisterInviteRoutes(r chi.Router) {
	r.Get("/invites", func(w http.ResponseWriter, req *http.Request) {
		params, err := invite.FetchInvitesParams(req)
		if err != nil {

		}
		response, err := invite.FetchInvites(context.Background(), params)
		if err != nil {
			fmt.Println("Error fetching invites:", err)
			http.Error(w, "Unable to fetch ivites", http.StatusInternalServerError)
		}

		w.WriteHeader(http.StatusCreated)
		json.NewEncoder(w).Encode(response)
	})
}
