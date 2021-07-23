package service

import (
	"time"
	"weact-backend/influxdb"
)

func WriteHotSearch() {
	measurement := "hot-search"
	tags := make(map[string]string)
	tags["type"] = "hot-search"
	fields := make(map[string]interface{})
	fields["avg"] = 50
	fields["max"] = 100

	influxdb.WritePointData(measurement, tags, fields, time.Now())
}
