package cfg

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"path/filepath"
	"strings"
)

type Config interface{
	Expand() ///this will be called after the config is read, and can be used to expand
				/// external variables or any other post processing
}

var cfgbase string

func Setup(basepath string){
	cfgbase = basepath
}

func Read(cfgname string, cfg Config)error{
	if cfgbase == ""{
		return errors.New(`Base path has not been set. cfg.Setup() must be called with a 
 config dir path, before trying to read config`)
	}
	if strings.Contains(cfgname, "../"){
		return errors.New("Config files must be under the config directory. Cannot use .. in config name")
	}

	fullpath := filepath.Join(cfgbase, cfgname + ".json")
	rawdata, err := ioutil.ReadFile(fullpath)
	if err != nil{
		return err
	}
	err = json.Unmarshal(rawdata, cfg)
	if err != nil{
		return err
	}
	cfg.Expand()

	return nil
}
