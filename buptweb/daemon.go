package buptweb

import (
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func daemonCheckAndLogin() {
	res, err := doCheck()
	if err != nil {
		log.Println(err)
		return
	}
	switch res {
	case s_success:
		log.Println("network is connected")
	case s_failed:
		log.Println("network is disconnected, need login")
		res1, err := doLogin()
		if err != nil {
			log.Println(err)
			return
		}
		switch res1 {
		case s_success:
			log.Println("login success")
		case s_failed:
			log.Println("login failed, invalid username and password")
		case s_unknown:
			log.Println("unknown status")
		}
	}

}

func Daemon() {
	ticker := time.NewTicker(time.Duration(config.Interval) * time.Second)
	defer ticker.Stop()
	termSign := make(chan os.Signal, 1)
	signal.Notify(termSign, syscall.SIGINT, syscall.SIGTERM)
	defer signal.Stop(termSign)

	daemonCheckAndLogin()

	for {
		select {
		case <-ticker.C:
			daemonCheckAndLogin()
		case <-termSign:
			return
		}
	}
}
