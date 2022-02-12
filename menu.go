package main

import (
	"math"
	"strconv"
	"strings"
	"time"
)

const (
	gr_lunch   = "Μεσημεριανό"
	gr_dinner  = "Βραδινό"
	gr_first   = "ΠΡΩΤΟ ΠΙΑΤΟ"
	gr_main    = "ΚΥΡΙΩΣ ΠΙΑΤΑ"
	gr_salad   = "ΣΑΛΑΤΑ"
	gr_bread   = "ΨΩΜΙ"
	gr_dessert = "ΓΛΥΚΟ"
	gr_fruit   = "ΦΡΟΥΤΟ"
)

type MenuArray []string

func (m MenuArray) findIndexWhereContains(value string) int {
	for i, v := range m {
		if strings.Contains(v, value) {
			return i
		}
	}
	return -1
}

type Menu struct {
	First   []string `json:"first"`
	Main    []string `json:"main"`
	Salad   []string `json:"salad"`
	Bread   []string `json:"bread"`
	Dessert []string `json:"dessert"`
	Fruit   bool     `json:"fruit"`
}

func getMaxInt(values ...int) int {
	max := math.MinInt
	for _, v := range values {
		if v > max {
			max = v
		}
	}
	return max
}

func NewMenuFromArray(array MenuArray) *Menu {
	menu := new(Menu)

	firstIndex := array.findIndexWhereContains(gr_first)
	mainIndex := array.findIndexWhereContains(gr_main)
	saladIndex := array.findIndexWhereContains(gr_salad)
	breadIndex := array.findIndexWhereContains(gr_bread)

	dessertIndex := array.findIndexWhereContains(gr_dessert)
	fruitIndex := array.findIndexWhereContains(gr_fruit)

	lastIndex := getMaxInt(dessertIndex, fruitIndex)

	menu.First = array[firstIndex+1 : mainIndex]
	menu.Main = array[mainIndex+1 : saladIndex]
	menu.Salad = array[saladIndex+1 : breadIndex]
	menu.Bread = array[breadIndex+1 : lastIndex]
	menu.Dessert = []string{}
	menu.Fruit = false

	if dessertIndex != -1 {
		menu.Dessert = array[dessertIndex+1:]
	}
	if fruitIndex != -1 {
		menu.Fruit = true
	}

	return menu
}

type DayMenu struct {
	Date   time.Time `json:"date"`
	Lunch  Menu      `json:"lunch"`
	Dinner Menu      `json:"dinner"`
}

func NewDayMenuFromArray(array MenuArray) *DayMenu {
	date := array[0]
	dateSplit := strings.Split(date, " ")
	date = dateSplit[len(dateSplit)-1]
	dateSplit = strings.Split(date, "/")
	year, _ := strconv.Atoi(dateSplit[2])
	month, _ := strconv.Atoi(dateSplit[1])
	day, _ := strconv.Atoi(dateSplit[0])

	newDate := time.Date(year, time.Month(month), day, 0, 0, 0, 0, time.UTC)

	lunchIndex := array.findIndexWhereContains(gr_lunch)
	dinnerIndex := array.findIndexWhereContains(gr_dinner)

	lunchArray := array[lunchIndex+1 : dinnerIndex]
	lunchMenu := NewMenuFromArray(lunchArray)

	dinner := array[dinnerIndex+1:]
	dinnerMenu := NewMenuFromArray(dinner)

	dayMenu := new(DayMenu)
	dayMenu.Date = newDate
	dayMenu.Lunch = *lunchMenu
	dayMenu.Dinner = *dinnerMenu
	return dayMenu
}

type WeekMenu struct {
	Menus     [7]DayMenu `json:"menus"`
	UpdatedAt time.Time  `json:"updated_at"`
}

func NewWeekMenuFromWeb() (*WeekMenu, bool) {
	var weekArrays [7]MenuArray = getMenuFromWeb()

	if len(weekArrays[0]) == 0 {
		return nil, false
	}

	var weekMenu = new(WeekMenu)
	weekMenu.UpdatedAt = time.Now()
	for i, day := range weekArrays {
		weekMenu.Menus[i] = *NewDayMenuFromArray(day)
	}
	return weekMenu, true
}

func (w WeekMenu) menuOfWeekday(weekDay Weekday) DayMenu {
	return w.Menus[weekDay-1]
}

func (w WeekMenu) menuOfDate(date time.Time) (DayMenu, bool) {
	for _, v := range w.Menus {
		if isSameDay(v.Date, date) {
			return v, true
		}
	}
	return DayMenu{}, false
}
