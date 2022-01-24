package main

import (
	"testing"

	"github.com/stretchr/testify/suite"
)

type TestSuite struct {
	suite.Suite
	database *Database
}

func (suite *TestSuite) SetupSuite() {
	suite.database = NewDatabase()
	suite.database.Connect()
}

func (suite *TestSuite) TearDownSuite() {
	suite.database.Disconnect()
}

func (suite *TestSuite) SetupTest() {
	suite.database.Set(1, "Item 1")
	suite.database.Set(2, "Item 2")
	suite.database.Set(3, "Item 3")
	suite.database.Set(4, "Item 4")
	suite.database.Set(5, "Item 5")
}

func (suite *TestSuite) TearDownTest() {
	suite.database.Clear()
}

func (suite *TestSuite) BeforeTest(suiteName, testName string) {
	println("Started testing", suiteName, testName)
}

func (suite *TestSuite) AfterTest(suiteName, testName string) {
	println("Finished testing", suiteName, testName)}

func TestExampleTestSuite(t *testing.T) {
	suite.Run(t, new(TestSuite))
}

func (suite *TestSuite) TestHashItem_ShouldReturnError_WhenIdIsNotFound() {
	// Arrange
	id := 10

	// Act
	hashed, err := HashItem(suite.database, id)

	// Assert
	suite.Empty(hashed)
	suite.NotNil(err)
}

func (suite *TestSuite) TestHashItem_ShouldReturnHashed_WhenIdIsFound() {
	// Arrange
	id := 3

	// Act
	hashed, err := HashItem(suite.database, id)

	// Assert
	suite.NotEmpty(hashed)
	suite.Nil(err)
}
