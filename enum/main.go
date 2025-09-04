package main //enumarated types
import "fmt"

type OrderStatus int

const ( // Enumerated type for order status
	Pending OrderStatus = iota
	Processing
	Shipped
	Delivered
)

func ChangeOrderStatus(status OrderStatus) {
	fmt.Println("order status changed to", status)
}
func main() {
	ChangeOrderStatus(Pending)
	ChangeOrderStatus(Processing)
	ChangeOrderStatus(Shipped)
	ChangeOrderStatus(Delivered)
}
