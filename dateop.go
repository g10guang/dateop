package dateop

import (
	"fmt"
	"time"
)

// Mark down how many days in a month.
// And Feb in leap year will add one more day.
var months = [13]int{0, 31, 28, 31, 30, 31, 30, 31, 31, 30, 31, 30, 31}

var days = [...]string{
	"Sunday",
	"Monday",
	"Tuesday",
	"Wednesday",
	"Thursday",
	"Friday",
	"Saturday",
}

var daysIndexMap = map[string]int {
	"Sunday": 0, "Monday": 1, "Tuesday": 2, "Wednesday": 3, "Thursday": 4, "Friday": 5, "Saturday": 6,
}

func isLeap(year int) bool {
	if year%4 == 0 && year%100 != 0 {
		return true
	}
	return year%400 == 0
}

func getDate(year, month int) int {
	if isLeap(year) && month == 2 {
		return months[2] + 1
	}
	return months[month]
}

func (d Date) GetDate() int {
	return getDate(d.Year, d.Month)
}

func (d Date) IsLeap() bool {
	return isLeap(d.Year)
}

type Date struct {
	Year  int
	Month int
	Day   int
}

func (d Date) Compare(o Date) int {
	switch {
	case d.Year > o.Year:
		return 1
	case d.Year == o.Year:
		switch {
		case d.Month > o.Month:
			return 1
		case d.Month == o.Month:
			switch {
			case d.Day == o.Day:
				return 0
			case d.Day > o.Day:
				return 1
			default:
				return -1
			}
		default:
			return -1
		}
	default:
		return -1
	}
}

// Calculate the interval from d to o.
func (d Date) Interval(o Date) (intervalDays int) {
	switch d.Compare(o) {
	case 0:
		return 0
	case -1:
		return o.Interval(d)
	}
	// Assume d > o
	leapYearsNum := 0
	for i := o.Year + 1; i < d.Year; {
		if isLeap(i) {
			leapYearsNum++
			i += 4
		} else {
			i++
		}
	}
	// interval days calculate by years gab.
	intervalDays += (d.Year-o.Year-1)*365 + leapYearsNum
	// calculate o.Year-o.Month-o.Day - o.Year-12-31
	intervalDays += o.GetDate() - o.Day + d.Day
	for i := o.Month + 1; i <= 12; i++ {
		intervalDays += getDate(o.Year, i)
	}
	for i := 1; i < d.Month; i++ {
		intervalDays += getDate(d.Year, i)
	}
	return intervalDays
}

// Calculate the week of d in d.Year.
func (d Date) Week() string {
	now := time.Now()
	dateNow := Date{now.Year(), int(now.Month()), now.Day()}
	interval := d.Interval(dateNow) % 7
	weekdayNow := int(now.Weekday())
	switch d.Compare(dateNow) {
	case 0:
		return days[weekdayNow]
	case 1:
		return days[(weekdayNow+interval)%7]
	default:
		t := 6 - (weekdayNow + interval) % 7
		return days[t]
	}
}

// Print d.Month in d.Year's calendar.
func (d Date) PrintCalendar() {
	//	print header
	fmt.Println("SUN\tMON\tTUE\tWED\tTHU\tFRI\tSAT")
	firstDay := Date{d.Year, d.Month, 1}
	firstDayWeekday := daysIndexMap[firstDay.Week()]
	pos := 0
	for ; pos < firstDayWeekday; pos++ {
		fmt.Print("\t")
	}
	for i := 1; i <= getDate(d.Year, d.Month); i++ {
		fmt.Printf("%d\t", i)
		pos++
		if pos == 7 {
			pos = 0
			fmt.Println()
		}
	}
	fmt.Println("\n")
}

func (d Date) String() string {
	return fmt.Sprintf("%d-%d-%d", d.Year, d.Month, d.Day)
}
