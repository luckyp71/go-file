package main

import (
	c "go-file/controllers"
	"testing"
)

//TestCreeateFile aims to test CreateFile method in go_file_controller
func TestCreateFile(t *testing.T) {
	var ExpectedResult, ActualResult string
	data := c.CreateFile("../assets/data.txt")
	ActualResult = data
	ExpectedResult = "Success"
	if ActualResult != ExpectedResult {
		t.Fatalf("Expected %s but got %s", ExpectedResult, ActualResult)
	}

}

//TestInsertData aims to test InsertData method in go_file_controller
func TestInsertData(t *testing.T) {
	var ExpectedResult, ActualResult string
	data := c.InsertData("../assets/data.txt")
	ActualResult = data
	ExpectedResult = "Success"
	if ActualResult != ExpectedResult {
		t.Fatalf("Expected %s but got %s", ExpectedResult, ActualResult)
	}
}

//TestReadFile aims to test ReadFile method in go_file_controller
func TestReadFile(t *testing.T) {
	var ExpectedResult, ActualResult []string
	data := c.ReadFile("../assets/data.txt")
	ActualResult = data
	ExpectedResult = []string{"1.data 1"}
	if ActualResult[0] != ExpectedResult[0] {
		t.Fatalf("Expected %s but got %s", ExpectedResult, ActualResult)
	}
}
