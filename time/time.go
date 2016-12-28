package time

import (
	"math"
	"time"
)

// Represents the number of elements in a given period
const (
	secondsPerMinute  = 60
	minutesPerHour    = 60
	hoursPerDay       = 24
	daysPerWeek       = 7
	monthsPerQuarter  = 3
	monthsPerYear     = 12
	yearsPerCenturies = 100
	yearsPerDecade    = 10
	weeksPerLongYear  = 53
	daysInLeapYear    = 366
	daysInNormalYear  = 365
	secondsInWeek     = 691200
	secondsInMonth    = 2678400
)

// Represents the different string formats for dates
const (
	DefaultFormat       = "2006-01-02 15:04:05"
	DateFormat          = "2006-01-02"
	FormattedDateFormat = "Jan 2, 2006"
	TimeFormat          = "15:04:05"
	HourMinuteFormat    = "15:04"
	HourFormat          = "15"
	DayDateTimeFormat   = "Mon, Aug 2, 2006 3:04 PM"
	CookieFormat        = "Monday, 02-Jan-2006 15:04:05 MST"
	RFC822Format        = "Mon, 02 Jan 06 15:04:05 -0700"
	RFC1036Format       = "Mon, 02 Jan 06 15:04:05 -0700"
	RFC2822Format       = "Mon, 02 Jan 2006 15:04:05 -0700"
	RFC3339Format       = "2006-01-02T15:04:05-07:00"
	RSSFormat           = "Mon, 02 Jan 2006 15:04:05 -0700"
)

// Parse returns a pointer to a new time instance from a string
// If the location is invalid, it returns an error instead.
func Parse(layout, value, location string) (time.Time, error) {
	l, _ := time.LoadLocation(location)
	t, err := time.ParseInLocation(layout, value, l)
	if err != nil {
		return t, err
	}

	return t, nil
}

// Today returns a pointer to a new carbon instance for today
// If the location is invalid, it returns an error instead.
func Today() time.Time {
	return time.Now()
}

// Tomorrow returns a pointer to a new carbon instance for tomorrow
// If the location is invalid, it returns an error instead.
func Tomorrow() time.Time {
	return AddDay(Today())
}

// Yesterday returns a pointer to a new carbon instance for yesterday
// If the location is invalid, it returns an error instead.
func Yesterday() time.Time {
	return SubDay(Today())
}

// unixTimeInSeconds represents the number of seconds between Year 1 and 1970
const unixTimeInSeconds = 62135596801

const maxNSecs = 999999999

// NowInLocation returns a new Carbon instance for right now in given location.
// The location is in IANA Time Zone database, such as "America/New_York".
func NowInLocation(loc string) (time.Time, error) {
	l, err := time.LoadLocation(loc)
	if err != nil {
		return nowIn(l), err
	}
	return nowIn(l), nil
}

func nowIn(loc *time.Location) time.Time {
	return time.Now().In(loc)
}

// Quarter gets the current quarter
func Quarter(t time.Time) int {
	month := t.Month()
	switch {
	case month < 4:
		return 1
	case month >= 4 && month < 7:
		return 2
	case month >= 7 && month < 10:
		return 3
	}
	return 4
}

// Age gets the age from the current instance time to now
func Age(t time.Time) int {
	return int(DiffInYears(t, time.Now(), true))
}

// DaysInMonth returns the number of days in the month
func DaysInMonth(t time.Time) int {
	return EndOfMonth(t).Day()
}

// DaysInYear returns the number of days in the year
func DaysInYear(t time.Time) int {
	if IsLeapYear(t) {
		return daysInLeapYear
	}

	return daysInNormalYear
}

// WeekOfMonth returns the week of the month
func WeekOfMonth(t time.Time) int {
	w := math.Ceil(float64(t.Day() / daysPerWeek))
	return int(w + 1)
}

// AddYears adds a year to the current time.
// Positive values travel forward while negative values travel into the past
func AddYears(t time.Time, y int) time.Time {
	return t.AddDate(y, 0, 0)
}

// AddYear adds a year to the current time
func AddYear(t time.Time) time.Time {
	return AddYears(t, 1)
}

// AddQuarters adds quarters to the current time.
// Positive values travel forward while negative values travel into the past
func AddQuarters(t time.Time, q int) time.Time {
	return t.AddDate(0, monthsPerQuarter*q, 0)
}

// AddQuarter adds a quarter to the current time
func AddQuarter(t time.Time) time.Time {
	return AddQuarters(t, 1)
}

// AddCenturies adds centuries to the time.
// Positive values travels forward while negative values travels into the past
func AddCenturies(t time.Time, cent int) time.Time {
	return t.AddDate(yearsPerCenturies*cent, 0, 0)
}

// AddCentury adds a century to the current time
func AddCentury(t time.Time) time.Time {
	return AddCenturies(t, 1)
}

