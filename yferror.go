package yfws

import (
	"fmt"
)

var YFErrors = map[string]error{
	"27": YFRoleNotFound,
	"31": YFEmptyRecipientList,
	"9":  YFLicenceBreach,
	"14": YFBirtObjectNull,
	"16": YFBirtSourceMissing,
	"37": YFUnknownStatusCode,
	"6":  YFLoginAlreadyInUse,
	"18": YFBirtCouldNotSaveBirtFile,
	"34": YFClientOrgsDisabled,
	"15": YFBirtObjectNoData,
	"19": YFCouldNotUpdatePassword,
	"25": YFCouldNotAuthenticateUser,
	"29": YFResponseIsTooLarge,
	"30": YFSourceNotFound,
	"35": YFDashboardTabNotFound,
	"0":  YFNoError,
	"7":  YFCouldNotDeletePerson,
	"20": YFUnknownWebserviceFunction,
	"36": YFScheduleNull,
	"-2": YFUnknownError,
	"10": YFCouldNotLoadReportAccess,
	"24": YFReportIsDraft,
	"5":  YFCouldNotReloadLicence,
	"8":  YFCouldNotFindPerson,
	"23": YFCouldNotFindReport,
	"33": YFFiltervaluesFailed,
	"39": YFLoginMaximumAttempts,
	"-1": YFCannotConnect,
	"1":  YFUserNotAuthenticated,
	"2":  YFNoWebserviceAccess,
	"21": YFInvalidClientReference,
	"22": YFClientExists,
	"26": YFUnsecureLogonNotEnabled,
	"28": YFCouldNotLoadFavourites,
	"32": YFBroadcastFailed,
	"3":  YFPersonRequired,
	"4":  YFCouldNotCreatePerson,
	"17": YFBirtCouldNotSave,
	"38": YFPasswordRequirementsNotMet,
	"11": YFCouldNotLoadReportList,
	"12": YFCouldNotFindGroup,
	"13": YFGroupExists,
	"100": YFInvalidGroup,
}

