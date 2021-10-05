package main

import (
	"fmt"
	"time"
)

func main() {
	Sample("Months", Months)
	Sample("Durations", Durations)
	Sample("Parse", Parse)
	Sample("ParseDuration", ParseDuration)
	Sample("ParseInLocation", ParseInLocation)
	Sample("Date", Date)
	Sample("BeforeAfter", BeforeAfter)
	Sample("Tick", Tick)
	Sample("NewTicker", NewTicker)
	Sample("Sleep", Sleep)
	Sample("Format", Format)
	Sample("Add", Add)
	Sample("Sub", Sub)
	Sample("Round", Round)
	Sample("Truncate", Truncate)
}

func Months() {
	// Month constants

	fmt.Printf("%[1]d %[1]v\n", time.January)
	fmt.Printf("%[1]d %[1]v\n", time.February)
	fmt.Printf("%[1]d %[1]v\n", time.March)
	fmt.Printf("%[1]d %[1]v\n", time.April)
	fmt.Printf("%[1]d %[1]v\n", time.May)
	fmt.Printf("%[1]d %[1]v\n", time.June)
	fmt.Printf("%[1]d %[1]v\n", time.July)
	fmt.Printf("%[1]d %[1]v\n", time.August)
	fmt.Printf("%[1]d %[1]v\n", time.September)
	fmt.Printf("%[1]d %[1]v\n", time.October)
	fmt.Printf("%[1]d %[1]v\n", time.November)
	fmt.Printf("%[1]d %[1]v\n", time.December)
}

func Durations() {
	// Duration constants

	fmt.Printf("%[1]d %[1]v\n", time.Nanosecond)
	fmt.Printf("%[1]d %[1]v\n", time.Microsecond)
	fmt.Printf("%[1]d %[1]v\n", time.Millisecond)
	fmt.Printf("%[1]d %[1]v\n", time.Second)
	fmt.Printf("%[1]d %[1]v\n", time.Minute)
	fmt.Printf("%[1]d %[1]v\n", time.Hour)
}

func Parse() {
	// Parses time

	const longForm = "Jan 2, 2006 at 3:04pm (MST)"
	t, _ := time.Parse(longForm, "Feb 3, 2013 at 7:54pm (PST)")
	fmt.Println(t)

	const shortForm = "2006-Jan-02"
	t, _ = time.Parse(shortForm, "2013-Feb-03")
	fmt.Println(t)

	t, _ = time.Parse(time.RFC3339, "2006-01-02T15:04:05Z")
	fmt.Println(t)

	t, _ = time.Parse(time.RFC3339, "2006-01-02T15:04:05+07:00")
	fmt.Println(t)

	_, err := time.Parse(time.RFC3339, time.RFC3339)
	fmt.Println("error", err) // Returns an error as the layout is not a valid time value
}

func ParseDuration() {
	// Parses duration

	hours, _ := time.ParseDuration("10h")
	comp, _ := time.ParseDuration("1h10m10s")
	micro, _ := time.ParseDuration("1µs")

	fmt.Println(hours)
	fmt.Println(comp)
	fmt.Printf("There are %.0f seconds in %v.\n", comp.Seconds(), comp)
	fmt.Printf("There are %d nanoseconds in %v.\n", micro.Nanoseconds(), micro)
}

func ParseInLocation() {
	// Parses time with a location

	loc, _ := time.LoadLocation("Europe/Istanbul")

	const longForm = "Jan 2, 2006 at 3:04pm"
	t, _ := time.ParseInLocation(longForm, "Jul 9, 2012 at 5:02am", loc)
	fmt.Println(t)

	const shortForm = "2006-Jan-02"
	t, _ = time.ParseInLocation(shortForm, "2012-Jul-09", loc)
	fmt.Println(t)
}

func Date() {
	date := time.Date(2000, 2, 1, 12, 30, 0, 0, time.UTC)

	dateYear, dateMonth, dateDay := date.Date()
	zoneName, zoneOffset := date.Zone()
	hour, min, sec := date.Clock()
	year, week := date.ISOWeek()

	fmt.Println("Date", dateYear, dateMonth, dateDay)
	fmt.Println("Location", date.Location())
	fmt.Println("Day", date.Day())
	fmt.Println("Year", date.Year())
	fmt.Println("YearDay", date.YearDay())
	fmt.Println("Weekday", date.Weekday())
	fmt.Println("Month", date.Month())
	fmt.Println("UTC", date.UTC())
	fmt.Println("ISOWeek", year, week)
	fmt.Println("Clock", hour, min, sec)
	fmt.Println("Minute", date.Minute())
	fmt.Println("Second", date.Second())
	fmt.Println("Nanosecond", date.Nanosecond())
	fmt.Println("String", date.String())
	fmt.Println("Zone", zoneName, zoneOffset)
	fmt.Println("Local", date.Local())
	fmt.Println("Unix", date.Unix())
	fmt.Println("UnixNano", date.UnixNano())
}

