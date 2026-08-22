package main

import "fmt"

type Notifier interface{
	Send(msg string) error
}

type EmailNotifier struct{
	email string
	password string
}

type SMSNotifier struct{
	phone int
}

func (e *EmailNotifier) Send(msg string) error{
	fmt.Println("Отправил ссылку для подтверждения почты на email!", e.email,
	 "ваш код проверки:",
	 e.password)

	return nil
}

func (s *SMSNotifier) Send(msg string) error{
	fmt.Println("Ваш код проверки выслан на номер: ", s.phone)

	return nil
}

func main(){
	var bit float64
	bit = 152.22

	PrintTypeDetails(bit)
}

func PrintTypeDetails(a any){
	switch tp := a.(type){
	case int:
		fmt.Println("Это целочисленное значение: ", tp)
	case string:
		fmt.Println("Это строка: ", tp)
	case float64:
		fmt.Println("Это число с плавающей точкой: ", tp)
	default:
		fmt.Println("Unknow type")
	} 
}


