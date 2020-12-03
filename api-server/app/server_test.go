package app

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/chauhanr/singlenetes/api-server/scheme"
	store "github.com/chauhanr/singlenetes/api-server/store"
	"github.com/chauhanr/singlenetes/api-server/util"
	"github.com/gorilla/mux"
	"github.com/stretchr/testify/assert"
)

func TestAddPodSuccessHandler(t *testing.T) {
	r := mux.NewRouter()
	mockClient := new(store.MockEtcdCtl)
	pod := scheme.PodV1{}
	pod.Metadata.Uid = "uid-test"

	p, _ := util.EncodeS8Object(pod)
	util.DecodeS8Object([]byte(p), &pod)

	mockClient.On("AddPod", pod).Return(nil)
	s := NewServer(r, mockClient)

	api := r.PathPrefix("/api/v1").Subrouter()
	api.HandleFunc("/default/pod", s.podCreateHandler())

	tserver := httptest.NewServer(s.Router)
	defer tserver.Close()

	tres, err := http.NewRequest("POST", tserver.URL+"/api/v1/default/pod", strings.NewReader(p))
	assert.NoError(t, err)
	_, err = http.DefaultClient.Do(tres)
	assert.NoError(t, err)
	/*	actualBody, err := ioutil.ReadAll(res.Body)
		assert.NoError(t, err)
	*/
}
