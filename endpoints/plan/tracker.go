package plan

import (
	"encoding/json"
	"github.com/anythingascode/api/models"
	"io/ioutil"
	"log"
	"time"
)

func remotePlanEpUsagesTracker(in models.RemotePlanInput, ep string) {
	data, err := ioutil.ReadFile("usages/usages.json")
	if err != nil {
		log.Println(err)
	}
	var usages []models.User
	err = json.Unmarshal(data, &usages)
	if err != nil {
		log.Println(err)
	}
	usages = append(usages, models.User{
		EndPoint:    ep,
		RequestType: "GET",
		Name:        in.Name,
		Email:       in.Email,
		Time:        time.Now(),
	})
	bd, err := json.Marshal(usages)
	if err != nil {
		log.Println(err)
	}
	err = ioutil.WriteFile("usages/usages.json", bd, 0644)
	if err != nil {
		log.Println(err)
	}
}

func localPlanEpUsagesTracker(name, email, ep string) {
	data, err := ioutil.ReadFile("usages/usages.json")
	if err != nil {
		log.Println(err)
	}
	var usages []models.User
	err = json.Unmarshal(data, &usages)
	if err != nil {
		log.Println(err)
	}
	usages = append(usages, models.User{
		EndPoint:    ep,
		RequestType: "POST",
		Name:        name,
		Email:       email,
		Time:        time.Now(),
	})
	bd, err := json.Marshal(usages)
	if err != nil {
		log.Println(err)
	}
	err = ioutil.WriteFile("usages/usages.json", bd, 0644)
	if err != nil {
		log.Println(err)
	}
}
