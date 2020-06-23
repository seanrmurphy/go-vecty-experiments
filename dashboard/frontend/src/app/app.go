package app

import (
	"github.com/seanrmurphy/go-vecty-experiments/dashboard/frontend/src/app/model"
	"github.com/seanrmurphy/go-vecty-experiments/dashboard/frontend/src/app/util"
)

var (
	State model.ApplicationData

	Listeners = util.NewListenerRegistry()
)
