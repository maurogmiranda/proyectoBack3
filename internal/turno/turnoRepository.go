package turno

import (
	"fmt"
	"odontologo/internal/domain"
	"odontologo/pkg/store"
)

type TurnoRepositoryInterface interface{
	FindTurnoById(id int) (domain.Turno, error)
	DeleteTurnoById(id int)  error
	PutTurno(turno domain.Turno) (domain.Turno, error)
	PostTurno(turno domain.Turno)  (domain.Turno, error)
//	PatchTurno(turno domain.Turno) (domain.Turno, error)
	FindTurnoByDNIPaciente(dni int) (domain.Turno, error)
	postTurnoDNIYMatricula(dni int, matricula string, fechaYHora, descripcion string) (domain.Turno, error)
}

type TurnoRepository struct {
	storage store.MysqlStorage
}

func NewTurnoRepository(db store.MysqlStorage) TurnoRepository{
	return TurnoRepository{storage:db}
}

func (tr TurnoRepository) FindTurnoById(id int) (domain.Turno, error){
	var turno domain.Turno
	fmt.Println(id)
	query := "SELECT * FROM turnos WHERE id = ?;"
	row := tr.storage.DB.QueryRow(query, id)
	err := row.Scan(&turno.Id,&turno.Paciente.Id, &turno.Odontologo.Id, &turno.FechaYHora, &turno.Descripcion)
	if err != nil {
		return domain.Turno{}, err
	}
	return turno, nil
}

func (tr TurnoRepository) DeleteTurnoById(id int)  error{
	query := "DELETE FROM turnos WHERE id = ?;"
	stmt, err := tr.storage.DB.Prepare(query)
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

func (tr TurnoRepository) PostTurno(turno domain.Turno) (domain.Turno, error){
	query := "INSERT INTO turnos (paciente_id, odontologo_id, fecha_y_hora, descripcion) VALUES (?, ?, ?, ?);"
	stmt, err := tr.storage.DB.Prepare(query)
	if err != nil {
		return domain.Turno{}, err
	}
	
	res, err := stmt.Exec(turno.Paciente.Id, turno.Odontologo.Id, turno.FechaYHora, turno.Descripcion)
	if err != nil {
		fmt.Println(err)
		return domain.Turno{}, err
	}
	_, err = res.RowsAffected()
	
	

	if err != nil {
		return domain.Turno{}, err
	}
	idTurno,_ := res.LastInsertId()

	turno.Id=int(idTurno)
	return turno,nil
}

func (tr TurnoRepository) PutTurno(turno domain.Turno) (domain.Turno, error) {
	query := "UPDATE turnos SET paciente_id = ?, odontologo_id = ?, fecha_y_hora = ?, descripcion = ? WHERE id = ?;"
	stmt, err := tr.storage.DB.Prepare(query)
	if err != nil {
		return domain.Turno{}, err
	}
	res, err := stmt.Exec(turno.Paciente.Id, turno.Odontologo.Id, turno.FechaYHora,turno.Descripcion, turno.Id)
	if err != nil {
		return domain.Turno{}, err
	}
	_, err = res.RowsAffected()
	if err != nil {
		return domain.Turno{}, err
	}
	
	return turno, nil
}

func (tr TurnoRepository) postTurnoDNIYMatricula(dni int, matricula string, fechaYHora string, descripcion string) (domain.Turno, error) {
	var turno domain.Turno
	paciente := tr.storage.DB.QueryRow("select * from pacientes where dni = ?", dni)
	err := paciente.Scan(&turno.Paciente.Id, &turno.Paciente.Nombre, &turno.Paciente.Apellido, &turno.Paciente.Domicilio, &turno.Paciente.DNI, &turno.Paciente.FechaDeAlta)
	if err != nil {
		return domain.Turno{}, err
	}
	odontologo := tr.storage.DB.QueryRow("select * from odontologos where matricula = ?", matricula)
	err = odontologo.Scan(&turno.Odontologo.Id, &turno.Odontologo.Apellido, &turno.Odontologo.Nombre, &turno.Odontologo.Matricula)
	if err != nil {
		return domain.Turno{}, err
	}
	query := "insert into turnos (paciente_id, odontologo_id, fecha_y_hora, descripcion) values (?, ?, ?, ?)"
	st, err := tr.storage.DB.Prepare(query)
	if err != nil {
		return domain.Turno{}, err
	}
	res, err := st.Exec(turno.Paciente.Id, turno.Odontologo.Id, fechaYHora, descripcion)
	if err != nil {
		return domain.Turno{}, err
	}

	id, err := res.LastInsertId()
	if err != nil {
		return domain.Turno{}, err
	}
	turno.Id = int(id)
	_, err = res.RowsAffected()
	if err != nil {
		return domain.Turno{}, err
	}
	return turno, nil
}

func (tr TurnoRepository) FindTurnoByDNIPaciente(dni int) (domain.Turno, error) {
	var turno domain.Turno
	row := tr.storage.DB.QueryRow("select turnos.id, turnos.paciente_id, turnos.odontologo_id, fecha_y_hora, descripcion from turnos inner join pacientes on turnos.paciente_id = pacientes.id where pacientes.dni = ?", dni)
	err := row.Scan(&turno.Id, &turno.Paciente.Id, &turno.Odontologo.Id, &turno.FechaYHora, &turno.Descripcion)
	if err != nil {
		return domain.Turno{}, err
	}


	return turno, nil
}