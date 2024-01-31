package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"mazad/services/auctions/model"
	"net/http"
	"time"
)

const serverURL = "http://localhost:3000" // Update with your server URL

func main() {
	auctions := []model.CreateAuctionPayload{
		{
			Item: model.Item{
				Name:         "Artwork 1",
				Description:  "Beautiful painting",
				Category:     "Art",
				Manufacturer: "Artist A",
				Condition:    "New",
				Images:       []string{"image_url_1", "image_url_2"},
			},
			ReservePrice: 500,
			AuctionEnd:   time.Now().Add(24 * time.Hour), // Auction ends in 24 hours
			Seller:       "seller_1",
		},
		{
			Item: model.Item{
				Name:         "Electronics 1",
				Description:  "Smartphone",
				Category:     "Electronics",
				Manufacturer: "Brand X",
				Condition:    "Used",
				Images:       []string{"image_url_3", "image_url_4"},
			},
			ReservePrice: 800,
			AuctionEnd:   time.Now().Add(48 * time.Hour), // Auction ends in 48 hours
			Seller:       "seller_2",
		},
		// Add more auctions as needed
		{
			Item: model.Item{
				Name:         "Furniture 1",
				Description:  "Wooden chair",
				Category:     "Furniture",
				Manufacturer: "Craftsman Y",
				Condition:    "Like New",
				Images:       []string{"image_url_5", "image_url_6"},
			},
			ReservePrice: 300,
			AuctionEnd:   time.Now().Add(72 * time.Hour), // Auction ends in 72 hours
			Seller:       "seller_3",
		},
		{
			Item: model.Item{
				Name:         "Fashion 1",
				Description:  "Designer dress",
				Category:     "Fashion",
				Manufacturer: "Designer Z",
				Condition:    "Brand New",
				Images:       []string{"image_url_7", "image_url_8"},
			},
			ReservePrice: 600,
			AuctionEnd:   time.Now().Add(96 * time.Hour), // Auction ends in 96 hours
			Seller:       "seller_4",
		},
		// Additional auctions
		{
			Item: model.Item{
				Name:         "Books Collection",
				Description:  "Set of classic novels",
				Category:     "Books",
				Manufacturer: "Various Authors",
				Condition:    "New",
				Images:       []string{"image_url_9", "image_url_10"},
			},
			ReservePrice: 200,
			AuctionEnd:   time.Now().Add(120 * time.Hour), // Auction ends in 120 hours
			Seller:       "seller_5",
		},
		{
			Item: model.Item{
				Name:         "Gaming Console",
				Description:  "Latest gaming console",
				Category:     "Electronics",
				Manufacturer: "GamingTech",
				Condition:    "Brand New",
				Images:       []string{"image_url_11", "image_url_12"},
			},
			ReservePrice: 1000,
			AuctionEnd:   time.Now().Add(144 * time.Hour), // Auction ends in 144 hours
			Seller:       "seller_6",
		},
		// Additional auctions
		{
			Item: model.Item{
				Name:         "Collectible Coins Set",
				Description:  "Rare coins from different eras",
				Category:     "Collectibles",
				Manufacturer: "Numismatic Artifacts",
				Condition:    "Like New",
				Images:       []string{"image_url_13", "image_url_14"},
			},
			ReservePrice: 1500,
			AuctionEnd:   time.Now().Add(168 * time.Hour), // Auction ends in 168 hours
			Seller:       "seller_7",
		},
		{
			Item: model.Item{
				Name:         "Home Theater System",
				Description:  "High-end audio and video setup",
				Category:     "Electronics",
				Manufacturer: "AudioVision",
				Condition:    "New",
				Images:       []string{"image_url_15", "image_url_16"},
			},
			ReservePrice: 1200,
			AuctionEnd:   time.Now().Add(192 * time.Hour), // Auction ends in 192 hours
			Seller:       "seller_8",
		},
		// Add more auctions as needed
	}

	for _, auction := range auctions {
		// Convert AuctionPayload to JSON
		payload, err := json.Marshal(auction)
		if err != nil {
			log.Printf("Error marshalling JSON: %v", err)
			continue
		}

		// Make HTTP POST request to create auction
		resp, err := http.Post(serverURL+"/api/auctions", "application/json", bytes.NewBuffer(payload))
		if err != nil {
			log.Printf("Error making HTTP POST request: %v", err)
			continue
		}
		defer resp.Body.Close()

		// Check the response status
		if resp.StatusCode != http.StatusCreated {
			log.Printf("Failed to create auction. Status code: %v", resp.StatusCode)
			continue
		}

		// Decode the response body
		var createdAuction model.Auction
		err = json.NewDecoder(resp.Body).Decode(&createdAuction)
		if err != nil {
			log.Printf("Error decoding response body: %v", err)
			continue
		}

		// Display the created auction ID
		fmt.Printf("Created auction with ID: %s\n", createdAuction.ID)
	}
}
