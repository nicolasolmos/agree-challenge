package entities

type DatabaseError struct {
	description string
}

func NewDatabaseError(paramMessage string) *DatabaseError {
	var newError = new(DatabaseError)
	newError.setDescription(paramMessage)
	return newError
}

func (baseType DatabaseError) Error() string {
	return baseType.description
}

func (baseType DatabaseError) setDescription(paramDescription string) {
	baseType.description = paramDescription
}
