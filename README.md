<h1 align="center">🕒 Kimai Notifier</h1>

<p align="center">
  <img src="https://img.shields.io/badge/Language-Go-blue.svg" alt="Language: Go">
  <img src="https://img.shields.io/badge/License-MIT-yellow.svg" alt="License: MIT">
</p>

<p align="center">
  <strong>🚀 A simple Go application to track time and receive notifications when you exceed the daily working hours threshold in the Kimai Time Tracking System.</strong>
</p>

## Idea

I had the idea to create this app because I always forget the time. I wanted to create a simple app that would notify me when I exceed the daily working hours threshold. I use the Kimai Time Tracking System to track my time, so I decided to use the Kimai API to fetch the total hours worked for the current day.

## Overview

The Kimai Notifier is a lightweight Go application designed to interact with the Kimai Time Tracking System's API. It fetches the total hours worked for the current day from the Kimai API at a predefined interval (5 minutes by default). If the total hours exceed a certain threshold (e.g., 8 hours), it sends a notification (user-defined) to alert the user to take a break.

## Dependencies

The following dependencies are required to build and run the Kimai Notifier:

- Go (Version 1.14 or later)
- GitHub.com/go-resty/resty/v2]
- Linux (Tested on Arch Linux and Pop!\_OS)
- notify-send (Tested with version 0.8.2)

## Installation

1. Install Go: Make sure you have Go installed on your system. You can download and install it from the official website: https://golang.org/

2. Install the Kimai Notifier: Clone the repository and build the executable. The executable will be created in the current directory.

```bash
git clone https://github.com/rabume/kimai-notifier.git
cd kimai-notifier
go get github.com/go-resty/resty/v2
go build
```

## Configuration

The `config.json` file contains the configuration for the Kimai Notifier. The following parameters can be configured:

- **apiEndpoint**: The URL of the Kimai API.
- **username**: The username of the Kimai user.
- **token**: The API token of the Kimai user.
- **interval**: The interval (in minutes) at which the Kimai Notifier will fetch the total hours worked for the current day from the Kimai API.
- **threshold**: The threshold (in hours) at which the Kimai Notifier will send a notification to alert the user to take a break.
