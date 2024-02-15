package types

import (
	"net/http"
	"strconv"
	"time"
)

type Message struct {
	Message     string     `db:"message"`
	Name        string     `db:"name"`
	Phone       string     `db:"phone"`
	Email       string     `db:"email"`
	MessageType string     `db:"message_type"`
	TourDate    *time.Time `db:"tour_date"`
	TourTime    *time.Time `db:"tour_time"`
	PropertyId  int        `db:"property_id"`
	UserId      int        `db:"user_id"`
}

func (m Message) GetRelationName() string {
	return "messages"
}

func (m Message) GetPrimaryKeyName() string {
	return "message_id"
}

type MessageFilter struct {
	Name        string     `db:"name"`
	Phone       string     `db:"phone"`
	Email       string     `db:"email"`
	MessageType string     `db:"message_type"`
	TourDate    *time.Time `db:"tour_date"`
	TourTime    *time.Time `db:"tour_time"`
	PropertyId  int        `db:"property_id"`
	UserId      int        `db:"user_id"`
}

func (m Message) GetRelationForFilterName() string {
	return "messages"
}

type MessageSummary struct {
}

func ParseMessageBody(r *http.Request) (*Message, error) {

	propertyId, err := strconv.Atoi(r.PostFormValue("property_id"))

	if err != nil {
		return nil, err
	}

	userId, err := strconv.Atoi(r.PostFormValue("user_id"))

	if err != nil {
		return nil, err
	}

	message := &Message{
		Message:     r.PostFormValue("message"),
		Name:        r.PostFormValue("name"),
		Phone:       r.PostFormValue("phone"),
		Email:       r.PostFormValue("email"),
		MessageType: r.PostFormValue("type"),
		TourDate:    nil,
		TourTime:    nil,
		PropertyId:  propertyId,
		UserId:      userId,
	}
	return message, nil
}

func ParseTourMessageBody(r *http.Request) (*Message, error) {
	propertyId, err := strconv.Atoi(r.PostFormValue("property_id"))
	const dateLayout = "2012-01-20"
	const timeLayout = "15:03:45"

	time_, err := time.Parse(timeLayout, r.PostFormValue("time"))

	if err != nil {
		return nil, err
	}

	date, err := time.Parse(dateLayout, r.PostFormValue("date"))

	if err != nil {
		return nil, err
	}

	message := &Message{
		Message:     r.PostFormValue("message"),
		Name:        r.PostFormValue("name"),
		Phone:       r.PostFormValue("phone"),
		Email:       r.PostFormValue("email"),
		MessageType: r.PostFormValue("type"),
		TourDate:    &date,
		TourTime:    &time_,
		PropertyId:  propertyId,
	}
	return message, nil
}
