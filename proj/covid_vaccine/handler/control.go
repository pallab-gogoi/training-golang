package handler

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"github.com/pallab-gogoi/covid_vaccine/database"
	"github.com/pallab-gogoi/covid_vaccine/model"
)

type Handler struct {
	DB *gorm.DB
}

func (h *Handler) GetData(in *gin.Context) {
	data, err := database.GetData(h.DB)
	if err != nil {
		log.Println(err)
		in.JSON(http.StatusInternalServerError, err)
	}
	in.JSON(http.StatusOK, data)

}

func (h *Handler) SaveData(in *gin.Context) {
	data := model.Details{}
	err := in.BindJSON(&data)
	if err != nil {
		log.Println(err)
		in.JSON(http.StatusInternalServerError, err)
	}
	err = database.SaveData(h.DB, &data)
	if err != nil {
		log.Println(err)
		in.JSON(http.StatusInternalServerError, err)
	}
	in.JSON(http.StatusOK, data)

}
