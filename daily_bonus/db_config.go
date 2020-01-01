package dailybonus

import (
	"encoding/json"
	"time"

	"github.com/gobuffalo/pop"
	"github.com/gobuffalo/uuid"
	"github.com/gobuffalo/validate"
)

// DbConfig holds Daily Bonus configuration
type DbConfig struct {
	ID                 uuid.UUID `json:"id" db:"id"`
	CreatedAt          time.Time `json:"created_at" db:"created_at"`
	UpdatedAt          time.Time `json:"updated_at" db:"updated_at"`
	ResetSeconds       int32     `json:"reset_seconds" db:"reset_seconds"`
	StreakBreakSeconds int32     `json:"streak_break_seconds" db:"streak_break_seconds"`
	WheelsJSON         string    `json:"wheels_json" db:"wheels_json"`
}

// String is not required by pop and may be deleted
func (d DbConfig) String() string {
	jd, _ := json.Marshal(d)
	return string(jd)
}

// DbConfigs is not required by pop and may be deleted
type DbConfigs []DbConfig

// String is not required by pop and may be deleted
func (d DbConfigs) String() string {
	jd, _ := json.Marshal(d)
	return string(jd)
}

// Validate gets run every time you call a "pop.Validate*" (pop.ValidateAndSave, pop.ValidateAndCreate, pop.ValidateAndUpdate) method.
// This method is not required and may be deleted.
func (d *DbConfig) Validate(tx *pop.Connection) (*validate.Errors, error) {
	return validate.NewErrors(), nil
}

// ValidateCreate gets run every time you call "pop.ValidateAndCreate" method.
// This method is not required and may be deleted.
func (d *DbConfig) ValidateCreate(tx *pop.Connection) (*validate.Errors, error) {
	return validate.NewErrors(), nil
}

// ValidateUpdate gets run every time you call "pop.ValidateAndUpdate" method.
// This method is not required and may be deleted.
func (d *DbConfig) ValidateUpdate(tx *pop.Connection) (*validate.Errors, error) {
	return validate.NewErrors(), nil
}
