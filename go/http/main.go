package main

import (
	"github.com/hashicorp/consul/api"
	"github.com/sirupsen/logrus"
)

func main() {
	config := api.DefaultConfig()

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

