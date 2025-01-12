package main

import (
	"fmt"
	"os"

	"github.com/gentlemanautomaton/winprint/driverinfo"
	"github.com/gentlemanautomaton/winprint/monitorinfo"
	"github.com/gentlemanautomaton/winprint/portinfo"
	"github.com/gentlemanautomaton/winprint/printerattr"
	"github.com/gentlemanautomaton/winprint/printerenum"
	"github.com/gentlemanautomaton/winprint/printerinfo"
	"github.com/gentlemanautomaton/winprint/printerstatus"
	"github.com/gentlemanautomaton/winprint/spoolerapi"
)

func usage() {
	fmt.Printf("Usage:\n  %s <command>\n    where <command> is one of: printers, drivers, ports, monitors\n", os.Args[0])
}

func main() {
	var cmd string
	switch len(os.Args) {
	case 1:
		cmd = "printers"
	case 2:
		cmd = os.Args[1]
	default:
		usage()
		os.Exit(1)
	}

	switch cmd {
	case "printers":
		printers, err := spoolerapi.EnumPrinters[printerinfo.Level2](printerenum.Local|printerenum.Connections, "")
		if err != nil {
			fmt.Printf("Failed to enumerate printers: %v\n", err)
		} else if len(printers) > 0 {
			fmt.Printf("Printers:\n")
			for i, printer := range printers {
				fmt.Printf("  %d: %s\n", i, printer.Name)
				fmt.Printf("    Server: %s\n", printer.Server)
				fmt.Printf("    ShareName: %s\n", printer.ShareName)
				fmt.Printf("    Port: %s\n", printer.Port)
				fmt.Printf("    Driver: %s\n", printer.Driver)
				fmt.Printf("    Comment: %s\n", printer.Comment)
				fmt.Printf("    Location: %s\n", printer.Location)
				fmt.Printf("    SeparatorFile: %s\n", printer.SeparatorFile)
				fmt.Printf("    PrintProcessor: %s\n", printer.PrintProcessor)
				fmt.Printf("    DataType: %s\n", printer.DataType)
				fmt.Printf("    Parameters: %s\n", printer.Parameters)
				fmt.Printf("    Attributes: %s\n", printer.Attributes.Join(", ", printerattr.FormatGo))
				fmt.Printf("    Priority: %d\n", printer.Priority)
				fmt.Printf("    DefaultPriority: %d\n", printer.DefaultPriority)
				fmt.Printf("    StartTime: %d\n", printer.StartTime)
				fmt.Printf("    UntilTime: %d\n", printer.UntilTime)
				fmt.Printf("    Status: %d\n", printer.Status)
				fmt.Printf("    Status: %s\n", printer.Status.Join(", ", printerstatus.FormatGo))
				fmt.Printf("    Jobs: %d\n", printer.Jobs)
				fmt.Printf("    AveragePPM: %d\n", printer.AveragePPM)
			}
		}
	case "drivers":
		drivers, err := spoolerapi.EnumPrinterDrivers[driverinfo.Level2]("", "")
		if err != nil {
			fmt.Printf("Failed to enumerate printer drivers: %v\n", err)
			os.Exit(1)
		} else if len(drivers) > 0 {
			fmt.Printf("Printer Drivers:\n")
			for i, driver := range drivers {
				fmt.Printf("  %d: %s (Type %d, %s)\n", i, driver.Name, driver.Version, driver.Environment)
			}
		}
	case "ports":
		ports, err := spoolerapi.EnumPorts[portinfo.Level2]("")
		if err != nil {
			fmt.Printf("Failed to enumerate printer ports: %v\n", err)
		} else if len(ports) > 0 {
			fmt.Printf("Printer Ports:\n")
			for i, port := range ports {
				fmt.Printf("  %d: %s (type: %s, monitor: \"%s\", description: \"%s\")\n", i, port.Name, port.Type, port.Monitor, port.Description)
			}
		}
	case "monitors":
		monitors, err := spoolerapi.EnumMonitors[monitorinfo.Level2]("")
		if err != nil {
			fmt.Printf("Failed to enumerate printer ports: %v\n", err)
		} else if len(monitors) > 0 {
			fmt.Printf("Printer Port Monitors:\n")
			for i, port := range monitors {
				fmt.Printf("  %d: %s (%s, \"%s\")\n", i, port.Name, port.Environment, port.Library)
			}
		}
	default:
		usage()
	}
}
