package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	eureka "github.com/xuanbo/eureka-client"
)

func main() {

	client := eureka.NewClient(&eureka.Config{
		DefaultZone: "http://localhost:8761/eureka/",
		App:         "go-example",
		// HostName:                     "localhost",
		Port:                         10000,
		RenewalIntervalInSecs:        10,
		RegistryFetchIntervalSeconds: 15,
		DurationInSecs:               30,
		Metadata: map[string]interface{}{
			"VERSION":              "0.1.0",
			"NODE_GROUP_ID":        0,
			"PRODUCT_CODE":         "DEFAULT",
			"PRODUCT_VERSION_CODE": "DEFAULT",
			"PRODUCT_ENV_CODE":     "DEFAULT",
			"SERVICE_VERSION_CODE": "DEFAULT",
		},
	})

	client.Start()

	http.HandleFunc("/v1/services", func(writer http.ResponseWriter, request *http.Request) {
		fmt.Println("/v1/services")
		apps := client.Applications
		b, _ := json.Marshal(apps)
		_, _ = writer.Write(b)
	})

	if err := http.ListenAndServe(":10000", nil); err != nil {
		fmt.Println(err)
	}
}
