package main

import (
	"distance_calc/model"
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"path/filepath"
	"strconv"
	"strings"
	"time"
)

type distance struct{}

func (d *distance) Help() string {
	txt := fmt.Sprintf(`
	Usage:
		%s %s

	Description:
		%s
	`,
		"timeline", "distance",
		d.Synopsis(),
	)

	return txt[1:]
}

func (d *distance) Synopsis() string {
	return "Aggregate distance of Google Maps timeline."
}

func (d *distance) Run(args []string) int {

	flagSet := flag.NewFlagSet("timeline distance", flag.ExitOnError)
	var (
		dayFrom = flagSet.String("from", "2019-11-01", "day from")
		dayTo   = flagSet.String("to", "2019-11-02", "day to")
		actType = flagSet.String("type", "IN_PASSENGER_VEHICLE MOTORCYCLING", "actitity type")
		root    = flagSet.String("dir", "./", "path to json file")
	)
	flagSet.Parse(args)

	from, err := time.ParseInLocation("2006-01-02", *dayFrom, time.Local)
	if err != nil {
		log.Fatal(err)
		return 1
	}
	to, err := time.ParseInLocation("2006-01-02", *dayTo, time.Local)
	if err != nil {
		log.Fatal(err)
		return 1
	}
	to = to.AddDate(0, 0, 1)

	actTypes := strings.Split(*actType, " ")

	files, err := filepath.Glob(*root + "/*")
	if err != nil {
		log.Fatal(err)
		return 1
	}

	var aggregate int
	for _, file := range files {
		aggregate += calc(file, from, to, actTypes)

	}

	fmt.Printf("Aggregate Distance: %v km\n", aggregate/1000)

	return 0
}

func calc(path string, from, to time.Time, actTypes []string) int {
	var aggregate int

	file, err := ioutil.ReadFile(path)
	if err != nil {
		log.Fatal(err)
	}

	timeLines := &model.Timeline{}

	err = json.Unmarshal(file, timeLines)
	if err != nil {
		log.Fatal(err)
	}

	for _, tl := range timeLines.TimelineObjects {
		startTimeMs, _ := strconv.ParseInt(tl.StartTimestampMs, 10, 64)
		start := startTimeMs / 1000

		if start >= from.Unix() {
			if start > to.Unix() {
				break
			}
			for _, at := range actTypes {
				if tl.ActivityType == at {
					aggregate += tl.Distance
				}
			}
			continue
		}
	}

	return aggregate
}
