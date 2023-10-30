package services

import (
	"Auth-API/models"
	"context"

	"go.mongodb.org/mongo-driver/mongo"
)

type UserService interface {
	GetAll() []models.UserModel
	GetById(id string) models.UserModel
	Create(user models.UserModel) models.UserModel
	Update(id string, user models.UserModel) models.UserModel
	GetByEmail(email string) (models.UserModel,error)
	// Delete(id string) models.UserModel
}

type UserServiceImplementation struct {
	userCollection *mongo.Collection
	ctx            context.Context
}

func NewUserService(userCollection *mongo.Collection, ctx context.Context) *UserServiceImplementation {
	return &UserServiceImplementation{
		userCollection: userCollection,
		ctx:            ctx,
	}
}

func (s *UserServiceImplementation) GetAll() []models.UserModel {
	var users []models.UserModel
	cursor, _ := s.userCollection.Find(s.ctx, models.UserModel{})
	for cursor.Next(s.ctx) {
		var user models.UserModel
		cursor.Decode(&user)
		users = append(users, user)
	}
	return users
}

func (s *UserServiceImplementation) GetById(id string) models.UserModel {
	var user models.UserModel
	_ = s.userCollection.FindOne(s.ctx, models.UserModel{Id: id}).Decode(&user)
	return user
}

func (s *UserServiceImplementation) Create(user models.UserModel) models.UserModel {
	_, err := s.userCollection.InsertOne(s.ctx, user)
	if err != nil {
		panic(err)
		// return
	}
	return user
}

func (s *UserServiceImplementation) Update(id string, user models.UserModel) models.UserModel {
	_, err := s.userCollection.UpdateOne(s.ctx, models.UserModel{Id: id}, user)
	if err != nil {
		panic(err)
	}
	return user
}

func (s *UserServiceImplementation) GetByEmail(Username string) (models.UserModel,error) {
	var user models.UserModel
	_ = s.userCollection.FindOne(s.ctx, models.UserModel{Username: Username}).Decode(&user)
	return user,nil
}
