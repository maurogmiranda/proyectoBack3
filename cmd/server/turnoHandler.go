package handler

import (
	"fmt"
	"odontologo/internal/domain"
	"odontologo/internal/turno"
	"odontologo/pkg/web"
	"strconv"

	"github.com/gin-gonic/gin"
)

type TurnoHandler struct{
	turnoService turno.TurnoServiceInterface
}

func NewTurnoHandler (t turno.TurnoServiceInterface) TurnoHandler{
	return TurnoHandler{turnoService: t}
}

func (th *TurnoHandler) GetByID() gin.HandlerFunc {
	return func(c *gin.Context) {
		idParam := c.Param("id")
		id, err := strconv.Atoi(idParam)
		if err != nil {
			web.Failure(c,500,err)
			return
		}
		fmt.Println(idParam)
		turno, err := th.turnoService.FindTurnoById(id)
		if err != nil {
			web.Failure(c,404,err)
			return
		}
		c.JSON(200, response{
			Data: turno,
		})
	}
}

func (th *TurnoHandler) Post() gin.HandlerFunc {
	return func(c *gin.Context) {
		var turno domain.Turno

		err := c.ShouldBindJSON(&turno)
		if err != nil {
			web.Failure(c,404,err)
			return
		}

		p, err := th.turnoService.PostTurno(turno)
		if err != nil {
			web.Failure(c,500,err)
			return
		}
		c.JSON(200, response{
			Data: p,
		})
	}
}

func (th *TurnoHandler) Delete() gin.HandlerFunc {
	return func(c *gin.Context) {
		idParam := c.Param("id")
		id, err := strconv.Atoi(idParam)
		if err != nil {
			web.Failure(c,500,err)
			return
		}
		err =th.turnoService.DeleteTurnoById(id)
		if err != nil {
			web.Failure(c,404,err)
			return
		}
			c.JSON(204, response{
			Data: nil,
		})
	}
}

func (th *TurnoHandler) Put() gin.HandlerFunc {
	return func(c *gin.Context) {

		var turno domain.Turno
		err := c.ShouldBindJSON(&turno)
		if err != nil {
			web.Failure(c,404,err)			
			return
		}
		newTurno, err := th.turnoService.PutTurno(turno)
		if err != nil {
			web.Failure(c,500,err)
			return
		}
		c.JSON(204, response{
			Data: newTurno,
		})
	}
}

func (th *TurnoHandler) GetByDNI() gin.HandlerFunc {
	return func(c *gin.Context) {
		dniParam := c.Param("dni")
		dni, err := strconv.Atoi(dniParam)
		if err != nil {
			web.Failure(c,400,err)
		}
		turno, err := th.turnoService.FindTurnoByDNIPaciente(dni)
		if err != nil {
			web.Failure(c,404,err)
			return
		}
		c.JSON(200, response{
			Data: turno,
		})
	}
}

func (th *TurnoHandler) Post2() gin.HandlerFunc {
	return func(c *gin.Context) {
		var turno domain.Turno

		err := c.ShouldBindJSON(&turno)
		if err != nil {
			web.Failure(c,404,err)
			return
		}

		p, err := th.turnoService.PostTurnoDNIYMatricula(turno.Paciente.DNI,turno.Odontologo.Matricula,turno.FechaYHora,turno.Descripcion)
		if err != nil {
			web.Failure(c,500,err)
			return
		}
		c.JSON(200, response{
			Data: p,
		})
	}
}

func (turnoHandler *TurnoHandler) Patch() gin.HandlerFunc {
	return func(c *gin.Context) {

		var turno domain.Turno
		err := c.ShouldBindJSON(&turno)
		if err != nil {
			web.Failure(c,404,err)			
			return
		}
		newTurno, err := turnoHandler.turnoService.PatchTurno(turno)
		if err != nil {
			web.Failure(c,500,err)
			return
		}
		c.JSON(204, response{
			Data: newTurno,
		})
	}
} 