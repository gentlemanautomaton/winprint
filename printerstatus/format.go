package printerstatus

// Format maps flags to their string representations.
type Format map[Value]string

// FormatGo maps values to Go-style constant strings.
var FormatGo = Format{
	Paused:                 "Paused",
	Error:                  "Error",
	PendingDeletion:        "PendingDeletion",
	PaperJam:               "PaperJam",
	PaperOut:               "PaperOut",
	ManualFeed:             "ManualFeed",
	PaperProblem:           "PaperProblem",
	Offline:                "Offline",
	ActiveIO:               "ActiveIO",
	Busy:                   "Busy",
	Printing:               "Printing",
	OutputBinFull:          "OutputBinFull",
	NotAvailable:           "NotAvailable",
	Waiting:                "Waiting",
	Processing:             "Processing",
	Initializing:           "Initializing",
	WarmingUp:              "WarmingUp",
	TonerLow:               "TonerLow",
	NoToner:                "NoToner",
	PageNotPrintable:       "PageNotPrintable",
	UserInterventionNeeded: "UserInterventionNeeded",
	OutOfMemory:            "OutOfMemory",
	DoorOpen:               "DoorOpen",
	ServerUnknown:          "ServerUnknown",
	PowerSave:              "PowerSave",
	ServerOffline:          "ServerOffline",
	DriverUpdateNeeded:     "DriverUpdateNeeded",
}
