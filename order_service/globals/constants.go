package globals

type OrderStatus string

const (
	Placed         OrderStatus = "placed"
	Delivered      OrderStatus = "delivered"
	OutForDelivery OrderStatus = "outForDelivery"
	Cancelled      OrderStatus = "cancelled"
	InProcessing   OrderStatus = "inProcessing"
	OnHold         OrderStatus = "onHold"
	Returned       OrderStatus = "returned"
)
