package serializers


// структура для выходных данных выпуска нового access токена
type VerifyTokenOut struct {
	Status 		string `json:"status" example:"ok"`
	StatusCode 	int `json:"statusCode" example:"200"`
}
