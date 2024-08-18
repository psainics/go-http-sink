package models

import (
	"time"

	"gorm.io/gorm"
)

type Room struct {
	gorm.Model
	ID int `json:"id" gorm:"primaryKey"`
}

type TestCase struct {
	gorm.Model
	ID   int    `json:"id" gorm:"primaryKey"`
	Name string `json:"name"`
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
