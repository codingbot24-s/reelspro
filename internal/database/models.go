package database

import (
	"time"
)

type User struct {
	ID        uint      `json:"id"`
	Username  string    `json:"username" validate:"required,min=3"`
	Password  string    `json:"password" validate:"required,min=6"`
	Email     string    `json:"email" validate:"required,email"`
	CreatedAt time.Time `json:"created_at" bson:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type Video struct {
	Id             uint   `json:"id"`
	Title          string `json:"title" validate:"required"`
	Description    string `json:"description" validate:"required"`
	VideoUrl       string `json:"video_url" validate:"required"`
	ThumbnailUrl   string `json:"thumbnail_url" validate:"required"`
	Controls       bool   `json:"controls"`
	Transformation struct {
		Width   int `json:"width"`
		Height  int `json:"height"`
		Quality int `json:"quality"`
	}
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

// VideoDimensions object for checking the video dimensions
// type VideoDimensions struct {
//
// }
