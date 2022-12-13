package controller

import (
	"cobaclean/internal/database"
	"cobaclean/internal/domain"
	"cobaclean/internal/model"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"gorm.io/gorm"
)

var ResponseJson = model.ResponseJson
var ResponseError = model.ResponseError

type UserController struct {
	Service domain.UserAdapterService
}

func (uc *UserController) Index(w http.ResponseWriter, r *http.Request) {
	// user := []model.User{}

	// if err := database.DB.Find(&user).Error; err != nil {
	// 	ResponseError(w, http.StatusInternalServerError, map[string]interface{}{
	// 		"code":    http.StatusInternalServerError,
	// 		"message": "Error",
	// 		"data":    err,
	// 	})
	// 	return
	// }

	var resUser []model.ResponseUser
	res := uc.Service.GetAllUsersService()

	if len(res) > 0 {
		for i := 0; i < len(res); i++ {
			var e = res[i]
			resUser = append(resUser, model.ResponseUser{
				Name:  e.Name,
				Email: e.Email,
			})
		}
	}
	ResponseJson(w, http.StatusOK, map[string]interface{}{
		"code":    http.StatusOK,
		"message": "Success",
		"data":    resUser,
	})
}
func Show(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)
	id, err := strconv.ParseInt(vars["id"], 10, 64)
	if err != nil {
		ResponseError(w, http.StatusBadRequest, map[string]interface{}{
			"code":    http.StatusBadRequest,
			"message": "Error",
		})
		return
	}

	var user model.User
	if err := database.DB.First(&user, id).Error; err != nil {
		switch err {
		case gorm.ErrRecordNotFound:
			ResponseError(w, http.StatusNotFound, map[string]interface{}{
				"code":    http.StatusNotFound,
				"message": "Error",
			})
			return
		default:
			ResponseError(w, http.StatusInternalServerError, map[string]interface{}{
				"code":    http.StatusInternalServerError,
				"message": "Error",
			})
			return
		}
	}

	ResponseJson(w, http.StatusOK, map[string]interface{}{
		"code":    http.StatusOK,
		"message": "Success",
		"data":    user,
	})
}
