package main

import (
	"fmt"
	"testing"
	"time"
)

func TestAddItem(tr Tracker, te testing.B) {

	// Print the tracker's items
	for _, item := range tr.Items {
		fmt.Printf("Name: %s\n", item.Name)
		fmt.Printf("Amount: %d\n", item.Amount)
		fmt.Printf("Expiration Date: %s\n", item.Expiration.Format("2006-01-02"))
		fmt.Printf("Order Arrival Time: %s\n", item.OrderArrival.Format("2006-01-02 15:04:05"))

		fmt.Println("Previous Dates:")
		for _, date := range item.PreviousDates {
			fmt.Println(date.Format("2006-01-02"))
		}

		// Calculate the number of days before expiration
		daysBeforeExpiration := int(item.Expiration.Sub(time.Now()).Hours() / 24)
		fmt.Printf("Days Before Expiration: %d\n", daysBeforeExpiration)


		// ...

		// Notify n days before expiration
		n := 3
		if daysBeforeExpiration <= n {
			fmt.Println("Expiring soon! Notify user.")
			fmt.Println("--------------------")

		}
		for _, item := range tr.Items {
			daysBeforeExpiration := int(item.Expiration.Sub(time.Now()).Hours() / 24)
			if daysBeforeExpiration <= n {
				// Send email notification
				SendEmailNotification(item.Name, daysBeforeExpiration)
			}
		}

		// ...
}
