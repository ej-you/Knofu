package serializers

import (
	echo "github.com/labstack/echo/v4"
	validate "github.com/gobuffalo/validate/v3"

	userErrors "github.com/ej-you/Knofu/user/errors"
	"github.com/ej-you/Knofu/user/models"
	"github.com/ej-you/Knofu/user/services"

	tokensServices "github.com/ej-you/Knofu/token/services"
	coreDB "github.com/ej-you/Knofu/core/db"
	coreErrors "github.com/ej-you/Knofu/core/errors"
	coreValidator "github.com/ej-you/Knofu/core/validator"
)


// структура для входных данных входа юзера
type LoginUserIn struct {
	Email 		string `json:"email" myvalid:"required|email" validate:"required" example:"user@example.com"`
	Password 	string `json:"password" myvalid:"required" validate:"required" example:"qwerty123"`
}

// базовая валидация полей по тегам
func (self *LoginUserIn) IsValid(errors *validate.Errors) {
	coreValidator.BaseValidator(self, errors)
}

// более глубокая валидация с возвратом ошибок валидации
func (self *LoginUserIn) Validate() (models.User, error) {
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
		return models.User{}, httpError
	}

	// получение соединения с БД
	dbConnect, err := coreDB.GetConnection()
	if err != nil {
		return models.User{}, coreErrors.DBConnectError
	}

	// проверка на наличие юзера в БД с таким email'ом
	var userFromDB models.User
	findResult := dbConnect.Where("email = ?", self.Email).First(&userFromDB)
	// если юзер с таким email'ом не найден
	if err = findResult.Error; err != nil {
		return models.User{}, userErrors.UserDoesNotExistError
	}

	// проверка пароля
	if !services.PasswordIsCorrect(self.Password, userFromDB.Password) {
		return models.User{}, userErrors.InvalidPasswordError
	}

	return userFromDB, nil
}


// структура для выходных данных входа юзера
type LoginUserOut struct {
	ID				uint64 `json:"id" example:"4354"`
	Email 			string `json:"email" example:"user@example.com"`
	FirstName 		string `json:"firstName" example:"John"`
	LastName 		string `json:"lastName" example:"Lennon"`
	AccessToken  	string `json:"accessToken" example:"c73gnetfhigcsi.gaes4inva4a.gcawinxbwi4"`
	RefreshToken 	string `json:"refreshToken" example:"ghvnkvg5ic73ea.hv567eke4n5.5ugkwe47hgv4"`
}

// формирование структуры для ответа
func GetLoginOutStruct(loggedUser models.User) (LoginUserOut, error) {
	// получение access токена для юзера
	accessToken, err := tokensServices.GetAccessToken(loggedUser.ID)
	if err != nil {
		return LoginUserOut{}, err
	}
	// получение refresh токена для юзера
	refreshToken, err := tokensServices.GetRefreshToken(loggedUser.ID)
	if err != nil {
		return LoginUserOut{}, err
	}

	userOut := LoginUserOut{
		ID: loggedUser.ID,
		Email: loggedUser.Email,
		FirstName: loggedUser.FirstName,
		LastName: loggedUser.LastName,
		AccessToken: accessToken,
		RefreshToken: refreshToken,
	}

	return userOut, nil
}
