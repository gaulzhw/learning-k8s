package jsondiff

import (
	"encoding/json"
	"log"

	"github.com/wI2L/jsondiff"
)

func diff(obj, newObj interface{}) {
	// 比较新旧deploy的不同，返回不同的bytes
	patch, err := jsondiff.Compare(obj, newObj)
	if err != nil {
		log.Fatalln(err)
		return
	}

	// 打patch，patchBytes就是我们需要的了
	patchBytes, err := json.MarshalIndent(patch, "", "    ")
	if err != nil {
		log.Fatalln(err)
		return
	}

	// 打印出来看一下
	log.Println(string(patchBytes))
}
