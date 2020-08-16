package chapter3

type FullStackError struct {
}

func (err *FullStackError) Error() string {
	return "Stack is Full"
}

type EmptyStackError struct {
}

func (err *EmptyStackError) Error() string {
	return "Stack is Empty"
}
