package entities

type DatabaseError struct {
	description string
}

func NewDatabaseError(paramMessage string) *DatabaseError {
	return &DatabaseError{description: paramMessage}
}

func (baseType DatabaseError) Error() string {
	return baseType.description
}