// AddMonths adds months to the current time.
// Positive value travels forward while negative values travels into the past
func AddMonths(t time.Time, m int) time.Time {
	return t.AddDate(0, m, 0)
}

// AddMonth adds a month to the current time
func AddMonth(t time.Time) time.Time {
	return AddMonths(t, 1)
}

// AddSeconds adds seconds to the current time.
// Positive values travels forward while negative values travels into the past.
func AddSeconds(t time.Time, s int) time.Time {
	d := time.Duration(s) * time.Second
	return t.Add(d)
}

// AddSecond adds a second to the time
func AddSecond(t time.Time) time.Time {
	return AddSeconds(t, 1)
}

// AddDays adds a day to the current time.
// Positive value travels forward while negative value travels into the past
func AddDays(t time.Time, d int) time.Time {
	return t.AddDate(0, 0, d)
}

// AddDay adds a day to the current time
func AddDay(t time.Time) time.Time {
	return AddDays(t, 1)
}

// AddWeeks adds a week to the current time.
// Positive value travels forward while negative value travels into the past.
func AddWeeks(t time.Time, w int) time.Time {
	return t.AddDate(0, 0, daysPerWeek*w)
}

// AddWeek adds a week to the current time
func AddWeek(t time.Time) time.Time {
	return AddWeeks(t, 1)
}

// AddHours adds an hour to the current time.
// Positive value travels forward while negative value travels into the past
func AddHours(t time.Time, h int) time.Time {
	d := time.Duration(h) * time.Hour

	return t.Add(d)
}

// AddHour adds an hour to the current time
func AddHour(t time.Time) time.Time {
	return AddHours(t, 1)
}

// AddMonthsNoOverflow adds a month to the current time, not overflowing in case the
// destination month has less days than the current one.
// Positive value travels forward while negative value travels into the past.
func AddMonthsNoOverflow(t time.Time, m int) time.Time {
	addedDate := t.AddDate(0, m, 0)
	if t.Day() != addedDate.Day() {
		return PreviousMonthLastDay(addedDate)
	}

	return addedDate
}

// PreviousMonthLastDay returns the last day of the previous month
func PreviousMonthLastDay(t time.Time) time.Time {
	return t.AddDate(0, 0, -t.Day())
}

// // AddMonthNoOverflow adds a month with no overflow to the current time
// func (c *Carbon) AddMonthNoOverflow() *Carbon {
// 	return c.AddMonthsNoOverflow(1)
// }

// // AddMinutes adds minutes to the current time.
// // Positive value travels forward while negative value travels into the past.
// func (c *Carbon) AddMinutes(m int) *Carbon {
// 	d := time.Duration(m) * time.Minute

// 	return NewCarbon(c.Add(d))
// }

// // AddMinute adds a minute to the current time
// func (c *Carbon) AddMinute() *Carbon {
// 	return c.AddMinutes(1)
// }

// // SubYear removes a year from the current time
// func (c *Carbon) SubYear() *Carbon {
// 	return c.SubYears(1)
// }

// // SubYears removes years from current time
// func (c *Carbon) SubYears(y int) *Carbon {
// 	return c.AddYears(-1 * y)
// }

// // SubQuarter removes a quarter from the current time
// func (c *Carbon) SubQuarter() *Carbon {
// 	return c.SubQuarters(1)
// }

// // SubQuarters removes quarters from current time
// func (c *Carbon) SubQuarters(q int) *Carbon {
// 	return c.AddQuarters(-q)
// }

// // SubCentury removes a century from the current time
// func (c *Carbon) SubCentury() *Carbon {
// 	return c.SubCenturies(1)
// }

// // SubCenturies removes centuries from the current time
// func (c *Carbon) SubCenturies(cent int) *Carbon {
// 	return c.AddCenturies(-cent)
// }

// // SubMonth removes a month from the current time
// func (c *Carbon) SubMonth() *Carbon {
// 	return c.SubMonths(1)
// }

// // SubMonths removes months from the current time
// func (c *Carbon) SubMonths(m int) *Carbon {
// 	return c.AddMonths(-m)
// }

// // SubMonthNoOverflow remove a month with no overflow from the current time
// func (c *Carbon) SubMonthNoOverflow() *Carbon {
// 	return c.SubMonthsNoOverflow(1)
// }

// // SubMonthsNoOverflow removes months with no overflow from the current time
// func (c *Carbon) SubMonthsNoOverflow(m int) *Carbon {
// 	return c.AddMonthsNoOverflow(-m)
// }

// SubDay removes a day from the current instance
func SubDay(t time.Time) time.Time {
	return SubDays(t, 1)
}

// SubDays removes days from the current time
func SubDays(t time.Time, d int) time.Time {
	return AddDays(t, -d)
}

// // SubWeekday removes a weekday from the current time
// func (c *Carbon) SubWeekday() *Carbon {
// 	return c.SubWeekdays(1)
// }

