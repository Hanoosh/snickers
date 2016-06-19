package rest

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/flavioribeiro/snickers/db/memory"
	"github.com/flavioribeiro/snickers/types"
)

func Index(w http.ResponseWriter, r *http.Request) {
	// TODO a great page for API root location
	fmt.Fprint(w, "Snickers")
}

func CreatePreset(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	dbInstance, err := memory.GetDatabase()
	if err != nil {
		fmt.Fprint(w, "error while creating preset")
	}

	var preset types.Preset
	respData, err := ioutil.ReadAll(r.Body)
	err = json.Unmarshal(respData, &preset)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	dbInstance.StorePreset(preset)

	w.WriteHeader(http.StatusOK)
}

func UpdatePreset(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "update preset")
}

func ListPresets(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	dbInstance, err := memory.GetDatabase()
	if err != nil {
		fmt.Fprint(w, "error while creating database")
	}

	var result []string
	for _, preset := range dbInstance.GetPresets() {
		presetJson, _ := json.Marshal(preset)
		result = append(result, string(presetJson))
	}

	fmt.Fprintf(w, "%s", result)
}

func GetPresetDetails(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "get preset details")
}

func CreateJob(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "create job")
}

func ListJobs(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "list jobs")
}

func GetJobDetails(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "get job details")
}