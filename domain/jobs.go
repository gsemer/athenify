package domain

type Job interface {
	Process() Result
}

type Result struct {
	User  User
	Error error
}
