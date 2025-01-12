package main

import (
	"fmt"
	"os"
	"strings"
	"syscall"

	"github.com/gentlemanautomaton/winprint"
	"github.com/gentlemanautomaton/winprint/accessoptions"
	"github.com/gentlemanautomaton/winprint/accessrights"
	"github.com/gentlemanautomaton/winprint/spoolerapi"
	"github.com/gentlemanautomaton/winprint/tcpipprinter"
	"github.com/gentlemanautomaton/winprint/tcpipprinter/portdata"
	"github.com/gentlemanautomaton/winprint/tcpipprinter/portproto"
)

func usage() {
	commands := []string{
		"install-printer <name>",
		"delete-printer  <name>",
		"install-driver  <inf-file-path>",
		"delete-driver   <driver-name>",
		"add-ip-port     <address>",
	}
	for i := range commands {
		commands[i] = os.Args[0] + " " + commands[i]
	}
	fmt.Printf("Usage:\n  %s\n", strings.Join(commands, "\n  "))
}

func main() {
	if len(os.Args) < 2 {
		usage()
		os.Exit(1)
	}

	switch cmd := os.Args[1]; cmd {
	case "delete-printer":
		if len(os.Args) != 3 {
			fmt.Printf("Expected printer name:\n  %s delete-printer <printer-name>\n", os.Args[0])
			os.Exit(1)
		}
		printerName := os.Args[2]
		if printerName == "" {
			fmt.Printf("An empty printer name was provided.\n")
			os.Exit(1)
		}
		fmt.Printf("Subcommand 'delete-printer': %s\n", printerName)
		printer, err := winprint.OpenPrinter(printerName, accessrights.FullPrinterAccess, accessoptions.NoCache)
		if err != nil {
			fmt.Printf("Failed to open printer \"%s\": %v\n", printerName, err)
			os.Exit(1)
		}
		defer printer.Close()
		if err := printer.Delete(); err != nil {
			fmt.Printf("Failed to remove printer \"%s\": %v\n", printerName, err)
			os.Exit(1)
		}
		fmt.Printf("Printer \"%s\" has been marked for deletion.\n", printerName)
	case "install-driver":
		if len(os.Args) != 3 {
			fmt.Printf("Expected inf file path:\n  %s install-driver <inf-file-path>\n", os.Args[0])
			os.Exit(1)
		}
		infPath := os.Args[2]
		if fi, err := os.Stat(infPath); err != nil {
			fmt.Printf("Unable to open inf file: %v\n", err)
			os.Exit(1)
		} else if !fi.Mode().IsRegular() {
			fmt.Printf("%s is not an inf file.\n", infPath)
			os.Exit(1)
		}
		fmt.Printf("subcommand 'install-driver': %s\n", infPath)
		if false {
			spoolerapi.UploadPrinterDriverPackage("", infPath, "Windows x64", 0, syscall.InvalidHandle)
		}
	case "delete-driver":
		if len(os.Args) != 3 {
			fmt.Printf("Expected driver name:\n  %s delete <driver-name>\n", os.Args[0])
			os.Exit(1)
		}
		driverName := os.Args[2]
		if driverName == "" {
			fmt.Printf("an empty driver name was provided.\n")
			os.Exit(1)
		}
		fmt.Printf("subcommand 'delete-driver': %s\n", driverName)
		if err := spoolerapi.DeletePrinterDriver("", "Windows x64", driverName, 0x04, 3); err != nil {
			fmt.Printf("Failed to remove \"%s\" printer driver: %v\n", driverName, err)
			os.Exit(1)
		}
		fmt.Printf("Printer driver \"%s\" removed successfully.\n", driverName)
	case "add-ip-port":
		if len(os.Args) != 3 {
			fmt.Printf("Expected port host address:\n  %s add-ip-port <address>\n", os.Args[0])
			os.Exit(1)
		}
		hostAddress := os.Args[2]
		if hostAddress == "" {
			fmt.Printf("an empty host address was provided.\n")
			os.Exit(1)
		}
		fmt.Printf("subcommand 'add-ip-port': %s\n", hostAddress)
		mon, err := tcpipprinter.OpenMonitor(accessrights.AdministerServer)
		if err != nil {
			fmt.Printf("Failed to add \"%s\" port: %v\n", hostAddress, err)
			os.Exit(1)
		}
		defer mon.Close()
		err = mon.AddPort(portdata.Level2{
			Name:            hostAddress,
			HostAddress:     hostAddress,
			Protocol:        portproto.RawTCP,
			PortNumber:      9100,
			SNMPEnabled:     true,
			SNMPCommunity:   "public",
			SNMPDeviceIndex: 1,
		})
		if err != nil {
			fmt.Printf("Failed to add \"%s\" port: %v\n", hostAddress, err)
			os.Exit(1)
		}
		fmt.Printf("TCP/IP printer port \"%s\" added successfully.\n", hostAddress)
	case "delete-ip-port":
		if len(os.Args) != 3 {
			fmt.Printf("Expected port host address:\n  %s delete-ip-port <address>\n", os.Args[0])
			os.Exit(1)
		}
		hostAddress := os.Args[2]
		if hostAddress == "" {
			fmt.Printf("an empty host address was provided.\n")
			os.Exit(1)
		}
		fmt.Printf("subcommand 'delete-ip-port': %s\n", hostAddress)
		mon, err := tcpipprinter.OpenMonitor(accessrights.AdministerServer)
		if err != nil {
			fmt.Printf("Failed to delete \"%s\" port: %v\n", hostAddress, err)
			os.Exit(1)
		}
		defer mon.Close()
		err = mon.DeletePort(hostAddress)
		if err != nil {
			fmt.Printf("Failed to delete \"%s\" port: %v\n", hostAddress, err)
			os.Exit(1)
		}
		fmt.Printf("TCP/IP printer port \"%s\" deleted successfully.\n", hostAddress)
	default:
		fmt.Printf("unrecognized command '%s'\n", cmd)
		os.Exit(1)
	}
}
