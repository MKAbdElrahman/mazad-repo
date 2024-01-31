package business

import "mazad/services/auctions/model"

// AuctionStore is an interface that defines methods for storing auctions.
type AuctionStore interface {
	CreateAuction(payload model.CreateAuctionPayload) (*model.Auction, error)
}

type AuctionService struct {
	AuctionStore AuctionStore
}

func NewAuctionService(auctionStore AuctionStore) *AuctionService {
	return &AuctionService{AuctionStore: auctionStore}
}
func (s *AuctionService) CreateAuction(payload model.CreateAuctionPayload) (*model.Auction, error) {
	return s.AuctionStore.CreateAuction(payload)
}
