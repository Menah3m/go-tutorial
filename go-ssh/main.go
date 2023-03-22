package main

import (
	"fmt"
	"golang.org/x/crypto/ssh"
	"golang.org/x/crypto/ssh/terminal"
	"log"
	"net"
	"os"
	"time"
)

/*
   @Auth: menah3m
   @Desc: go 使用 ssh 连接
*/

type SSHTarget struct {
	Host string
	Port int
}

func (s *SSHTarget) GetAddr() string {
	addr := fmt.Sprintf("%s:%d", s.Host, s.Port)
	return addr
}

type SSHConnector struct {
	SSHTarget    *SSHTarget
	Auth         []ssh.AuthMethod
	Addr         string
	ClientConfig *ssh.ClientConfig
	Client       *ssh.Client
	Session      *ssh.Session
	User         string
	Password     string
	err          error
}

func (s *SSHConnector) Connect(host string, port int) {
	s.Auth = append(s.Auth, ssh.Password(s.Password))
	s.ClientConfig = &ssh.ClientConfig{
		User:    s.User,
		Auth:    s.Auth,
		Timeout: 30 * time.Second,
		HostKeyCallback: func(hostname string, remote net.Addr, key ssh.PublicKey) error {
			return nil
		},
	}
	s.SSHTarget = &SSHTarget{
		Host: host,
		Port: port,
	}
	s.Addr = s.SSHTarget.GetAddr()
	if s.Client, s.err = ssh.Dial("tcp", s.Addr, s.ClientConfig); s.err != nil {
		fmt.Println(s.err)
	}
	if s.Session, s.err = s.Client.NewSession(); s.err != nil {
		fmt.Println(s.err)
	}

}

func main() {
	sshConnector := &SSHConnector{
		User:     "root",
		Password: "1!P@ssword",
	}
	sshConnector.Connect("192.168.108.112", 22)
	defer sshConnector.Session.Close()

	fd := int(os.Stdin.Fd())
	oldState, err := terminal.MakeRaw(fd)
	if err != nil {
		fmt.Println(err)
	}
	defer terminal.Restore(fd, oldState)

	// 重定向到stdout\stderr\stdin
	sshConnector.Session.Stdout = os.Stdout
	sshConnector.Session.Stderr = os.Stderr
	sshConnector.Session.Stdin = os.Stdin

	// 配置终端的外观
	termWidth, termHeight, err := terminal.GetSize(fd)
	if err != nil {
		fmt.Println(err)
	}
	//配置终端模式
	termModes := ssh.TerminalModes{
		ssh.ECHO:          1,
		ssh.TTY_OP_ISPEED: 14400,
		ssh.TTY_OP_OSPEED: 14400,
	}

	if err := sshConnector.Session.RequestPty("xterm-256color", termHeight, termWidth, termModes); err != nil {
		log.Fatal(err)
	}
	cmd := "ls -l"
	fmt.Println("执行命令:", cmd)
	fmt.Println("结果:")
	sshConnector.Session.Run(cmd)

}
