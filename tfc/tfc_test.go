package tfc

import (
	"os"
	"testing"
)

func TestTfcClient(t *testing.T) {
	con := &Config{
		Token:     os.Getenv("TFE_TOKEN"),
		Workspace: "terraform",
		Org:       "anythingascode",
		Address:   "https://app.terraform.io",
	}
	t.Run("NewClient", func(t *testing.T) {
		_, err := NewClient(con)
		if err != nil {
			t.Fatalf("failed to create tfe client: %s", err)

		}
	})
	t.Run("GetRunId", func(t *testing.T) {
		client, err := NewClient(con)
		if err != nil {
			t.Fatalf("%s", err)

		}
		_, err = client.GetRunId()
		if err != nil {
			t.Fatalf("%s", err)

		}
	})
	t.Run("GetRun", func(t *testing.T) {
		client, err := NewClient(con)
		if err != nil {
			t.Fatalf("failed to create tfe client: %s", err)

		}
		_, err = client.GetRun()
		if err != nil {
			t.Fatalf("%s", err)

		}
	})
	t.Run("GetPlanId", func(t *testing.T) {
		client, err := NewClient(con)
		if err != nil {
			t.Fatalf("failed to create tfe client: %s", err)

		}
		_, err = client.GetPlanId()
		if err != nil {
			t.Fatalf("%s", err)

		}
	})
	t.Run("GetPlan", func(t *testing.T) {
		client, err := NewClient(con)
		if err != nil {
			t.Fatalf("failed to create tfe client: %s", err)

		}
		_, err = client.GetPlan()
		if err != nil {
			t.Fatalf("%s", err)

		}
	})
}
