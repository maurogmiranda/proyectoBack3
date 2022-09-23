package handler

import (
	"odontologo/internal/domain"
	"odontologo/internal/paciente"
	"odontologo/pkg/web"
	"strconv"

	"github.com/gin-gonic/gin"
)

type PacienteHandler struct{
	pacienteService paciente.PacienteService
}


func NewPacienteHandler (p paciente.PacienteService) PacienteHandler{
	return PacienteHandler{pacienteService: p}
}

func (pacienteHandler *PacienteHandler) GetByID() gin.HandlerFunc {
	return func(c *gin.Context) {
		idParam := c.Param("id")
		id, err := strconv.Atoi(idParam)
		if err != nil {
			web.Failure(c,500,err)
			return
		}
		paciente, err := pacienteHandler.pacienteService.FindPacienteById(id)
		if err != nil {
			web.Failure(c,404,err)
			return
		}
		c.JSON(200, response{
			Data: paciente,
		})
	}
}

func (pacienteHandler *PacienteHandler) Post() gin.HandlerFunc {
	return func(c *gin.Context) {
		var paciente domain.Paciente

		err := c.ShouldBindJSON(&paciente)
		if err != nil {
			web.Failure(c,404,err)
			return
		}

		p, err := pacienteHandler.pacienteService.PostPaciente(paciente)
		if err != nil {
			web.Failure(c,500,err)
			return
		}
		c.JSON(200, response{
			Data: p,
		})
	}
}

func (pacienteHandler *PacienteHandler) Delete() gin.HandlerFunc {
	return func(c *gin.Context) {
		idParam := c.Param("id")
		id, err := strconv.Atoi(idParam)
		if err != nil {
			web.Failure(c,500,err)
			return
		}
		err =pacienteHandler.pacienteService.DeletePacienteById(id)
		if err != nil {
			web.Failure(c,404,err)
			return
		}
			c.JSON(204, response{
			Data: nil,
		})
	}
}

func (pacienteHandler *PacienteHandler) Put() gin.HandlerFunc {
	return func(c *gin.Context) {

		var paciente domain.Paciente
		err := c.ShouldBindJSON(&paciente)
		if err != nil {
			web.Failure(c,404,err)			
			return
		}
		newPaciente, err := pacienteHandler.pacienteService.PutPaciente(paciente)
		if err != nil {
			web.Failure(c,500,err)
			return
		}
		c.JSON(204, response{
			Data: newPaciente,
		})
	}
}

 func (pHandler *PacienteHandler) Patch() gin.HandlerFunc {
	return func(c *gin.Context) {

		var paciente domain.Paciente
		err := c.ShouldBindJSON(&paciente)
		if err != nil {
			web.Failure(c,404,err)			
			return
		}
		newPaciente, err := pHandler.pacienteService.PatchPaciente(paciente)
		if err != nil {
			web.Failure(c,500,err)
			return
		}
		c.JSON(204, response{
			Data: newPaciente,
		})
	}
} 