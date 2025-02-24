package repositories

import (
    "bidding-system/internal/models"
    "gorm.io/gorm"
)

type ItemRepository struct {
    db *gorm.DB
}

func NewItemRepository(db *gorm.DB) *ItemRepository {
    return &ItemRepository{db: db}
}

// CreateItem creates a new item
func (r *ItemRepository) CreateItem(item *models.Item) error {
    return r.db.Create(item).Error
}

// GetItems returns all available items for bidding
func (r *ItemRepository) GetItems() ([]models.Item, error) {
    var items []models.Item
    if err := r.db.Find(&items).Error; err != nil {
        return nil, err
    }
    return items, nil
}

// GetItemByID returns an item by its ID
func (r *ItemRepository) GetItemByID(itemID uint) (*models.Item, error) {
    var item models.Item
    if err := r.db.First(&item, itemID).Error; err != nil {
        return nil, err
    }
    return &item, nil
}