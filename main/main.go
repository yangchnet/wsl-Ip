package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"os/exec"
	"regexp"
	"strings"
)

const (
	host_file = "/mnt/c/Windows/System32/drivers/etc/hosts"
	// host_file = "./hosts"
)

func exec_shell(s string) (string, error) {
	cmd := exec.Command("/bin/bash", "-c", s)

	var out bytes.Buffer
	cmd.Stdout = &out

	err := cmd.Run()

	return out.String(), err
}

func getIp() (ip string, err error) {
	output, err := exec_shell("ifconfig eth0")
	if err != nil {
		fmt.Println(err)
	}

	re, err := regexp.Compile(`inet \d+\.\d+\.\d+\.\d+`)
	if err != nil {
		return "", err
	}
	return re.FindString(output)[5:], nil
}

func join() (string, error) {
	ip, err := getIp()
	if err != nil {
		return "", err
	}
	return fmt.Sprintf("%s wsl # wsl2 ip proxy\n", ip), nil
}

func getHosts() (string, error) {
	hosts_content, err := ioutil.ReadFile(host_file)
	if err != nil {
		return "", err
	}
	return string(hosts_content), nil
}

func mdyHosts() error {
	proxy_str, err := join()
	if err != nil {
		return err
	}

	origin_hosts, err := getHosts()
	if err != nil {
		fmt.Println(err)
	}
	host_slice := strings.Split(origin_hosts, "\n")
	re, _ := regexp.Compile(`. # wsl2 ip proxy`)
	modified := false
	for i, h := range host_slice {
		if s := re.FindString(h); s != "" {
			modified = true
			host_slice[i] = proxy_str
		}
	}
	if !modified {
		host_slice = append(host_slice, proxy_str)
	}

	modified_hosts := strings.Join(host_slice, "\n")

	err = ioutil.WriteFile(host_file, []byte(modified_hosts), 0644)
	fmt.Printf("已配置DNS：%s", proxy_str)
	return err
}

func main() {
	mdyHosts()
}