var YFRoleNotFound = fmt.Errorf("Yellowfin Error (%v) : %s", 27, "ROLE_NOT_FOUND")
var YFEmptyRecipientList = fmt.Errorf("Yellowfin Error (%v) : %s", 31, "EMPTY_RECIPIENT_LIST")
var YFLicenceBreach = fmt.Errorf("Yellowfin Error (%v) : %s", 9, "LICENCE_BREACH")
var YFBirtObjectNull = fmt.Errorf("Yellowfin Error (%v) : %s", 14, "BIRT_OBJECT_NULL")
var YFBirtSourceMissing = fmt.Errorf("Yellowfin Error (%v) : %s", 16, "BIRT_SOURCE_MISSING")
var YFUnknownStatusCode = fmt.Errorf("Yellowfin Error (%v) : %s", 37, "UNKNOWN_STATUS_CODE")
var YFLoginAlreadyInUse = fmt.Errorf("Yellowfin Error (%v) : %s", 6, "LOGIN_ALREADY_IN_USE")
var YFBirtCouldNotSaveBirtFile = fmt.Errorf("Yellowfin Error (%v) : %s", 18, "BIRT_COULD_NOT_SAVE_BIRT_FILE")
var YFClientOrgsDisabled = fmt.Errorf("Yellowfin Error (%v) : %s", 34, "CLIENT_ORGS_DISABLED")
var YFBirtObjectNoData = fmt.Errorf("Yellowfin Error (%v) : %s", 15, "BIRT_OBJECT_NO_DATA")
var YFCouldNotUpdatePassword = fmt.Errorf("Yellowfin Error (%v) : %s", 19, "COULD_NOT_UPDATE_PASSWORD")
var YFCouldNotAuthenticateUser = fmt.Errorf("Yellowfin Error (%v) : %s", 25, "COULD_NOT_AUTHENTICATE_USER")
var YFResponseIsTooLarge = fmt.Errorf("Yellowfin Error (%v) : %s", 29, "RESPONSE_IS_TOO_LARGE")
var YFSourceNotFound = fmt.Errorf("Yellowfin Error (%v) : %s", 30, "SOURCE_NOT_FOUND")
var YFDashboardTabNotFound = fmt.Errorf("Yellowfin Error (%v) : %s", 35, "DASHBOARD_TAB_NOT_FOUND")
var YFNoError = fmt.Errorf("Yellowfin Error (%v) : %s", 0, "NO_ERROR")
var YFCouldNotDeletePerson = fmt.Errorf("Yellowfin Error (%v) : %s", 7, "COULD_NOT_DELETE_PERSON")
var YFUnknownWebserviceFunction = fmt.Errorf("Yellowfin Error (%v) : %s", 20, "UNKNOWN_WEBSERVICE_FUNCTION")
var YFScheduleNull = fmt.Errorf("Yellowfin Error (%v) : %s", 36, "SCHEDULE_NULL")
var YFUnknownError = fmt.Errorf("Yellowfin Error (%v) : %s", -2, "UNKNOWN_ERROR")
var YFCouldNotLoadReportAccess = fmt.Errorf("Yellowfin Error (%v) : %s", 10, "COULD_NOT_LOAD_REPORT_ACCESS")
var YFReportIsDraft = fmt.Errorf("Yellowfin Error (%v) : %s", 24, "REPORT_IS_DRAFT")
var YFCouldNotReloadLicence = fmt.Errorf("Yellowfin Error (%v) : %s", 5, "COULD_NOT_RELOAD_LICENCE")
var YFCouldNotFindPerson = fmt.Errorf("Yellowfin Error (%v) : %s", 8, "COULD_NOT_FIND_PERSON")
var YFCouldNotFindReport = fmt.Errorf("Yellowfin Error (%v) : %s", 23, "COULD_NOT_FIND_REPORT")
var YFFiltervaluesFailed = fmt.Errorf("Yellowfin Error (%v) : %s", 33, "FILTERVALUES_FAILED")
var YFLoginMaximumAttempts = fmt.Errorf("Yellowfin Error (%v) : %s", 39, "LOGIN_MAXIMUM_ATTEMPTS")
var YFCannotConnect = fmt.Errorf("Yellowfin Error (%v) : %s", -1, "CANNOT_CONNECT")
var YFUserNotAuthenticated = fmt.Errorf("Yellowfin Error (%v) : %s", 1, "USER_NOT_AUTHENTICATED")
var YFNoWebserviceAccess = fmt.Errorf("Yellowfin Error (%v) : %s", 2, "NO_WEBSERVICE_ACCESS")
var YFInvalidClientReference = fmt.Errorf("Yellowfin Error (%v) : %s", 21, "INVALID_CLIENT_REFERENCE")
var YFClientExists = fmt.Errorf("Yellowfin Error (%v) : %s", 22, "CLIENT_EXISTS")
var YFUnsecureLogonNotEnabled = fmt.Errorf("Yellowfin Error (%v) : %s", 26, "UNSECURE_LOGON_NOT_ENABLED")
var YFCouldNotLoadFavourites = fmt.Errorf("Yellowfin Error (%v) : %s", 28, "COULD_NOT_LOAD_FAVOURITES")
var YFBroadcastFailed = fmt.Errorf("Yellowfin Error (%v) : %s", 32, "BROADCAST_FAILED")
var YFPersonRequired = fmt.Errorf("Yellowfin Error (%v) : %s", 3, "PERSON_REQUIRED")
var YFCouldNotCreatePerson = fmt.Errorf("Yellowfin Error (%v) : %s", 4, "COULD_NOT_CREATE_PERSON")
var YFBirtCouldNotSave = fmt.Errorf("Yellowfin Error (%v) : %s", 17, "BIRT_COULD_NOT_SAVE")
var YFPasswordRequirementsNotMet = fmt.Errorf("Yellowfin Error (%v) : %s", 38, "PASSWORD_REQUIREMENTS_NOT_MET")
var YFCouldNotLoadReportList = fmt.Errorf("Yellowfin Error (%v) : %s", 11, "COULD_NOT_LOAD_REPORT_LIST")
var YFCouldNotFindGroup = fmt.Errorf("Yellowfin Error (%v) : %s", 12, "COULD_NOT_FIND_GROUP")
var YFGroupExists = fmt.Errorf("Yellowfin Error (%v) : %s", 13, "GROUP_EXISTS")
var YFInvalidGroup = fmt.Errorf("Yellowfin Error (%v) : %s", 100, "INVALID_GROUP")
