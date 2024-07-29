package response

type ResponseStatus struct {
	Status int    `json:"status"`
	Error  string `json:"error, omitempty"`
}

const (
	StatusOK    = 200
	StatusError = 500
)

func Ok() ResponseStatus {
	return ResponseStatus{
		Status: StatusOK,
	}
}

func Error(errorMsg string) ResponseStatus {
	return ResponseStatus{
		Status: StatusError,
		Error:  errorMsg,
	}
}
