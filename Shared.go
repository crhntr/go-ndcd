package nationaldrugcodedirectory

type Error int8

const (
	DoneError Error = iota
)

func (err Error) Error() string {
	if err == DoneError {
		return "done"
	}
	return "err"
}
