package domain


type Turno struct {
	Id int `json:"id"`
	Paciente   Paciente   `json:"paciente,omitempty"`
	Odontologo Odontologo `json:"odontologo,omitempty"`
	FechaYHora string `json:"fechaYHora,omitempty"`
	Descripcion string     `json:"descripcion,omitempty"`
}