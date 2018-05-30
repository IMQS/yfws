package yfws

import (
	"errors"
	"fmt"
	"net/http"
	"strings"

	"github.com/clbanning/mxj"
)

// Make request to Yellowfin service.
func SendRequest(url, msg string, params map[string]string) ([]mxj.Map, error) {
	yfrequest, ok := yfRequests[msg]
	if !ok {
		return nil, fmt.Errorf("invalid request '%s'", msg)
	}

	local := yfrequest.Request

	// Check that all vars are filled
	for _, name := range yfrequest.Params {
		if value, ok := params[name]; ok {
			local = strings.Replace(local, name, value, -1)
		} else {
			return nil, fmt.Errorf("incomplete request. No value for param '%s'", name)
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

	m, err := mxj.NewMapXmlReader(resp.Body)
	resp.Body.Close()
	if err != nil {
		return nil, err
	}

	if resp.StatusCode != 200 {
		return nil, fmt.Errorf("request error HTTP %d ->\n%s", resp.StatusCode, parseError(m))
	}

	return parseResponse(m, yfrequest.Call, yfrequest.Resource)
}

// Parse response from Yellowfin service after making request.
func parseResponse(m mxj.Map, responsename, idmapname string) ([]mxj.Map, error) {
	path := fmt.Sprintf("Envelope.Body.%sResponse.%sReturn.-href", responsename, responsename)
	mainid, err := m.ValueForPathString(path)
	if err != nil {
		return nil, err
	}
	value, err := m.ValuesForKey("multiRef", fmt.Sprintf("-id:%s", mainid[1:]))
	if err != nil {
		return nil, err
	}
	if len(value) == 0 {
		return nil, errors.New("Main reponse not found")
	}

	valuemap := mxj.Map(value[0].(map[string]interface{}))

	statuscode, err := valuemap.ValueForPathString("statusCode.#text")
	if err != nil {
		return nil, err
	}

	if statuscode != "SUCCESS" {
		errcode, _ := valuemap.ValueForPathString("errorCode.#text")
		return nil, YFErrors[errcode]
	}

	response := make([]mxj.Map, 0, 1)

	if idmapname != "" {
		ids, err := valuemap.ValuesForPath(idmapname + "." + idmapname)
		if err != nil {
			return nil, fmt.Errorf("Could not find %s.%s %v", idmapname, idmapname, err)
		}
		for _, id := range ids {
			idmap := mxj.Map(id.(map[string]interface{}))
			idvalue, err := idmap.ValueForPathString("-href")
			if err != nil {
				return nil, err
			}

			multiref, err := m.ValuesForKey("multiRef", "-id:"+idvalue[1:])
			if err != nil {
				return nil, fmt.Errorf("Could not find multiref for id %s, %v", idvalue[1:], err)
			}
			if len(multiref) == 0 {
				return nil, fmt.Errorf("mutilRef element not found for id : %s", idvalue[1:])
			}
			response = append(response, mxj.Map(multiref[0].(map[string]interface{})))
		}
	} else {
		// Add the original back, this is to get responses such as sessionid etc.
		response = append(response, valuemap)
	}

	return response, nil
}

// Parse error response from Yellowfin service after failed request.
func parseError(m mxj.Map) string {
	errMsg, err := m.ValueForPathString("Envelope.Body.Fault.faultstring")
	if err != nil {
		return m.StringIndentNoTypeInfo(0)
	}

	return errMsg
}
