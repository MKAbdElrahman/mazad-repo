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
func (s *AuctionStore) CreateAuction(payload model.CreateAuctionPayload) (*model.Auction, error) {
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

// GetAllAuctions retrieves all auctions from the PostgreSQL database.
func (s *AuctionStore) GetAllAuctions() ([]*model.Auction, error) {
	query := `
		SELECT id, item_name, item_description, item_category, item_manufacturer, item_condition, item_images, reserve_price, current_high_bid, seller, winner, created_at, updated_at, end_at, status
		FROM auctions
	`

	rows, err := s.db.Query(query)
	if err != nil {
		return nil, fmt.Errorf("failed to get all auctions: %v", err)
	}
	defer rows.Close()

	var auctions []*model.Auction

	for rows.Next() {
		var auction model.Auction
		var imageUrlsBytes []byte

		err := rows.Scan(
			&auction.ID,
			&auction.Item.Name,
			&auction.Item.Description,
			&auction.Item.Category,
			&auction.Item.Manufacturer,
			&auction.Item.Condition,
			&imageUrlsBytes,
			&auction.ReservePrice,
			&auction.CurrentHighBid,
			&auction.Seller,
			&auction.Winner,
			&auction.CreatedAt,
			&auction.UpdatedAt,
			&auction.EndAt,
			&auction.Status,
		)
		if err != nil {
			return nil, fmt.Errorf("failed to scan row: %v", err)
		}

		// Convert the byte array to a string and then split it into an array of strings
		auction.Item.Images = strings.Split(string(imageUrlsBytes), ",")

		auctions = append(auctions, &auction)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("error reading rows: %v", err)
	}

	return auctions, nil
}

// GetAuctionByID retrieves an auction from the PostgreSQL database by ID.
func (s *AuctionStore) GetAuctionByID(auctionID string) (*model.Auction, error) {
	query := `
		SELECT id, item_name, item_description, item_category, item_manufacturer, item_condition, item_images, reserve_price, current_high_bid, seller, winner, created_at, updated_at, end_at, status
		FROM auctions
		WHERE id = $1
	`

	var auction model.Auction
	var imageUrlsBytes []byte

	err := s.db.QueryRow(query, auctionID).Scan(
		&auction.ID,
		&auction.Item.Name,
		&auction.Item.Description,
		&auction.Item.Category,
		&auction.Item.Manufacturer,
		&auction.Item.Condition,
		&imageUrlsBytes,
		&auction.ReservePrice,
		&auction.CurrentHighBid,
		&auction.Seller,
		&auction.Winner,
		&auction.CreatedAt,
		&auction.UpdatedAt,
		&auction.EndAt,
		&auction.Status,
	)

	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil // Auction not found
		}
		return nil, fmt.Errorf("failed to get auction by ID: %v", err)
	}

	// Convert the byte array to a string and then split it into an array of strings
	auction.Item.Images = strings.Split(string(imageUrlsBytes), ",")

	return &auction, nil
}
