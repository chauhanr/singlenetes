package app

import (
	"testing"

	"github.com/chauhanr/singlenetes/api-server/scheme"
	store "github.com/chauhanr/singlenetes/api-server/store"
	"github.com/gorilla/mux"
)

func TestAddPodSuccessHandler(t *testing.T) {
	r := mux.NewRouter()
	mockClient := new(store.MockEtcdCtl)
	pod := scheme.PodV1{}
	mockClient.On("AddPod", pod).Return(nil)
	s := NewServer(r, &mockClient)

	api := r.PathPrefix("/api/v1").Subrouter()
	api.HandleFunc("/default/pod", s.podCreateHandler())

}
