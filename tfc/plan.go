package tfc

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
)

type Plan struct {
	FormatVersion    string          `json:"format_version"`
	TerraformVersion string          `json:"terraform_version"`
	Variables        interface{}     `json:"variables"`
	ResourceDrift    []ResourceState `json:"resource_drift"`
	ResourceChanges  []ResourceState `json:"resource_changes"`
	PriorState       struct {
		FormatVersion    string      `json:"format_version"`
		TerraformVersion string      `json:"terraform_version"`
		Values           interface{} `json:"values"`
	} `json:"prior_state"`
}

type ResourceState struct {
	Address      string  `json:"address"`
	Mode         string  `json:"mode"`
	Type         string  `json:"type"`
	Name         string  `json:"name"`
	ProviderName string  `json:"provider_name"`
	Change       Changes `json:"change"`
}
type Changes struct {
	Actions []string `json:"actions"`
	Delta
	BeforeSensitive interface{} `json:"before_sensitive"`
	AfterSensitive  interface{} `json:"after_sensitive"`
}
type Delta struct {
	Before interface{} `json:"before"`
	After  interface{} `json:"after"`
}
type Diff struct {
	Delta
}
type Drift struct {
	Delta
}
type resourceChanges map[string][]Diff
type resourceDrifts map[string][]Drift

func plan2BeProcessed(con *Config) ([]byte, error) {
	var jsonPlan []byte
	var err error
	if con.PlanFile != "" {
		jsonPlan, err = ioutil.ReadFile(con.PlanFile)
		if err != nil {
			return nil, errors.New(fmt.Sprintf("error during reading plan file :%v", err))
		}
	}
	if con.PlanFile == "" && con.RunId != "" || con.Workspace != "" {
		client, err := NewClient(con)
		if err != nil {
			return nil, errors.New(fmt.Sprintf("error creating tfe client :%v", err))
		}
		jsonPlan, err = client.GetPlan()
		if err != nil {
			return nil, errors.New(fmt.Sprintf("error getting plan :%v", err))
		}

	}
	return jsonPlan, nil
}

func RChanges(con *Config) (string, error) {
	jsonPlan, err := plan2BeProcessed(con)
	if err != nil {
		return "", err
	}
	var plan Plan
	json.Unmarshal(jsonPlan, &plan)
	var m = make(resourceChanges)
	for _, v := range plan.ResourceChanges {
		var d []Diff
		before := toMap(v.Change.Before)
		after := toMap(v.Change.After)
		if before == nil {
			d = append(d, Diff{
				Delta{
					After: "New Deployment",
				},
			})
		}
		if after == nil {
			d = append(d, Diff{
				Delta: Delta{
					Before: func() interface{} {
						if _, ok := before["id"]; ok {
							return before["id"]
						}
						return nil
					}(),
					After: "Destroy",
				},
			})
		}
		if before != nil && after != nil {
			for key, _ := range before {
				if !isEqual(key, before, after) {
					d = append(d, Diff{
						Delta: Delta{
							After: map[string]interface{}{
								key: after[key],
							},
							Before: map[string]interface{}{
								key: before[key],
							},
						},
					})
				}
			}
		}
		if len(d) > 0 {
			m[v.Type] = append(m[v.Type], d...)
		}
	}
	jdata, err := json.Marshal(m)
	if err != nil {
		return "", err
	}
	return string(jdata), err
}

func RDrifts(con *Config) (string, error) {
	jsonPlan, err := plan2BeProcessed(con)
	if err != nil {
		return "", err
	}
	var plan Plan
	json.Unmarshal(jsonPlan, &plan)
	var m = make(resourceDrifts)
	for _, v := range plan.ResourceDrift {
		var d []Drift
		before := toMap(v.Change.Before)
		after := toMap(v.Change.After)
		if before != nil && after != nil {
			for key, _ := range before {
				if !isEqual(key, before, after) {
					d = append(d, Drift{
						Delta: Delta{
							After: map[string]interface{}{
								key: after[key],
							},
							Before: map[string]interface{}{
								key: before[key],
							},
						},
					})
				}
			}
		}
		if len(d) > 0 {
			m[v.Type] = append(m[v.Type], d...)
		}
	}
	jdata, err := json.Marshal(m)
	if err != nil {
		return "", err
	}
	return string(jdata), nil
}
