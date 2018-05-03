package yfws

import (
	"bytes"
	"errors"
	"fmt"
	"github.com/clbanning/mxj"
	"io/ioutil"
	"net/http"
	"strings"
)

func SendRequest(url, msg string, params map[string]string) ([]mxj.Map, error) {

	yfrequest, ok := yfRequests[msg]
	if !ok {
		return nil, fmt.Errorf("Could not find request for %s", msg)
	}

	local := yfrequest.Request

	// Check that all vars are filled
	for _, name := range yfrequest.Params {
		if value, ok := params[name]; ok {
			local = strings.Replace(local, name, value, -1)
		} else {
			return nil, fmt.Errorf("Could not find value for param %s", name)
		}
	}
	
	req, err := http.NewRequest("POST", url, strings.NewReader(local))
	if err != nil {
		return nil, err
	}
	req.Header.Add("Content-Type", "text/xml;charset=UTF-8")
	req.Header.Add("SOAPAction", `""`)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	if resp.StatusCode != 200 {
		bdy, _ := ioutil.ReadAll(resp.Body)
		fmt.Println(string(bdy))
		resp.Body.Close()
		return nil, fmt.Errorf("Yellowfin Error: HTTP Code %v", resp.StatusCode)
	}
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	resp.Body.Close()
	m, err := mxj.NewMapXmlReader(bytes.NewReader(body))
	if err != nil {
		return nil, err
	}

	response := make([]mxj.Map, 0)

	err = parseResponse(&m, &response, yfrequest.Call, yfrequest.Resource)
	if err != nil {
		return nil, err
	}
	return response, nil

}

func parseResponse(m *mxj.Map, response *[]mxj.Map, responsename, idmapname string) error {
	path := fmt.Sprintf("Envelope.Body.%sResponse.%sReturn.-href", responsename, responsename)
	mainid, err := m.ValueForPathString(path)
	if err != nil {
		return err
	}
	value, err := m.ValuesForKey("multiRef", fmt.Sprintf("-id:%s", mainid[1:]))
	if err != nil {
		return err
	}
	if len(value) == 0 {
		return errors.New("Main reponse not found")
	}

	valuemap := mxj.Map(value[0].(map[string]interface{}))

	statuscode, err := valuemap.ValueForPathString("statusCode.#text")
	if err != nil {
		return err
	}

	if statuscode != "SUCCESS" {
		errcode, _ := valuemap.ValueForPathString("errorCode.#text")
		return YFErrors[errcode]
	}

	if idmapname != "" {
		ids, err := valuemap.ValuesForPath(idmapname + "." + idmapname)
		if err != nil {
			return fmt.Errorf("Could not find %s.%s %v", idmapname, idmapname, err)
		}
		for _, id := range ids {
			idmap := mxj.Map(id.(map[string]interface{}))
			idvalue, err := idmap.ValueForPathString("-href")
			if err != nil {
				return err
			}

			multiref, err := m.ValuesForKey("multiRef", "-id:"+idvalue[1:])
			if err != nil {
				return fmt.Errorf("Could not find multiref for id %s, %v", idvalue[1:], err)
			}
			if len(multiref) == 0 {
				return fmt.Errorf("mutilRef element not found for id : %s", idvalue[1:])
			}
			*response = append(*response, mxj.Map(multiref[0].(map[string]interface{})))
		}
	} else {
		// Add the original back, this is to get responses such as sessionid etc.
		*response = append(*response, valuemap)
	}
	return nil

}
