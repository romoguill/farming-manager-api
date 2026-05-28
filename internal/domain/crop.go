package domain

import (
	"errors"
	"fmt"

	"github.com/google/uuid"
)

type Crop struct {
	ID      uuid.UUID
	Species CropSpecies
	Variety CropVariety
}

type CropSpecies string

const (
	Maize     = "maize"
	Soybean   = "soybean"
	Sunflower = "sunflower"
	Wheat     = "wheat"
)

type CropVariety string

const (
	FirstVariety  = "first"
	SecondVariety = "second"
)

func NewSpecies(species string) (CropSpecies, error) {
	switch species {
	case string(Maize):
		return Maize, nil
	case string(Soybean):
		return Soybean, nil
	case string(Sunflower):
		return Sunflower, nil
	case string(Wheat):
		return Wheat, nil
	default:
		return "", errors.New("invalid species")
	}
}

func (s CropSpecies) String() string {
	return string(s)
}

func NewVariety(variety string) (CropVariety, error) {
	switch variety {
	case string(FirstVariety):
		return FirstVariety, nil
	case string(SecondVariety):
		return SecondVariety, nil
	default:
		return "", errors.New("invalid variety")
	}
}

func (v CropVariety) String() string {
	return string(v)
}

func NewCrop(id uuid.UUID, species CropSpecies, variety CropVariety) (*Crop, error) {
	return &Crop{
		ID:      id,
		Species: species,
		Variety: variety,
	}, nil
}

func (c *Crop) String() string {
	return fmt.Sprintf("Crop{ID: %s, Species: %s, Variety: %s}", c.ID, c.Species, c.Variety)
}
