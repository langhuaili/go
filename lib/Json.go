package lib

import (
	"encoding/json"
	"fmt"
)

func JsonFomat() string {
	var woman = Woman{"李嫣",15}
	v,_ := json.Marshal(woman)
	fmt.Println(v)
	jsonsa := string(v)
	return jsonsa
}

func JsonDeFomat() Woman {
	var str = []byte(`{"Name":"李嫣","Age":121215}`)
	var wom = Woman{}
	json.Unmarshal(str,&wom)
	return wom
}