package models

type Webhook struct {
	URL   string `json:"url"   bson:"url"   validate:"required"`
	Games []Game `json:"games" bson:"games"`
}
