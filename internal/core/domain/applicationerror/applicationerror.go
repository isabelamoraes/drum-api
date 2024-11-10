package applicationerror

type ErrorType string

const (
	NotFoundApplication            ErrorType = "notfound"
	UnprocessableEntityApplication ErrorType = "unprocessableentity"
	UnathorizedApplication         ErrorType = "unathorized"
	ForbiddenApplication           ErrorType = "forbidden"
	InternalServerError            ErrorType = "internalservererror"
	BadRequestApplication          ErrorType = "badrequest"
)

type Error struct {
	ErrorType ErrorType
	Message   string
}

func (e Error) Error() string {
	return e.Message
}

func MakeNotFoundApplicationError(entity string) Error {
	return Error{
		ErrorType: NotFoundApplication,
		Message:   entity,
	}
}

func MakeUnprocessableEntityApplicationError(message string) Error {
	return Error{
		ErrorType: UnprocessableEntityApplication,
		Message:   message,
	}
}

func MakeUnathorizedApplicationError(message string) Error {
	return Error{
		ErrorType: UnathorizedApplication,
		Message:   message,
	}
}

func MakeForbiddenApplicationError(message string) Error {
	return Error{
		ErrorType: ForbiddenApplication,
		Message:   message,
	}
}

func MakeInternalServerError(message string) Error {
	return Error{
		ErrorType: InternalServerError,
		Message:   message,
	}
}
