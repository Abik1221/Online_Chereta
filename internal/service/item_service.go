package services

import (
    "bidding-system/internal/models"
    "bidding-system/internal/repositories"
)

type ItemService struct {
    itemRepo *repositories.ItemRepository
}

func NewItemService(itemRepo *repositories.ItemRepository) *ItemService {
    return &ItemService{itemRepo: itemRepo}
}

// CreateItem creates a new item
func (s *ItemService) CreateItem(item *models.Item) error {
    return s.itemRepo.CreateItem(item)
}

// GetItems returns all available items for bidding
func (s *ItemService) GetItems() ([]models.Item, error) {
    return s.itemRepo.GetItems()
}

// GetItemByID returns an item by its ID
func (s *ItemService) GetItemByID(itemID uint) (*models.Item, error) {
    return s.itemRepo.GetItemByID(itemID)
}