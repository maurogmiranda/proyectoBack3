package turno

import (
	"fmt"
	"odontologo/internal/domain"
	"odontologo/internal/odontologo"
	"odontologo/internal/paciente"
)

type TurnoServiceInterface interface {
	FindTurnoById(id int) (domain.Turno, error)
	DeleteTurnoById(id int) error
	PutTurno(turno domain.Turno) (domain.Turno, error)
	PostTurno(turno domain.Turno) (domain.Turno, error)
	PatchTurno(turnoVar domain.Turno) (domain.Turno, error)
	FindTurnoByDNIPaciente(dni int) (domain.Turno, error)
	PostTurnoDNIYMatricula(dni int, matricula string, fechaYHora string, descripcion string) (domain.Turno, error)
}

type TurnoService struct{
	turnoRepository TurnoRepository
	pacienteService paciente.PacienteServiceInterface
	odontologoService odontologo.OdontologoServiceInterface
}

func NewTurnoService(tR TurnoRepository, ps paciente.PacienteServiceInterface,os odontologo.OdontologoServiceInterface) TurnoService{
	return TurnoService{turnoRepository: tR, pacienteService: ps,odontologoService: os}
}

func (ts TurnoService) FindTurnoById(id int) (domain.Turno, error){
	turno, err := ts.turnoRepository.FindTurnoById(id)
	odontologo,err:=ts.odontologoService.FindOdontologoById(turno.Odontologo.Id)
	if err!=nil {
		fmt.Println("Hubo un error:",err.Error())
		return domain.Turno{}, err
	}
	turno.Odontologo = odontologo
	paciente,err := ts.pacienteService.FindPacienteById(turno.Paciente.Id)
	if err!=nil {
		fmt.Println("Hubo un error:",err.Error())
		return domain.Turno{}, err
	}
	turno.Paciente = paciente
	if err != nil {
		fmt.Println("Hubo un error:",err.Error())
		return domain.Turno{}, err
	}
	return turno, nil
}

func (ts TurnoService) DeleteTurnoById(id int) error{
	err := ts.turnoRepository.DeleteTurnoById(id)
	if err != nil {
		fmt.Println("Hubo un error:",err.Error())
		return err
	}
	return nil
}

func (ts TurnoService) PostTurno(turnoVar domain.Turno) (domain.Turno, error){
	turno, err := ts.turnoRepository.PostTurno(turnoVar)
	odontologo,err:=ts.odontologoService.FindOdontologoById(turno.Odontologo.Id)
	if err!=nil {
		fmt.Println("Hubo un error:",err.Error())
		return domain.Turno{}, err
	}
	turno.Odontologo = odontologo
	paciente,err := ts.pacienteService.FindPacienteById(turno.Paciente.Id)
	if err!=nil {
		fmt.Println("Hubo un error:",err.Error())
		return domain.Turno{}, err
	}
	turno.Paciente = paciente
	if err != nil {
		fmt.Println("Hubo un error:",err.Error())
		return domain.Turno{}, err
	}
	return turno, nil
}

func (ts TurnoService) PutTurno(turnoVar domain.Turno) (domain.Turno, error){
	id := turnoVar.Id
	_, err := ts.FindTurnoById(id)
	if err != nil {
		fmt.Println("Hubo un error:",err.Error())
		return domain.Turno{}, err
	}
	turno,err := ts.turnoRepository.PutTurno(turnoVar)
	if err != nil {
		fmt.Println("Hubo un error:",err.Error())
		return domain.Turno{}, err
	}
	return turno, nil
}

func (ts TurnoService) FindTurnoByDNIPaciente(dni int) (domain.Turno, error){
	turno, err := ts.turnoRepository.FindTurnoByDNIPaciente(dni)
	odontologo,err:=ts.odontologoService.FindOdontologoById(turno.Odontologo.Id)
	if err!=nil {
		fmt.Println("Hubo un error:",err.Error())
		return domain.Turno{}, err
	}
	turno.Odontologo = odontologo
	paciente,err := ts.pacienteService.FindPacienteById(turno.Paciente.Id)
	if err!=nil {
		fmt.Println("Hubo un error:",err.Error())
		return domain.Turno{}, err
	}
	turno.Paciente = paciente
	if err != nil {
		fmt.Println("Hubo un error:",err.Error())
		return domain.Turno{}, err
	}
	return turno, nil
}

func (ts TurnoService) PostTurnoDNIYMatricula(dni int, matricula string, fechaYHora string, descripcion string) (domain.Turno, error){
	turno, err := ts.turnoRepository.postTurnoDNIYMatricula(dni,matricula,fechaYHora,descripcion)
	if err != nil {
		fmt.Println("Hubo un error:",err.Error())
		return domain.Turno{}, err
	}
	return turno, nil
}

func (t TurnoService) PatchTurno(turnoVar domain.Turno) (domain.Turno, error) {
	id := turnoVar.Id
	turno, err := t.FindTurnoById(id)
	if err != nil {
		fmt.Println("Hubo un error:",err.Error())
		return domain.Turno{}, err
	}
	if turnoVar.Odontologo.Id != 0 {
		turno.Odontologo.Id = turnoVar.Odontologo.Id
	}
	if turnoVar.Paciente.Id != 0 {
		turno.Paciente.Id = turnoVar.Paciente.Id
	}
	if  turnoVar.FechaYHora!= "" {
		turno.FechaYHora = turnoVar.FechaYHora
	}
	if turnoVar.Descripcion != "" {
		turno.Descripcion = turnoVar.Descripcion
	}

	turno,err = t.turnoRepository.PutTurno(turno)
		if err != nil {
		fmt.Println("Hubo un error:",err.Error())
		return domain.Turno{}, err
	}
	return turno, nil 
}