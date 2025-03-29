package models

// Match representa un partido de fútbol con sus eventos principales
type Match struct {
	ID          uint   `json:"id" gorm:"primaryKey"` // Identificador único del partido (clave primaria)
	HomeTeam    string `json:"homeTeam"`             // Nombre del equipo local
	AwayTeam    string `json:"awayTeam"`             // Nombre del equipo visitante
	MatchDate   string `json:"matchDate"`            // Fecha del partido (en formato string)
	TotalGoals  int    `json:"totalGoals"`           // Goles totales registrados en el partido
	YellowCards int    `json:"yellowCards"`          // Cantidad de tarjetas amarillas
	RedCards    int    `json:"redCards"`             // Cantidad de tarjetas rojas
	ExtraTime   bool   `json:"extraTime"`            // Indica si el partido tuvo tiempo extra
}
