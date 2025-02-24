package repositories

import (
    "bidding-system/internal/models"
    "gorm.io/gorm"
)

type BidRepository struct {
    db *gorm.DB
}

func NewBidRepository(db *gorm.DB) *BidRepository {
    return &BidRepository{db: db}
}

// CreateBid creates a new bid
func (r *BidRepository) CreateBid(bid *models.Bid) error {
    return r.db.Create(bid).Error
}

// GetBidsByUserID returns all bids placed by a user
func (r *BidRepository) GetBidsByUserID(userID uint) ([]models.Bid, error) {
    var bids []models.Bid
    if err := r.db.Where("user_id = ?", userID).Find(&bids).Error; err != nil {
        return nil, err
    }
    return bids, nil
}

// GetHighestBidForItem returns the highest bid for an item
func (r *BidRepository) GetHighestBidForItem(itemID uint) (*models.Bid, error) {
    var bid models.Bid
    if err := r.db.Where("item_id = ?", itemID).Order("bid_amount DESC").First(&bid).Error; err != nil {
        return nil, err
    }
    return &bid, nil
}