// // SubWeekdays removes a weekday from the current time
// func (c *Carbon) SubWeekdays(wd int) *Carbon {
// 	return c.AddWeekdays(-wd)
// }

// // SubWeek removes a week from the current time
// func (c *Carbon) SubWeek() *Carbon {
// 	return c.SubWeeks(1)
// }

// // SubWeeks removes weeks to the current time
// func (c *Carbon) SubWeeks(w int) *Carbon {
// 	return c.AddWeeks(-w)
// }

// // SubHour removes an hour from the current time
// func (c *Carbon) SubHour() *Carbon {
// 	return c.SubHours(1)
// }

// // SubHours removes hours from the current time
// func (c *Carbon) SubHours(h int) *Carbon {
// 	return c.AddHours(-h)
// }

// // SubMinute removes a minute from the current time
// func (c *Carbon) SubMinute() *Carbon {
// 	return c.SubMinutes(1)
// }

// // SubMinutes removes minutes from the current time
// func (c *Carbon) SubMinutes(m int) *Carbon {
// 	return c.AddMinutes(-m)
// }

// // SubSecond removes a second from the current time
// func (c *Carbon) SubSecond() *Carbon {
// 	return c.SubSeconds(1)
// }

// // SubSeconds removes seconds from the current time
// func (c *Carbon) SubSeconds(s int) *Carbon {
// 	return c.AddSeconds(-s)
// }

// // DateString return the current time in Y-m-d format
// func (c *Carbon) DateString() string {
// 	return c.Format(DateFormat)
// }

// // FormattedDateString returns the current time as a readable date
// func (c *Carbon) FormattedDateString() string {
// 	return c.Format(FormattedDateFormat)
// }

// // TimeString returns the current time in hh:mm:ss format
// func (c *Carbon) TimeString() string {
// 	return c.Format(TimeFormat)
// }

// // DateTimeString returns the current time in Y-m-d hh:mm:ss format
// func (c *Carbon) DateTimeString() string {
// 	return c.Format(DefaultFormat)
// }

// // DayDateTimeString returns the current time with a day, date and time format
// func (c *Carbon) DayDateTimeString() string {
// 	return c.Format(DayDateTimeFormat)
// }

// // AtomString formats the current time to a Atom date format
// func (c *Carbon) AtomString() string {
// 	return c.Format(RFC3339Format)
// }

// // CookieString formats the current time to a Cookie date format
// func (c *Carbon) CookieString() string {
// 	return c.Format(CookieFormat)
// }

// // ISO8601String returns the current time in ISO8601 format
// func (c *Carbon) ISO8601String() string {
// 	return c.Format(RFC3339Format)
// }

// // RFC822String returns the current time in RFC 822 format
// func (c *Carbon) RFC822String() string {
// 	return c.Format(RFC822Format)
// }

// // RFC850String returns the current time in RFC 850 format
// func (c *Carbon) RFC850String() string {
// 	return c.Format(time.RFC850)
// }

// // RFC1036String returns the current time in RFC 1036 format
// func (c *Carbon) RFC1036String() string {
// 	return c.Format(RFC1036Format)
// }

// // RFC1123String returns the current time in RFC 1123 format
// func (c *Carbon) RFC1123String() string {
// 	return c.Format(time.RFC1123Z)
// }

// // RFC2822String returns the current time in RFC 2822 format
// func (c *Carbon) RFC2822String() string {
// 	return c.Format(RFC2822Format)
// }

// // RFC3339String returns the current time in RFC 3339 format
// func (c *Carbon) RFC3339String() string {
// 	return c.Format(RFC3339Format)
// }

// // RSSString returns the current time for RSS format
// func (c *Carbon) RSSString() string {
// 	return c.Format(RSSFormat)
// }

// // W3CString returns the current time for WWW Consortium format
// func (c *Carbon) W3CString() string {
// 	return c.Format(RFC3339Format)
// }

// // IsWeekday determines if the current time is a weekday
// func (c *Carbon) IsWeekday() bool {
// 	return !c.IsWeekend()
// }

// // IsWeekend determines if the current time is a weekend day
// func (c *Carbon) IsWeekend() bool {
// 	d := c.Weekday()
// 	for _, wd := range c.WeekendDays() {
// 		if d == wd {
// 			return true
// 		}
// 	}

// 	return false
// }

// // IsYesterday determines if the current time is yesterday
// func (c *Carbon) IsYesterday() bool {
// 	n := Now().SubDay()

// 	return c.IsSameDay(n)
// }

// // IsToday determines if the current time is today
// func (c *Carbon) IsToday() bool {
// 	return c.IsSameDay(Now())
// }

// // IsTomorrow determines if the current time is tomorrow
// func (c *Carbon) IsTomorrow() bool {
// 	n := Now().AddDay()

// 	return c.IsSameDay(n)
// }

