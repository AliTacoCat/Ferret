package main

import (
	"fmt"
	"os"
	"time"
)

// type Animal interface {
//Speak() string
//}

//type Dog struct{}

//func (d Dog) Speak() string {
//	return "Woof!"
//}

//func (d Dog) Fetch() string {
//	return "Fetching the ball!"
//}

//type Cat struct{}

//func (c Cat) Speak() string {
//	return "Meow!"
//}

//func SPEAKNEW(a Animal) {
//	fmt.Println(a.Speak())
//}

func main() {

	// Create instances of Dog and Cat
	// myDog := Dog{}
	// myCat := Cat{}

	//SPEAKNEW(myDog)

	start := time.Now()

	files, err := os.ReadDir("/Users/alize/downloads")

	if err != nil {
		fmt.Println("Error reading directory:", err)
		return
	}

	count := 0

	for _, file := range files {
		info, _ := file.Info() // gets detailed info
		fmt.Printf("%s - Size: %d bytes\n", file.Name(), info.Size())
		if file.IsDir() {
			fmt.Println("\n 📁", file.Name(), "(folder)")
		} else {
			fmt.Println("\n 📄", file.Name(), "(file)")
		}
		if !file.IsDir() {
			count++
		}
	}

	elapsed := time.Since(start)

	fmt.Printf("\n Total files: %d\n", count)
	fmt.Printf("Execution time: %s\n", elapsed)
}
