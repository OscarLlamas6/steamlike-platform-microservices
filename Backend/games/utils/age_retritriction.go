package utils

import (
	"fmt"
	"strconv"
	"strings"
	"time"

	age "github.com/bearbin/go-age"
)

func CalculateAge(birthDate string) int {

	dateComponents := strings.Split(birthDate, "-")

	dayString := dateComponents[2]
	day, err := strconv.Atoi(dayString)
	if err != nil {
		fmt.Println(err)
		return 100
	}

	monthString := dateComponents[1]
	month, err := strconv.Atoi(monthString)
	if err != nil {
		fmt.Println(err)
		return 100
	}

	yearString := dateComponents[0]
	year, err := strconv.Atoi(yearString)
	if err != nil {
		fmt.Println(err)
		return 100
	}

	dob := time.Date(year, time.Month(month), day, 0, 0, 0, 0, time.UTC)

	currentAge := age.Age(dob)

	return currentAge

}

func GameAllowed(age int, restriction string) bool {

	if age < 10 {

		switch restriction {
		case "EC":
			return true
		case "E":
			return true
		case "E10+":
			return false
		case "T":
			return false
		case "M":
			return false
		case "AO":
			return false
		default:
			return false
		}

	} else if age >= 10 && age < 13 {

		switch restriction {
		case "EC":
			return true
		case "E":
			return true
		case "E10+":
			return true
		case "T":
			return false
		case "M":
			return false
		case "AO":
			return false
		default:
			return false
		}

	} else if age >= 13 && age < 17 {

		switch restriction {
		case "EC":
			return true
		case "E":
			return true
		case "E10+":
			return true
		case "T":
			return true
		case "M":
			return false
		case "AO":
			return false
		default:
			return false
		}

	} else if age == 17 {

		switch restriction {
		case "EC":
			return true
		case "E":
			return true
		case "E10+":
			return true
		case "T":
			return true
		case "M":
			return true
		case "AO":
			return false
		default:
			return false
		}

	} else if age >= 18 {

		switch restriction {
		case "EC":
			return true
		case "E":
			return true
		case "E10+":
			return true
		case "T":
			return true
		case "M":
			return true
		case "AO":
			return true
		default:
			return true
		}

	}

	return false
}
