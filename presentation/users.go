package presentation

import (
	"athenify/domain"
	"encoding/json"
	"net/http"

	"github.com/google/uuid"
)

type UserHandler struct {
	us domain.UserService
}

func NewUserHandler(us domain.UserService) *UserHandler {
	return &UserHandler{us: us}
}

// Create user
func (uh UserHandler) Create(w http.ResponseWriter, r *http.Request) {
	// Convert JSON to User structure
	var user domain.User
	json.NewDecoder(r.Body).Decode(&user)

	// Create user
	user, err := uh.us.Create(user)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}

	// Convert User structure to []byte
	userBytes, err := json.Marshal(user)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}
	w.WriteHeader(http.StatusAccepted)
	w.Write(userBytes)
}

func (uh UserHandler) Get(w http.ResponseWriter, r *http.Request) {
	// Get user ID from Header and convert it to UUID format
	userIDString := r.Header.Get("user_id")
	if userIDString == "" {
		http.Error(w, "Unable to find user ID", http.StatusForbidden)
	}
	userID, err := uuid.Parse(userIDString)
	if err != nil {
		http.Error(w, "Invalid UUID format", http.StatusBadRequest)
	}

	// Get user by ID
	user, err := uh.us.GetByID(userID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}

	// Convert User structure to []byte
	userBytes, err := json.Marshal(user)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}
	w.WriteHeader(http.StatusAccepted)
	w.Write(userBytes)
}
