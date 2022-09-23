package domain

type Odontologo struct{
	Id int `json:"id"`
	Nombre string `json:"nombre,omitempty"`
	Apellido string `json:"apellido,omitempty"`
	Matricula string `json:"matricula,omitempty"`
}