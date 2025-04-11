package main

import (
	"net/http"
)

func (cfg *apiConfig) handlerChirpsList(w http.ResponseWriter, r *http.Request) {
	authorId := r.URL.Query().Get("author_id")

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

	if authorId != "" {
		var filteredChirps []Chirp
		for _, chirp := range chirps {
			if chirp.UserID.String() == authorId {
				filteredChirps = append(filteredChirps, chirp)
			}
		}
		chirps = filteredChirps
	}

	respondWithJSON(w, http.StatusOK,
		chirps,
	)

}
