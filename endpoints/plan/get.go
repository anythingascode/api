package plan

import (
	"fmt"
	"github.com/anythingascode/api/models"
	"github.com/anythingascode/api/tfc"
	"github.com/gin-gonic/gin"
	"net/http"
)

func GetChangesInRemotePlan(c *gin.Context) {
	var in models.RemotePlanInput
	if err := c.BindJSON(&in); err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": fmt.Sprintf("%v", err)})
		return
	}
	remotePlanEpUsagesTracker(in, "/planreview")
	if in.RunId == "" && in.Workspace == "" {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": fmt.Sprintf("runid or workspace name required.")})
		return
	}
	con := &tfc.Config{
		Workspace: in.Workspace,
		RunId:     in.RunId,
		Token:     in.Token,
	}
	changes, err := tfc.RChanges(con)
	if err != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": fmt.Sprintf("%v", err)})
		return
	}
	if changes != "" {
		c.String(http.StatusOK, string(changes))
		return
	} else {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": fmt.Sprintf("%v", err)})
		return
	}
}

func GetDriftsInRemotePlan(c *gin.Context) {
	var in models.RemotePlanInput
	if err := c.BindJSON(&in); err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": fmt.Sprintf("%v", err)})
		return
	}
	remotePlanEpUsagesTracker(in, "/drift")
	if in.RunId == "" && in.Workspace == "" {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": fmt.Sprintf("runid or workspace name required.")})
		return
	}
	con := &tfc.Config{
		Workspace: in.Workspace,
		RunId:     in.RunId,
		Token:     in.Token,
	}
	drifts, err := tfc.RDrifts(con)
	if err != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": fmt.Sprintf("%v", err)})
		return
	}
	if drifts != "" {
		c.String(http.StatusOK, string(drifts))
		return
	} else {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": fmt.Sprintf("%v", err)})
		return
	}

}
