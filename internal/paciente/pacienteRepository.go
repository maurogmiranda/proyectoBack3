package paciente

import (
	"fmt"
	"odontologo/internal/domain"
	"odontologo/pkg/store"
)

type PacienteRepositoryInterface interface{
	FindPacienteById(id int) (domain.Paciente, error)
	DeletePacienteById(id int)  error
	PutPaciente(odontologo domain.Paciente) error
	PostPaciente(odontologo domain.Paciente)  (domain.Paciente, error)
}

type PacienteRepository struct{
	storege store.MysqlStorage
}

func NewPacienteRepository(storageVar store.MysqlStorage) PacienteRepository{
	return PacienteRepository{storege:storageVar}
}

func (o PacienteRepository) FindPacienteById(id int) (domain.Paciente, error){
	var paciente domain.Paciente
	query := "SELECT * FROM pacientes WHERE id = ?;"
	row := o.storege.DB.QueryRow(query, id)
	err := row.Scan(&paciente.Id, &paciente.Nombre, &paciente.Apellido, &paciente.Domicilio,&paciente.DNI,&paciente.FechaDeAlta)
	if err != nil {
		return domain.Paciente{}, err
	}
	return paciente, nil
}

func (o PacienteRepository) DeletePacienteById(id int)  error{
	query := "DELETE FROM pacientes WHERE id = ?;"
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

func (o PacienteRepository) PostPaciente(paciente domain.Paciente) (domain.Paciente, error){
	query := "INSERT INTO pacientes (nombre, apellido, domicilio, dni,fecha_de_alta) VALUES (?, ?, ?, ?, ?);"
	stmt, err := o.storege.DB.Prepare(query)
	if err != nil {
		return domain.Paciente{}, err
	}
	
	res, err := stmt.Exec(paciente.Nombre, paciente.Apellido, paciente.Domicilio,paciente.DNI,paciente.FechaDeAlta)
	if err != nil {
		fmt.Println(err)
		return domain.Paciente{}, err
	}
	_, err = res.RowsAffected()
	
	

	if err != nil {
		return domain.Paciente{}, err
	}
	idPaciente,_ := res.LastInsertId()

	paciente.Id=int(idPaciente)
	return paciente,nil
}

func (o PacienteRepository) PutPaciente(paciente domain.Paciente) error {
	query := "UPDATE pacientes SET nombre = ?, apellido = ?, domicilio = ?, dni = ?, fecha_de_alta = ? WHERE id = ?;"
	stmt, err := o.storege.DB.Prepare(query)
	if err != nil {
		return err
	}
	res, err := stmt.Exec(paciente.Nombre, paciente.Apellido, paciente.Domicilio,paciente.DNI,paciente.FechaDeAlta, paciente.Id)
	if err != nil {
		return err
	}
	_, err = res.RowsAffected()
	if err != nil {
		return err
	}
	
	return nil
}