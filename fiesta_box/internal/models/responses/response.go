package responses

type StatusCode int

const (
	Success StatusCode = 200
	Processing StatusCode = 201
	InvalidMessage StatusCode = 400
	UnknownMessageType StatusCode = 404
	Error   StatusCode = 500
)

func (s StatusCode) String() string {
	switch s {
	case Success:
		return "Success"
	case Processing:
		return "Processing"
	case InvalidMessage:
		return "InvalidMessage"
	case UnknownMessageType:
		return "UnknownMessageType"
	case Error:
		return "Error"
	default:
		return "Unknown"
	}
}

type SocketResponse struct {
	Status StatusCode `json:"status"`
	Message string `json:"message"`
	Content interface{} `json:"content"`
}