func BeforeAfter() {
	// Check the time before or after than given time

	year2000 := time.Date(2000, 1, 1, 0, 0, 0, 0, time.UTC)
	year3000 := time.Date(3000, 1, 1, 0, 0, 0, 0, time.UTC)

	isYear2000BeforeYear3000 := year2000.Before(year3000) // True
	isYear3000BeforeYear2000 := year3000.Before(year2000) // False

	fmt.Printf("year2000.Before(year3000) = %v\n", isYear2000BeforeYear3000)
	fmt.Printf("year3000.Before(year2000) = %v\n", isYear3000BeforeYear2000)

	isYear2000AfterYear3000 := year2000.After(year3000) // False
	isYear3000AfterYear2000 := year3000.After(year2000) // True

	fmt.Printf("year2000.After(year3000) = %v\n", isYear2000AfterYear3000)
	fmt.Printf("year3000.After(year2000) = %v\n", isYear3000AfterYear2000)
}

func Tick() {
	// Ticks every given time with channels

	count := 1

	c := time.Tick(2 * time.Second)
	for next := range c {
		if count == 5 {
			c = nil
			break
		}
		fmt.Printf("%v\n", next)
		count++
	}
}

func NewTicker() {
	// Ticks every given time with channels

	ticker := time.NewTicker(time.Second)
	defer ticker.Stop()
	done := make(chan bool)

	go func() {
		time.Sleep(3 * time.Second)
		done <- true
	}()

	for {
		select {
		case <-done:
			return
		case t := <-ticker.C:
			fmt.Println("Current time: ", t)
		}
	}
}

func Sleep() {
	// Sleeps given time

	fmt.Println("Sleep started")

	time.Sleep(2 * time.Second)

	fmt.Println("Sleep finished")
}

func Format() {
	// Formats time with given format

	t, err := time.Parse(time.UnixDate, "Wed Feb 25 11:06:39 PST 2015")
	if err != nil { // Always check errors even if they should not happen.
		panic(err)
	}

	// time.Time's Stringer method is useful without any format.
	fmt.Println("default format:", t)

	// Predefined constants in the package implement common layouts.
	fmt.Println("Unix format:", t.Format(time.UnixDate))
}

func Add() {
	// Adds date or duration to time

	start := time.Date(2009, 1, 1, 0, 0, 0, 0, time.UTC)

	oneDayLater := start.AddDate(0, 0, 1)
	oneMonthLater := start.AddDate(0, 1, 0)
	oneYearLater := start.AddDate(1, 0, 0)

	oneHourLater := start.Add(1 * time.Hour)

	fmt.Printf("oneDayLater: start.AddDate(0, 0, 1) = %v\n", oneDayLater)
	fmt.Printf("oneMonthLater: start.AddDate(0, 1, 0) = %v\n", oneMonthLater)
	fmt.Printf("oneYearLater: start.AddDate(1, 0, 0) = %v\n", oneYearLater)
	fmt.Printf("oneHourLater: start.Add(1 * time.Hour) = %v\n", oneHourLater)
}

func Sub() {
	// Subs two times

	start := time.Date(2000, 1, 1, 0, 0, 0, 0, time.UTC)
	end := time.Date(2000, 1, 1, 12, 0, 0, 0, time.UTC)

	difference := end.Sub(start)

	fmt.Printf("Diff = %v\n", difference)
}

func Round() {
	// Rounds the time with given duration

	t := time.Date(0, 0, 0, 12, 15, 30, 918273645, time.UTC)

	round := []time.Duration{
		time.Nanosecond,
		time.Microsecond,
		time.Millisecond,
		time.Second,
		2 * time.Second,
		time.Minute,
		10 * time.Minute,
		time.Hour,
	}

	for _, d := range round {
		fmt.Printf("t.Round(%6s) = %s\n", d, t.Round(d).Format("15:04:05.999999999"))
	}
}

func Truncate() {
	// Truncates the time with given duration

	t, _ := time.Parse("2006 Jan 02 15:04:05", "2012 Dec 07 12:15:30.918273645")

	trunc := []time.Duration{
		time.Nanosecond,
		time.Microsecond,
		time.Millisecond,
		time.Second,
		2 * time.Second,
		time.Minute,
		10 * time.Minute,
	}

	for _, d := range trunc {
		fmt.Printf("t.Truncate(%5s) = %s\n", d, t.Truncate(d).Format("15:04:05.999999999"))
	}

	midnight := time.Date(t.Year(), t.Month(), t.Day(), 0, 0, 0, 0, time.Local)
	fmt.Println("midnight:", midnight)
}

func Sample(name string, fn func())  {
	fmt.Println(">", name)
	fn()
	fmt.Println()
}