// // IsFuture determines if the current time is in the future, ie. greater (after) than now
// func (c *Carbon) IsFuture() bool {
// 	return c.After(time.Now())
// }

// // IsPast determines if the current time is in the past, ie. less (before) than now
// func (c *Carbon) IsPast() bool {
// 	return c.Before(time.Now())
// }

// IsLeapYear determines if current current time is a leap year
func IsLeapYear(t time.Time) bool {
	y := t.Year()
	if (y%4 == 0 && y%100 != 0) || y%400 == 0 {
		return true
	}

	return false
}

// // IsLongYear determines if the instance is a long year
// func (c *Carbon) IsLongYear() bool {
// 	carb := create(c.Year(), time.December, 31, 0, 0, 0, 0, c.Location())
// 	_, w := carb.WeekOfYear()

// 	return w == weeksPerLongYear
// }

// // IsSameAs compares the formatted values of the two dates.
// // If passed date is nil, compares against today
// func (c *Carbon) IsSameAs(format string, carb *Carbon) bool {
// 	if carb == nil {
// 		return c.Format(DefaultFormat) == Now().Format(DefaultFormat)
// 	}

// 	return c.Format(DefaultFormat) == carb.Format(DefaultFormat)
// }

// // IsCurrentYear determines if the current time is in the current year
// func (c *Carbon) IsCurrentYear() bool {
// 	return c.Year() == Now().Year()
// }

// // IsSameYear checks if the passed in date is in the same year as the current time year.
// // If passed date is nil, compares against today
// func (c *Carbon) IsSameYear(carb *Carbon) bool {
// 	if carb == nil {
// 		return c.Year() == nowIn(c.Location()).Year()
// 	}

// 	return c.Year() == carb.Year()
// }

// // IsCurrentMonth determines if the current time is in the current month
// func (c *Carbon) IsCurrentMonth() bool {
// 	return c.Month() == Now().Month()
// }

// // IsSameMonth checks if the passed in date is in the same month as the current month
// // If passed date is nil, compares against today
// func (c *Carbon) IsSameMonth(carb *Carbon, sameYear bool) bool {
// 	m := nowIn(c.Location()).Month()
// 	if carb != nil {
// 		m = carb.Month()
// 	}
// 	if sameYear {
// 		return c.IsSameYear(carb) && c.Month() == m
// 	}

// 	return c.Month() == m
// }

// // IsSameDay checks if the passed in date is the same day as the current day.
// // If passed date is nil, compares against today
// func (c *Carbon) IsSameDay(carb *Carbon) bool {
// 	n := nowIn(c.Location())
// 	if carb != nil {
// 		n = carb
// 	}

// 	return c.Year() == n.Year() && c.Month() == n.Month() && c.Day() == n.Day()
// }

// // IsSunday checks if this day is a Sunday.
// func (c *Carbon) IsSunday() bool {
// 	return c.Weekday() == time.Sunday
// }

// // IsMonday checks if this day is a Monday.
// func (c *Carbon) IsMonday() bool {
// 	return c.Weekday() == time.Monday
// }

// // IsTuesday checks if this day is a Tuesday.
// func (c *Carbon) IsTuesday() bool {
// 	return c.Weekday() == time.Tuesday
// }

// // IsWednesday checks if this day is a Wednesday.
// func (c *Carbon) IsWednesday() bool {
// 	return c.Weekday() == time.Wednesday
// }

// // IsThursday checks if this day is a Thursday.
// func (c *Carbon) IsThursday() bool {
// 	return c.Weekday() == time.Thursday
// }

// // IsFriday checks if this day is a Friday.
// func (c *Carbon) IsFriday() bool {
// 	return c.Weekday() == time.Friday
// }

// // IsSaturday checks if this day is a Saturday.
// func (c *Carbon) IsSaturday() bool {
// 	return c.Weekday() == time.Saturday
// }

// // IsLastWeek returns true is the date is within last week
// func (c *Carbon) IsLastWeek() bool {
// 	secondsInWeek := float64(secondsInWeek)
// 	difference := time.Now().Sub(c.Time)
// 	if difference.Seconds() > 0 && difference.Seconds() < secondsInWeek {
// 		return true
// 	}

// 	return false
// }

// // IsLastMonth returns true is the date is within last month
// func (c *Carbon) IsLastMonth() bool {
// 	secondsInMonth := float64(secondsInMonth)
// 	difference := time.Now().Sub(c.Time)
// 	if difference.Seconds() > 0 && difference.Seconds() < secondsInMonth {
// 		return true
// 	}

// 	return false
// }

// // Eq determines if the current carbon is equal to another
// func (c *Carbon) Eq(carb *Carbon) bool {
// 	return c.Equal(carb.Time)
// }

// // EqualTo determines if the current carbon is equal to another
// func (c *Carbon) EqualTo(carb *Carbon) bool {
// 	return c.Eq(carb)
// }

