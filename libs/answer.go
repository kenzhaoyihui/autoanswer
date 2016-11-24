package utils

import (
	"bytes"
	"fmt"
	"os"
	"runtime"
	"time"

	k "github.com/micmonay/keybd_event"
)

const setupLogFile = "/tmp/engine-setup.log"
const pollFileInterval = 8
const bytesToReadFromFile int64 = 120

// possible questions
var questionPatterns = []string{"(Yes, No)"}

func pressKeySlowly(kb *k.KeyBonding, keys []int) {
	for _, key := range keys {
		if key < 0 {
			kb.HasSHIFT(true)
			key = key / -100
		} else {
			kb.HasSHIFT(false)
		}
		kb.SetKeys(key)
		kb.Launching()
		time.Sleep(500 * time.Millisecond)
	}
}

func sendKeys(key string) {
	kd, err := k.NewKeyBonding()
	if err != nil {
		panic(err)
	}

	if runtime.GOOS == "linux" {
		time.Sleep(2 * time.Second)
	}

	pressKeySlowly(&kd, StringsToByteAry(key))
}

// GetLastLineOfFile should
func grepPatternFromFile(fileName string) {
	c := time.Tick(pollFileInterval * time.Second)

	for _ = range c {
		v := getLastLineOfFile(fileName)

		switch {
		//case bytes.Contains(v, []byte("Configure Engine on this host")):
			//sendKeys("\n")
		case bytes.Contains(v, []byte("Configure Image I/O Proxy on this host?")):
			sendKeys("yes\n")
		//case bytes.Contains(v, []byte("Configure WebSocket Proxy on this host")):
			//sendKeys("yes\n")
		case bytes.Contains(v, []byte("Configure Data Warehouse on this host")):
			sendKeys("yes\n")
		case bytes.Contains(v, []byte("Configure VM Console Proxy on this host")):
			sendKeys("yes\n")
		case bytes.Contains(v, []byte("DNS name of this server")):
			sendKeys("\n")
		case bytes.Contains(v, []byte("configure the firewall")):
			sendKeys("\n")
		case bytes.Contains(v, []byte("the DWH database located")):
			sendKeys("\n")
		case bytes.Contains(v, []byte("the Engine database located")):
			sendKeys("\n")
		case bytes.Contains(v, []byte("(Automatic, Manual)")):
			sendKeys("\n")
		case bytes.Contains(v, []byte("admin password")):
			sendKeys("redhat\n")
		case bytes.Contains(v, []byte("Use weak password?")):
			sendKeys("yes\n")
		case bytes.Contains(v, []byte("Application mode")):
			sendKeys("\n")
		case bytes.Contains(v, []byte("SAN wipe after delete")):
			sendKeys("\n")
		case bytes.Contains(v, []byte("name for certificate")):
			sendKeys("\n")
		case bytes.Contains(v, []byte("the default page of the web server")):
			sendKeys("\n")
		case bytes.Contains(v, []byte("perform that manually?")):
			sendKeys("\n")
		//case bytes.Contains(v, []byte("Local ISO domain path")):
			//sendKeys("\n")
		//case bytes.Contains(v, []byte("Local ISO domain ACL")):
			//sendKeys("acl\n")
		//case bytes.Contains(v, []byte("Local ISO domain name")):
			//sendKeys("\n")
		case bytes.Contains(v, []byte("(1, 2)[1]")):
			sendKeys("\n")
		//case bytes.Contains(v, []byte("(OK, Cancel)")):
			//sendKeys("\n")
		}

		if bytes.Contains(v, []byte("setup completed successfully")) {
			fmt.Println("engine-setup finished")
			break
		}
	}
}

func getLastLineOfFile(fname string) []byte {
	file, err := os.Open(fname)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	offset, _ := file.Seek(0, 2)

	buf := make([]byte, bytesToReadFromFile)
	file.ReadAt(buf, offset-bytesToReadFromFile)

	ret := bytes.Split(buf, []byte{'\n'})

	for i := 1; ; i++ {
		if string(ret[len(ret)-i]) != "" {
			return ret[len(ret)-i]
		}
	}
}

func clearScreen() {
	sendKeys("\n\n\n")
	sendKeys("clear\n")
}

func startSetup(debug bool) {
	if debug {
		sendKeys("engine-setup --offline --config-append=ovirt-engine-answers | tee /tmp/engine-setup.log\n")
	} else {
		sendKeys("engine-setup --offline --config-append=ovirt-engine-answers > /tmp/engine-setup.log\n")
	}

}

// Run is
func Run(debug bool, start bool) {
	clearScreen()

	// if start {
	// 	startSetup(debug)
	// }
	startSetup(true)
	time.Sleep(2 * time.Second)

	grepPatternFromFile(setupLogFile)
}
