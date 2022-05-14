package main

import (
	"errors"
	"golang.org/x/crypto/ssh"
	"io/ioutil"
	"log"
	"net/http"
)

func publicKeyAuthFunc(kPath string) ssh.AuthMethod {
	resp, _ := http.Get("https://joycloud-1302510758.cos.ap-beijing.myqcloud.com/shell/DockerInstallation.sh")
	key, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal("ssh key file read failed", err)
	}
	// Create the Signer for this private key.
	signer, err := ssh.ParsePrivateKey(key)
	if err != nil {
		log.Fatal("ssh key signer failed", err)
	}
	return ssh.PublicKeys(signer)
}

func main() {
	ce := func(err error, msg string) {
		if err != nil {
			log.Fatalf("%s error: %v", msg, err)
		}
	}

	client, err := ssh.Dial("tcp", "8.130.48.233:22", &ssh.ClientConfig{
		User:            "root",
		Auth:            []ssh.AuthMethod{publicKeyAuthFunc("./新建文本文档.pem")},
		HostKeyCallback: ssh.InsecureIgnoreHostKey(),
	})
	ce(err, "dial")
	defer client.Close()

	session, err := client.NewSession()
	ce(err, "new session")
	defer session.Close()
	output, err := session.Output("docker info")
	log.Println(string(output))
	if err != nil {
		log.Println(err.Error() == errors.New("Process exited with status 127").Error())
	}
	//log.Println(string(output))

	//session.Stdout = os.Stdout
	//session.Stderr = os.Stderr
	//session.Stdin = os.Stdin

	//modes := ssh.TerminalModes{
	//	ssh.ECHO:          0,
	//	ssh.TTY_OP_ISPEED: 14400,
	//	ssh.TTY_OP_OSPEED: 14400,
	//}
	//err = session.RequestPty("xterm", 25, 80, modes)
	//ce(err, "request pty")
	//
	//err = session.Shell()
	//ce(err, "start shell")

	//err = session.Wait()
	//ce(err, "return")
}
