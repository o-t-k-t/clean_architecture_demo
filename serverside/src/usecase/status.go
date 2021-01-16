package usecase

type Status int

func (s Status) Code() int {
	return int(s)
}
