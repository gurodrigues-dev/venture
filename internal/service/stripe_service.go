package service

import (
	"fmt"
	"gin/types"
	"log"
	"time"

	"github.com/stripe/stripe-go/v79"
	"github.com/stripe/stripe-go/v79/customer"
	"github.com/stripe/stripe-go/v79/invoice"
	"github.com/stripe/stripe-go/v79/paymentintent"
	"github.com/stripe/stripe-go/v79/paymentlink"
	"github.com/stripe/stripe-go/v79/paymentmethod"
	"github.com/stripe/stripe-go/v79/price"
	"github.com/stripe/stripe-go/v79/product"
	"github.com/stripe/stripe-go/v79/subscription"
)

func CreatePrice(access types.StripeEnv, productId string, driver types.Driver) (*stripe.Price, error) {

	stripe.Key = access.SecretKey

	params := &stripe.PriceParams{
		Currency: stripe.String(string(stripe.CurrencyBRL)),
		Product:  stripe.String(productId),
		Recurring: &stripe.PriceRecurringParams{
			Interval: stripe.String("month"),
		},
		UnitAmount: stripe.Int64(driver.Amount * 100),
	}

	pr, err := price.New(params)

	if err != nil {
		return nil, err
	}

	return pr, err

}

func CreateProduct(driver types.Driver, client types.Responsible, access types.StripeEnv) (*stripe.Product, error) {

	stripe.Key = access.SecretKey

	params := &stripe.ProductParams{
		Name:        stripe.String(fmt.Sprintf("Assinatura - %s & %s", driver.CNH, client.CPF)),
		Description: stripe.String(fmt.Sprintf("Cobrança Mensal de Assinatura entre Motorista %s & Responsável %s", driver.Name, client.Name)),
	}

	prod, err := product.New(params)
	if err != nil {
		fmt.Println("Erro ao criar produto:", err)
		return nil, err
	}

	return prod, nil

}

func RegisterCardToCustomer(client types.Responsible, access types.StripeEnv) (*stripe.PaymentMethod, error) {

	stripe.Key = access.SecretKey

	params := &stripe.PaymentMethodParams{
		Type: stripe.String(string(stripe.PaymentMethodTypeCard)),
		Card: &stripe.PaymentMethodCardParams{
			Token: stripe.String(client.Token),
		},
	}
	pm, err := paymentmethod.New(params)
	if err != nil {
		fmt.Println("Erro ao criar método de pagamento:", err)
		return nil, err
	}
	return pm, nil

}

func CreateCustomer(client types.Responsible, access types.StripeEnv) (*stripe.Customer, error) {

	stripe.Key = access.SecretKey

	pm, err := RegisterCardToCustomer(client, access)

	if err != nil {
		return nil, err
	}

	params := &stripe.CustomerParams{
		Name:          stripe.String(client.Name),
		Email:         stripe.String(client.Email),
		PaymentMethod: stripe.String(pm.ID),
		InvoiceSettings: &stripe.CustomerInvoiceSettingsParams{
			DefaultPaymentMethod: stripe.String(pm.ID),
		},
	}

	cust, err := customer.New(params)

	if err != nil {
		return nil, err
	}

	return cust, nil

}

func CreateSubscription(customerId, priceId string, driver types.Driver, access types.StripeEnv) (*stripe.Subscription, error) {

	stripe.Key = access.SecretKey

	params := &stripe.SubscriptionParams{
		CancelAt: stripe.Int64(time.Now().AddDate(0, driver.ExpMonthBilling, 0).Unix()),
		Customer: stripe.String(customerId),
		Items: []*stripe.SubscriptionItemsParams{
			{
				Price: stripe.String(priceId),
			},
		},
	}

	subs, err := subscription.New(params)
	if err != nil {
		log.Print(err)
		return nil, err
	}

	return subs, nil

}

func ListSubscriptions(customerId *string, access *types.StripeEnv) ([]types.SubscriptionInfo, error) {

	stripe.Key = access.SecretKey

	params := &stripe.SubscriptionListParams{
		Customer: stripe.String(*customerId),
	}
	params.Filters.AddFilter("limit", "", "10")

	var subscriptions []types.SubscriptionInfo

	i := subscription.List(params)

	for i.Next() {
		s := i.Subscription()
		subscriptions = append(subscriptions, types.SubscriptionInfo{
			ID:     s.ID,
			Status: string(s.Status),
		})
	}

	if err := i.Err(); err != nil {
		return nil, err
	}

	return subscriptions, nil

}

func ListInvoces(subscriptionId *string, access *types.StripeEnv) ([]types.InvoiceInfo, error) {

	stripe.Key = access.SecretKey

	params := &stripe.InvoiceListParams{
		Subscription: stripe.String(*subscriptionId),
	}
	params.Filters.AddFilter("limit", "", "1")

	var invoices []types.InvoiceInfo

	i := invoice.List(params)
	for i.Next() {
		inv := i.Invoice()
		invoices = append(invoices, types.InvoiceInfo{
			ID:              inv.ID,
			Status:          string(inv.Status),
			AmountDue:       inv.AmountDue,
			AmountRemaining: inv.AmountRemaining,
		})
	}

	if err := i.Err(); err != nil {
		return nil, err
	}

	return invoices, nil

}

func CalculateRemainingValueSubscription(invoices []types.InvoiceInfo) *types.InvoiceRemaining {

	invoice := types.InvoiceRemaining{
		InvoiceValue: float64(invoices[0].AmountDue / 100),
		Quantity:     float64(len(invoices)),
	}

	invoice.Remaining = invoice.InvoiceValue * (12 - invoice.Quantity)

	invoice.Fines = invoice.Remaining * 0.40

	return &invoice

}

func CreateLinkPreSigned(priceId string, stripeEnv *types.StripeEnv) (*stripe.PaymentLink, error) {

	stripe.Key = stripeEnv.SecretKey

	params := &stripe.PaymentLinkParams{
		LineItems: []*stripe.PaymentLinkLineItemParams{
			{
				Price:    stripe.String(priceId),
				Quantity: stripe.Int64(1),
			},
		},
	}

	pay, err := paymentlink.New(params)

	if err != nil {
		return nil, err
	}

	return pay, nil

}

func CreateSingleInvoice(amount int64, client types.Responsible, currency, description string, stripeEnv types.StripeEnv) (*stripe.PaymentIntent, error) {

	stripe.Key = stripeEnv.SecretKey

	params := &stripe.PaymentIntentParams{
		Customer:      stripe.String(client.CustomerID),
		Amount:        stripe.Int64(amount * 100),
		Currency:      stripe.String(currency),
		PaymentMethod: stripe.String(client.PaymentMethod),
		OffSession:    stripe.Bool(true),
		Confirm:       stripe.Bool(true),
	}

	paym, err := paymentintent.New(params)

	if err != nil {
		return nil, err
	}

	return paym, err

}
