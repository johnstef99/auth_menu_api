package main

import (
	"encoding/json"
	"log"
	"os"
)

const cacheFileName string = "cache.json"

var loadedMenu *WeekMenu

func main() {
	loadedMenu = loadWeekMenu()
	serveAPI()
}

func loadWeekMenu() *WeekMenu {
	weekMenu, err := readWeekMenuFromFile(cacheFileName)
	if os.IsNotExist(err) {
		log.Println("Cache file doesn't exists")
		ok := updateMenu(weekMenu)
		if ok {
			log.Println("Cache file was updated")
		} else {
			log.Fatal("No data to work with...Exiting")
		}
	} else if err != nil {
		log.Fatal(err)
	}
	return weekMenu
}

func updateMenu(weekMenu *WeekMenu) bool {
	weekMenu, ok := NewWeekMenuFromWeb()
	if ok {
		writeWeekMenuToFile(weekMenu, cacheFileName)
	} else {
		log.Println("Cache was not updated because no new data was received")
	}
	return ok
}

func writeWeekMenuToFile(weekMenu *WeekMenu, filename string) error {
	js, err := json.Marshal(weekMenu)
	if err != nil {
		log.Fatal(err)
	}
	return os.WriteFile(filename, js, 0644)
}

func readWeekMenuFromFile(filename string) (*WeekMenu, error) {
	cache, err := os.ReadFile(filename)
	weekMenu := new(WeekMenu)

	if err != nil {
		return weekMenu, err
	}

	err = json.Unmarshal(cache, &weekMenu)

	if err != nil {
		return weekMenu, err
	}

	return weekMenu, err
}
