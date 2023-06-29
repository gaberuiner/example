package methods

import (
	"context"
	loms "route256/checkout/internal/loms_client"

	loms_createOrder "route256/checkout/internal/other_server_connection/createOrder"
	check "route256/checkout/internal/protoc/checkout"
)

func (m Method) PurchaseMethod(client check.CheckoutClient, user uint64) (check.PurchaseResponse, error) {
	purchaseReq := &check.PurchaseRequest{
		User: user,
	}
	Method := New()
	listCartRes, err := Method.ListCartMethod(client, user)
	if err != nil {
		return check.PurchaseResponse{}, err
	}
	var lomsItems []*loms.Items
	for _, item := range listCartRes.Items {
		newitem := &loms.Items{
			Sku:   item.Sku,
			Count: item.Count,
		}
		lomsItems = append(lomsItems, newitem)
	}
	err = loms_createOrder.LomsCreateOrder(user, lomsItems)
	if err != nil {
		return check.PurchaseResponse{}, err
	}

	purchaseRes, err := client.Purchase(context.Background(), purchaseReq)
	if err != nil {
		return check.PurchaseResponse{}, err
	}
	return check.PurchaseResponse{OrderID: purchaseRes.OrderID}, nil
}
