package moex_iss

import "net/http"

type Client struct {
  Http *http.Client
}

func NewClient(client *http.Client) *Client {
  return &Client{Http: client}
}

func (moex *Client) GetSecInfo(ticker string) (*MoexSecInfo, error) {
  return nil, nil
}

