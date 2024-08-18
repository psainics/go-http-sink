package models

import (
	"time"

	"gorm.io/gorm"
)

type Room struct {
	gorm.Model
	ID int `json:"id" gorm:"primaryKey"`
}

func (r *Room) TableName() string {
	return "rooms"
}

type TestCase struct {
	gorm.Model
	ID   int    `json:"id" gorm:"primaryKey"`
	Name string `json:"name"`
}

func (t *TestCase) TableName() string {
	return "test_cases"
}

type TestResult struct {
	gorm.Model
	ID         int       `json:"id" gorm:"primaryKey"`
	TestCase   TestCase  `json:"testCase" gorm:"foreignKey:TestCaseID"`
	Room       Room      `json:"room" gorm:"foreignKey:RoomID"`
	StatusCode int       `json:"statusCode"`
	HTTPMethod string    `json:"httpMethod"`
	Body       string    `json:"body"`
	CreatedAt  time.Time `json:"createdAt"`
	TestCaseID int       `json:"testCaseId"`
	RoomID     int       `json:"roomId"`
}

func (t *TestResult) TableName() string {
	return "test_results"
}