// // Ne determines if the current carbon is not equal to another
// func (c *Carbon) Ne(carb *Carbon) bool {
// 	return !c.Eq(carb)
// }

// // NotEqualTo determines if the current carbon is not equal to another
// func (c *Carbon) NotEqualTo(carb *Carbon) bool {
// 	return c.Ne(carb)
// }

// // Gt determines if the current carbon is greater (after) than another
// func (c *Carbon) Gt(carb *Carbon) bool {
// 	return c.After(carb.Time)
// }

// // GreaterThan determines if the current carbon is greater (after) than another
// func (c *Carbon) GreaterThan(carb *Carbon) bool {
// 	return c.Gt(carb)
// }

// // Gte determines if the instance is greater (after) than or equal to another
// func (c *Carbon) Gte(carb *Carbon) bool {
// 	return c.Gt(carb) || c.Eq(carb)
// }

// // GreaterThanOrEqualTo determines if the instance is greater (after) than or equal to another
// func (c *Carbon) GreaterThanOrEqualTo(carb *Carbon) bool {
// 	return c.Gte(carb) || c.Eq(carb)
// }

// // Lt determines if the instance is less (before) than another
// func (c *Carbon) Lt(carb *Carbon) bool {
// 	return c.Before(carb.Time)
// }

// // LessThan determines if the instance is less (before) than another
// func (c *Carbon) LessThan(carb *Carbon) bool {
// 	return c.Lt(carb)
// }

// // Lte determines if the instance is less (before) or equal to another
// func (c *Carbon) Lte(carb *Carbon) bool {
// 	return c.Lt(carb) || c.Eq(carb)
// }

// // LessThanOrEqualTo determines if the instance is less (before) or equal to another
// func (c *Carbon) LessThanOrEqualTo(carb *Carbon) bool {
// 	return c.Lte(carb)
// }

// // Between determines if the current instance is between two others
// // eq Indicates if a > and < comparison should be used or <= or >=
// func (c *Carbon) Between(a, b *Carbon, eq bool) bool {
// 	if a.Gt(b) {
// 		a, b = swap(a, b)
// 	}
// 	if eq {
// 		return c.Gte(a) && c.Lte(b)
// 	}

// 	return c.Gt(a) && c.Lt(b)
// }

// // Closest returns the closest date from the current time
// func (c *Carbon) Closest(a, b *Carbon) *Carbon {
// 	if c.DiffInSeconds(a, true) < c.DiffInSeconds(b, true) {
// 		return a
// 	}

// 	return b
// }

// // Farthest returns the farthest date from the current time
// func (c *Carbon) Farthest(a, b *Carbon) *Carbon {
// 	if c.DiffInSeconds(a, true) > c.DiffInSeconds(b, true) {
// 		return a
// 	}

// 	return b
// }

// // Min returns the minimum instance between a given instance and the current instance
// func (c *Carbon) Min(carb *Carbon) *Carbon {
// 	if carb == nil {
// 		carb = nowIn(c.Location())
// 	}
// 	if c.Lt(carb) {
// 		return c
// 	}

// 	return carb
// }

// // Minimum returns the minimum instance between a given instance and the current instance
// func (c *Carbon) Minimum(carb *Carbon) *Carbon {
// 	return c.Min(carb)
// }

// // Max returns the maximum instance between a given instance and the current instance
// func (c *Carbon) Max(carb *Carbon) *Carbon {
// 	if carb == nil {
// 		carb = nowIn(c.Location())
// 	}

// 	if c.Gt(carb) {
// 		return c
// 	}

// 	return carb
// }

// // Maximum returns the maximum instance between a given instance and the current instance
// func (c *Carbon) Maximum(carb *Carbon) *Carbon {
// 	return c.Max(carb)
// }

// DiffInYears returns the difference in years
func DiffInYears(t time.Time, ti time.Time, abs bool) int64 {
	if t.Year() == ti.Year() {
		return 0
	}

	diffHr := DiffInHours(t, ti, abs)
	hrLastYear := int64(DaysInYear(t) * hoursPerDay)

	if (diffHr - hrLastYear) >= 0 {
		diff := int64(ti.In(time.UTC).Year() - t.In(time.UTC).Year())

		return absValue(abs, diff)
	}

	return 0
}

// // DiffInMonths returns the difference in months
// func (c *Carbon) DiffInMonths(carb *Carbon, abs bool) int64 {
// 	if carb == nil {
// 		carb = nowIn(c.Location())
// 	}

// 	if c.Month() == carb.Month() && c.Year() == carb.Year() {
// 		return 0
// 	}

// 	diffHr := c.DiffInHours(carb, abs)
// 	hrLastMonth := int64(c.DaysInMonth() * hoursPerDay)

