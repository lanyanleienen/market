package common

import "encoding/json"

func SwapTo(request,target interface{})error{
	dataBytes,err := json.Marshal(request)
	if err != nil{
		return err
	}
	return json.Unmarshal(dataBytes,target)
}
