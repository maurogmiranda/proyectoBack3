package paciente

import (
	"fmt"
	"odontologo/internal/domain"
)

type PacienteServiceInterface interface {
	FindPacienteById(id int) (domain.Paciente, error)
	DeletePacienteById(id int) error
	PostPaciente(paciente domain.Paciente) (domain.Paciente, error)
	PutPaciente(paciente domain.Paciente) (domain.Paciente, error)
	PatchPaciente(pacienteVar domain.Paciente) (domain.Paciente, error)
}

type PacienteService struct {
	pacienteRepository PacienteRepositoryInterface
}

func NewPacienteService(p PacienteRepository) PacienteService {
	return PacienteService{pacienteRepository: p}
}

func (p *PacienteService) FindPacienteById(id int) (domain.Paciente, error) {
	paciente, err := p.pacienteRepository.FindPacienteById(id)
	if err != nil {
		fmt.Println("Hubo un error:",err.Error())
		return domain.Paciente{}, err
	}
	return paciente, nil
}

func (p *PacienteService) DeletePacienteById(id int) error {
	err := p.pacienteRepository.DeletePacienteById(id)
	if err != nil {
		fmt.Println("Hubo un error:",err.Error())
		return err
	}
	return nil
}

func (p *PacienteService) PostPaciente(pacienteVar domain.Paciente) (domain.Paciente, error) {
	paciente, err := p.pacienteRepository.PostPaciente(pacienteVar)
	if err != nil {
		fmt.Println("Hubo un error:",err.Error())
		return domain.Paciente{}, err
	}
	return paciente, nil
}

func (p *PacienteService) PutPaciente(pacienteVar domain.Paciente) (domain.Paciente, error) {
	id := pacienteVar.Id
	_, err := p.FindPacienteById(id)
	if err != nil {
		fmt.Println("Hubo un error:",err.Error())
		return domain.Paciente{}, err
	}
	err = p.pacienteRepository.PutPaciente(pacienteVar)
	if err != nil {
		fmt.Println("Hubo un error:",err.Error())
		return domain.Paciente{}, err
	}
	return pacienteVar, nil
}

func (p *PacienteService) PatchPaciente(pacienteVar domain.Paciente) (domain.Paciente, error) {
	id := pacienteVar.Id
	paciente, err := p.FindPacienteById(id)
	if err != nil {
		fmt.Println("Hubo un error:",err.Error())
		return domain.Paciente{}, err
	}
	if pacienteVar.Nombre != "" {
		paciente.Nombre = pacienteVar.Nombre
	}
	if pacienteVar.Apellido != "" {
		paciente.Apellido = pacienteVar.Apellido
	}
	if pacienteVar.Domicilio != "" {
		paciente.Domicilio = pacienteVar.Domicilio
	}
	if pacienteVar.DNI != 0 {
		paciente.DNI = pacienteVar.DNI
	}
	if pacienteVar.FechaDeAlta != "" {
		paciente.FechaDeAlta = pacienteVar.FechaDeAlta
	}
	err = p.pacienteRepository.PutPaciente(paciente)
	if err != nil {
		return domain.Paciente{}, err
	}
	return paciente, nil 
}