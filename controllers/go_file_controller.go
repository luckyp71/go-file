package controllers

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
	"strings"
)

//CreateFile is used to create a file
func CreateFile(path string) string {
	file, err := os.Create(path)
	if err != nil {
		fmt.Println(err)
		return "Error"
	}
	fmt.Println("File is created successfully")
	defer file.Close()
	return "Success"
}

//InsertData is used to insert data
func InsertData(path string) string {
	sfile, err := os.OpenFile(path, os.O_APPEND|os.O_WRONLY, 0777)
	file, err := ioutil.ReadFile(path)
	//If file does not exists
	if err != nil {
		CreateFile(path)
		sfile, _ = os.OpenFile(path, os.O_APPEND|os.O_WRONLY, 0777)
		noString := "1"
		fmt.Fprintf(sfile, "%s", noString+"."+"data "+noString)
		return "Success"
	}
	data := strings.Split(string(file), "\n")
	lastNumber := len(data)

	if data[0] != "" {
		line := strings.Split(data[lastNumber-1], ".")
		no, _ := strconv.Atoi(line[0])
		fmt.Println(no)
		noString := strconv.Itoa(no + 1)
		fmt.Fprintf(sfile, "\n%s", noString+"."+"data "+noString)
	} else {
		noString := "1"
		fmt.Fprintf(sfile, "%s", noString+"."+"data "+noString)
	}
	return "Success"
}

//ReadFile is used to read data from data.txt file
func ReadFile(path string) []string {
	file, err := os.Open(path)
	if err != nil {
		fmt.Println(err)
		panic(err)
	} else {
		defer file.Close()
		fileScanner := bufio.NewScanner(file)

		var data []string
		for fileScanner.Scan() {
			data = append(data, fileScanner.Text())
		}
		return data
	}
}
