package main

import (

	"fmt"

	"strconv"
	"strings"
)

type LogStat interface{
	Summary() string
}

type AccessError struct{
	Line string
	Message string
}

func (e *AccessError) Error() string{
	//	var ve *AccessError
	//if errors.As(err, &ve){
	//	log.Printf("ошибка в строке %s: %s", ve.Line, ve.Message)
	//}

	return e.Summary()
}

func (e *AccessError) Summary() string{
	return fmt.Sprintf("ошибка в строке '%s': %s", e.Line, e.Message)
}

func ProcessLogs(logs []string) (map[int]int, error){
	maps := make(map[int]int, len(logs))

	defer func(){
		if r := recover(); r != nil{
			fmt.Println("error from ", r)
		}

		//panic("критическая ошибка сервера")
	}()

	for _, logEntry:= range logs{
		parts := strings.Split(logEntry, ":")
		if len(parts) != 2{
			fmt.Println("panic suka blyat")
		}
		id1 := parts[0]
		req1 := parts[1]

		id, err := strconv.Atoi(id1)
		if err != nil{
			fmt.Printf("error: ", err)
			return nil, &AccessError{
				Line: logEntry,
				Message: "invalid format",
			}
		}

		req, err1 := strconv.Atoi(req1)
		if err1 != nil{
			fmt.Printf("error: ", err1)
			return nil, &AccessError{
				Line: logEntry,
				Message: "invalid format",
			}
		}
		maps[id] = req



	}
	return maps, nil
}
func main() {
    logs := []string{"42:200", "13:500", "invalid_log"}
    res, err := ProcessLogs(logs)
    if err != nil {
        fmt.Println("Ошибка:", err)
    } else {
        fmt.Println("Результат:", res)
    }
}
	
	

