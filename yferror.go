package yfws

import (
	"fmt"
)

var YFErrors = map[string]error{
	"29":  ErrYFResponseIsTooLarge,
	"31":  ErrYFEmptyRecipientList,
	"14":  ErrYFBirtObjectNull,
	"19":  ErrYFCouldNotUpdatePassword,
	"34":  ErrYFClientOrgsDisabled,
	"0":   ErrYFNoError,
	"8":   ErrYFCouldNotFindPerson,
	"21":  ErrYFInvalidClientReference,
	"26":  ErrYFUnsecureLogonNotEnabled,
	"-1":  ErrYFCannotConnect,
	"9":   ErrYFLicenceBreach,
	"15":  ErrYFBirtObjectNoData,
	"35":  ErrYFDashboardTabNotFound,
	"39":  ErrYFLoginMaximumAttempts,
	"100": ErrYFInvalidGroup,
	"5":   ErrYFCouldNotReloadLicence,
	"6":   ErrYFLoginAlreadyInUse,
	"33":  ErrYFFiltervaluesFailed,
	"37":  ErrYFUnknownStatusCode,
	"3":   ErrYFPersonRequired,
	"23":  ErrYFCouldNotFindReport,
	"7":   ErrYFCouldNotDeletePerson,
	"10":  ErrYFCouldNotLoadReportAccess,
	"18":  ErrYFBirtCouldNotSaveBirtFile,
	"25":  ErrYFCouldNotAuthenticateUser,
	"27":  ErrYFRoleNotFound,
	"30":  ErrYFSourceNotFound,
	"-2":  ErrYFUnknownError,
	"2":   ErrYFNoWebserviceAccess,
	"12":  ErrYFCouldNotFindGroup,
	"13":  ErrYFGroupExists,
	"22":  ErrYFClientExists,
	"24":  ErrYFReportIsDraft,
	"28":  ErrYFCouldNotLoadFavourites,
	"32":  ErrYFBroadcastFailed,
	"4":   ErrYFCouldNotCreatePerson,
	"11":  ErrYFCouldNotLoadReportList,
	"36":  ErrYFScheduleNull,
	"38":  ErrYFPasswordRequirementsNotMet,
	"17":  ErrYFBirtCouldNotSave,
	"20":  ErrYFUnknownWebserviceFunction,
	"1":   ErrYFUserNotAuthenticated,
	"16":  ErrYFBirtSourceMissing,
}
var ErrYFResponseIsTooLarge = fmt.Errorf("Yellowfin Error (%v) : %s", 29, "RESPONSE_IS_TOO_LARGE")
var ErrYFEmptyRecipientList = fmt.Errorf("Yellowfin Error (%v) : %s", 31, "EMPTY_RECIPIENT_LIST")
var ErrYFBirtObjectNull = fmt.Errorf("Yellowfin Error (%v) : %s", 14, "BIRT_OBJECT_NULL")
var ErrYFCouldNotUpdatePassword = fmt.Errorf("Yellowfin Error (%v) : %s", 19, "COULD_NOT_UPDATE_PASSWORD")
var ErrYFClientOrgsDisabled = fmt.Errorf("Yellowfin Error (%v) : %s", 34, "CLIENT_ORGS_DISABLED")
var ErrYFNoError = fmt.Errorf("Yellowfin Error (%v) : %s", 0, "NO_ERROR")
var ErrYFCouldNotFindPerson = fmt.Errorf("Yellowfin Error (%v) : %s", 8, "COULD_NOT_FIND_PERSON")
var ErrYFInvalidClientReference = fmt.Errorf("Yellowfin Error (%v) : %s", 21, "INVALID_CLIENT_REFERENCE")
var ErrYFUnsecureLogonNotEnabled = fmt.Errorf("Yellowfin Error (%v) : %s", 26, "UNSECURE_LOGON_NOT_ENABLED")
var ErrYFCannotConnect = fmt.Errorf("Yellowfin Error (%v) : %s", -1, "CANNOT_CONNECT")
var ErrYFLicenceBreach = fmt.Errorf("Yellowfin Error (%v) : %s", 9, "LICENCE_BREACH")
var ErrYFBirtObjectNoData = fmt.Errorf("Yellowfin Error (%v) : %s", 15, "BIRT_OBJECT_NO_DATA")
var ErrYFDashboardTabNotFound = fmt.Errorf("Yellowfin Error (%v) : %s", 35, "DASHBOARD_TAB_NOT_FOUND")
var ErrYFLoginMaximumAttempts = fmt.Errorf("Yellowfin Error (%v) : %s", 39, "LOGIN_MAXIMUM_ATTEMPTS")
var ErrYFInvalidGroup = fmt.Errorf("Yellowfin Error (%v) : %s", 100, "INVALID_GROUP")
var ErrYFCouldNotReloadLicence = fmt.Errorf("Yellowfin Error (%v) : %s", 5, "COULD_NOT_RELOAD_LICENCE")
var ErrYFLoginAlreadyInUse = fmt.Errorf("Yellowfin Error (%v) : %s", 6, "LOGIN_ALREADY_IN_USE")
var ErrYFFiltervaluesFailed = fmt.Errorf("Yellowfin Error (%v) : %s", 33, "FILTERVALUES_FAILED")
var ErrYFUnknownStatusCode = fmt.Errorf("Yellowfin Error (%v) : %s", 37, "UNKNOWN_STATUS_CODE")
var ErrYFPersonRequired = fmt.Errorf("Yellowfin Error (%v) : %s", 3, "PERSON_REQUIRED")
var ErrYFCouldNotFindReport = fmt.Errorf("Yellowfin Error (%v) : %s", 23, "COULD_NOT_FIND_REPORT")
var ErrYFCouldNotDeletePerson = fmt.Errorf("Yellowfin Error (%v) : %s", 7, "COULD_NOT_DELETE_PERSON")
var ErrYFCouldNotLoadReportAccess = fmt.Errorf("Yellowfin Error (%v) : %s", 10, "COULD_NOT_LOAD_REPORT_ACCESS")
var ErrYFBirtCouldNotSaveBirtFile = fmt.Errorf("Yellowfin Error (%v) : %s", 18, "BIRT_COULD_NOT_SAVE_BIRT_FILE")
var ErrYFCouldNotAuthenticateUser = fmt.Errorf("Yellowfin Error (%v) : %s", 25, "COULD_NOT_AUTHENTICATE_USER")
var ErrYFRoleNotFound = fmt.Errorf("Yellowfin Error (%v) : %s", 27, "ROLE_NOT_FOUND")
var ErrYFSourceNotFound = fmt.Errorf("Yellowfin Error (%v) : %s", 30, "SOURCE_NOT_FOUND")
var ErrYFUnknownError = fmt.Errorf("Yellowfin Error (%v) : %s", -2, "UNKNOWN_ERROR")
var ErrYFNoWebserviceAccess = fmt.Errorf("Yellowfin Error (%v) : %s", 2, "NO_WEBSERVICE_ACCESS")
var ErrYFCouldNotFindGroup = fmt.Errorf("Yellowfin Error (%v) : %s", 12, "COULD_NOT_FIND_GROUP")
var ErrYFGroupExists = fmt.Errorf("Yellowfin Error (%v) : %s", 13, "GROUP_EXISTS")
var ErrYFClientExists = fmt.Errorf("Yellowfin Error (%v) : %s", 22, "CLIENT_EXISTS")
var ErrYFReportIsDraft = fmt.Errorf("Yellowfin Error (%v) : %s", 24, "REPORT_IS_DRAFT")
var ErrYFCouldNotLoadFavourites = fmt.Errorf("Yellowfin Error (%v) : %s", 28, "COULD_NOT_LOAD_FAVOURITES")
var ErrYFBroadcastFailed = fmt.Errorf("Yellowfin Error (%v) : %s", 32, "BROADCAST_FAILED")
var ErrYFCouldNotCreatePerson = fmt.Errorf("Yellowfin Error (%v) : %s", 4, "COULD_NOT_CREATE_PERSON")
var ErrYFCouldNotLoadReportList = fmt.Errorf("Yellowfin Error (%v) : %s", 11, "COULD_NOT_LOAD_REPORT_LIST")
var ErrYFScheduleNull = fmt.Errorf("Yellowfin Error (%v) : %s", 36, "SCHEDULE_NULL")
var ErrYFPasswordRequirementsNotMet = fmt.Errorf("Yellowfin Error (%v) : %s", 38, "PASSWORD_REQUIREMENTS_NOT_MET")
var ErrYFBirtCouldNotSave = fmt.Errorf("Yellowfin Error (%v) : %s", 17, "BIRT_COULD_NOT_SAVE")
var ErrYFUnknownWebserviceFunction = fmt.Errorf("Yellowfin Error (%v) : %s", 20, "UNKNOWN_WEBSERVICE_FUNCTION")
var ErrYFUserNotAuthenticated = fmt.Errorf("Yellowfin Error (%v) : %s", 1, "USER_NOT_AUTHENTICATED")
var ErrYFBirtSourceMissing = fmt.Errorf("Yellowfin Error (%v) : %s", 16, "BIRT_SOURCE_MISSING")
