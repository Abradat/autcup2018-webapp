package routers

import (
	"autcup/controllers"
	"github.com/astaxie/beego"
)

func init() {

    beego.Router("/", &controllers.MainController{}, "get:MainPage")
    beego.Router("/index", &controllers.MainController{}, "get:MainPage" )

    beego.Router("/teams", &controllers.MainController{}, "get:TeamPage")
	beego.Router("/add-team", &controllers.MainController{}, "get:AddTeamPage")

	beego.Router("/challenges", &controllers.MainController{}, "get:ChallengePage")
	beego.Router("/add-challenge", &controllers.MainController{}, "get:AddChallengePage")

	beego.Router("/scores", &controllers.MainController{}, "get:ScorePage")
	beego.Router("/add-score", &controllers.MainController{}, "get:AddScorePage")

	beego.Router("/results", &controllers.MainController{}, "get:ResultPage")

	beego.Router("/about", &controllers.MainController{}, "get:AboutPage" )
	beego.Router("/contact", &controllers.MainController{}, "get:ContactPage" )
	beego.Router("/portfolio", &controllers.MainController{}, "get:PortfolioPage" )


	//Team table management
	beego.Router("/manager/team/delete", &controllers.ManageController{}, "*:DeleteTeam")
	beego.Router("/manager/team/insert", &controllers.ManageController{}, "*:InsertTeam")
	//beego.Router("/manager/team/retrieve/:id([0-9]+)", &controllers.ManageController{}, "*:RetrieveTeam")

	// Challenge table management
	beego.Router("/manager/challenge/delete", &controllers.ManageController{}, "*:DeleteChallenge")
	beego.Router("/manager/challenge/insert", &controllers.ManageController{}, "*:InsertChallenge")

	//Score table management
	beego.Router("/manager/score/delete", &controllers.ManageController{}, "*:DeleteScore")
	beego.Router("/manager/score/insert", &controllers.ManageController{}, "*:InsertScore")
	//beego.Router("/manager/score/delete", &controllers.ManageController{}, "*:DeleteScore")
}
