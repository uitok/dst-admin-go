package entity

import "gorm.io/gorm"

type ModInfo struct {
	gorm.Model
	Auth          string                 `json:"auth"`
	ConsumerAppid float64                `json:"consumer_appid"`
	CreatorAppid  float64                `json:"creator_appid"`
	Description   string                 `json:"description"`
	FileUrl       string                 `json:"file_url"`
	Modid         string                 `json:"modid"`
	Img           string                 `json:"img"`
	LastTime      float64                `json:"last_time"`
	ModConfig     map[string]interface{} `gorm:"TYPE:json" json:"mod_config"`
	Name          string                 `json:"name"`
	V             string                 `json:"v"`
}
