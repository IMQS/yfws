package yfws

import (
	"errors"
	"fmt"
	"github.com/clbanning/mxj"
	"net/http"
	"strconv"
	"strings"
)

func SendRequest(url, msg string, params map[string]string) ([]mxj.Map, error) {
	response := make([]mxj.Map, 0)

	// Check that all vars are filled
	yfrequest, ok := YFRequests[msg]
	if !ok {
		return response, fmt.Errorf("Could not find soap for %s", msg)
	}
	for _, param := range yfrequest.Params {
		if _, ok := params[param]; !ok {
			return response, fmt.Errorf("Could not find value for param %s", param)
		}
	}
	local := yfrequest.Request
	for name, value := range params {
		local = strings.Replace(local, name, value, -1)
	}

	req, err := http.NewRequest("POST", url, strings.NewReader(local))
	if err != nil {
		return response, err
	}
	req.Header.Add("Content-Type", "text/xml;charset=UTF-8")
	req.Header.Add("SOAPAction", `""`)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return response, err
	}
	m, err := mxj.NewMapXmlReader(resp.Body)
	resp.Body.Close()
	if err != nil {
		return response, err
	}

	err = parseResponse(&m, &response, "remoteAdministrationCall", "contentResources")
	if err != nil {
		return response, fmt.Errorf("Parse Response failed : %v", err)
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
		errcodeint, _ := strconv.Atoi(errcode)
		return fmt.Errorf("Yellowfin Error (%s) : %s", errcode, yferrors[errcodeint])
	}
	if idmapname != "" {
		ids, err := valuemap.ValuesForPath(idmapname + "." + idmapname)
		if err != nil {
			return err
		}
		for _, id := range ids {
			idmap := mxj.Map(id.(map[string]interface{}))
			idvalue, err := idmap.ValueForPathString("-href")
			if err != nil {
				return err
			}

			multiref, err := m.ValuesForKey("multiRef", "-id:"+idvalue[1:])
			if err != nil {
				return err
			}
			if len(multiref) == 0 {
				return fmt.Errorf("mutilRef element not found for id : %s", idvalue[1:])
			}
			*response = append(*response, mxj.Map(multiref[0].(map[string]interface{})))
		}
	}
	return nil

}
