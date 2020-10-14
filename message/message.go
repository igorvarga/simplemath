package message

// TODO add error message const
const (
	Add      = "add"
	Subtract = "subtract"
	Divide   = "divide"
	Multiply = "multiply"
)

type ResultMessage struct {
	Action string  `json:"action"`
	X      float64 `json:"x"`
	Y      float64 `json:"y"`
	Answer float64 `json:"answer"`
	Cached bool    `json:"cached"`
}

type ErrorMessage struct {
	Code  string `json:"code"`
	Error string `json:"message"`
}
