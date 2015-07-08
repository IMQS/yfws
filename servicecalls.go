package yfws

type YFRequest struct {
	Request   string
	Params []string
	Call      string
	Resource  string
}

var YFRequests = map[string]*YFRequest{
	"importcontent": &YFRequest{
		Request: importcontent,
		Params: []string{
			"%ADMIN%",
			"%PASSWORD%",
			"%DATA%",
		},
		Call:     "remoteAdministrationCall",
		Resource: "",
	},
	"getcontent": &YFRequest{
		Request: getcontent,
		Params: []string{
			"%ADMIN%",
			"%PASSWORD%",
		},
		Call:     "remoteAdministrationCall",
		Resource: "contentResources",
	},
	"deletecontent": &YFRequest{
		Request: deletecontent,
		Params: []string{
			"%ADMIN%",
			"%PASSWORD%",
			"%FUNCTION%",
			"%ID%",
		},
		Call:     "remoteAdministrationCall",
		Resource: "",
	},
	"login": &YFRequest{
		Request: login,
		Params: []string{
			"%ADMIN%",
			"%PASSWORD%",
			"%USER%",
		},
		Call:     "remoteAdministrationCall",
		Resource: "",
	},
	"schema": &YFRequest{
		Request: schema,
		Params: []string{
			"%ADMIN%",
			"%PASSWORD%",
			"%REPORTID%",
		},
		Call:     "remoteReportCall",
		Resource: "columns",
	},
	"getreport": &YFRequest{
		Request: getreport,
		Params: []string{
			"%ADMIN%",
			"%PASSWORD%",
			"%REPORTID%",
			"%FILTERID%",
			"%FILTERVALUE%",
		},
		Call:     "remoteReportCall",
		Resource: "charts",
	},
}

var yferrors = map[int]string{
	-2: "UNKNOWN_ERROR",
	-1: "CANNOT_CONNECT",
	0:  "NO_ERROR",
	1:  "USER_NOT_AUTHENTICATED",
	2:  "NO_WEBSERVICE_ACCESS",
	3:  "PERSON_REQUIRED",
	4:  "COULD_NOT_CREATE_PERSON",
	5:  "COULD_NOT_RELOAD_LICENCE",
	6:  "LOGIN_ALREADY_IN_USE",
	7:  "COULD_NOT_DELETE_PERSON",
	8:  "COULD_NOT_FIND_PERSON",
	9:  "LICENCE_BREACH",
	10: "COULD_NOT_LOAD_REPORT_ACCESS",
	11: "COULD_NOT_LOAD_REPORT_LIST",
	12: "COULD_NOT_FIND_GROUP",
	13: "GROUP_EXISTS",
	14: "BIRT_OBJECT_NULL",
	15: "BIRT_OBJECT_NO_DATA",
	16: "BIRT_SOURCE_MISSING",
	17: "BIRT_COULD_NOT_SAVE",
	18: "BIRT_COULD_NOT_SAVE_BIRT_FILE",
	19: "COULD_NOT_UPDATE_PASSWORD",
	20: "UNKNOWN_WEBSERVICE_FUNCTION",
	21: "INVALID_CLIENT_REFERENCE",
	22: "CLIENT_EXISTS",
	23: "COULD_NOT_FIND_REPORT",
	24: "REPORT_IS_DRAFT",
	25: "COULD_NOT_AUTHENTICATE_USER",
	26: "UNSECURE_LOGON_NOT_ENABLED",
	27: "ROLE_NOT_FOUND",
	28: "COULD_NOT_LOAD_FAVOURITES",
	29: "RESPONSE_IS_TOO_LARGE",
	30: "SOURCE_NOT_FOUND",
	31: "EMPTY_RECIPIENT_LIST",
	32: "BROADCAST_FAILED",
	33: "FILTERVALUES_FAILED",
	34: "CLIENT_ORGS_DISABLED",
	35: "DASHBOARD_TAB_NOT_FOUND",
	36: "SCHEDULE_NULL",
	37: "UNKNOWN_STATUS_CODE",
	38: "PASSWORD_REQUIREMENTS_NOT_MET",
	39: "LOGIN_MAXIMUM_ATTEMPTS",
}

