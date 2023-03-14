package HttpUtil

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
)

func PostWithBody(url string, values map[string]string) map[string]interface{} {
	payload, err := json.Marshal(values)

	resp, err := http.Post(url, "application/json",
		bytes.NewBuffer(payload))

	if err != nil {
		fmt.Println(err)
		return nil
	}

	var res map[string]interface{}

	err = json.NewDecoder(resp.Body).Decode(&res)
	if err != nil {
		fmt.Println(err)
		return nil
	}

	return res
}
