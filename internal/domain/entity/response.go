package entity

type Response struct {
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

type ResponsePaginate struct {
	Message    string      `json:"message"`
	Pagination interface{} `json:"pagination"`
}
