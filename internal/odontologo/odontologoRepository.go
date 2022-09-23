package odontologo

import (
	"fmt"

	"odontologo/internal/domain"
	"odontologo/pkg/store"
)

type OdontologoRepositoryInterface interface{
	FindOdontologoById(id int) (domain.Odontologo, error)
	DeleteOdontologoById(id int)  error
	PutOdontologo(odontologo domain.Odontologo) error
	PostOdontologo(odontologo domain.Odontologo)  (domain.Odontologo, error)
}

type OdontologoRepository struct{
	storege store.MysqlStorage
}

func NewOdontologoRepository(storageVar store.MysqlStorage) OdontologoRepository{
	return OdontologoRepository{storege:storageVar}
}

func (o OdontologoRepository) FindOdontologoById(id int) (domain.Odontologo, error){
	var odontologo domain.Odontologo
	query := "SELECT * FROM odontologos WHERE id = ?;"
	row := o.storege.DB.QueryRow(query, id)
	err := row.Scan(&odontologo.Id, &odontologo.Nombre, &odontologo.Apellido, &odontologo.Matricula)
	if err != nil {
		return domain.Odontologo{}, err
	}
	return odontologo, nil
}

func (o OdontologoRepository) 	DeleteOdontologoById(id int)  error{
	query := "DELETE FROM odontologos WHERE id = ?;"
	stmt, err := o.storege.DB.Prepare(query)
	if err != nil {
		return err
	}
	res, err := stmt.Exec(id)
	if err != nil {
		return err
	}
	_, err = res.RowsAffected()
	if err != nil {
		return err
	}
	return nil
}

func (o OdontologoRepository) PostOdontologo(odontologo domain.Odontologo) (domain.Odontologo, error){
	query := "INSERT INTO odontologos (nombre, apellido, matricula) VALUES (?, ?, ?);"
	stmt, err := o.storege.DB.Prepare(query)
	if err != nil {
		return domain.Odontologo{}, err
	}
	
	res, err := stmt.Exec(odontologo.Nombre, odontologo.Apellido, odontologo.Matricula)
	if err != nil {
		fmt.Println(err)
		return domain.Odontologo{}, err
	}
	_, err = res.RowsAffected()
	
	

	if err != nil {
		return domain.Odontologo{}, err
	}
	idOdontologo,_ := res.LastInsertId()

	odontologo.Id=int(idOdontologo)
	return odontologo,nil
}

func (odontologoRepository OdontologoRepository) PutOdontologo(odontologo domain.Odontologo) error {
	query := "UPDATE odontologos SET nombre = ?, apellido = ?, matricula = ? WHERE id = ?;"
	stmt, err := odontologoRepository.storege.DB.Prepare(query)
	if err != nil {
		return err
	}
	res, err := stmt.Exec(odontologo.Nombre, odontologo.Apellido, odontologo.Matricula, odontologo.Id)
	if err != nil {
		return err
	}
	_, err = res.RowsAffected()
	if err != nil {
		return err
	}
	
	return nil
}
