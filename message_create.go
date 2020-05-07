package main

import "github.com/go-xorm/xorm"

type Service struct {
	db *xorm.Engine
}

type Message struct {
	Id             int32  `xorm:"'id' pk autoincr"`
	Name           string `xorm:"name"`
	Email          string `xorm:"email"`
	MessengerType  string `xorm:"messenger_type"`
	MessengerValue string `xorm:"messenger_value"`
	Comment        string `xorm:"comment"`
}

func (s *Message) TableName() string {
	return "messages"
}

func (s *Service) createMessage(request Message) (Message,  error) {
	_, err := s.db.Omit("id").InsertOne(&request)
	return request, err
}
