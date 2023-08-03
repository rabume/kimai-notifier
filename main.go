package main

import (
	"encoding/json"
	"fmt"
	"log"
	"math"
	"os"
	"os/exec"
	"runtime"
	"time"

	"github.com/go-resty/resty/v2"
)

type Config struct {
	ApiEndpoint string  `json:"apiEndpoint"`
	Username    string  `json:"username"`
	Token       string  `json:"token"`
	Interval    float64 `json:"interval"`
	Threshold   float64 `json:"threshold"`
}

var config Config
var InformedAboutLunch = false

func main() {

	// Init config variables
	file, err := os.ReadFile("./config.json")
	if err != nil {
		log.Fatal("Error reading config file:", err)
	}

	err = json.Unmarshal(file, &config)
	if err != nil {
		log.Fatal("Error parsing config file:", err)
	}

	fmt.Println("Welcome to work, " + config.Username + " 👋")
	fmt.Println("-------------------------------------------")
	fmt.Println("")

	for {
		totalHours, err := getTotalHours()
		if err != nil {
			fmt.Println("Error fetching total hours:", err)
			continue
		}

		// Delete last line before printing new total hours
		fmt.Print("\033[1A")
		fmt.Printf("Total hours worked until now: %.2f\n", totalHours)

		// Check if it is between 12:00 PM and 1:00 PM and notify user to take a break
		// TODO: Only inform user if set in config -> add config option
		now := time.Now()
		if now.Hour() == 12 && now.Minute() >= 0 && now.Minute() <= 59 && !InformedAboutLunch {
			InformedAboutLunch = true
			notifyUser("It's lunch time! \nTake a break. 😋")
		}

		if totalHours > config.Threshold {
			notifyUser(fmt.Sprintf("You have worked %.2f hours today. \nTake a break! ⏰", totalHours))
			break
		}

		time.Sleep(time.Duration(config.Interval) * time.Minute)
	}
}

func getTotalHours() (float64, error) {
	now := time.Now()
	begin := time.Date(now.Year(), now.Month(), now.Day(), 0, 0, 0, 0, time.Local)
	end := now

	client := resty.New()
	resp, err := client.R().
		SetHeader("accept", "application/json").
		SetHeader("X-AUTH-USER", config.Username).
		SetHeader("X-AUTH-TOKEN", config.Token).
		Get(config.ApiEndpoint + fmt.Sprintf("?project=&begin=%s", begin.Format("2006-01-02T15:04:05")))

	if err != nil {
		return 0, err
	}

	var response []map[string]interface{}
	json.Unmarshal(resp.Body(), &response)

	// Filter timesheets within the current day and calculate total hours
	var totalHours float64
	for _, timesheet := range response {

		beginTime, err := time.Parse("2006-01-02T15:04:05-0700", timesheet["begin"].(string))
		if err != nil {
			return 0, err
		}

		if beginTime.Before(end) {
			duration := timesheet["duration"].(float64)
			duration = math.Round(duration)
			totalHours += duration
		}
	}

  // Get the active timesheet and calculate the duration
	resp, err = client.R().
		SetHeader("accept", "application/json").
		SetHeader("X-AUTH-USER", config.Username).
		SetHeader("X-AUTH-TOKEN", config.Token).
		Get(config.ApiEndpoint + "/active")

	if err != nil {
		return 0, err
	}

	var activeTimesheet []map[string]interface{}
	json.Unmarshal(resp.Body(), &activeTimesheet)

  if len(activeTimesheet) > 0 && activeTimesheet[0]["begin"] != nil {
		beginTime, err := time.Parse("2006-01-02T15:04:05-0700", activeTimesheet[0]["begin"].(string))
		if err != nil {
			return 0, err
		}

		duration := time.Since(beginTime).Seconds()
		duration = math.Round(duration)
		totalHours += duration
	}

	totalMinutes := math.Floor(totalHours / 60)
	hours := math.Floor(totalMinutes / 60)
	minutes := totalMinutes - (hours * 60)

	totalHours = hours + (minutes / 100)
	return totalHours, nil
}

func notifyUser(message string) {
	if runtime.GOOS == "windows" {
		fmt.Println("Can't Execute this on a windows machine")
	} else {
		cmd := exec.Command("notify-send", message)
		err := cmd.Run()
		if err != nil {
			fmt.Println("Error:", err)
		}
	}
}
