package domain

import (
	"fmt"

	"github.com/google/uuid"
)

type Crop struct {
	ID      uuid.UUID
	Species CropSpecies
	Variety int // 0 is first variety, 1 is second variety
	// TODO: Add more fields
	// Field   Field
	// PlantingDate time.Time
	// HarvestDate time.Time
	// Yield float64
}

func NewCrop(id uuid.UUID, species CropSpecies, variety int) (*Crop, error) {
	return &Crop{
		ID:      id,
		Species: species,
		Variety: variety,
	}, nil
}

func (c *Crop) String() string {
	return fmt.Sprintf("Crop{ID: %s, Species: %s, Variety: %d}", c.ID, c.Species, c.Variety)
}
