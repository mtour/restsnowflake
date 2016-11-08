package models

type ReturnData struct {
	Code 	int		`json:"code"`
	Message	string 		`json:"message"`
	Data	interface{}	`json:"data"`
}

func GetReturnData(code int, message string,data interface{})(rd *ReturnData){
	rd = &ReturnData{code,message,data}
	return
}
