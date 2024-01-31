package handler

import (
	"encoding/json"
	"mazad/services/auctions/model"
	"net/http"

	"github.com/go-chi/chi/v5"
)

// AuctionStore is an interface that defines methods for storing auctions.
type AuctionService interface {
	CreateAuction(payload model.CreateAuctionPayload) (*model.Auction, error)
	GetAllAuctions() ([]*model.Auction, error)
	GetAuctionByID(auctionID string) (*model.Auction, error)
}

type AuctionHandler struct {
	AuctionService AuctionService
}

// NewAuctionHandler creates a new instance of AuctionHandler.
func NewAuctionHandler(auctionService AuctionService) *AuctionHandler {
	return &AuctionHandler{AuctionService: auctionService}
}

func (h *AuctionHandler) CreateAuction(w http.ResponseWriter, r *http.Request) {
	var payload model.CreateAuctionPayload

	// Decode the incoming JSON payload
	err := json.NewDecoder(r.Body).Decode(&payload)
	if err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	// Store the new auction and get the created auction
	createdAuction, storeErr := h.AuctionService.CreateAuction(payload)
	if storeErr != nil {
		http.Error(w, "Failed to store the auction", http.StatusInternalServerError)
		return
	}

	// Respond with the created auction
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(createdAuction)
}

func (h *AuctionHandler) GetAuctionByID(w http.ResponseWriter, r *http.Request) {
	// Extract the auction ID from the request parameters or URL path
	// For example, assuming the auction ID is passed as a query parameter:
	auctionID := chi.URLParam(r, "id")

	// Implement logic to retrieve the auction by ID (you'll need to define this in your AuctionService)
	// For example:
	auction, err := h.AuctionService.GetAuctionByID(auctionID)
	if err != nil {
		http.Error(w, "Failed to get the auction by ID", http.StatusInternalServerError)
		return
	}

	// Check if the auction is not found
	if auction == nil {
		http.Error(w, "Auction not found", http.StatusNotFound)
		return
	}

	// Respond with the auction
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(auction)
}

func (h *AuctionHandler) GetAllAuctions(w http.ResponseWriter, r *http.Request) {

	allAuctions, err := h.AuctionService.GetAllAuctions()
	if err != nil {
		http.Error(w, "Failed to get all auctions", http.StatusInternalServerError)
		return
	}

	// Respond with the list of all auctions
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(allAuctions)
}
