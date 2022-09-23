package odontologo

import (
	"fmt"
	"odontologo/internal/domain"
)

type OdontologoServiceInterface interface {
	FindOdontologoById(id int) (domain.Odontologo, error)
	DeleteOdontologoById(id int) error
	PostOdontologo(odontologo domain.Odontologo)  (domain.Odontologo, error)
	PutOdontologo(odontologo domain.Odontologo) (domain.Odontologo,error)
	PatchOdontologo(odontologoVar domain.Odontologo) (domain.Odontologo,error)
}

type OdontologoService struct{
	odontologoRepository OdontologoRepositoryInterface
}

func NewOdontologoService(o OdontologoRepository) OdontologoService{
	return OdontologoService{odontologoRepository: o}
}

func (o *OdontologoService) FindOdontologoById(id int) (domain.Odontologo, error){
	odontologo, err := o.odontologoRepository.FindOdontologoById(id)
	if err != nil {
		fmt.Println("Hubo un error:",err.Error())
		return domain.Odontologo{}, err
	}
	return odontologo, nil
}

func (o *OdontologoService) DeleteOdontologoById(id int) error{
	err := o.odontologoRepository.DeleteOdontologoById(id)
	if err != nil {
		fmt.Println("Hubo un error:",err.Error())
		return err
	}
	return nil
}

func (o *OdontologoService) PostOdontologo(odontologo domain.Odontologo)  (domain.Odontologo, error){
	odontologo, err := o.odontologoRepository.PostOdontologo(odontologo)
	if err != nil {
		fmt.Println("Hubo un error:",err.Error())
		return domain.Odontologo{}, err
	}
	return odontologo, nil
}

func (o *OdontologoService) PutOdontologo(odontologoVar domain.Odontologo) (domain.Odontologo,error){
	id := odontologoVar.Id
	odontologo, err := o.FindOdontologoById(id)
	if err != nil {
		fmt.Println("Hubo un error:",err.Error())
		return domain.Odontologo{}, err
	}
	 err = o.odontologoRepository.PutOdontologo(odontologoVar)
	if err != nil {
		fmt.Println("Hubo un error:",err.Error())
		return domain.Odontologo{}, err
	}
	return odontologo, nil
}

func (o *OdontologoService) PatchOdontologo(odontologoVar domain.Odontologo) (domain.Odontologo,error){
	id := odontologoVar.Id
	odontologo, err := o.FindOdontologoById(id)
	if err != nil {
		fmt.Println("Hubo un error:",err.Error())
		return domain.Odontologo{}, err
	}
	if odontologoVar.Nombre != "" {
		odontologo.Nombre = odontologoVar.Nombre
	}
	if odontologoVar.Apellido != "" {
		odontologo.Apellido = odontologoVar.Apellido
	}
	if odontologoVar.Matricula != "" {
		odontologo.Matricula = odontologoVar.Matricula
	}
	 err = o.odontologoRepository.PutOdontologo(odontologo)
	if err != nil {
		fmt.Println("Hubo un error:",err.Error())
		return domain.Odontologo{}, err
	}
	return odontologo, nil
}