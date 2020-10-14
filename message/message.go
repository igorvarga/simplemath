package message

const (
	Add     = "add"
	Subtract    = "subtract"
	Divide    = "divide"
	Multiply     = "multiply"
)

type ResultMessage struct {
	Action string  `json:"action"`
	X      float64 `json:"x"`
	Y      float64 `json:"y"`
	Answer float64 `json:"answer"`
	Cached bool    `json:"cached"`
}
