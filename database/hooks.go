package database

import (
	"fmt"
	"net/mail"

	"gorm.io/gorm"
)

// BeforeCreate hook to check if `StatsSend: Email` field is valid.
func (stats *SendStats) BeforeCreate(tx *gorm.DB) (err error) {
	_, err = mail.ParseAddress(stats.Email)
	if err != nil {
		return RaiseGenericError(fmt.Sprintf("\nEmail is not valid.\nError: %s\nEmail: %s", err, stats.Email))
	}
	return
}