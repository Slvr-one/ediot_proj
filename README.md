# EDIOT!

## An Experation Date I/O Tracker.

* create a Tracker for each new item block that arrives,
* enter an Amount and E.D. 
* Saves previous E.D. - Order Arrival Time(OAT) by Name of Product(NoP).

<!-- * Subscription - including managed DB (Mongo, mysql, Postgres) -->

* Web based version - EDT


## Structure:

src/main.go -


---

## Here's a brief explanation of the code:

We save some previous dates for the items.
Finally, we print out the information for each item, including the previous dates.


reating a Golang application for tracking the expiration dates of supermarket products would involve several key functions. Here's a high-level overview of what those functions might look like:
Product Input Function: This function would be responsible for adding new products to the system. It would take details like the product name, category, date of purchase, and expiration date as input.
go

Copy

func addProduct(name string, category string, purchaseDate time.Time, expiryDate time.Time) error {
    // Implementation goes here
}
Expiration Check Function: This function would check if a product is expired or not. It would take the product ID as input and return whether the product is expired.
go

Copy

func isExpired(productId int) (bool, error) {
    // Implementation goes here
}
Product Retrieval Function: This function would retrieve the details of a product. It would take the product ID as input and return the product details.
go

Copy

func getProduct(productId int) (*Product, error) {
    // Implementation goes here
}
Product Update Function: This function would update the details of a product. It would take the product ID and the new details as input.
go

Copy

func updateProduct(productId int, newDetails *Product) error {
    // Implementation goes here
}
Product Deletion Function: This function would delete a product from the system. It would take the product ID as input.
go

Copy

func deleteProduct(productId int) error {
    // Implementation goes here
}
Notification Function: This function would send notifications to the user when a product is about to expire. It would take the product ID and the number of days before expiration to send the notification as input.
go

Copy

func sendNotification(productId int, daysBeforeExpiry int) error {
    // Implementation goes here
}
Remember, these are just the basic functions. Depending on the specific requirements of your application, you might need to add more functions or modify these. For example, you might want to add functions to handle user accounts, or to categorize products.