const importcontent = `<soapenv:Envelope xmlns:xsi="http://www.w3.org/2001/XMLSchema-instance" xmlns:xsd="http://www.w3.org/2001/XMLSchema" xmlns:soapenv="http://schemas.xmlsoap.org/soap/envelope/" xmlns:ser="http://service.web.mi.hof.com" xmlns:soapenc="http://schemas.xmlsoap.org/soap/encoding/">
   <soapenv:Header/>
   <soapenv:Body>
      <ser:remoteAdministrationCall soapenv:encodingStyle="http://schemas.xmlsoap.org/soap/encoding">
         <in0 xsi:type="ser:AdministrationServiceRequest">
            <function xsi:type="xsd:string">IMPORTCONTENT</function>
            <loginId xsi:type="xsd:string">%ADMIN%</loginId>
            <password xsi:type="xsd:string">%PASSWORD%</password>
            <orgId xsi:type="xsd:int">1</orgId>
            <parameters xsi:type="ser:ArrayOf_soapenc_string" soapenc:arrayType="xsd:string[1]">
              <item xsd:type="soapenc:base64">%DATA%</item>
            </parameters>
         </in0>
      </ser:remoteAdministrationCall>
   </soapenv:Body>
</soapenv:Envelope>`

const getcontent = `<soapenv:Envelope xmlns:xsi="http://www.w3.org/2001/XMLSchema-instance" xmlns:xsd="http://www.w3.org/2001/XMLSchema" xmlns:soapenv="http://schemas.xmlsoap.org/soap/envelope/" xmlns:ser="http://service.web.mi.hof.com" xmlns:soapenc="http://schemas.xmlsoap.org/soap/encoding/">
   <soapenv:Header/>
   <soapenv:Body>
      <ser:remoteAdministrationCall soapenv:encodingStyle="http://schemas.xmlsoap.org/soap/encoding/">
         <in0 xsi:type="ser:AdministrationServiceRequest">
            <function xsi:type="xsd:string">GETCONTENT</function>
            <loginId xsi:type="xsd:string">%ADMIN%</loginId>
            <password xsi:type="xsd:string">%PASSWORD%</password>
            <orgId xsi:type="xsd:int">1</orgId>
         </in0>
      </ser:remoteAdministrationCall>
   </soapenv:Body>
</soapenv:Envelope>`

const deletecontent = `<soapenv:Envelope xmlns:xsi="http://www.w3.org/2001/XMLSchema-instance" xmlns:xsd="http://www.w3.org/2001/XMLSchema" xmlns:soapenv="http://schemas.xmlsoap.org/soap/envelope/" xmlns:ser="http://service.web.mi.hof.com" xmlns:soapenc="http://schemas.xmlsoap.org/soap/encoding/">
   <soapenv:Header/>
   <soapenv:Body>
      <ser:remoteAdministrationCall soapenv:encodingStyle="http://schemas.xmlsoap.org/soap/encoding/">
         <in0 xsi:type="ser:AdministrationServiceRequest">
            <function xsi:type="xsd:string">%FUNCTION%</function>
            <loginId xsi:type="xsd:string">%ADMIN%</loginId>
            <orgId xsi:type="xsd:int">1</orgId>
            <password xsi:type="xsd:string">%PASSWORD%</password>
            <parameters xsi:type="ser:ArrayOf_soapenc_string" soapenc:arrayType="xsd:string[]">
           <item xsd:type="xsd:string">%ID%</item>
            </parameters>
         </in0>
      </ser:remoteAdministrationCall>
   </soapenv:Body>
</soapenv:Envelope>`

