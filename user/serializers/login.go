package serializers

import (
	"fmt"

	echo "github.com/labstack/echo/v4"
	validate "github.com/gobuffalo/validate/v3"

	"github.com/Danil-114195722/Knofu/user/models"
	"github.com/Danil-114195722/Knofu/user/services"
	tokensServices "github.com/Danil-114195722/Knofu/tokens/services"
	coreDB "github.com/Danil-114195722/Knofu/core/db"
	coreValidator "github.com/Danil-114195722/Knofu/core/validator"
)


type LoginUserIn struct {
	Email 		string `json:"email" validate:"required|email"`
	Password 	string `json:"password" validate:"required"`
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
		return models.User{}, echo.NewHTTPError(500, map[string]string{"dbConnect": "Failed to connect to DB"})
	}

	// проверка на наличие юзера в БД с таким email'ом
	var userFromDB models.User
	findResult := dbConnect.Where("email = ?", self.Email).First(&userFromDB)
	// если юзер с таким email'ом не найден
	if err = findResult.Error; err != nil {
		httpError := echo.NewHTTPError(400, map[string]string{"email": "User with given email does not exist"})
		return models.User{}, httpError
	}

	// проверка пароля
	if !services.PasswordIsCorrect(self.Password, userFromDB.Password) {
		httpError := echo.NewHTTPError(400, map[string]string{"password": "Invalid password"})
		return models.User{}, httpError
	}

	return userFromDB, nil
}


type LoginUserOut struct {
	ID				uint64 `json:"id"`
	Email 			string `json:"email"`
	FirstName 		string `json:"firstName"`
	LastName 		string `json:"lastName"`
	AccessToken  	string `json:"accessToken"`
	RefreshToken 	string `json:"refreshToken"`
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
