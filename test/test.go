package main

import (
	"fmt"
	"transporte/src/library/date"
	"transporte/src/library/lib"
)

func main() {

	var fact []map[string]interface{}
	fact = append(fact, map[string]interface{}{
		"periodo": "2022-04",
	})
	fact = append(fact, map[string]interface{}{
		"periodo": "2022-05",
	})
	var newFact []string
	for _, v := range fact {
		newFact = append(newFact, v["periodo"].(string))

	}
	fmt.Println(newFact)
	index := lib.IndexOfStrings(newFact, "2022-08")
	fmt.Println(index)

	date_fact := date.GetDate("02/04/2020")
	date_now := date.GetDateLocation()
	//date_end := date.GetLastDateOfMonth(date_fact)
	/*if date_fact == date_end {
		is_last_of_month = true
	}*/

	month_init := int64(date_fact.Month())
	year_init := date_fact.Year()
	month_now := int64(date_now.Month())
	year_now := date_now.Year()

	//var data_facturaciones []map[string]interface{}
	var month = int64(12)

	for i := year_init; i <= year_now; i++ {
		if i == year_now {
			month = month_now
		}
		for e := month_init; e <= month; e++ {
			// fmt.Println(i, e)
			year := fmt.Sprintf("%v", i)
			month := fmt.Sprintf("%02d", e)
			if lib.IndexOfStrings(newFact, year+"-"+month) == -1 {
				fmt.Println(year, month)
			}
		}

		month_init = 1
	}
}
