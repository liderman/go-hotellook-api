# go-hotellook-api
Golang implementation Hotellook API for static data access

[![GoDoc](https://godoc.org/github.com/liderman/go-hotellook-api?status.svg)](https://godoc.org/github.com/liderman/go-hotellook-api)

Installation
-----------
	go get github.com/liderman/go-hotellook-api

Usage
-----------
Creates a new instance HotellookApi:
```go
hotelApi := hotellook.NewHotellookApi("YOUR_TOKEN")
```

Getting a list of hotels for locationId `895`:
```go
hotels, err := hotelApi.Hotels(895)
```

Requirements
-----------

* Need at least `go1.5` or newer.

Documentation
-----------

You can read package documentation [here](http:godoc.org/github.com/liderman/go-hotellook-api).