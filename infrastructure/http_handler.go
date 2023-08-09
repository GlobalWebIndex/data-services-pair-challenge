package infrastructure

import (
	"encoding/json"
	"log"
	"net/http"
	"strings"

	"github.com/GlobalWebIndex/data-services-pair-challenge/domain"
)

// CreateAudienceHandler handles the HTTP request for creating an audience.
type CreateAudienceHandler struct {
	audienceService *domain.AudienceService
}

// NewCreateAudienceHandler creates a new CreateAudienceHandler with the provided AudienceService.
func NewCreateAudienceHandler(audienceService *domain.AudienceService) *CreateAudienceHandler {
	return &CreateAudienceHandler{
		audienceService: audienceService,
	}
}

// ServeHTTP handles the HTTP request for creating an audience.
// Option 2: fix this http handler to return correctly ID of newly
// created record in the JSON response.
func (handler *CreateAudienceHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	var audience domain.Audience
	err := json.NewDecoder(r.Body).Decode(&audience)
	if err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	// Call the AudienceService to create the audience.
	createdAudience, err := handler.audienceService.CreateAudience(r.Context(), audience.Name, audience.Expression)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Marshal the created audience as JSON and send it in the response.
	w.Header().Set("Content-Type", "application/json")
	err = json.NewEncoder(w).Encode(createdAudience)
	if err != nil {
		http.Error(w, "Failed to encode response", http.StatusInternalServerError)
		return
	}
}

// AudienceByIDHandler handles HTTP requests related to the audience entity.
type AudienceByIDHandler struct {
	AudienceService *domain.AudienceService
}

// NewGetAudienceByIDHandler creates a new instance of AudienceHandler.
func NewGetAudienceByIDHandler(audienceService *domain.AudienceService) *AudienceByIDHandler {
	return &AudienceByIDHandler{
		AudienceService: audienceService,
	}
}

// GetAudienceByID handles the request to get an audience by its ID.
func (h *AudienceByIDHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	id := strings.TrimPrefix(r.URL.Path, "/audiences/")
	audienceID := domain.AudienceID(id)

	audience, err := h.AudienceService.GetAudienceByID(r.Context(), audienceID)
	if err != nil {
		log.Printf("Error while fetching audience by ID: %v", err)
		http.Error(w, "Failed to fetch audience", http.StatusInternalServerError)
		return
	}

	if audience == nil {
		http.NotFound(w, r)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(audience)
}
