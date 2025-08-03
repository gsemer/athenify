package domain

type Job interface {
	Process()
}

type Result struct {
	User  User
	Error error
}
