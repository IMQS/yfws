package yfws

type yfRequest struct {
	Request  string
	Params   []string
	Call     string
	Resource string
}

var yfRequests = map[string]*yfRequest{
	"changepassword": &yfRequest{
		Request: changepassword,
		Params: []string{
			"%ADMIN%",
			"%PASSWORD%",
			"%USERID%",
			"%USERPASSWORD%",
		},
		Call:     "remoteAdministrationCall",
		Resource: "",
	},
	"createuser": &yfRequest{
		Request: createuser,
		Params: []string{
			"%ADMIN%",
			"%PASSWORD%",
			"%EMAIL%",
			"%FIRSTNAME%",
			"%LASTNAME%",
			"%USERPASSWORD%",
			"%USERID%",
		},
		Call:     "remoteAdministrationCall",
		Resource: "",
	},
	"deletecontent": &yfRequest{
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
	"getcontent": &yfRequest{
		Request: getcontent,
		Params: []string{
			"%ADMIN%",
			"%PASSWORD%",
		},
		Call:     "remoteAdministrationCall",
		Resource: "contentResources",
	},
	"getreport": &yfRequest{
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
	"getallreports": &yfRequest{
		Request: getallreports,
		Params: []string{
			"%ADMIN%",
			"%PASSWORD%",
			"%USER%",
		},
		Call:     "remoteAdministrationCall",
		Resource: "reports",
	},
	"importcontent": &yfRequest{
		Request: importcontent,
		Params: []string{
			"%ADMIN%",
			"%PASSWORD%",
			"%DATA%",
		},
		Call:     "remoteAdministrationCall",
		Resource: "",
	},
	"login": &yfRequest{
		Request: login,
		Params: []string{
			"%ADMIN%",
			"%PASSWORD%",
			"%USER%",
			"%CONTENTCATEGORY%",
			"%SCENARIOFILTER%",
			"%GLOBALFILTERS%",
		},
		Call:     "remoteAdministrationCall",
		Resource: "",
	},
	"logout": &yfRequest{
		Request: logout,
		Params: []string{
			"%ADMIN%",
			"%PASSWORD%",
			"%SESSIONID%",
			"%USER%",
		},
		Call:     "remoteAdministrationCall",
		Resource: "",
	},
	"schema": &yfRequest{
		Request: schema,
		Params: []string{
			"%ADMIN%",
			"%PASSWORD%",
			"%REPORTID%",
		},
		Call:     "remoteReportCall",
		Resource: "columns",
	},
	"updateuser": &yfRequest{
		Request: updateuser,
		Params: []string{
			"%ADMIN%",
			"%PASSWORD%",
			"%ROLE%",
			"%USERID%",
		},
		Call:     "remoteAdministrationCall",
		Resource: "",
	},
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
            <parameters xsi:type="ser:ArrayOf_soapenc_string" soapenc:arrayType="xsd:string[]">
               <item xsd:type="xsd:string">ENTRY=BROWSE</item>
               <item xsd:type="xsd:string">%CONTENTCATEGORY%</item>
               <item xsd:type="xsd:string">%SCENARIOFILTER%</item>
               %GLOBALFILTERS%
            </parameters>
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

const getallreports = `<soapenv:Envelope xmlns:xsi="http://www.w3.org/2001/XMLSchema-instance" xmlns:xsd="http://www.w3.org/2001/XMLSchema" xmlns:soapenv="http://schemas.xmlsoap.org/soap/envelope/" xmlns:ser="http://service.web.mi.hof.com" xmlns:soapenc="http://schemas.xmlsoap.org/soap/encoding/">
   <soapenv:Header/>
   <soapenv:Body>
      <ser:remoteAdministrationCall soapenv:encodingStyle="http://schemas.xmlsoap.org/soap/encoding/">
         <in0 xsi:type="ser:AdministrationServiceRequest">
            <function xsi:type="xsd:string">GETALLUSERREPORTS</function>
            <loginId xsi:type="xsd:string">%ADMIN%</loginId>
            <orgId xsi:type="xsd:int">1</orgId>
            <password xsi:type="xsd:string">%PASSWORD%</password>
            <person xsi:type="ser:AdministrationPerson">
               <userId xsi:type="xsd:string">%USER%</userId>
            </person>
         </in0>
      </ser:remoteAdministrationCall>
   </soapenv:Body>
</soapenv:Envelope>`

const createuser = `<soapenv:Envelope xmlns:xsi="http://www.w3.org/2001/XMLSchema-instance" xmlns:xsd="http://www.w3.org/2001/XMLSchema" xmlns:soapenv="http://schemas.xmlsoap.org/soap/envelope/" xmlns:ser="http://service.web.mi.hof.com" xmlns:soapenc="http://schemas.xmlsoap.org/soap/encoding/">
   <soapenv:Header/>
   <soapenv:Body>
      <ser:remoteAdministrationCall soapenv:encodingStyle="http://schemas.xmlsoap.org/soap/encoding/">
         <in0 xsi:type="ser:AdministrationServiceRequest">
            <function xsi:type="xsd:string">ADDUSER</function>
            <loginId xsi:type="xsd:string">%ADMIN%</loginId>
            <orgId xsi:type="xsd:int">1</orgId>
            <password xsi:type="xsd:string">%PASSWORD%</password>
            <person xsi:type="ser:AdministrationPerson">
               <emailAddress xsi:type="xsd:string">%EMAIL%</emailAddress>
               <firstName xsi:type="xsd:string">%FIRSTNAME%</firstName>
               <initial xsi:type="xsd:string"></initial>
               <ipId xsi:type="xsd:int"></ipId>
               <languageCode xsi:type="xsd:string"></languageCode>
               <lastName xsi:type="xsd:string">%LASTNAME%</lastName>
               <password xsi:type="xsd:string">%USERPASSWORD%</password>
               <roleCode xsi:type="xsd:string">YFREPORTCONSUMER</roleCode>
               <salutationCode xsi:type="xsd:string">DR</salutationCode>
               <timeZoneCode xsi:type="xsd:string"></timeZoneCode>
               <userId xsi:type="xsd:string">%USERID%</userId>
            </person>
         </in0>
      </ser:remoteAdministrationCall>
   </soapenv:Body>
</soapenv:Envelope>`

const changepassword = `<soapenv:Envelope xmlns:xsi="http://www.w3.org/2001/XMLSchema-instance" xmlns:xsd="http://www.w3.org/2001/XMLSchema" xmlns:soapenv="http://schemas.xmlsoap.org/soap/envelope/" xmlns:ser="http://service.web.mi.hof.com" xmlns:soapenc="http://schemas.xmlsoap.org/soap/encoding/">
   <soapenv:Header/>
   <soapenv:Body>
      <ser:remoteAdministrationCall soapenv:encodingStyle="http://schemas.xmlsoap.org/soap/encoding/">
         <in0 xsi:type="ser:AdministrationServiceRequest">
            <function xsi:type="xsd:string">CHANGEPASSWORD</function>
            <loginId xsi:type="xsd:string">%ADMIN%</loginId>
            <orgId xsi:type="xsd:int">1</orgId>
            <password xsi:type="xsd:string">%PASSWORD%</password>
            <person xsi:type="ser:AdministrationPerson">
               <userId xsi:type="xsd:string">%USERID%</userId>
               <password xsi:type="xsd:string">%USERPASSWORD%</password>
            </person>
         </in0>
      </ser:remoteAdministrationCall>
   </soapenv:Body>
</soapenv:Envelope>`

const updateuser = `<soapenv:Envelope xmlns:xsi="http://www.w3.org/2001/XMLSchema-instance" xmlns:xsd="http://www.w3.org/2001/XMLSchema" xmlns:soapenv="http://schemas.xmlsoap.org/soap/envelope/" xmlns:ser="http://service.web.mi.hof.com" xmlns:soapenc="http://schemas.xmlsoap.org/soap/encoding/">
   <soapenv:Header/>
   <soapenv:Body>
      <ser:remoteAdministrationCall soapenv:encodingStyle="http://schemas.xmlsoap.org/soap/encoding/">
         <in0 xsi:type="ser:AdministrationServiceRequest">
            <function xsi:type="xsd:string">UPDATEUSER</function>
            <loginId xsi:type="xsd:string">%ADMIN%</loginId>
            <orgId xsi:type="xsd:int">1</orgId>
            <password xsi:type="xsd:string">%PASSWORD%</password>
            <person xsi:type="ser:AdministrationPerson">
               <roleCode xsi:type="xsd:string">%ROLE%</roleCode>
               <userId xsi:type="xsd:string">%USERID%</userId>
            </person>
         </in0>
      </ser:remoteAdministrationCall>
   </soapenv:Body>
</soapenv:Envelope>`

const logout = `<soapenv:Envelope xmlns:xsi="http://www.w3.org/2001/XMLSchema-instance" xmlns:xsd="http://www.w3.org/2001/XMLSchema" xmlns:soapenv="http://schemas.xmlsoap.org/soap/envelope/" xmlns:ser="http://service.web.mi.hof.com" xmlns:soapenc="http://schemas.xmlsoap.org/soap/encoding/">
   <soapenv:Header/>
   <soapenv:Body>
      <ser:remoteAdministrationCall soapenv:encodingStyle="http://schemas.xmlsoap.org/soap/encoding/">
         <in0 xsi:type="ser:AdministrationServiceRequest">
            <function xsi:type="xsd:string">LOGOUTUSER</function>
            <loginId xsi:type="xsd:string">%ADMIN%</loginId>
            <orgId xsi:type="xsd:int">1</orgId>
            <orgRef xsi:type="xsd:string">Default</orgRef>
            <password xsi:type="xsd:string">%PASSWORD%</password>
            <loginSessionId xsi:type="xsd:string">%SESSIONID%</loginSessionId>
            <person xsi:type="ser:AdministrationPerson">
               <userId xsi:type="xsd:string">%USER%</userId>
            </person>
         </in0>
      </ser:remoteAdministrationCall>
   </soapenv:Body>
</soapenv:Envelope>`
