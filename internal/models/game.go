package models

type Game struct {
	ID       string `bson:"game_id"   json:"id"        validate:"required"`
	Name     string `bson:"name"      json:"name"      validate:"required"`
	ImageURL string `bson:"image_url" json:"image_url" validate:"required"`
}
