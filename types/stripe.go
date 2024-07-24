package types

type StripeEnv struct {
	PublicKey string
	SecretKey string
}

type SubscriptionInfo struct {
	ID     string
	Status string
}

type InvoiceInfo struct {
	ID              string
	Status          string
	AmountDue       int64
	AmountRemaining int64
}

type InvoiceRemaining struct {
	InvoiceValue float64
	Quantity     float64
	Remaining    float64
	Fines        float64
}
