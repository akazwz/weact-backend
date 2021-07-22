package influxdb

import (
	"github.com/influxdata/influxdb-client-go/v2"
	"time"
)

// You can generate a Token from the "Tokens Tab" in the UI
const token = "h74rCcT9oAvtpx9H00rpmjmwDy_0htRqovKuHkg1dXTn7tJuCDdh19uJGId3QeOTpp-QsVIu-_DGZBrCpIVNIw=="
const bucket = "weibo"
const org = "akazwz"

func initClient() influxdb2.Client {
	client := influxdb2.NewClient("http://hellozwz.com:8086", token)
	// always close client at the end
	defer client.Close()
	return client
}

func WriteLineData(str string) {
	client := initClient()
	writeAPI := client.WriteAPI(org, bucket)
	writeAPI.WriteRecord(str)
	writeAPI.Flush()
}

type InfluxPoint struct {
	tags   map[string]string
	fields map[string]interface{}
	ts     time.Time
}

func WritePointData(measurement string, tags map[string]string, fields map[string]interface{}, ts time.Time) {
	client := initClient()
	writeAPI := client.WriteAPI(org, bucket)
	point := influxdb2.NewPoint(measurement, tags, fields, ts)
	writeAPI.WritePoint(point)
	writeAPI.Flush()
}
