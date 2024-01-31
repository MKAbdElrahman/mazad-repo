package model

import (
	"time"
)

type Item struct {
	Name         string   `json:"item_name"`
	Description  string   `json:"item_description"`
	Category     string   `json:"item_category"`
	Manufacturer string   `json:"item_manufacturer"`
	Condition    string   `json:"item_condition"`
	Images       []string `json:"item_images"`
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
type AuctionPayload struct {
	Item         Item      `json:"item"`
	ReservePrice int       `json:"reserve_price"`
	AuctionEnd   time.Time `json:"auction_end"`
	Seller       string    `json:"seller"`
}
