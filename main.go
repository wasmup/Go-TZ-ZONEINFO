package main

import (
	"fmt"
	"os"
	"strings"
	"time"
)

func main() {
	fmt.Println("Local =", time.Local.String())
	fmt.Println("Now =", time.Now())
	fmt.Println(strings.Repeat("_", 80))
	if tz := os.Getenv("TZ"); len(tz) > 0 {
		z := os.Getenv("ZONEINFO")
		if _, err := os.Stat(z); err != nil {
			fmt.Println(err)
		}
		fmt.Printf("TZ =%q\n", tz)
		fmt.Printf("ZONEINFO =%q\n", z)

		// Local represents the system's local time zone.
		// On Unix systems, Local consults the TZ environment
		// variable to find the time zone to use. No TZ means
		// use the system default /etc/localtime.
		// TZ="" means use UTC.
		// TZ="foo" means use file foo in the system timezone directory.
		fmt.Println("Local =", time.Local.String())
		fmt.Println("Now =", time.Now())

		// LoadLocation returns the Location with the given name.
		//
		// If the name is "" or "UTC", LoadLocation returns UTC.
		// If the name is "Local", LoadLocation returns Local.
		//
		// Otherwise, the name is taken to be a location name corresponding to a file
		// in the IANA Time Zone database, such as "America/New_York".
		//
		// LoadLocation looks for the IANA Time Zone database in the following
		// locations in order:
		//
		//   - the directory or uncompressed zip file named by the ZONEINFO environment variable
		//   - on a Unix system, the system standard installation location
		//   - $GOROOT/lib/time/zoneinfo.zip
		//   - the time/tzdata package, if it was imported
		Location, err := time.LoadLocation(tz)
		if err != nil {
			fmt.Println(err)
		} else {
			time.Local = Location
		}
		fmt.Println("Local =", time.Local.String())
		fmt.Println("Now =", time.Now())
	}
}