// 	if (diffHr - hrLastMonth) >= 0 {
// 		var m int64
// 		if c.Year() < carb.Year() {
// 			m = (int64(monthsPerYear) - int64(c.In(time.UTC).Month())) + (int64(carb.In(time.UTC).Month()) - 1)
// 			totalHr := int64(c.DaysInMonth() * hoursPerDay)
// 			cHr := c.StartOfMonth().DiffInHours(c, abs)
// 			remainHr := totalHr - cHr
// 			spentInHr := carb.StartOfMonth().DiffInHours(carb, abs)
// 			if (remainHr + spentInHr) >= totalHr {
// 				m = m + 1
// 			}
// 		} else if c.Year() > carb.Year() {
// 			m = (int64(monthsPerYear) - int64(carb.In(time.UTC).Month())) + (int64(c.In(time.UTC).Month()) - 1)
// 			totalHr := int64(carb.DaysInMonth() * hoursPerDay)
// 			carbHr := carb.StartOfMonth().DiffInHours(carb, abs)
// 			remainHr := totalHr - carbHr
// 			spentInHr := c.StartOfMonth().DiffInHours(c, abs)
// 			if (remainHr + spentInHr) >= totalHr {
// 				m = m + 1
// 			}
// 		} else {
// 			m = int64(carb.In(time.UTC).Month() - c.In(time.UTC).Month())
// 		}

// 		diffYr := c.Year() - carb.Year()
// 		if diffYr > 1 {
// 			diff := c.DiffInYears(carb, abs)*monthsPerYear + m

// 			return absValue(abs, diff)
// 		}

// 		diff := m

// 		return absValue(abs, diff)
// 	}

// 	return 0
// }

// // DiffDurationInString returns the duration difference in string format
// func (c *Carbon) DiffDurationInString(carb *Carbon) string {
// 	if carb == nil {
// 		carb = nowIn(c.Location())
// 	}

// 	return strings.Replace(carb.Sub(c.Time).String(), "-", "", 1)
// }

// // DiffInWeeks returns the difference in weeks
// func (c *Carbon) DiffInWeeks(carb *Carbon, abs bool) int64 {
// 	if carb == nil {
// 		carb = nowIn(c.Location())
// 	}
// 	return c.DiffInDays(carb, abs) / daysPerWeek
// }

// // DiffInDays returns the difference in days
// func (c *Carbon) DiffInDays(carb *Carbon, abs bool) int64 {
// 	if carb == nil {
// 		carb = nowIn(c.Location())
// 	}
// 	return c.DiffInHours(carb, abs) / hoursPerDay
// }

// // DiffInNights returns the difference in nights
// func (c *Carbon) DiffInNights(carb *Carbon, abs bool) int64 {
// 	if carb == nil {
// 		carb = nowIn(c.Location())
// 	}
// 	return c.DiffInDays(carb, abs)
// }

// // Filter represents a predicate used for filtering diffs
// type Filter func(*Carbon) bool

// // dayDuration reprensets a day in time.Duration format
// const dayDuration = time.Hour * hoursPerDay

// // DiffInDaysFiltered returns the difference in days using a filter
// func (c *Carbon) DiffInDaysFiltered(f Filter, carb *Carbon, abs bool) int64 {
// 	return c.DiffFiltered(dayDuration, f, carb, abs)
// }

// // DiffInHoursFiltered returns the difference in hours using a filter
// func (c *Carbon) DiffInHoursFiltered(f Filter, carb *Carbon, abs bool) int64 {
// 	return c.DiffFiltered(time.Hour, f, carb, abs)
// }

// // DiffInWeekdays returns the difference in weekdays
// func (c *Carbon) DiffInWeekdays(carb *Carbon, abs bool) int64 {
// 	f := func(t *Carbon) bool {
// 		return t.IsWeekday()
// 	}

// 	return c.DiffFiltered(dayDuration, f, carb, abs)
// }

// // DiffInWeekendDays returns the difference in weekend days using a filter
// func (c *Carbon) DiffInWeekendDays(carb *Carbon, abs bool) int64 {
// 	f := func(t *Carbon) bool {
// 		return t.IsWeekend()
// 	}

// 	return c.DiffFiltered(dayDuration, f, carb, abs)
// }

// // DiffFiltered returns the difference by the given duration using a filter
// func (c *Carbon) DiffFiltered(duration time.Duration, f Filter, carb *Carbon, abs bool) int64 {
// 	if carb == nil {
// 		carb = nowIn(c.Location())
// 	}
// 	if c.IsSameDay(carb) {
// 		return 0
// 	}

// 	inverse := false
// 	var counter int64
// 	s := int64(duration.Seconds())
// 	start, end := c.Copy(), carb.Copy()
// 	if start.Gt(end) {
// 		start, end = swap(start, end)
// 		inverse = true
// 	}
// 	for start.DiffInSeconds(end, true)/s > 0 {
// 		if f(end) {
// 			counter++
// 		}
// 		end = NewCarbon(end.Add(-duration))
// 	}
// 	if inverse {
// 		counter = -counter
// 	}

