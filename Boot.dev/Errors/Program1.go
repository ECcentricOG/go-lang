package main

import (
	"errors"
	"fmt"
)

func validateStatus(status string) error {
    if status == "" {
        return errors.New("status cannot be empty")
    }
    if len(status) > 140 {
        return errors.New("status exceeds 140 characters")
    }
    return nil
}

func main() {
    fmt.Println(validateStatus("This is a valid status update that is well within the character limit."))
    fmt.Println(validateStatus(""))
    fmt.Println(validateStatus("This status update is way too long. Lorem ipsum dolor sit amet, consectetur adipiscing elit, sed do eiusmod tempor incididunt ut labore et dolore magna aliqua. Ut enim ad minim veniam, quis nostrud exercitation ullamco."))
}
