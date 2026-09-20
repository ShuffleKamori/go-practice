package events

type request_t struct {
	Type string
}

var (
	Request = make(chan request_t, 3)
	Result  = make(chan any, 3)
)

func SendRequest(command string) { //GUI uses for make request in backend
	select {
	case Request <- request_t{Type: command}:
	default:
	}
}

func CheckResult() (any, bool) { //GUI uses for checkenh result from request
	select {
	case data := <-Result:
		return data, true
	default:
		return nil, false
	}
}

/*
Моя памятка:

вызовы функций
events.SendRequest() создаёт запрос GUI -> backend для выполнения задачи, пример: events.SendRequest("GET_CATALOG")
events.CheckResult() вызывается GUI для того, чтобы проверить выполнилась ли задача



*/