const login = `<soapenv:Envelope xmlns:xsi="http://www.w3.org/2001/XMLSchema-instance" xmlns:xsd="http://www.w3.org/2001/XMLSchema" xmlns:soapenv="http://schemas.xmlsoap.org/soap/envelope/" xmlns:ser="http://service.web.mi.hof.com" xmlns:soapenc="http://schemas.xmlsoap.org/soap/encoding/">
   <soapenv:Header/>
   <soapenv:Body>
      <ser:remoteAdministrationCall soapenv:encodingStyle="http://schemas.xmlsoap.org/soap/encoding/">
         <in0 xsi:type="ser:AdministrationServiceRequest">
            <function xsi:type="xsd:string">LOGINUSERNOPASSWORD</function>
            <loginId xsi:type="xsd:string">%ADMIN%</loginId>
            <orgId xsi:type="xsd:int">1</orgId>
            <orgRef xsi:type="xsd:string">Default</orgRef>
            <password xsi:type="xsd:string">%PASSWORD%</password>
            <person xsi:type="ser:AdministrationPerson">
               <userId xsi:type="xsd:string">%USER%</userId>
            </person>
         </in0>
      </ser:remoteAdministrationCall>
   </soapenv:Body>
</soapenv:Envelope>`

const schema = `<soapenv:Envelope xmlns:xsi="http://www.w3.org/2001/XMLSchema-instance" xmlns:xsd="http://www.w3.org/2001/XMLSchema" xmlns:soapenv="http://schemas.xmlsoap.org/soap/envelope/" xmlns:ser="http://service.web.mi.hof.com" xmlns:soapenc="http://schemas.xmlsoap.org/soap/encoding/">
   <soapenv:Header/>
   <soapenv:Body>
      <ser:remoteReportCall soapenv:encodingStyle="http://schemas.xmlsoap.org/soap/encoding/">
         <in0 xsi:type="ser:ReportServiceRequest">
            <loginId xsi:type="xsd:string">%ADMIN%</loginId>
            <orgId xsi:type="xsd:int">1</orgId>
            <password xsi:type="xsd:string">%PASSWORD%</password>
            <reportId xsi:type="xsd:int">%REPORTID%</reportId>
            <reportRequest xsi:type="xsd:string">SCHEMA</reportRequest>
         </in0>
      </ser:remoteReportCall>
   </soapenv:Body>
</soapenv:Envelope>`

const getreport = `<soapenv:Envelope xmlns:xsi="http://www.w3.org/2001/XMLSchema-instance" xmlns:xsd="http://www.w3.org/2001/XMLSchema" xmlns:soapenv="http://schemas.xmlsoap.org/soap/envelope/" xmlns:ser="http://service.web.mi.hof.com" xmlns:soapenc="http://schemas.xmlsoap.org/soap/encoding/">
   <soapenv:Header/>
   <soapenv:Body>
      <ser:remoteReportCall soapenv:encodingStyle="http://schemas.xmlsoap.org/soap/encoding/">
         <in0 xsi:type="ser:ReportServiceRequest">
            <loginId xsi:type="xsd:string">%ADMIN%</loginId>
            <orgId xsi:type="xsd:int">1</orgId>
            <password xsi:type="xsd:string">%PASSWORD%</password>
            <reportId xsi:type="xsd:int">%REPORTID%</reportId>
            <reportRequest xsi:type="xsd:string">HTMLCHARTONLY</reportRequest>
            <filters xsi:type="rep:ArrayOf_tns1_ReportFilter" soapenc:arrayType="ser:ReportFilter[]" xmlns:rep="ReportFilter">
             <item xsi:type="ser:ReportFilter">
             	<filterId xsi:type="xsd:int">%FILTERID%</filterId>
             	<isOmitted xsi:type="xsd:bool">false</isOmitted>
             	<dataValue xsi:type="xsd:string">%FILTERVALUE%</dataValue>
             </item>
            </filters>
         </in0>
      </ser:remoteReportCall>
   </soapenv:Body>
</soapenv:Envelope>`
