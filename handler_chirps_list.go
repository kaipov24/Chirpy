package main

import (
	"net/http"
)

func (cfg *apiConfig) handlerChirpsList(w http.ResponseWriter, r *http.Request) {

	users, err := cfg.db.GetChirps(r.Context())

	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "Couldn't get users", err)
		return
	}

	var chirps []Chirp
	for _, user := range users {
		chirps = append(chirps, Chirp{
			ID:        user.ID,
			CreatedAt: user.CreatedAt,
			UpdatedAt: user.UpdatedAt,
			Body:      user.Body,
			UserID:    user.UserID,
		})
	}

	respondWithJSON(w, http.StatusOK,
		chirps,
	)

}
