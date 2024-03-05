package httpmetrics

import (
	"fmt"
	"regexp"
	"strconv"
)

// https://en.wikipedia.org/wiki/Common_Log_Format
// 127.0.0.1 user-identifier frank [10/Oct/2000:13:55:36 -0700] "GET /apache_pb.gif HTTP/1.0" 200 2326
var CommonLogFmt = regexp.MustCompile("^(?<ip>\\S+) (?<id>\\S+) (?<user>\\S+) \\[(?<ts>[^]]+)\\] \"(?<verb>\\S+) (?<path>\\S+) (?<version>\\S+)\" (?<statusCode>\\d+) (?<bytes>\\d+) (?<referer>\\S+) (?<agent>.+)")

type Log struct {
	ip         string
	id         string
	user       string
	ts         string
	verb       string
	path       string
	version    string
	statusCode int
	bytes      int
	referer    string
	agent      string
}

func Parse(line string) (*Log, error) {
	if !CommonLogFmt.MatchString(line) {
		return nil, fmt.Errorf("invalid line %#v", line)
	}

	match := CommonLogFmt.FindStringSubmatch(line)
	log := Log{}
	for i, name := range CommonLogFmt.SubexpNames() {
		if i != 0 {
			val := match[i]
			switch name {
			case "ip":
				log.ip = val
			case "id":
				log.id = val
			case "user":
				log.user = val
			case "ts":
				log.ts = val
			case "verb":
				log.verb = val
			case "path":
				log.path = val
			case "version":
				log.version = val
			case "statusCode":
				i, err := strconv.Atoi(val)
				if err != nil {
					return nil, fmt.Errorf("invalid statusCode %#v", val)
				}
				log.statusCode = i
			case "bytes":
				i, err := strconv.Atoi(val)
				if err != nil {
					return nil, fmt.Errorf("invalid bytes %#v", val)
				}
				log.bytes = i
			case "referer":
				log.referer = val
			case "agent":
				log.agent = val
			default:
			}
		}
	}
	return &log, nil
}
