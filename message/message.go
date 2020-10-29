package message

// TODO add error message const
const (
	ActionAdd      = "add"
	ActionSubtract = "subtract"
	ActionDivide   = "divide"
	ActionMultiply = "multiply"
)

type ResultMessage struct {
	Action string  `json:"action"`
	X      float64 `json:"x"`
	Y      float64 `json:"y"`
	Answer float64 `json:"answer"`
	Cached string    `json:"cached,omitempty"`
}

type ErrorMessage struct {
	Code  string `json:"code"`
	Error string `json:"message"`
}
