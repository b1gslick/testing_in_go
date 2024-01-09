package data

import "time"

// UserImage is the type for user profile images.
type UserImage struct {
	CreatedAt time.Time `json:"-"`
	UpdatedAt time.Time `json:"-"`
	UserID    string    `json:"user_id"`
	FileName  string    `json:"file_name"`
	ID        int       `json:"id"`
}
