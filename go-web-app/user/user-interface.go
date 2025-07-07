package user

type BaseUser interface {
	PrintFormatted()
	GetFromStdin() error
}
