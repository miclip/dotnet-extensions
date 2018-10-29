package ui

// UI ...
type UI interface {
	PrintLinef(pattern string, args ...interface{})
	ErrorLinef(pattern string, args ...interface{})

	Printf(pattern string, args ...interface{})
	Errorf(pattern string, args ...interface{})

}
