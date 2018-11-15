package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"path"
	//"sync"
)

func main() {
	c := make(map[string]interface{})
	c["name"] = "Gopher"
	c["title"] = "programmer"
	c["contact"] = map[string]interface{}{
		"home": "415.333.3333",
		"cell": "415.555.5555",
	}

	data, err := json.MarshalIndent(c, "", "      ") //这里返回的data值，类型是[]byte
	if err != nil {
		log.Println("ERROR:", err)
	}

	fmt.Println(string(data))

	WriteConfig("a.json", data)
}

func WriteConfig(cfg string, jsonByte []byte) { //这里的cfg就是我要写到的目标文件 ./host.json
	if cfg == "" {
		log.Fatalln("use -c to specify configuration file")
	}

	_, err := WriteBytes(cfg, jsonByte)
	if err != nil {
		log.Fatalln("write config file:", cfg, "fail:", err)
	}

	//lock.Lock()
	//defer lock.Unlock()

	log.Println("write config file:", cfg, "successfully")

}

func WriteBytes(filePath string, b []byte) (int, error) {
	os.MkdirAll(path.Dir(filePath), os.ModePerm)
	fw, err := os.Create(filePath)
	if err != nil {
		return 0, err
	}
	defer fw.Close()
	return fw.Write(b)
}
