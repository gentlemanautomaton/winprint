package printerattr

// Format maps flags to their string representations.
type Format map[Value]string

// FormatGo maps values to Go-style constant strings.
var FormatGo = Format{
	Queued:                        "Queued",
	Direct:                        "Direct",
	Default:                       "Default",
	Shared:                        "Shared",
	Network:                       "Network",
	Hidden:                        "Hidden",
	Local:                         "Local",
	EnableDevQ:                    "EnableDevQ",
	KeepPrintedJobs:               "KeepPrintedJobs",
	DoCompleteFirst:               "DoCompleteFirst",
	WorkOffline:                   "WorkOffline",
	BidirectionalCommunication:    "BidirectionalCommunication",
	RawOnly:                       "RawOnly",
	Published:                     "Published",
	Fax:                           "Fax",
	Redirected:                    "Redirected",
	PushedUser:                    "PushedUser",
	PushedMachine:                 "PushedMachine",
	Machine:                       "Machine",
	FriendlyName:                  "FriendlyName",
	TerminalServicesGenericDriver: "TerminalServicesGenericDriver",
	PerUser:                       "PerUser",
	EnterpriseCloud:               "EnterpriseCloud",
}
