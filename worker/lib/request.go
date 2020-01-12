package lib

import (
	"net/http"
	"io/ioutil"
	"log"
	"os"
)

func Request(service, url,  method string) []byte{

    client := &http.Client{}
	req, err := http.NewRequest(method, url, nil)
	if err != nil{
		log.Fatal(err)
	}

    resp, err := client.Do(req)
    if err != nil{
        log.Fatal(err)
	}
	if resp.StatusCode != 200 {
		log.Printf("[ERRO] Não foi possível se comunicar com o serviço %s. Status Code: %d", service, resp.StatusCode)
		os.Exit(1)
	}

	response, err := ioutil.ReadAll(resp.Body)

	if err != nil{
		log.Fatal(err)
	}
	return response
}