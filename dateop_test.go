package dateop_test

import (
	"testing"
	"github.com/g10guang/dateop"
	"time"
	"math"
)

// Use golang build-in time to calculate the day interval.
func getInterval(d, o string) int {
	format := "2006-01-02 15:04:05"
	dt, _ := time.Parse(format, d+" 00:00:00")
	ot, _ := time.Parse(format, o+" 00:00:00")
	// Note: time.Sub function will have a maximum duration neck.
	// If maximum meets, this function will return 106751.
	diff := dt.Sub(ot)
	return int(math.Abs(diff.Hours() / 24))
}

func TestDate_Compare(t *testing.T) {
	suits := []struct {
		d      dateop.Date
		o      dateop.Date
		expect int
	}{
		{d: dateop.Date{1100, 11, 11}, o: dateop.Date{2018, 3, 28}, expect: -1},
		{d: dateop.Date{2018, 2, 28}, o: dateop.Date{1999, 1, 1}, expect: 1},
		{d: dateop.Date{2018, 3, 28}, o: dateop.Date{2018, 3, 28}, expect: 0},
	}
	for _, s := range suits {
		if r := s.d.Compare(s.o); r != s.expect {
			t.Errorf("input: d=%s o=%s expect=%d return=%d", s.d, s.o, s.expect, r)
		}
	}
}

func TestDate_Interval(t *testing.T) {
	suits := []struct {
		d      dateop.Date
		o      dateop.Date
		expect int
	}{
		{d: dateop.Date{1997, 2, 28}, o: dateop.Date{2018, 3, 28}, expect: 7698},
		{d: dateop.Date{2018, 3, 28}, o: dateop.Date{1997, 2, 28}, expect: getInterval("2018-03-28", "1997-02-28")},
		{d: dateop.Date{2016, 5, 4}, o: dateop.Date{1969, 5, 20}, expect: getInterval("2016-05-04", "1969-05-20")},
		{d: dateop.Date{1800, 11, 11}, o: dateop.Date{2018, 3, 28}, expect: getInterval("1800-11-11", "2018-03-28")},
	}
	for _, s := range suits {
		if r := s.d.Interval(s.o); r != s.expect {
			t.Errorf("input: d=%s o=%s expect=%d return=%d", s.d, s.o, s.expect, r)
		}
	}
}

func TestDate_Week(t *testing.T) {
	suits := []struct {
		d      dateop.Date
		expect string
	}{
		{d: dateop.Date{2018, 3, 28}, expect: "Wednesday"},
		{d: dateop.Date{2018, 4, 30}, expect: "Monday"},
		{d: dateop.Date{2018, 3, 1}, expect: "Thursday"},
		{d: dateop.Date{1997, 2, 28}, expect: "Friday"},
		{d: dateop.Date{1972, 3, 28}, expect: "Tuesday"},
		{d: dateop.Date{2030, 3, 28}, expect: "Thursday"},
	}
	for _, s := range suits {
		if r := s.d.Week(); r != s.expect {
			t.Errorf("input: d=%s expect=%s return=%s", s.d, s.expect, r)
		}
	}

}

func TestDate_PrintCalendar(t *testing.T) {
	suits := []struct {
		d dateop.Date
	}{
		{d: dateop.Date{2018, 3, 28}},
		{d: dateop.Date{2018, 4, 30}},
		{d: dateop.Date{2018, 3, 1}},
		{d: dateop.Date{1997, 2, 28}},
		{d: dateop.Date{1996, 2, 29}},
		{d: dateop.Date{2030, 3, 28}},
	}
	for _, s := range suits {
		s.d.PrintCalendar()
	}
}
