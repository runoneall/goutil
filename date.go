package goutil

import (
	"strconv"
	"strings"
	"time"
)

func isHave(haystack string, needle string) bool {
	contains := strings.Contains(haystack, needle)
	return contains
}

func Date_Now() map[string]int {
	now := time.Now()
	year := now.Year()
	month := int(now.Month())
	week := int(now.Weekday())
	day := now.Day()
	hour := now.Hour()
	minute := now.Minute()
	second := now.Second()
	return map[string]int{
		"year":   year,
		"month":  month,
		"week":   week,
		"day":    day,
		"hour":   hour,
		"minute": minute,
		"second": second,
	}
}

func Date_Format(formatQuery string, useLong bool, useAlias bool) string {
	now := time.Now()
	longYear := strconv.Itoa(now.Year())
	year := longYear[2:]
	month := strconv.Itoa(int(now.Month()))
	week := strconv.Itoa(int(now.Weekday()))
	day := strconv.Itoa(now.Day())
	hour := strconv.Itoa(now.Hour())
	minute := strconv.Itoa(now.Minute())
	second := strconv.Itoa(now.Second())
	if useLong {
		year = longYear
		if len(month) == 1 {
			month = "0" + month
		}
		if len(week) == 1 {
			week = "0" + week
		}
		if len(day) == 1 {
			day = "0" + day
		}
		if len(hour) == 1 {
			hour = "0" + hour
		}
		if len(minute) == 1 {
			minute = "0" + minute
		}
		if len(second) == 1 {
			second = "0" + second
		}
	}
	if useAlias {
		month = now.Month().String()
		week = now.Weekday().String()
	}
	if isHave(formatQuery, "{{year}}") {
		formatQuery = strings.Replace(formatQuery, "{{year}}", year, -1)
	}
	if isHave(formatQuery, "{{month}}") {
		formatQuery = strings.Replace(formatQuery, "{{month}}", month, -1)
	}
	if isHave(formatQuery, "{{week}}") {
		formatQuery = strings.Replace(formatQuery, "{{week}}", week, -1)
	}
	if isHave(formatQuery, "{{day}}") {
		formatQuery = strings.Replace(formatQuery, "{{day}}", day, -1)
	}
	if isHave(formatQuery, "{{hour}}") {
		formatQuery = strings.Replace(formatQuery, "{{hour}}", hour, -1)
	}
	if isHave(formatQuery, "{{minute}}") {
		formatQuery = strings.Replace(formatQuery, "{{minute}}", minute, -1)
	}
	if isHave(formatQuery, "{{second}}") {
		formatQuery = strings.Replace(formatQuery, "{{second}}", second, -1)
	}
	return formatQuery
}
