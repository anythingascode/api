package plan

import (
	"fmt"
	"github.com/anythingascode/api/tfc"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"strings"
	"time"
)

func PostChangesInLocalPlan(c *gin.Context) {
	name := c.PostForm("name")
	email := c.PostForm("email")
	localPlanEpUsagesTracker(name, email, "/localplanreview")
	if name == "" || email == "" {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "name and email field is required."})
		return
	}
	// Source
	file, err := c.FormFile("file")
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": fmt.Sprintf("%v", err)})
		return
	}

	filename := "endpoints/plan/plans_to_be_processed/" + filepath.Base(strings.ReplaceAll(time.Now().Format(time.UnixDate), " ", "-")+"-"+file.Filename)
	if err := c.SaveUploadedFile(file, filename); err != nil {
		c.String(http.StatusBadRequest, "upload file err: %s", err.Error())
		return
	} else {
		log.Printf("File %s uploaded successfully with fields name=%s and email=%s.", file.Filename, name, email)
	}
	con := &tfc.Config{
		PlanFile: filename,
	}
	//defer os.Remove(filename)
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

func PostDriftsInLocalPlan(c *gin.Context) {
	name := c.PostForm("name")
	email := c.PostForm("email")
	localPlanEpUsagesTracker(name, email, "/localplanreview")
	if name == "" || email == "" {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "name and email field is required."})
		return
	}
	// Source
	file, err := c.FormFile("file")
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": fmt.Sprintf("%v", err)})
		return
	}

	filename := "endpoints/plan/plans_to_be_processed/" + filepath.Base(strings.ReplaceAll(time.Now().Format(time.UnixDate), " ", "-")+"-"+file.Filename)

	if err := c.SaveUploadedFile(file, filename); err != nil {
		c.String(http.StatusBadRequest, "upload file err: %s", err.Error())
		return
	} else {
		log.Printf("File %s uploaded successfully with fields name=%s and email=%s.", file.Filename, name, email)
	}
	con := &tfc.Config{
		PlanFile: filename,
	}
	defer os.Remove(filename)
	changes, err := tfc.RDrifts(con)
	if err != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": fmt.Sprintf("%v", err)})
		return
	}
	if changes != "" {
		c.String(http.StatusOK, string(changes))
		return
	} else {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": fmt.Sprintf("%v", err)})
	}

}
