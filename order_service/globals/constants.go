package globals

type OrderStatus string

const (
	Placed         OrderStatus = "placed"
	Delivered      OrderStatus = "delivered"
	OutForDelivery OrderStatus = "outForDelivery"
	Cancelled      OrderStatus = "cancelled"
	Active         OrderStatus = "active"
	InProcessing   OrderStatus = "inProcessing"
	OnHold         OrderStatus = "onHold"
	Returned       OrderStatus = "returned"
)
