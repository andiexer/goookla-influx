# GoOokla Speedtest to InfluxDB

this is a small golang app that will monitor continuously your internet speed with ookla speedtest and saves the  data to influxDB.

> WORK IN PROGRESS! this is a project to get in touch with golang

## Usage

| Environment Variable 	| Description                                                      	| Default Value 	|
|----------------------	|------------------------------------------------------------------	|---------------	|
| SINK                 	| where the data should be sent. possible values: <br> **console, influxdbv1, influxdbv2** | console            |
| SERVER_ID             | ookla speedtest server id if you want to use always the same server for tests. if empty, speedtest will auto negotiate the nearest server                                | ""           |
| INTERVAL             |  wait between each speedtest in seconds                            | 900          |
| INFLUX_HOST            | **only for influxdb sinks** <br> influxdb host                            | http://localhost:8086         |
| INFLUX_V1_USER            | **only for influxdbv1 sink** <br> influxdb user to connect                          | ""        |
| INFLUX_V1_PASSWORD            | **only for influxdbv1 sink** <br> influxdb password to connect                           | ""        |
| INFLUX_V2_ORG            | **only for influxdbv2 sink** <br> organization for influxdb >= 1.8                         | "devlabs"        |
| INFLUX_V2_BUCKET          | **only for influxdbv2 sink** <br> data bucket for influxdb >= 1.8                         | "speedtest"        |
| INFLUX_V2_AUTHTOKEN         | **only for influxdbv2 sink** <br> auth token for influxdb >= 1.8                         | ""        |