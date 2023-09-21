package scan

// data from https://www.iana.org/assignments/service-names-port-numbers/service-names-port-numbers.csv
var knownPorts = map[int]string{
	// 80: "http",
	554: "rtsp",
	// 8080: "http-alt",
	8554: "rtsp-alt",
}
