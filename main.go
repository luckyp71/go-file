package main

import (
	//	"bufio"
	//	"fmt"
	c "go-file/controllers"
	//	"os"
)

func main() {
	//	c.CreateFile("./assets/data.txt")

	//	lines := c.ReadFile("./assets/data.txt")

	//	for lines.Scan() {
	//		fmt.Println(lines.Text())
	//	}

	//	for _, v := range lines {
	//		fmt.Println(v)
	//	}

	c.InsertData("./assets/data.txt")

}
