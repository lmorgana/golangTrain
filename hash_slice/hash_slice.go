package main

import (
	"bufio"
	"encoding/gob"
	"fmt"
	"io"
	"os"
	"regexp"
	"strings"
)

type userBook struct {
	Name    string
	Surname string
	Phone   string
	Id      int
}

type phoneBook struct {
	Mp map[string]userBook
}

func isPhoneNumber(phoneNumber string) bool {
	result, _ := regexp.MatchString("(8[0-9]{10})", phoneNumber)
	return result
}

func (book *phoneBook) Add(args []string) {
	phoneNumber := args[0]
	if !isPhoneNumber(phoneNumber) {
		fmt.Println("wrong phone number:", args[0])
		return
	}
	_, isExist := book.Mp[phoneNumber]
	if isExist {
		fmt.Println("number in use")
		return
	}
	var user userBook
	user.Name = args[1]
	user.Surname = args[2]
	user.Phone = phoneNumber
	book.Mp[phoneNumber] = user
}

func lookUp(mp map[string]userBook, findBy, find string) []userBook {
	var slice []userBook

	switch findBy {
	case "name":
		for _, user := range mp {
			if user.Name == find {
				slice = append(slice, user)
			}
		}
	case "surname":
		for _, user := range mp {
			if user.Surname == find {
				slice = append(slice, user)
			}
		}
	case "ALL":
		for _, user := range mp {
			slice = append(slice, user)
		}
	case "number":
		slice = append(slice, mp[find])
	}
	return slice
}

func (book *phoneBook) Find(args []string) {
	var find string
	findBy := "name"
	if len(book.Mp) < 1 {
		fmt.Println("error: data is empty")
		return
	} else if len(args) == 1 && args[0] == "ALL" {
		findBy = "ALL"
	} else if args[0] == "BY" && len(args) >= 3 {
		findBy = args[1]
		find = args[2]
	} else {
		find = args[0]
	}
	users := lookUp(book.Mp, findBy, find)
	for i, user := range users {
		fmt.Printf("№%d %s %s %s\n", i, user.Name, user.Surname, user.Phone)
	}
}

func (book *phoneBook) Delete(args []string) {
	delete(book.Mp, args[0])
}

func process(line string, book *phoneBook) {
	line = strings.Trim(line, "\n")
	arguments := strings.Split(line, " ")
	if len(arguments) < 2 {
		fmt.Println("need more arguments")
		return
	}
	switch arguments[0] {
	case "ADD":
		book.Add(arguments[1:])
	case "FIND":
		book.Find(arguments[1:])
	case "DELETE":
		book.Delete(arguments[1:])

	default:
		fmt.Println("unknown command: ", arguments[0])

	}
}

func loadData(fileName string) (*phoneBook, error) {
	var book phoneBook
	book.Mp = make(map[string]userBook)
	file, err := os.Open(fileName)
	if err != nil {
		return &book, err
	}
	defer file.Close()
	decoder := gob.NewDecoder(file)
	err = decoder.Decode(&book)
	if err != nil {
		fmt.Println("error: can't decode file:", err)
	}
	return &book, nil
}

func saveData(fileName string, book *phoneBook) error {
	err := os.Remove(fileName)
	file, err := os.OpenFile(fileName, os.O_CREATE|os.O_WRONLY, 0667)
	if err != nil {
		return err
	}
	defer file.Close()
	encoder := gob.NewEncoder(file)
	err = encoder.Encode(book)
	if err != nil {
		fmt.Println("error: can't encode file", err)
	}
	return nil
}

func main() {
	saveFile := "saveFile.gob"
	fmt.Print("Hello you in phoneBook, some command:\n" +
		"ADD phoneNumber name surname - add new number,(phone number begin by 8)\n" +
		"DELETE phoneNumber - delete number by name,\n" +
		"FIND BY name - show information about user ('BY' is identificator, example FIND BY name andrey\n" +
		"To EXIT press ctrl + D\n")
	buff := bufio.NewReader(os.Stdin)
	book, err := loadData(saveFile)
	if err != nil {
		fmt.Println("Some problem with save file, we will create new one")
	}
	for {
		line, err := buff.ReadString('\n')
		if err == io.EOF {
			break
		} else if err != nil {
			fmt.Printf("error: error when reading file %v", err)
			return
		}
		process(line, book)
	}
	err = saveData(saveFile, book)
	if err != nil {
		fmt.Println("error: can't write new logs")
		return
	}
	fmt.Println("Bye bye")
}
