package main

import (
	"encoding/json"

	"net/http"

	"github.com/google/uuid"
	"github.com/kaipov24/Chirpy/internal/auth"
)

func (cfg *apiConfig) handlerUserUpgrade(w http.ResponseWriter, r *http.Request) {
	type parameters struct {
		Event string `json:"event"`
		Data  struct {
			UserId string `json:"user_id"`
		} `json:"data"`
	}
	type response struct {
	}

	decoder := json.NewDecoder(r.Body)
	params := parameters{}
	err := decoder.Decode(&params)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "Couldn't decode parameters", err)
		return
	}

	if params.Event != "user.upgraded" {
		respondWithError(w, http.StatusNoContent, "Invalid event", nil)
		return
	}
	stringId, err := uuid.Parse(params.Data.UserId)
	if err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid user ID", err)
		return
	}
	_, err = cfg.db.RedChirpUser(r.Context(), stringId)
	if err != nil {
		if err.Error() == "sql: no rows in result set" {
			respondWithError(w, http.StatusNotFound, "User not found", err)
			return
		}
		respondWithError(w, http.StatusInternalServerError, "Couldn't update user", err)
		return
	}

	apiKey, err := auth.GetAPIKey(r.Header)

	if err != nil {
		respondWithError(w, http.StatusUnauthorized, "Invalid API key", err)
		return
	}

	apiKeyMatch := cfg.polkaKey == apiKey

	if !apiKeyMatch {
		respondWithError(w, http.StatusUnauthorized, "Invalid API key", nil)
		return
	}

	respondWithJSON(w, http.StatusNoContent, response{})
}
