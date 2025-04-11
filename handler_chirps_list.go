package main

import (
	"net/http"
	"sort"
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

	sortDirection := "asc"
	sortDirectionParam := r.URL.Query().Get("sort")
	if sortDirectionParam == "desc" {
		sortDirection = "desc"
	}

	sort.Slice(chirps, func(i, j int) bool {
		if sortDirection == "desc" {
			return chirps[i].CreatedAt.After(chirps[j].CreatedAt)
		}
		return chirps[i].CreatedAt.Before(chirps[j].CreatedAt)
	})

	respondWithJSON(w, http.StatusOK,
		chirps,
	)

}
