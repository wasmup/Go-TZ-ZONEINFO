# Code

```go
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
```

---

# Build and run

```sh
	time go build -o ./app -ldflags "-s -w" main.go
# real    0m0.156s
# user    0m0.139s
# sys     0m0.053s

	time ./app
# Local = Local
# Now = 2023-02-23 13:43:47.815432846 +0000 GMT m=+0.000149244
# ________________________________________________________________________________

# real    0m0.014s
# user    0m0.003s
# sys     0m0.001s

    time docker image build -t app:1.0.2 .
# real    0m3.601s
# user    0m0.136s
# sys     0m0.091s

	time docker run app:1.0.2
# Local = UTC
# Now = 2023-02-23 13:44:18.755858882 +0000 UTC m=+0.000113243
# ________________________________________________________________________________
# TZ ="America/New_York"
# ZONEINFO ="/zoneinfo.zip"
# Local = UTC
# Now = 2023-02-23 13:44:18.755921024 +0000 UTC m=+0.000175371
# Local = America/New_York
# Now = 2023-02-23 08:44:18.75601801 -0500 EST m=+0.000272364

# real    0m0.761s
# user    0m0.065s
# sys     0m0.034s

	time docker image build -t app:1.0.3 -f Dockerfile_scratch .
# real    0m3.885s
# user    0m0.183s
# sys     0m0.056s
	time docker run app:1.0.3
# Local = UTC
# Now = 2023-02-23 13:45:12.919103053 +0000 UTC m=+0.000091008
# ________________________________________________________________________________
# TZ ="America/New_York"
# ZONEINFO ="/zoneinfo.zip"
# Local = UTC
# Now = 2023-02-23 13:45:12.919164237 +0000 UTC m=+0.000152190
# Local = America/New_York
# Now = 2023-02-23 08:45:12.919240378 -0500 EST m=+0.000228329

# real    0m0.623s
# user    0m0.052s
# sys     0m0.040s

	time docker image build -t app:1.0.4 -f Dockerfile_no_ENV .
	time docker run app:1.0.4
# Local = UTC
# Now = 2023-02-23 13:52:57.688322309 +0000 UTC m=+0.000103221
# ________________________________________________________________________________

# real    0m0.664s
# user    0m0.061s
# sys     0m0.025s
```
