package buptweb

import (
	"log"
)

func Check() {
	state, err := doCheck()
	if err != nil {
		log.Fatalln(err)
	}
	switch state {
	case s_success:
		log.Println("network is connected")
	case s_failed:
		log.Println("network is disconnected")
	}
}

func doCheck() (Status, error) {
	_, status_code, err := HttpGet(CONN_TEST_URL, nil, HEADERS, false)
	if err != nil {
		return s_error, err
	}
	if status_code == CONN_SUCCESS_CODE {
		return s_success, nil
	} else {
		return s_failed, nil
	}
}
