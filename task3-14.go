package main

import "fmt"

func printGetHolidaysByIdOfMonthResult() {
	fmt.Println("Task 3 - variant 14")
	fmt.Println("Test case - 1")
	fmt.Println(getHolidaysByIdOfMonth(1))
}

func getHolidaysByIdOfMonth(idOfMonth int) []string {
	switch idOfMonth {
	case 1:
		return []string{"New Year's Day (January 1st)"}
	case 2:
		return []string{"Defender of Ukraine Day (February 23rd)"}
	case 3:
		return []string{"International Women's Day (March 8th)"}
	case 4:
		return []string{"Easter (variable date in April)"}
	case 5:
		return []string{"Victory Day (May 9th)"}
	case 6:
		return []string{"Constitution Day (June 28th)"}
	case 7:
		return []string{"Baptism of Kyivan Rus (July 28th)"}
	case 8:
		return []string{"Independence Day (August 24th)"}
	case 9:
		return []string{"Day of Knowledge (September 1st)"}
	case 10:
		return []string{"Ukrainian Cossack Day (October 14th)"}
	case 11:
		return []string{"Day of Dignity and Freedom (November 21st)"}
	case 12:
		return []string{"Christmas (December 25th)"}
	default:
		return []string{"There is no such month!"}
	}
}
