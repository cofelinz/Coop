package config

import (
	"os"
	"io/ioutil"
	"strings"
	"encoding/xml"
	"encoding/json"
	"errors"
)

type HostInfo struct {
	Ip  string `json:"ip"`
	Isp int64  `json:"isp"`
}

type Config struct {
	host	HostInfo	`json:"is_test"`
}

var config Config

func GetConfig() Config {
	return config
}

func init() {
	LoadAuto(&config, "server.json")
}

// 根据指定路径加载
func LoadByFile(file string, config interface{}) error {
	fd, err := os.Open(file)
	if err != nil {
		//Log.Notice("open file[%s] fail. err=[%v]", file, err)
		return err
	}
	defer fd.Close()

	b, _ := ioutil.ReadAll(fd)
	if strings.Index(file, ".xml")>0 {
		err = xml.Unmarshal(b, config)
	} else if strings.Index(file, ".json")>0 {
		err = json.Unmarshal(b, config)
	} else {
		return errors.New("unknow file type. file="+file)
	}
	if err!=nil {
		//Log.Error("unmarshal file[%s] fail. err=[%v]", file, err)
	}
	return err
}

func LoadAuto(config interface{}, fname string) (err error) {
	if fname=="" {
		fname="server.xml"
	}
	for {
		err = LoadByFile("../conf/"+fname, config)
		if nil==err {
			break
		} else {
			//Log.Debug("develop envirment.")
		}

		// 开发环境 TODO 替换####
		err = LoadByFile("./conf/test/"+fname, config)
		break
	}

	if nil==err {
		//Log.Info("config info=[%v]", config)
	} else {
		//Log.Error("load config fail. error=[%v]", err)
	}
	return
}
