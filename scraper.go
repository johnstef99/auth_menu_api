package main

import (
	"log"
	"strconv"

	"github.com/gocolly/colly/v2"
)

func getMenuFromWeb() [7]MenuArray {
	c := colly.NewCollector()

	var weekArrays [7]MenuArray

	day := 1                  // start from monday
	for i := 1; i <= 7; i++ { // loop to call 7 times the OnHTML function
		c.OnHTML(".kt-accordion-inner-wrap > .kt-accordion-pane-"+strconv.Itoa(i)+" .kt-accordion-panel-inner", func(h *colly.HTMLElement) {
			weekArrays[day-1] = h.ChildTexts("p, h3, h4, h2, li")
			day++ // increment day by one after this OnHTML executes
		})
	}

	err := c.Visit("https://www.auth.gr/weekly-menu/")
	if err != nil {
		log.Println(err)
	}
	return weekArrays
}
