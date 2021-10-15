package utils

import (
	"fmt"
	"time"
)

func GetDatetimeString(countryArea string) string {
	now := time.Now()

	location, err := time.LoadLocation(countryArea)
	if err != nil {
		panic(err)
	}

	result := now.In(location)

	return result.Format("2006-01-02 15:04:05")
}

func GetDateString(countryArea string) string {
	now := time.Now()

	location, err := time.LoadLocation(countryArea)
	if err != nil {
		panic(err)
	}

	result := now.In(location)

	return result.Format("2006-01-02")
}

func GetDatetimeStringFromInt(t int64, countryArea string) string {
	now := time.Unix(t, 0)
	location, err := time.LoadLocation(countryArea)
	if err != nil {
		fmt.Println(err)
	}

	result := now.In(location)

	return result.Format("2006-01-02 15:04:05")
}

func GetUnixTimeToday() float64 {
	dt := time.Now().UTC()
	fmt.Println(dt)
	fmt.Println(dt)
	input := dt.Format("2006-01-02")
	layout := "2006-01-02"
	t, _ := time.Parse(layout, input)
	fmt.Println(t.Unix())
	return float64(t.Unix())
}

func GetUnixTimeNow() int64 {
	dt := time.Now().UTC()
	fmt.Println(dt)
	fmt.Println(dt)
	input := dt.Format("2006-01-02")
	layout := "2006-01-02"
	t, _ := time.Parse(layout, input)
	fmt.Println(t.Unix())
	return t.Unix()
}

func AddUnixTimeNow(seconds int64) int64 {
	t := GetUnixTimeNow()
	return t + seconds
}

func GetUnixTimeFromString(datetime string, countryArea string) int64 {
	location, err := time.LoadLocation(countryArea)
	if err != nil {
		panic(err)
	}

	layout := "2006-01-02 15:04:05"
	t, _ := time.ParseInLocation(layout, datetime, location)
	return int64(t.Unix())
}

func GetUnixTimeBefore(hours int64) int64 {
	today := time.Now()
	result := today.Add(-time.Duration(hours) * time.Hour).Unix()
	return int64(result)
}
