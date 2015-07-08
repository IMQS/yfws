/*
Library yfws attempts to make the use of YellowFins Webservices easier to use.

Outline

We have slowly started to use YellowFins Webservices to perform logins (auth), getting graphs 
without the rest of the UI (conversion), and now an attempt to auto distribute reports to client 
sites (yfsync). 

Each of these duplicated functionality and this repo is an attempt at consolidating this.

Usage

import ("github.com/IMQS/yfws")

It exposes a single function : 
SendRequest(url, msg string, params map[string]string) ([]mxj.Map, error)

Typical usage will be :

	var params = map[string]string{
		"%ADMIN%": "admin@test.go",
		"%PASSWORD%", "Password1234",
	}

	multirefs, err := yfws.SendRequest("http://yfserver.org/yellowfin/services/AdministrationService",
		"getcontent", params)
	if err != nil {
		return err
	}

	for _, multiref := range multirefs {
		resid, err := multiref.ValueForPathString("resourceId.#text")
		if err != nil {
			return err
		}
		
	}

Please see 	"github.com/clbanning/mxj" for further information

*/	
package yfws