// 	return absValue(abs, counter)
// }

// DiffInHours returns the difference in hours
func DiffInHours(t time.Time, d time.Time, abs bool) int64 {
	return DiffInMinutes(t, d, abs) / minutesPerHour
}

// DiffInMinutes returns the difference in minutes
func DiffInMinutes(t time.Time, d time.Time, abs bool) int64 {
	return DiffInSeconds(t, d, abs) / secondsPerMinute
}

// DiffInSeconds returns the difference in seconds
func DiffInSeconds(t time.Time, d time.Time, abs bool) int64 {
	diff := d.Unix() - t.Unix()
	return absValue(abs, diff)
}

// // SecondsSinceMidnight returns the number of seconds since midnight.
// func (c *Carbon) SecondsSinceMidnight() int {
// 	startOfDay := c.StartOfDay()

// 	return int(c.DiffInSeconds(startOfDay, true))
// }

// // SecondsUntilEndOfDay returns the number of seconds until 23:59:59.
// func (c *Carbon) SecondsUntilEndOfDay() int {
// 	dayEnd := c.EndOfDay()

// 	return int(c.DiffInSeconds(dayEnd, true))
// }

// absValue returns the abs value if needed
func absValue(needsAbs bool, value int64) int64 {
	if needsAbs && value < 0 {
		return -value
	}

	return value
}

// func swap(a, b *Carbon) (*Carbon, *Carbon) {
// 	return b, a
// }

// // DiffForHumans returns the difference in a human readable format in the current locale.
// // When comparing a value in the past to default now:
// // 1 hour ago
// // 5 months ago
// // When comparing a value in the future to default now:
// // 1 hour from now
// // 5 months from now
// // When comparing a value in the past to another value:
// // 1 hour before
// // 5 months before
// // When comparing a value in the future to another value:
// // 1 hour after
// // 5 months after
// func (c *Carbon) DiffForHumans(d *Carbon, abs, absolute, short bool) (string, error) {
// 	isNow := (d == nil)
// 	if isNow {
// 		d = Now()
// 	}

// 	var unit string
// 	var count int64

// 	switch true {
// 	case c.DiffInYears(d, abs) > 0:
// 		if short {
// 			unit = "y"
// 		} else {
// 			unit = "year"
// 		}
// 		count = c.DiffInYears(d, abs)
// 		break

// 	case c.DiffInMonths(d, abs) > 0:
// 		if short {
// 			unit = "m"
// 		} else {
// 			unit = "month"
// 		}
// 		count = c.DiffInMonths(d, abs)
// 		break

// 	case c.DiffInDays(d, abs) > 0:
// 		if short {
// 			unit = "d"
// 		} else {
// 			unit = "day"
// 		}
// 		count = c.DiffInDays(d, abs)
// 		if count >= daysPerWeek {
// 			if short {
// 				unit = "w"
// 			} else {
// 				unit = "week"
// 			}
// 			count = int64(count / daysPerWeek)
// 		}
// 		break

// 	case c.DiffInHours(d, abs) > 0:
// 		if short {
// 			unit = "h"
// 		} else {
// 			unit = "hour"
// 		}
// 		count = c.DiffInHours(d, abs)
// 		break

// 	case c.DiffInMinutes(d, abs) > 0:
// 		if short {
// 			unit = "min"
// 		} else {
// 			unit = "minute"
// 		}
// 		count = c.DiffInMinutes(d, abs)
// 		break

// 	default:
// 		if short {
// 			unit = "s"
// 		} else {
// 			unit = "second"
// 		}
// 		count = c.DiffInSeconds(d, abs)
// 		break
// 	}

// 	if count == 0 {
// 		count = 1
// 	}

// 	t, err := c.Translator.chooseUnit(unit, count)
// 	if err != nil {
// 		return "", err
// 	}

// 	if absolute {
// 		return t, nil
// 	}

// 	isFuture := c.GreaterThan(d)

// 	var transID string
// 	if isNow {
// 		if isFuture {
// 			transID = "from_now"
// 		} else {
// 			transID = "ago"
// 		}
// 	} else {
// 		if isFuture {
// 			transID = "after"
// 		} else {
// 			transID = "before"
// 		}
// 	}

// 	/* TODO
// 	// Some langs have special pluralization for past and future tense.
// 	// tryKeyExists := unit + "_" + transID
// 	if tryKeyExists != c.Translator.Choose(tryKeyExists, count) {
// 		time, _ = c.Translator.Choose(tryKeyExists, count)
// 	}
// 	*/

// 	return c.Translator.chooseTrans(transID, t), nil
// }

// // StartOfDay returns the time at 00:00:00 of the same day
// func (c *Carbon) StartOfDay() *Carbon {
// 	return create(c.Year(), c.Month(), c.Day(), 0, 0, 0, 0, c.Location())
// }

