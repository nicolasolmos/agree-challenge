package entities

type DatabaseError struct {
	description string
}

func NewDatabaseError(paramMessage string) *DatabaseError {
	myError := DatabaseError{description: paramMessage}
	return &myError
}

func (baseType DatabaseError) Error() string {
	return baseType.description
}
