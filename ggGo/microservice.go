package main

import (
	"context"
	"net/http"
)

type Service interface {
	getCatFact(context.Context) (*CatFact, error)
}

type CatFactService struct {
	url string
}

func NewCatFactService(url string) Service {
	return &CatFactService{
		url: url,
	}
}

func (s *CatFactService) getCatFact(ctx context.Context) (*CatFact, error) {
	resp, err := http.Get(s.url)
	if err != nil {
		return nil, err
	}
}
