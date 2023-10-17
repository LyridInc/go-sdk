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

// // go test -v -run ^TestHarborListProjects$
// func TestHarborListProjects(t *testing.T) {
// 	harborBaseURL := os.Getenv("HARBOR_BASE_URL")
// 	harborUsername := os.Getenv("HARBOR_USERNAME")
// 	harborPassword := os.Getenv("HARBOR_PASSWORD")
// 	client := model.NewHarborClient(harborBaseURL, harborUsername, harborPassword)
// 	b, err := client.ListProjects()
// 	if err != nil {
// 		t.Fatal(err)
// 	}

// 	t.Log(string(b))
// }

// // go test -v -run ^TestHarborGetProject$
// func TestHarborGetProject(t *testing.T) {
// 	harborBaseURL := os.Getenv("HARBOR_BASE_URL")
// 	harborUsername := os.Getenv("HARBOR_USERNAME")
// 	harborPassword := os.Getenv("HARBOR_PASSWORD")
// 	client := model.NewHarborClient(harborBaseURL, harborUsername, harborPassword)
// 	b, err := client.GetProject("build")
// 	if err != nil {
// 		t.Fatal(err)
// 	}

// 	t.Log(string(b))
// }

// // go test -v -run ^TestHarborGetProjectRepositories$
// func TestHarborGetProjectRepositories(t *testing.T) {
// 	harborBaseURL := os.Getenv("HARBOR_BASE_URL")
// 	harborUsername := os.Getenv("HARBOR_USERNAME")
// 	harborPassword := os.Getenv("HARBOR_PASSWORD")
// 	client := model.NewHarborClient(harborBaseURL, harborUsername, harborPassword)
// 	b, err := client.GetProjectRepositories("build")
// 	if err != nil {
// 		t.Fatal(err)
// 	}

// 	t.Log(string(b))
// }

// // go test -v -run ^TestHarborGetRepository$
// func TestHarborGetRepository(t *testing.T) {
// 	harborBaseURL := os.Getenv("HARBOR_BASE_URL")
// 	harborUsername := os.Getenv("HARBOR_USERNAME")
// 	harborPassword := os.Getenv("HARBOR_PASSWORD")
// 	client := model.NewHarborClient(harborBaseURL, harborUsername, harborPassword)
// 	b, err := client.GetRepository("build", "vega")
// 	if err != nil {
// 		t.Fatal(err)
// 	}

// 	t.Log(string(b))
// }

// // go test -v -run ^TestHarborGetRepositoryArtifacts$
// func TestHarborGetRepositoryArtifacts(t *testing.T) {
// 	harborBaseURL := os.Getenv("HARBOR_BASE_URL")
// 	harborUsername := os.Getenv("HARBOR_USERNAME")
// 	harborPassword := os.Getenv("HARBOR_PASSWORD")
// 	client := model.NewHarborClient(harborBaseURL, harborUsername, harborPassword)
// 	b, err := client.GetRepositoryArtifacts("build", "vega", "develop")
// 	if err != nil {
// 		t.Fatal(err)
// 	}

// 	t.Log(string(b))
// }
