package handler

import (
	"encoding/json"
	"mazad/services/auctions/model"
	"net/http"
)

// AuctionStore is an interface that defines methods for storing auctions.
type AuctionService interface {
	CreateAuction(payload model.AuctionPayload) (*model.Auction, error)
}

type AuctionHandler struct {
	AuctionService AuctionService
}

// NewAuctionHandler creates a new instance of AuctionHandler.
func NewAuctionHandler(auctionService AuctionService) *AuctionHandler {
	return &AuctionHandler{AuctionService: auctionService}
}

func (h *AuctionHandler) HandleCreateAuction(w http.ResponseWriter, r *http.Request) {
	var payload model.AuctionPayload

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
