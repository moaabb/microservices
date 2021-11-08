package server

import (
	"context"

	"github.com/hashicorp/go-hclog"
	protos "github.com/moaabb/microservices/currency/protos/currency"
)

type Currency struct {
	l hclog.Logger
	protos.UnimplementedCurrencyServer
}

func NewCurrency(l hclog.Logger) *Currency {
	return &Currency{l: l}
}

func (c *Currency) GetRate(ctx context.Context, rr *protos.RateRequest) (*protos.RateResponse, error) {
	c.l.Info("Handle Get Rate", "base", rr.GetBase(), "destination", rr.GetDestination())

	return &protos.RateResponse{Rate: 0.5}, nil
}
