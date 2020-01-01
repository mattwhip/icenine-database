package userdata

import (
	"encoding/json"
	"time"

	"github.com/gobuffalo/pop"
	"github.com/gobuffalo/uuid"
	"github.com/gobuffalo/validate"
	"github.com/gobuffalo/validate/validators"
)

// UdUser is a user with data in the system
type UdUser struct {
	ID               uuid.UUID `json:"id" db:"id"`
	CreatedAt        time.Time `json:"created_at" db:"created_at"`
	UpdatedAt        time.Time `json:"updated_at" db:"updated_at"`
	UID              string    `json:"u_id" db:"u_id"`
	Coins            int64     `json:"coins" db:"coins"`
	Rating           float64   `json:"rating" db:"rating"`
	RatingDeviation  float64   `json:"rating_deviation" db:"rating_deviation"`
	RatingVolatility float64   `json:"rating_volatility" db:"rating_volatility"`
}

// String is not required by pop and may be deleted
func (u UdUser) String() string {
	ju, _ := json.Marshal(u)
	return string(ju)
}

// UdUsers is not required by pop and may be deleted
type UdUsers []UdUser

// String is not required by pop and may be deleted
func (u UdUsers) String() string {
	ju, _ := json.Marshal(u)
	return string(ju)
}

// Validate gets run every time you call a "pop.Validate*" (pop.ValidateAndSave, pop.ValidateAndCreate, pop.ValidateAndUpdate) method.
// This method is not required and may be deleted.
func (u *UdUser) Validate(tx *pop.Connection) (*validate.Errors, error) {
	return validate.Validate(
		&validators.StringIsPresent{Field: u.UID, Name: "UID"},
	), nil
}

// ValidateCreate gets run every time you call "pop.ValidateAndCreate" method.
// This method is not required and may be deleted.
func (u *UdUser) ValidateCreate(tx *pop.Connection) (*validate.Errors, error) {
	return validate.NewErrors(), nil
}

// ValidateUpdate gets run every time you call "pop.ValidateAndUpdate" method.
// This method is not required and may be deleted.
func (u *UdUser) ValidateUpdate(tx *pop.Connection) (*validate.Errors, error) {
	return validate.NewErrors(), nil
}
