package endpoints

import (
	"github.com/anythingascode/api/endpoints/plan"
	"github.com/anythingascode/api/models"
)

var (
	PostEndPoints = models.EndPoints{
		"/localplanreview": plan.PostChangesInLocalPlan,
		"/localplandrift":  plan.PostDriftsInLocalPlan,
	}
)
