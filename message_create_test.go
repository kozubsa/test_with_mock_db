package main

import (
	"fmt"
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/go-xorm/xorm"
	_ "github.com/lib/pq"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCreateMessage(t *testing.T) {
	asrt := assert.New(t)

	s := &Service{}

	engine, err := xorm.NewEngine("postgres", "dbname=test sslmode=disable")
	if !asrt.NoError(err) {
		return
	}

	engine.ShowSQL()

	db, mock, err := sqlmock.New()
	if !asrt.NoError(err) {
		return
	}

	engine.DB().DB = db
	s.db = engine

	tests := []struct {
		name string
		args Message
	}{
		{
			args: Message{
				Name:           "Name",
				Email:          "Email",
				MessengerType:  "MessengerType",
				MessengerValue: "MessengerValue",
				Comment:        "Comment",
			},
		},
		{
			args: Message{
				Name:           "Name 2",
				Email:          "Email 2",
				MessengerType:  "MessengerType 2",
				MessengerValue: "MessengerValue 2",
				Comment:        "Comment 2",
			},
		},
	}

	for key, test := range tests {
		name := fmt.Sprintf("%v %v", key, test.name)
		mock.
			ExpectQuery(`INSERT INTO "messages"`).
			WithArgs(
				test.args.Name,
				test.args.Email,
				test.args.MessengerType,
				test.args.MessengerValue,
				test.args.Comment,
			).
			WillReturnRows(sqlmock.NewRows([]string{"id"}).AddRow(1))

		msg, err := s.createMessage(test.args)
		if !asrt.NoError(err, name) {
			return
		}
		if !asrt.NotZero(msg.Id, name) {
			return
		}
		asrt.Equal(msg.Name, test.args.Name, name)
		asrt.Equal(msg.Email, test.args.Email, name)
		asrt.Equal(msg.MessengerType, test.args.MessengerType, name)
		asrt.Equal(msg.MessengerValue, test.args.MessengerValue, name)
		asrt.Equal(msg.Comment, test.args.Comment, name)
	}
}
