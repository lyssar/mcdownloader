package errors

import (
	"github.com/lyssar/msdcli/logger"
	log "github.com/sirupsen/logrus"
)

type ApplicationError struct {
	Err      string
	ErrLevel log.Level
}

func (e ApplicationError) Error() string {
	return e.Err
}

func NewWarning(text string) *ApplicationError {
	return &ApplicationError{Err: text, ErrLevel: log.WarnLevel}
}

func NewError(text string) *ApplicationError {
	return &ApplicationError{Err: text, ErrLevel: log.ErrorLevel}
}

func NewFatal(text string) *ApplicationError {
	return &ApplicationError{Err: text, ErrLevel: log.FatalLevel}
}

// Errs is an errors that collects other errors, for when you want to do
// several things and then report all of them.
type Errs struct {
	errors []ApplicationError
}

// Add will add the errors to the errors list
func (e *Errs) Add(err ApplicationError) {
	e.errors = append(e.errors, err)
}

// IsEmpty will return true, if the errors list is empty.
func (e *Errs) IsEmpty() bool {
	return len(e.errors) == 0
}

// Errors returns all errors as array.
func (e *Errs) Errors() []ApplicationError {
	return e.errors
}

// Print will log all errors.
func (e *Errs) Print() {
	if !e.IsEmpty() {
		for _, err := range e.errors {
			if err.ErrLevel == log.WarnLevel {
				logger.Warning(err)
			} else {
				logger.Error(err)
			}
		}
	}
}

// Check will exit the program if the errors list is not empty.
func (e *Errs) Check() {
	e.Print()
	if !e.IsEmpty() {
		logger.Fatal("errors found, exiting")
	}
}

// Check will exit the program if the errors list is not empty.
func Check(e *ApplicationError) {
	if e != nil {
		if e.ErrLevel == log.WarnLevel {
			logger.Warning(e)
		} else if e.ErrLevel == log.FatalLevel {
			logger.Fatal(e)
		} else {
			logger.Error(e)
		}
		logger.Exit(1)
	}
}

func CheckStandardErr(err error) {
	if err != nil {
		Check(NewError(err.Error()))
	}
}
