package store

import (
	"database/sql"
	"fmt"
	"mazad/services/auctions/model"
	"strings"
	"time"

	"github.com/lib/pq"
)

// PostgreSQLStore is an implementation of the AuctionStore interface using PostgreSQL.
type AuctionStore struct {
	db *sql.DB
}

// NewPostgreSQLStore creates a new instance of PostgreSQLStore.
func NewPostgresAuctionStore(db *sql.DB) *AuctionStore {
	return &AuctionStore{db: db}
}

// CreateAuction inserts a new auction into the PostgreSQL database and returns the created auction.
func (s *AuctionStore) CreateAuction(payload model.AuctionPayload) (*model.Auction, error) {
	// Convert []string to pq.Array for handling TEXT ARRAY in PostgreSQL
	imageURLs := pq.Array(payload.Item.Images)

	// Insert the auction into the PostgreSQL database
	query := `
		INSERT INTO auctions (item_name, item_description, item_category, item_manufacturer, item_condition, item_images, reserve_price, current_high_bid, seller, winner, created_at, updated_at, end_at, status)
		VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13, $14)
		RETURNING id, item_name, item_description, item_category, item_manufacturer, item_condition, item_images, reserve_price, current_high_bid, seller, winner, created_at, updated_at, end_at, status
	`

	var createdAuction model.Auction
	var imageUrlsBytes []byte
	err := s.db.QueryRow(
		query,
		payload.Item.Name,
		payload.Item.Description,
		payload.Item.Category,
		payload.Item.Manufacturer,
		payload.Item.Condition,
		imageURLs,
		payload.ReservePrice,
		0, // Initial current high bid, adjust as needed
		payload.Seller,
		"", // Initial winner, adjust as needed
		time.Now(),
		time.Now(),
		payload.AuctionEnd,
		model.Live,
	).Scan(
		&createdAuction.ID,
		&createdAuction.Item.Name,
		&createdAuction.Item.Description,
		&createdAuction.Item.Category,
		&createdAuction.Item.Manufacturer,
		&createdAuction.Item.Condition,
		&imageUrlsBytes,
		&createdAuction.ReservePrice,
		&createdAuction.CurrentHighBid,
		&createdAuction.Seller,
		&createdAuction.Winner,
		&createdAuction.CreatedAt,
		&createdAuction.UpdatedAt,
		&createdAuction.EndAt,
		&createdAuction.Status,
	)

	if err != nil {
		return nil, fmt.Errorf("failed to create auction: %v", err)
	}

	// Convert the byte array to a string and then split it into an array of strings
	createdAuction.Item.Images = strings.Split(string(imageUrlsBytes), ",")

	return &createdAuction, nil
}