// // EndOfDay returns the time at 23:59:59 of the same day
// func (c *Carbon) EndOfDay() *Carbon {
// 	return create(c.Year(), c.Month(), c.Day(), 23, 59, 59, maxNSecs, c.Location())
// }

// // StartOfMonth returns the date on the first day of the month and the time to 00:00:00
// func (c *Carbon) StartOfMonth() *Carbon {
// 	return create(c.Year(), c.Month(), 1, 0, 0, 0, 0, c.Location())
// }

// // EndOfMonth returns the date at the end of the month and time at 23:59:59
func EndOfMonth(t time.Time) time.Time {
	return time.Date(t.Year(), t.Month()+1, 0, 23, 59, 59, maxNSecs, t.Location())
}

// // StartOfQuarter returns the date at the first day of the quarter and time at 00:00:00
// func (c *Carbon) StartOfQuarter() *Carbon {
// 	month := time.Month((c.Quarter()-1)*monthsPerQuarter + 1)

// 	return create(c.Year(), time.Month(month), 1, 0, 0, 0, 0, c.Location())
// }

// // EndOfQuarter returns the date at end of the quarter and time at 23:59:59
// func (c *Carbon) EndOfQuarter() *Carbon {
// 	return c.StartOfQuarter().AddMonths(monthsPerQuarter - 1).EndOfMonth()
// }

// // StartOfYear returns the date at the first day of the year and the time at 00:00:00
// func (c *Carbon) StartOfYear() *Carbon {
// 	return create(c.Year(), time.January, 1, 0, 0, 0, 0, c.Location())
// }

// // EndOfYear returns the date at end of the year and time to 23:59:59
// func (c *Carbon) EndOfYear() *Carbon {
// 	return create(c.Year(), time.December, 31, 23, 59, 59, maxNSecs, c.Location())
// }

// // StartOfDecade returns the date at the first day of the decade and time at 00:00:00
// func (c *Carbon) StartOfDecade() *Carbon {
// 	year := c.Year() - c.Year()%yearsPerDecade

// 	return create(year, time.January, 1, 0, 0, 0, 0, c.Location())
// }

// // EndOfDecade returns the date at the end of the decade and time at 23:59:59
// func (c *Carbon) EndOfDecade() *Carbon {
// 	year := c.Year() - c.Year()%yearsPerDecade + yearsPerDecade - 1

// 	return create(year, time.December, 31, 23, 59, 59, maxNSecs, c.Location())
// }

// // StartOfCentury returns the date of the first day of the century at 00:00:00
// func (c *Carbon) StartOfCentury() *Carbon {
// 	year := c.Year() - c.Year()%yearsPerCenturies

// 	return create(year, time.January, 1, 0, 0, 0, 0, c.Location())
// }

// // EndOfCentury returns the date of the end of the century at 23:59:59
// func (c *Carbon) EndOfCentury() *Carbon {
// 	year := c.Year() - 1 - c.Year()%yearsPerCenturies + yearsPerCenturies

// 	return create(year, time.December, 31, 23, 59, 59, maxNSecs, c.Location())
// }

// // StartOfWeek returns the date of the first day of week at 00:00:00
// func (c *Carbon) StartOfWeek(t time.Time) time.Time {
// 	return Previous(t, WeekStartsAt(t))
// }

// // EndOfWeek returns the date of the last day of the week at 23:59:59
// func EndOfWeek(t time.Time) time.Time {
// 	return EndOfDay(Next(t, WeekEndsAt(t)))
// }

// // Next changes the time to the next occurrence of a given day of the week
// func Next(t time.Time, wd time.Weekday) time.Time {
// 	t = AddDay(t)
// 	for t.Weekday() != wd {
// 		t = AddDay(t)
// 	}

// 	return StartOfDay(t)
// }

// // NextWeekday goes forward to the next weekday
// func NextWeekday(t time.Time) time.Time {
// 	return AddWeekday(t)
// }

// // PreviousWeekday goes back to the previous weekday
// func PreviousWeekday(t time.Time) time.Time {
// 	return SubWeekday(t)
// }

// // NextWeekendDay goes forward to the next weekend day
// func NextWeekendDay(t time.Time) time.Time {
// 	t = AddDay(t)
// 	for !IsWeekend(t) {
// 		t = AddDay(t)
// 	}

// 	return c
// }

// // PreviousWeekendDay goes back to the previous weekend day
// func PreviousWeekendDay(t time.Time) time.Time {
// 	t = SubDay(t)
// 	for !IsWeekend(t) {
// 		t = SubDay(t)
// 	}

// 	return t
// }

// // Previous changes the time to the previous occurrence of a given day of the week
// func Previous(t time.Time, wd time.Weekday) time.Time {
// 	t = SubDay(t)
// 	for t.Weekday() != wd {
// 		t = SubDay(t)
// 	}

// 	return StartOfDay(t)
// }
