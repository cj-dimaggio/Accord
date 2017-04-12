package main

import (
	"io/ioutil"
	"os"
	"path"
	"strings"

	"github.com/Ssawa/accord/accord"
	"github.com/Ssawa/accord/components"
	"github.com/sirupsen/logrus"
)

type MyManager struct {
}

type Data struct {
	File  string
	Value string
}

func ParseData(bytes []byte) Data {
	str := string(bytes)
	parts := strings.Split(str, ":")
	return Data{
		File:  parts[0],
		Value: parts[1],
	}
}

func (manager MyManager) Process(msg accord.Message, fromRemote bool) error {
	data := ParseData(msg.Payload)
	filename := path.Join(os.Args[5], data.File)
	os.Remove(filename)
	return ioutil.WriteFile(filename, []byte(data.Value), 0777)
}

func (manager MyManager) ShouldProcess(msg accord.Message, history *accord.HistoryStack) bool {
	return true
}

func main() {
	logger := logrus.New()
	logger.Level = logrus.DebugLevel

	comps := []accord.Component{
		&components.WebReceiver{
			BindAddress: os.Args[1],
		},

		&components.PollListener{
			BindAddress: os.Args[2],
		},

		&components.PollRequestor{
			ConnectAddress: os.Args[3],
		},
	}

	accord := accord.NewAccord(MyManager{}, comps, os.Args[4], logrus.NewEntry(logger))

	accord.StartAndListen(os.Interrupt)
}
