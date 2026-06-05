# Go Conference Booking App

A simple console-based booking application written in Go. This app simulates ticket booking for a conference, validates user input, tracks remaining tickets, and sends a delayed ticket confirmation.

## Project Description

This Booking App is a beginner-friendly Go project that demonstrates how to build a command-line ticket reservation system. It accepts user details, validates the input, tracks available tickets, and uses Go routines to simulate sending a confirmation email asynchronously.

The app is designed to showcase basic Go features, including:
- Structs for data modeling
- Functions for input handling and validation
- Concurrency with goroutines and wait groups
- Modular code organization using packages

## Features

- Console user interface for booking tickets
- Input validation for name, email, and ticket quantity
- Tracks total and remaining tickets
- Stores booking details in memory
- Simulates ticket sending asynchronously using Go routines

## Project Structure

- `main.go` - application entry point, user input handling, booking logic, and ticket sending
- `helper/helper.go` - helper functions for greeting and validating user input
- `go.mod` - Go module definition

## How to Run

1. Open a terminal in the project directory.
2. Run the app with:

```bash
go run main.go
```

3. Follow the prompts to enter:
- First name
- Last name
- Email address
- Number of tickets to purchase

## Validation Rules

- First name and last name must each have at least 2 characters
- Email must contain `@`
- Ticket quantity must be greater than 0 and less than or equal to the remaining tickets

## Notes

- The app starts with `50` tickets available for the conference.
- Bookings are stored in a slice of `UserData` structures.
- Ticket confirmation is printed after a 10-second delay to simulate sending.

## Improvements

Possible enhancements for the future:

- Add persistent storage for bookings
- Support multiple booking sessions in a loop
- Add better email validation
- Implement a web or GUI interface
- Add unit tests for validation and booking logic

## License

This project is open source and available to use as you like.
