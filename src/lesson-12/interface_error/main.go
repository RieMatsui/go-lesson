package main

import "fmt"


type Myerror struct {
	Message string
	ErrCode int
}

func (e *Myerror) Error() string {
	return e.Message
}

func RaiseError() error {
	return &Myerror{
        Message: "カスタムエラーが発生しました",
        ErrCode: 1,
    }
}

func main() {
	err := RaiseError()
	fmt.Println(err.Error())

	e, ok := err.(*Myerror)
	if ok {
		fmt.Println(e.ErrCode)
	}
}
