package services

import (
	"bidding-system/internal/models"
	"bidding-system/internal/repositories"
	"errors"
)

type BidService struct {
	bidRepo *repositories.BidRepository
}

func NewBidService(bidRepo *repositories.BidRepository) *BidService {
	return &BidService{bidRepo: bidRepo}
}

// PlaceBid places a bid on an item
func (s *BidService) PlaceBid(userID, itemID uint, bidAmount float64) (*models.Bid, error) {
	// Check if the bid amount is greater than the current highest bid
	highestBid, err := s.bidRepo.GetHighestBidForItem(itemID)
	if err == nil && bidAmount <= highestBid.BidAmount {
		return nil, errors.New("bid amount must be higher than the current highest bid")
	}

	// Create the bid
	bid := &models.Bid{
		UserID:    userID,
		ItemID:    itemID,
		BidAmount: bidAmount,
		Status:    "active",
	}

	if err := s.bidRepo.CreateBid(bid); err != nil {
		return nil, err
	}

	// Notify the previous highest bidder
	if highestBid != nil {
		notificationService.SendBidNotification(highestBid.UserID, "You have been outbid!")
	}

	return bid, nil
}

// GetUserBids returns all bids placed by a user
func (s *BidService) GetUserBids(userID uint) ([]models.Bid, error) {
	return s.bidRepo.GetBidsByUserID(userID)
}
