package main

type Status int

const (
	StatusSuccess Status = 200
	StatusCreated Status = 201

	StatusInternalServerError Status = 500

    StatusBadRequest    Status = 400
    StatusNotFound      Status = 404
	StatusConflict      Status = 409
)

type Response struct {
    Status Status `json:"status"`
    Value interface{} `json:"value"`
}

type ClientResponse struct {
    Response
    Client *Client
}

