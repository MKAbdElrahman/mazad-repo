package model

import (
	"time"
)

type Item struct {
	Name         string   `json:"name"`
	Description  string   `json:"description"`
	Category     string   `json:"category"`
	Manufacturer string   `json:"manufacturer"`
	Condition    string   `json:"condition"`
	Images       []string `json:"images"`
}

type Auction struct {
	ID             string        `json:"id"`
	Item           Item          `json:"item"`
	ReservePrice   int           `json:"reserve_price"`
	CurrentHighBid int           `json:"current_high_bid"`
	Seller         string        `json:"seller"`
	Winner         string        `json:"winner"`
	CreatedAt      time.Time     `json:"created_at"`
	UpdatedAt      time.Time     `json:"updated_at"`
	EndAt          time.Time     `json:"auction_end"`
	Status         AuctionStatus `json:"status"`
}

// AuctionStatus represents the possible states of an auction.
type AuctionStatus int

const (
	Live AuctionStatus = iota
	Finished
	Cancelled
	ReserveNotMet
)

// AuctionPayload represents the data expected when creating a new auction.
type CreateAuctionPayload struct {
	Item         Item      `json:"item"`
	ReservePrice int       `json:"reserve_price"`
	AuctionEnd   time.Time `json:"auction_end"`
	Seller       string    `json:"seller"`
}

// UpdateAuctionPayload represents the data expected when updating an existing auction.
type UpdateAuctionPayload struct {
	AuctionID      string        `json:"auction_id"`
	ReservePrice   int           `json:"reserve_price"`
	CurrentHighBid int           `json:"current_high_bid"`
	AuctionEnd     time.Time     `json:"auction_end"`
	Status         AuctionStatus `json:"status"`
}


