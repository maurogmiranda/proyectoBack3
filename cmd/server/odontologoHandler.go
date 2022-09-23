package handler

import (
	"odontologo/internal/domain"
	"odontologo/internal/odontologo"
	"odontologo/pkg/web"
	"strconv"

	"github.com/gin-gonic/gin"
)

type OdontologoHandler struct{
	odontologoService odontologo.OdontologoService
}

type response struct {
	Data interface{} `json:"data"`
}

func NewOdontologoHandler (o odontologo.OdontologoService) OdontologoHandler{
	return OdontologoHandler{odontologoService: o}
}

func (odontologoHandler *OdontologoHandler) GetByID() gin.HandlerFunc {
	return func(c *gin.Context) {
		idParam := c.Param("id")
		id, err := strconv.Atoi(idParam)
		if err != nil {
			web.Failure(c,500,err)
			return
		}
		odontologo, err := odontologoHandler.odontologoService.FindOdontologoById(id)
		if err != nil {
			web.Failure(c,404,err)
			return
		}
		c.JSON(200, response{
			Data: odontologo,
		})
	}
}

func (odontologoHandler *OdontologoHandler) Post() gin.HandlerFunc {
	return func(c *gin.Context) {
		var odontologo domain.Odontologo

		err := c.ShouldBindJSON(&odontologo)
		if err != nil {
			web.Failure(c,400,err)
			return
		}

		o, err := odontologoHandler.odontologoService.PostOdontologo(odontologo)
		if err != nil {
			web.Failure(c,500,err)
			return
		}
		c.JSON(200, response{
			Data: o,
		})
	}
}

func (odontologoHandler *OdontologoHandler) Delete() gin.HandlerFunc {
	return func(c *gin.Context) {
		idParam := c.Param("id")
		id, err := strconv.Atoi(idParam)
		if err != nil {
			web.Failure(c,500,err)
			return
		}
		err =odontologoHandler.odontologoService.DeleteOdontologoById(id)
		if err != nil {
			web.Failure(c,400,err)
			return
		}
			c.JSON(204, response{
			Data: nil,
		})
	}
}

func (odonHandler *OdontologoHandler) Put() gin.HandlerFunc {
	return func(c *gin.Context) {

		var odontologo domain.Odontologo
		err := c.ShouldBindJSON(&odontologo)
		if err != nil {
			web.Failure(c,400,err)			
			return
		}
		newOdontologo, err := odonHandler.odontologoService.PutOdontologo(odontologo)
		if err != nil {
			web.Failure(c,500,err)
			return
		}
		c.JSON(204, response{
			Data: newOdontologo,
		})
	}
}

func (odonHandler *OdontologoHandler) Patch() gin.HandlerFunc {
	return func(c *gin.Context) {

		var odontologo domain.Odontologo
		err := c.ShouldBindJSON(&odontologo)
		if err != nil {
			web.Failure(c,400,err)			
			return
		}
		newOdontologo, err := odonHandler.odontologoService.PatchOdontologo(odontologo)
		if err != nil {
			web.Failure(c,500,err)
			return
		}
		c.JSON(204, response{
			Data: newOdontologo,
		})
	}
}