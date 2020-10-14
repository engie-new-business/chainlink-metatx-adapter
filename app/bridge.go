package app

import (
	"bytes"
	"io/ioutil"
	"net/http"

	"github.com/linkpoolio/bridges"
	"github.com/sirupsen/logrus"
)

func Handler(w http.ResponseWriter, r *http.Request) {
	rockside, err := NewRelayerAdapterFromEnv()
	if err != nil {
		logrus.Fatal(err)
	}

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		logrus.WithError(err).Error("could not read body")
	}
	logrus.WithField("body", string(body)).Info("request body")

	r.Body = ioutil.NopCloser(bytes.NewReader(body))
	bridges.NewServer(rockside).Handler(w, r)
}
