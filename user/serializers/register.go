package serializers

import (
	"fmt"

	echo "github.com/labstack/echo/v4"
	validate "github.com/gobuffalo/validate/v3"

	"github.com/ej-you/Knofu/user/models"
	"github.com/ej-you/Knofu/user/services"
	userErrors "github.com/ej-you/Knofu/user/errors"

	tokensServices "github.com/ej-you/Knofu/token/services"
	coreDB "github.com/ej-you/Knofu/core/db"
	coreErrors "github.com/ej-you/Knofu/core/errors"
	coreValidator "github.com/ej-you/Knofu/core/validator"
)


// структура для входных данных регистрации юзера
type RegisterUserIn struct {
	Email 		string `json:"email" myvalid:"required|email" validate:"required" example:"user@example.com"`
	FirstName 	string `json:"firstName" myvalid:"required" validate:"required" example:"John"`
	LastName 	string `json:"lastName" myvalid:"required" validate:"required" example:"Lennon"`
	Password 	string `json:"password" myvalid:"required|min:8|max:50" validate:"required" example:"qwerty123" minLength:"8" maxLength:"50"`
}

// базовая валидация полей по тегам
func (self *RegisterUserIn) IsValid(errors *validate.Errors) {
	coreValidator.BaseValidator(self, errors)
}

// более глубокая валидация с возвратом ошибок валидации
func (self *RegisterUserIn) Validate() error {
	// базовая валидация полей по тегам
	var validateErrors *validate.Errors = validate.Validate(self)

	if len(validateErrors.Errors) > 0 {
		// словарь для ошибок
		errMap := make(map[string]string, len(validateErrors.Errors))

		for key, value := range validateErrors.Errors {
			errMap[key] = value[0]
		}
		// возвращаем *echo.HTTPError
		httpError := echo.NewHTTPError(400, errMap)
		return httpError
	}

	// получение соединения с БД
	dbConnect, err := coreDB.GetConnection()
	if err != nil {
		return coreErrors.DBConnectError
	}

	// проверка на уже существование юзера в БД с таким email'ом
	var userFromDB models.User
	findResult := dbConnect.Where("email = ?", self.Email).First(&userFromDB)
	// если юзер с таким email'ом найден
	if err = findResult.Error; err == nil {
		return userErrors.UserAlreadyExistsError
	}

	return nil
}

// создание нового юзера в БД
func (self *RegisterUserIn) Create() (models.User, error) {
	// получаем хэш пароля
	hashPasswd, err := services.EncodePassword(self.Password)
	if err != nil {
		return models.User{}, err
	}

	newUser := models.User{
		Email: self.Email,
		FirstName: self.FirstName,
		LastName: self.LastName,
		Password: hashPasswd,
	}

	// получение соединения с БД
	dbConnect, err := coreDB.GetConnection()
	if err != nil {
		return models.User{}, coreErrors.DBConnectError
	}

	createResult := dbConnect.Create(&newUser)
	if err = createResult.Error; err != nil {
		fmt.Println(err)
		return models.User{}, err
	}


	return newUser, nil
}


// структура для выходных данных регистрации юзера
type RegisterUserOut struct {
	ID				uint64 `json:"id" example:"4354"`
	Email 			string `json:"email" example:"user@example.com"`
	FirstName 		string `json:"firstName" example:"John"`
	LastName 		string `json:"lastName" example:"Lennon"`
	AccessToken  	string `json:"accessToken" example:"c73gnetfhigcsi.gaes4inva4a.gcawinxbwi4"`
	RefreshToken 	string `json:"refreshToken" example:"ghvnkvg5ic73ea.hv567eke4n5.5ugkwe47hgv4"`
}

// формирование структуры для ответа
func GetRegisterOutStruct(newUser models.User) (RegisterUserOut, error) {
	// получение access токена для юзера
	accessToken, err := tokensServices.GetAccessToken(newUser.ID)
	if err != nil {
		return RegisterUserOut{}, err
	}
	// получение refresh токена для юзера
	refreshToken, err := tokensServices.GetRefreshToken(newUser.ID)
	if err != nil {
		return RegisterUserOut{}, err
	}

	userOut := RegisterUserOut{
		ID: newUser.ID,
		Email: newUser.Email,
		FirstName: newUser.FirstName,
		LastName: newUser.LastName,
		AccessToken: accessToken,
		RefreshToken: refreshToken,
	}

	return userOut, nil
}
