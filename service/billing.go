package service

import (
	"context"
	"errors"
	"log"
	"strateegy/billing-service/controller/dto"
	"strateegy/billing-service/grpc"
)

type BillingService struct {
}

func NewBillingService() *BillingService {
	return &BillingService{}
}

func (s *BillingService) ChangePlan(token string, data dto.Billing) error {
	connJwt := grpc.GetConnJwt()
	connBilling := grpc.GetConnBilling()

	clientJwt := grpc.NewSendIDClient(connJwt)
	clientUser := grpc.NewChangePlanClient(connBilling)

	tokenRequest := &grpc.Token{
		Token: token,
	}

	res, err := clientJwt.RequestID(context.Background(), tokenRequest)
	if err != nil {
		return err
	}

	ID := res.GetID()

	billingRequest := &grpc.Billing{
		ID:   ID,
		Plan: "",
	}

	if data.Value < 10 {
		return errors.New("choose a valid plan")
	} else if data.Value == 10 {
		billingRequest.Plan = "BASIC"
		clientUser.RequestBilling(context.Background(), billingRequest)
	} else if data.Value == 20 {
		billingRequest.Plan = "PRO"
		req, err := clientUser.RequestBilling(context.Background(), billingRequest)
		if err != nil {
			log.Println(err)
			return err
		}
		log.Println(req.GetStatus())
	} else if data.Value == 30 {
		billingRequest.Plan = "TEAM"
		clientUser.RequestBilling(context.Background(), billingRequest)
	} else {
		return errors.New("choose a valid plan")
	}

	return nil

}
