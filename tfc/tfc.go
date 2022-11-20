package tfc

import (
	"context"
	"errors"
	"github.com/hashicorp/go-tfe"
)

type Client struct {
	TfeClient *tfe.Client
	Config    *Config
}

type Config struct {
	Workspace string `json:"workspace"`
	RunId     string `json:"runid"`
	PlanFile  string `json:"plan-file"`
	Token     string `json:"token"`
	Address   string `json:"address"`
	Org       string `json:"org"`
}

var ctx = context.Background()

func NewClient(con *Config) (*Client, error) {
	config := &tfe.Config{
		Token:   con.Token,
		Address: "https://app.terraform.io",
	}
	if config.Token == "" || config.Address == "" {
		return nil, errors.New("tfe token or address should not be empty.")
	}
	nc, err := tfe.NewClient(config)
	if err != nil {
		return nil, err
	}
	return &Client{
		TfeClient: nc,
		Config:    con,
	}, nil
}

func (c *Client) GetRunId() (string, error) {
	workspace, err := c.TfeClient.Workspaces.Read(ctx, c.Config.Org, c.Config.Workspace)
	if err != nil {
		return "", err
	}
	return workspace.CurrentRun.ID, nil

}
func (c *Client) GetRun() (*tfe.Run, error) {
	var runId string
	var err error
	if c.Config.RunId == "" {
		runId, err = c.GetRunId()
		if err != nil {
			return nil, err
		}
	} else {
		runId = c.Config.RunId
	}
	run, err := c.TfeClient.Runs.Read(ctx, runId)
	if err != nil {
		return nil, err
	}
	return run, nil
}

func (c *Client) GetPlanId() (string, error) {
	run, err := c.GetRun()
	if err != nil {
		return "", err
	}
	return run.Plan.ID, nil
}

func (c *Client) GetPlan() ([]byte, error) {
	planId, err := c.GetPlanId()
	if err != nil {
		return nil, err
	}
	jsonPlan, err := c.TfeClient.Plans.ReadJSONOutput(ctx, planId)
	if err != nil {
		return nil, err
	}
	return jsonPlan, nil
}
