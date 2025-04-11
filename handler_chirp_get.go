package main

import (
	"fmt"
	"net/http"
	"path"

	"github.com/google/uuid"
)

func (cfg *apiConfig) handlerChirpGet(w http.ResponseWriter, r *http.Request) {
	chirpID, err := uuid.Parse(path.Base(r.URL.Path))
	if err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid chirp ID", err)
		return
	}

	chirp, err := cfg.db.GetChirpByID(r.Context(), chirpID)
	if err != nil {
		respondWithError(w, http.StatusNotFound, "Couldn't get chirp", err)
		return
	}
	fmt.Printf("Chirp: %+v\n", chirp)
	respondWithJSON(w, http.StatusOK, Chirp{
		ID:        chirp.ID,
		CreatedAt: chirp.CreatedAt,
		UpdatedAt: chirp.UpdatedAt,
		Body:      chirp.Body,
		UserID:    chirp.UserID,
	})
}
