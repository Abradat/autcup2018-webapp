package models

type Team struct {
	Id int `form:"-"`
	TeamName string `form:"team_name,text,team_name:"`
	DepName string `form:"dep_name,text,dep_name:"`
}

type Challenge struct {
	Id int `form:"-"`
	ChallengeName string `form:"challenge_name,text,challenge_name"`
	Ratio float32 `form:"ratio,number,ratio"`
}

type Score struct {
	Id int `form:"-"`
	TeamName string `form:"team_name,text,team_name:"`
	ChallengeName string `form:"challenge_name,text,challenge_name"`
	Score float32 `form:"score,number,score"`
}

type Draw struct {
	Id int
	TeamName string
}

func (a *Team) TableName() string {
	return "teams"
}

func (a *Challenge) TableName() string {
	return "challenges"
}

func (a *Score) TableName() string {
	return "scores"
}

func (a *Draw) TableName() string {
	return "draws"
}

