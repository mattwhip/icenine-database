package userdata

import (
	"encoding/json"
	"time"

	"github.com/gobuffalo/pop"
	"github.com/gobuffalo/uuid"
	"github.com/gobuffalo/validate"
)

// UdConfig holds user data service config
type UdConfig struct {
	ID                      uuid.UUID `json:"id" db:"id"`
	CreatedAt               time.Time `json:"created_at" db:"created_at"`
	UpdatedAt               time.Time `json:"updated_at" db:"updated_at"`
	InitialCoins            int64     `json:"init_coins" db:"init_coins"`
	InitialRating           float64   `json:"init_rating" db:"init_rating"`
	InitialRatingDeviation  float64   `json:"init_rating_deviation" db:"init_rating_deviation"`
	InitialRatingVolatility float64   `json:"init_rating_volatility" db:"init_rating_volatility"`
}

// String is not required by pop and may be deleted
func (u UdConfig) String() string {
	ju, _ := json.Marshal(u)
	return string(ju)
}

// UdConfigs is not required by pop and may be deleted
type UdConfigs []UdConfig

// String is not required by pop and may be deleted
func (u UdConfigs) String() string {
	ju, _ := json.Marshal(u)
	return string(ju)
}

// Validate gets run every time you call a "pop.Validate*" (pop.ValidateAndSave, pop.ValidateAndCreate, pop.ValidateAndUpdate) method.
// This method is not required and may be deleted.
func (u *UdConfig) Validate(tx *pop.Connection) (*validate.Errors, error) {
	return validate.NewErrors(), nil
}

// ValidateCreate gets run every time you call "pop.ValidateAndCreate" method.
// This method is not required and may be deleted.
func (u *UdConfig) ValidateCreate(tx *pop.Connection) (*validate.Errors, error) {
	return validate.NewErrors(), nil
}

// ValidateUpdate gets run every time you call "pop.ValidateAndUpdate" method.
// This method is not required and may be deleted.
func (u *UdConfig) ValidateUpdate(tx *pop.Connection) (*validate.Errors, error) {
	return validate.NewErrors(), nil
}
