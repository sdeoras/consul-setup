package main

import (
	"net/http"
	"os"
	"path/filepath"

	"github.com/hashicorp/consul/api"
	"github.com/sirupsen/logrus"
)

func main() {
	execPath, err := os.Executable()
	if err != nil {
		logrus.Fatal(err)
	}

	execPath, _ = filepath.Split(execPath)
	execPath = filepath.Join(execPath, "ssl")

	tlsConfig := &api.TLSConfig{
		Address:            "server.ip.1",
		CAFile:             filepath.Join(execPath, "ca.cert"),
		CertFile:           filepath.Join(execPath, "consul.cert"),
		KeyFile:            filepath.Join(execPath, "consul.key"),
		InsecureSkipVerify: true,
	}

	consulTLSConfig, err := api.SetupTLSConfig(tlsConfig)
	if err != nil {
		logrus.Fatal(err)
	}

	config := api.DefaultConfig()
	config.Address = "127.0.0.1:8080"
	config.Scheme = "https"
	config.HttpClient = new(http.Client)

	config.HttpClient.Transport = &http.Transport{
		TLSClientConfig: consulTLSConfig,
	}

	logrus.Info("user app initiating connection with:", config.Address)
	logrus.Info("use ctrl-c to exit")

	// Get a new client
	client, err := api.NewClient(config)
	if err != nil {
		logrus.Fatal(err)
	}

	// Get a handle to the KV API
	kv := client.KV()

	for i := 0; ; i++ {
		// PUT a new KV pair
		p := &api.KVPair{Key: "foo", Value: []byte("test")}
		_, err = kv.Put(p, nil)
		if err != nil {
			logrus.Error(err)
		}

		// Lookup the pair
		_, _, err := kv.Get("foo", nil)
		if err != nil {
			logrus.Error(err)
		}

		if i%100 == 0 {
			logrus.Info("ran 100 more put-get commands")
		}
	}
}
