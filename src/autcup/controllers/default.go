package controllers

import (
	"github.com/astaxie/beego"
	"strconv"
	"github.com/astaxie/beego/orm"
	models "autcup/models"
	"github.com/astaxie/beego/validation"
	"fmt"

	"sort"
)



type MainController struct {
	beego.Controller
}

type ManageController struct {
	beego.Controller
}

func (c *MainController) MainPage() {

	c.Data["PageName"] = "Main Page"
	c.TplName = "index.html"

}

func (c *MainController) TeamPage() {

	o := orm.NewOrm()
	o.Using("default")

	var teams []*models.Team
	num, err := o.QueryTable("teams").All(&teams)


	if err != orm.ErrNoRows && num > 0 {

		c.Data["records"] = teams
	}

	c.Data["PageName"] = "Teams Page"

	c.TplName = "teams.html"

}

func (c *MainController) AddTeamPage() {
	c.Data["PageName"] = "Add Team Page"
	c.TplName = "add-team.html"
}

func (c *MainController) ChallengePage() {

	o := orm.NewOrm()
	o.Using("default")

	var challenges []*models.Challenge
	num, err := o.QueryTable("challenges").All(&challenges)


	if err != orm.ErrNoRows && num > 0 {

		c.Data["records"] = challenges
	}

	c.TplName = "challenges.html"

}

func (c *MainController) AddChallengePage() {
	c.TplName = "add-challenge.html"

}

func isInChallenges (challenge string, challenges []string) bool {

	flag := false

	for _, myChallenge := range challenges {
		if(myChallenge == challenge) {
			flag = true
		}
	}
	return flag

}

func (c *MainController) ScorePage() {

	o := orm.NewOrm()
	o.Using("default")

	var scores []*models.Score
	num, err := o.QueryTable("scores").All(&scores)


	if err != orm.ErrNoRows && num > 0 {

		sort.Slice(scores, func(i, j int) bool {return scores[i].Score > scores[j].Score})
		c.Data["records"] = scores
	}

	var challengesTotal []*models.Challenge
	o.QueryTable("challenges").All(&challengesTotal)

	//var challenges []string


	//for _, challengeTotal := range challengesTotal {

	//	challenges = append(challenges, challengeTotal.ChallengeName)
	//}


	/*for _, score := range scores {

		if(isInChallenges(score.ChallengeName, challenges) == false) {
			challenges = append(challenges, score.ChallengeName)
		}
	}
	*/

	c.Data["challenges"] = challengesTotal

	c.TplName = "scores.html"

}

func (c *MainController) AddScorePage() {
	c.TplName = "add-score.html"

}

func (c *MainController) ResultPage() {

	o := orm.NewOrm()
	o.Using("default")

	var teamsQuery []*models.Team
	var teams []string
	teamsQueryNumber, teamsQueryErr := o.QueryTable("teams").All(&teamsQuery)


	if teamsQueryErr != orm.ErrNoRows && teamsQueryNumber > 0 {

		for _, teamQuery := range teamsQuery {
			teams = append(teams, teamQuery.TeamName)
		}
	}


	var challengesQuery []*models.Challenge
	challengesRatios := make(map[string] float32)

	challengesQueryNumber, challengesQueryErr := o.QueryTable("challenges").All(&challengesQuery)

	if challengesQueryErr != orm.ErrNoRows && challengesQueryNumber > 0 {

		for _, challengeQuery := range challengesQuery {
			challengesRatios[challengeQuery.ChallengeName] = challengeQuery.Ratio
		}
	}


	var scoresQuery []*models.Score
	//var teamsTotalScore map[string]float32
	teamsTotalScore := make(map[string] float32)
	teamsTotalScoreWithRatio := make(map[string] float32)
	var sum, sumWithRatio float32

	scoresQueryNumber, scoresQueryErr := o.QueryTable("scores").All(&scoresQuery)

	if scoresQueryErr != orm.ErrNoRows && scoresQueryNumber > 0 {

		for _, team := range teams {

			//var sum float32 = 0.0
			sum = 0.0
			sumWithRatio = 0.0
			for _, scoreQuery := range scoresQuery {

				if(team == scoreQuery.TeamName) {

					sum +=  scoreQuery.Score
					sumWithRatio += (scoreQuery.Score * challengesRatios[scoreQuery.ChallengeName])
				}

			}

			teamsTotalScore[team] = sum
			teamsTotalScoreWithRatio[team] = sumWithRatio

		}
	}


	c.Data["scores"] = scoresQuery
	c.Data["teams"] = teams
	c.Data["teamsTotalScore"] = teamsTotalScore
	c.Data["teamsTotalScoreWithRatio"] = teamsTotalScoreWithRatio
	c.Data["challenges"] = challengesRatios


	c.TplName = "results.html"

}

func (c *MainController) PortfolioPage() {
	c.TplName = "portfolio.html"
}

func (c *MainController) AboutPage() {
	c.TplName = "about.html"
}

func (c *MainController) ContactPage() {
	c.TplName = "contact.html"
}


func (manage *ManageController) DeleteTeam() {
	o := orm.NewOrm()
	o.Using("default")

	teamId := manage.GetString("id")
	id, _:= strconv.Atoi(teamId)
	team := models.Team{Id : id}

	if err := manage.ParseForm(&team); err != nil {
		beego.Error("Couldn't parse the form. Reason: ", err)
	} else {

		if manage.Ctx.Input.Method() == "POST" {
			valid := validation.Validation{}
			isValid, _ := valid.Valid(team)
			if !isValid {
				manage.Data["Errors"] = valid.ErrorsMap
				beego.Error("Form didn't validate.")
			} else {

				id, err := o.Delete(&team)
				if err == nil {
					msg := fmt.Sprintf("Team deleted with id:", id)
					beego.Debug(msg)
				} else {
					msg := fmt.Sprintf("Couldn't delete new team. Reason: ", err)
					beego.Debug(msg)
				}
			}
		}
	}

	manage.Redirect("/teams", 302)


}

