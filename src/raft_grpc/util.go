package main

import (
	"io/ioutil"
	"log"
	"net"

	"golang.org/x/exp/slices"
	"gopkg.in/yaml.v3"
)

type conf struct {
	Address       string   `yaml:"address"`
	Port          int      `yaml:"port"`
	RaftId        string   `yaml:"raftId"`
	RaftDir       string   `yaml:"raftDir"`
	RaftBootstrap bool     `yaml:"raftBootstrap"`
	Life          int      `yaml:"life"`
	Players       []string `yaml:"players"`
}

func getConf(filepath string) (c *conf) {
	yamlFile, err := ioutil.ReadFile(filepath)
	if err != nil {
		log.Fatalf("failed to read config %s: %v", filepath, err)
	}

	err = yaml.Unmarshal(yamlFile, &c)
	if err != nil {
		log.Fatalf("failed to parse config %s: %v", filepath, err)
	}

	log.Printf("config %s: %+v\n", filepath, *c)

	if !slices.Contains(c.Players, c.RaftId) {
		log.Fatalf("raftId %s not in players %v", c.RaftId, c.Players)
	}

	return c
}

func GetOutboundIP() net.IP {
	conn, err := net.Dial("udp", "8.8.8.8:80")
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	localAddr := conn.LocalAddr().(*net.UDPAddr)

	return localAddr.IP
}
