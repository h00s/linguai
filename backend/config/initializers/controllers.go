package initializers

import (
	"github.com/go-raptor/raptor/v2"
	"github.com/h00s/linguai/app/controllers"
)

func Controllers() raptor.Controllers {
	return raptor.Controllers{
		&controllers.GrammarController{},
		&controllers.SPAController{},
	}
}