func (manage *ManageController) InsertTeam() {

	o := orm.NewOrm()
	o.Using("default")
	team := models.Team{}

	if err := manage.ParseForm(&team); err != nil {
		beego.Error("Couldn't parse the form. Reason: ", err)
	} else {

		if manage.Ctx.Input.Method() == "POST" {
			valid := validation.Validation{}
			isValid, _ := valid.Valid(team)
			if !isValid {
				manage.Data["Errors"] = valid.ErrorsMap
				beego.Error("Form didn't validate.")
			} else {

				id, err := o.Insert(&team)
				if err == nil {
					msg := fmt.Sprintf("Team inserted with id:", id)
					beego.Debug(msg)
				} else {
					msg := fmt.Sprintf("Couldn't insert new team. Reason: ", err)
					beego.Debug(msg)
				}
			}
		}
	}
	manage.Redirect("/teams", 302)


}

func (manage *ManageController) RetrieveTeam() {

	o := orm.NewOrm()
	o.Using("default")
	teamId, _ := strconv.Atoi(manage.Ctx.Input.Param(":id"))

	team := models.Team{Id : teamId}
	err := o.Read(&team)

	if err == orm.ErrNoRows {
		fmt.Println("No result found.")
	} else if err == orm.ErrMissPK {
		fmt.Println("No primary key found.")
	} else {
		fmt.Println(team.TeamName, team.DepName)
	}

}


// Challenges Controllers

func (manage *ManageController) InsertChallenge() {

	o := orm.NewOrm()
	o.Using("default")
	challenge := models.Challenge{}

	if err := manage.ParseForm(&challenge); err != nil {
		beego.Error("Couldn't parse the form. Reason: ", err)
	} else {

		if manage.Ctx.Input.Method() == "POST" {
			valid := validation.Validation{}
			isValid, _ := valid.Valid(challenge)
			if !isValid {
				manage.Data["Errors"] = valid.ErrorsMap
				beego.Error("Form didn't validate.")
			} else {

				id, err := o.Insert(&challenge)
				if err == nil {
					msg := fmt.Sprintf("Challenge inserted with id:", id)
					beego.Debug(msg)
				} else {
					msg := fmt.Sprintf("Couldn't insert new challenge. Reason: ", err)
					beego.Debug(msg)
				}
			}
		}

	}


	manage.Redirect("/challenges", 302)


}

func (manage *ManageController) DeleteChallenge() {
	o := orm.NewOrm()
	o.Using("default")

	challengeId := manage.GetString("id")
	id, _:= strconv.Atoi(challengeId)
	challenge := models.Challenge{Id : id}

	if err := manage.ParseForm(&challenge); err != nil {
		beego.Error("Couldn't parse the form. Reason: ", err)
	} else {

		if manage.Ctx.Input.Method() == "POST" {
			valid := validation.Validation{}
			isValid, _ := valid.Valid(challenge)
			if !isValid {
				manage.Data["Errors"] = valid.ErrorsMap
				beego.Error("Form didn't validate.")
			} else {

				id, err := o.Delete(&challenge)
				if err == nil {
					msg := fmt.Sprintf("Challenge deleted with id : ", id)
					beego.Debug(msg)
				} else {
					msg := fmt.Sprintf("Couldn't delete challenge with reason: ", err)
					beego.Debug(msg)
				}
			}
		}
	}

	manage.Redirect("/challenges", 302)


}

// Score Controller

func (manage *ManageController) InsertScore() {

	o := orm.NewOrm()
	o.Using("default")
	score := models.Score{}

	if err := manage.ParseForm(&score); err != nil {
		beego.Error("Couldn't parse the form. Reason: ", err)
	} else {

		if manage.Ctx.Input.Method() == "POST" {
			valid := validation.Validation{}
			isValid, _ := valid.Valid(score)
			if !isValid {
				manage.Data["Errors"] = valid.ErrorsMap
				beego.Error("Form didn't validate.")
			} else {

				id, err := o.Insert(&score)
				if err == nil {
					msg := fmt.Sprintf("Score inserted with id:", id)
					beego.Debug(msg)
				} else {
					msg := fmt.Sprintf("Couldn't insert new score. Reason: ", err)
					beego.Debug(msg)
				}
			}
		}

	}


	manage.Redirect("/scores", 302)


}


func (manage *ManageController) DeleteScore() {
	o := orm.NewOrm()
	o.Using("default")

	scoreId := manage.GetString("id")
	id, _:= strconv.Atoi(scoreId)
	score := models.Score{Id : id}

	if err := manage.ParseForm(&score); err != nil {
		beego.Error("Couldn't parse the form. Reason: ", err)
	} else {

		if manage.Ctx.Input.Method() == "POST" {
			valid := validation.Validation{}
			isValid, _ := valid.Valid(score)
			if !isValid {
				manage.Data["Errors"] = valid.ErrorsMap
				beego.Error("Form didn't validate.")
			} else {

				id, err := o.Delete(&score)
				if err == nil {
					msg := fmt.Sprintf("Score deleted with id : ", id)
					beego.Debug(msg)
				} else {
					msg := fmt.Sprintf("Couldn't delete score with reason: ", err)
					beego.Debug(msg)
				}
			}
		}
	}

	manage.Redirect("/scores", 302)


}