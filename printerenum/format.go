package printerenum

// Format maps flags to their string representations.
type Format map[Flags]string

// FormatGo maps values to Go-style constant strings.
var FormatGo = Format{
	Default:            "Default",
	Local:              "Local",
	Connections:        "Connections",
	Name:               "Name",
	Remote:             "Remote",
	Shared:             "Shared",
	Network:            "Network",
	ExpansionSupported: "ExpansionSupported",
	Container:          "Container",
	Icon1:              "Icon1",
	Icon2:              "Icon2",
	Icon3:              "Icon3",
	Icon4:              "Icon4",
	Icon5:              "Icon5",
	Icon6:              "Icon6",
	Icon7:              "Icon7",
	Icon8:              "Icon8",
	Hide:               "Hide",
	CategoryAll:        "CategoryAll",
	Category3D:         "Category3D",
}
