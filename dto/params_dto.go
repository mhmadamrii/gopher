package dto

type Response struct {
	StatusCode int
	Message    string
	Paginate   *Paginate
	Data       any
}
