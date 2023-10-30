package models

type UserModel struct {
	Id       string    `json:"id" bson:"_id" `
	Username string `json:"username" bson:"username" binding:"required"`
	Password string `json:"password" bson:"password" binding:"required"`
}

