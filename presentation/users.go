package presentation

import (
	"athenify/app/services"
	"athenify/domain"
	"athenify/persistence"
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
func (uh UserHandler) Create(wp *persistence.WorkerPool) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// Convert JSON to User structure
		var user domain.User
		json.NewDecoder(r.Body).Decode(&user)

		resultCh := make(chan domain.Result, 1)
		wp.Jobs <- &services.CreateUserJob{User: user, UserService: uh.us, Result: resultCh}

		result := <-resultCh
		defer close(resultCh)

		// Handle result
		if result.Error != nil {
			http.Error(w, result.Error.Error(), http.StatusBadRequest)
			return
		}

		// Convert User structure to []byte
		userBytes, err := json.Marshal(result.User)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
		}
		w.WriteHeader(http.StatusAccepted)
		w.Write(userBytes)
	}
}

func (uh UserHandler) Get(wp *persistence.WorkerPool) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// Get user ID from Header and convert it to UUID format
		userIDString := r.Header.Get("user_id")
		if userIDString == "" {
			http.Error(w, "Unable to find user ID", http.StatusForbidden)
		}
		userID, err := uuid.Parse(userIDString)
		if err != nil {
			http.Error(w, "Invalid UUID format", http.StatusBadRequest)
		}

		resultCh := make(chan domain.Result, 1)
		wp.Jobs <- &services.GetUserJob{UserID: userID, UserService: uh.us, Result: resultCh}

		result := <-resultCh
		defer close(resultCh)

		// Handle result
		if result.Error != nil {
			http.Error(w, result.Error.Error(), http.StatusBadRequest)
			return
		}

		// Convert User structure to []byte
		userBytes, err := json.Marshal(result.User)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
		}
		w.WriteHeader(http.StatusAccepted)
		w.Write(userBytes)
	}
}
