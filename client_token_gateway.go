package braintree

import (
	"context"
	"encoding/xml"
)

const clientTokenVersion = 2

type ClientTokenGateway struct {
	*Braintree
}

func (g *ClientTokenGateway) Generate(ctx context.Context) (string, error) {
	return g.generate(ctx, &ClientTokenRequest{
		Version: clientTokenVersion,
	})
}

func (g *ClientTokenGateway) GenerateWithCustomer(ctx context.Context, customerId string) (string, error) {
	return g.generate(ctx, &ClientTokenRequest{
		Version:    clientTokenVersion,
		CustomerID: customerId,
	})
}

func (g *ClientTokenGateway) GenerateWithCustomerMerchantIdentifier(ctx context.Context, customerId, merchantAccountId string) (string, error) {
	return g.generate(ctx, &ClientTokenRequest{
		Version:           clientTokenVersion,
		CustomerID:        customerId,
		MerchantAccountId: merchantAccountId,
	})
}

func (g *ClientTokenGateway) generate(ctx context.Context, req *ClientTokenRequest) (string, error) {
	resp, err := g.execute(ctx, "POST", "client_token", req)
	if err != nil {
		return "", err
	}
	switch resp.StatusCode {
	case 201:
		var b clientToken
		if err := xml.Unmarshal(resp.Body, &b); err != nil {
			return "", err
		}
		return b.ClientToken, nil
	}
	return "", &invalidResponseError{resp}
}
