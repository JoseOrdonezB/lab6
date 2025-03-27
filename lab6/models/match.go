package models

type Match struct {
    ID        uint   `json:"id" gorm:"primaryKey"`
    HomeTeam  string `json:"homeTeam"`
    AwayTeam  string `json:"awayTeam"`
    MatchDate string `json:"matchDate"`
}