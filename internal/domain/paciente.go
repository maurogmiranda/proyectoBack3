package domain


type Paciente struct {
	Id int `json:"id"`
	Nombre      string `json:"nombre,omitempty" `
	Apellido    string `json:"apellido,omitempty" `
	Domicilio   string `json:"domicilio,omitempty" `
	DNI         int `json:"dni,omitempty" `
	FechaDeAlta string `json:"fechaDeAlta,omitempty" `
}