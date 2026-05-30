package domain

import "errors"

var (
	ErrInvalidSpecies = errors.New("invalid species")
)

type CropSpecies struct {
	value string
}

var (
	Maize     CropSpecies = CropSpecies{value: "maize"}
	Soybean   CropSpecies = CropSpecies{value: "soybean"}
	Sunflower CropSpecies = CropSpecies{value: "sunflower"}
	Wheat     CropSpecies = CropSpecies{value: "wheat"}
)

func (s CropSpecies) String() string {
	return s.value
}
