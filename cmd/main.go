package main

import (
	"database/sql"

	"odontologo/cmd/server"
	"odontologo/internal/odontologo"
	"odontologo/internal/paciente"
	"odontologo/internal/turno"
	"odontologo/pkg/store"

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
)

func main() {
	usuario :="root"
	contrasenia := "root"

	
	db, err := sql.Open("mysql",usuario + ":"+contrasenia+"@/")
	if err != nil {
		panic(err.Error())
	}
	err = db.Ping()
	if err != nil {
		panic(err.Error())
	}
	db.Exec("CREATE DATABASE IF NOT EXISTS my_db")
	db.Exec("CREATE TABLE IF NOT EXISTS `my_db`.`odontologos` (`id` INT NOT NULL AUTO_INCREMENT,`nombre` VARCHAR(45) NULL,`apellido` VARCHAR(45) NULL,`matricula` VARCHAR(45) NULL,PRIMARY KEY (`id`));")
	db.Exec("CREATE TABLE IF NOT EXISTS `my_db`.`pacientes` (`id` INT NOT NULL AUTO_INCREMENT,`nombre` VARCHAR(45) NULL,`apellido` VARCHAR(45) NULL,`domicilio` VARCHAR(45) NULL,`dni` INT NULL,`fecha_de_alta` VARCHAR(45) NULL,PRIMARY KEY (`id`));")
	db.Exec("CREATE TABLE IF NOT EXISTS `my_db`.`turnos` (`id` INT NOT NULL AUTO_INCREMENT,`paciente_id` INT NULL,`odontologo_id` INT NULL,`fecha_y_hora` VARCHAR(45) NULL,`descripcion` VARCHAR(1000) NULL,PRIMARY KEY (`id`),INDEX `odontologoID_idx` (`odontologo_id` ASC) VISIBLE,INDEX `pacienteID_idx` (`paciente_id` ASC) VISIBLE,CONSTRAINT `odontologoID` FOREIGN KEY (`odontologo_id`) REFERENCES `my_db`.`odontologos` (`id`) ON DELETE NO ACTION ON UPDATE NO ACTION, CONSTRAINT `pacienteID` FOREIGN KEY (`paciente_id`) REFERENCES `my_db`.`pacientes` (`id`) ON DELETE NO ACTION ON UPDATE NO ACTION);")

	db.Close()


	db, err = sql.Open("mysql", usuario + ":"+contrasenia+"@/my_db")
	if err != nil {
		panic(err.Error())
	}
	err = db.Ping()
	if err != nil {
		panic(err.Error())
	}

	storage := store.NewMySQLStorage(db)

	repoOdontologo := odontologo.NewOdontologoRepository(storage)
	serviceOdontologo := odontologo.NewOdontologoService(repoOdontologo)
	odonHandler := handler.NewOdontologoHandler(serviceOdontologo)

	repoPaciente := paciente.NewPacienteRepository(storage)
	servicePaciente := paciente.NewPacienteService(repoPaciente)
	pacienteHandler := handler.NewPacienteHandler(servicePaciente)

	repoTurno := turno.NewTurnoRepository(storage)
	serviceTurno := turno.NewTurnoService(repoTurno,&servicePaciente,&serviceOdontologo)
	turnoHandler := handler.NewTurnoHandler(serviceTurno)

	r := gin.Default()

	r.GET("/ping", func(c *gin.Context) { c.String(200, "pong") })
	odontologos := r.Group("/odontologos")
	{
		odontologos.GET(":id", odonHandler.GetByID())
		odontologos.POST("", odonHandler.Post())
		odontologos.DELETE(":id", odonHandler.Delete())
		odontologos.PUT("",odonHandler.Put())
		odontologos.PATCH("", odonHandler.Patch())
	}

	pacientes := r.Group("/pacientes")
	{
		pacientes.GET(":id", pacienteHandler.GetByID())
		pacientes.POST("", pacienteHandler.Post())
		pacientes.DELETE(":id", pacienteHandler.Delete())
		pacientes.PUT("",pacienteHandler.Put())
		pacientes.PATCH("", pacienteHandler.Patch())
	}

	turnos := r.Group("/turnos")
	{
		turnos.GET(":id", turnoHandler.GetByID())
		turnos.POST("", turnoHandler.Post())
		turnos.DELETE(":id", turnoHandler.Delete())
		turnos.PUT("",turnoHandler.Put())
		turnos.GET("/findByDNI/:dni",turnoHandler.GetByDNI())
		turnos.POST("/dniYMatricula",turnoHandler.Post2())
		turnos.PATCH("", turnoHandler.Patch())
	}

	

	r.Run(":8080")
}