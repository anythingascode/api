package endpoints

import (
	"github.com/anythingascode/api/endpoints/plan"
	"github.com/anythingascode/api/models"
)

var (
	GetEndPoints = models.EndPoints{
		"/planreview": plan.GetChangesInRemotePlan,
		"/drift":      plan.GetDriftsInRemotePlan,
	}
)
