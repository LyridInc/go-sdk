package sdk

import (
	"os"
	"testing"
)

func TestInitialize(t *testing.T) {
	if got := GetInstance().Initialize(os.Getenv("LYRID_ACCESS"), os.Getenv("LYRID_SECRET")); got != nil {
		t.Errorf("Error Initializing() = %s", got)
	}
}

func TestGetUser(t *testing.T) {
	got := GetInstance().GetUserProfile()
	if got == nil {
		t.Errorf("GetUserProfile doesn't return any value")
	}

	t.Log(got)
}

func TestGetAccount(t *testing.T) {
	got := GetInstance().GetAccountProfile()
	if got == nil {
		t.Errorf("GetAccountProfile doesn't return any value")
	}
	for _, account := range got {
		t.Log(account)
	}
}

func TestGetApps(t *testing.T) {
	got := GetInstance().GetApps()
	if got == nil {
		t.Errorf("GetAccountProfile doesn't return any value")
	}
	for _, app := range got {
		t.Log(app)
	}
}

func TestGetModules(t *testing.T) {
	got := GetInstance().GetModules("")
	if got == nil {
		t.Errorf("GetAccountProfile doesn't return any value")
	}
	for _, module := range got {
		t.Log(module)
	}
}

func TestGetRevisions(t *testing.T) {
	got := GetInstance().GetRevisions("", "")
	if got == nil {
		t.Errorf("GetAccountProfile doesn't return any value")
	}
	for _, revision := range got {
		t.Log(revision)
	}
}

func TestGetFunctions(t *testing.T) {
	got := GetInstance().GetFunctions("", "", "")
	if got == nil {
		t.Errorf("GetAccountProfile doesn't return any value")
	}
	for _, functions := range got {
		t.Log(functions)
	}
}
func TestGetFunctionExecute(t *testing.T) {
	got, _ := GetInstance().ExecuteFunction("", "LYR", "{\"InputSample\":\"Hello\"}")
	t.Log(string(got))
}
