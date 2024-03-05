package httpmetrics

import (
	"reflect"
	"testing"
)

func TestParse(t *testing.T) {
	type args struct {
		line string
	}
	tests := []struct {
		name    string
		args    args
		want    *Log
		wantErr bool
	}{
		{
			name: "a",
			args: args{
				line: `127.0.0.1 - frank [10/Oct/2000:13:55:36 -0700] "GET /apache_pb.gif HTTP/1.0" 200 2326 "http://www.example.com/start.html" "Mozilla/4.08 [en] (Win98; I ;Nav)"`,
			},
			want: &Log{
				ip:         "127.0.0.1",
				id:         "-",
				user:       "frank",
				ts:         "10/Oct/2000:13:55:36 -0700",
				verb:       "GET",
				path:       "/apache_pb.gif",
				version:    "HTTP/1.0",
				statusCode: 200,
				bytes:      2326,
				referer:    `"http://www.example.com/start.html"`,
				agent:      `"Mozilla/4.08 [en] (Win98; I ;Nav)"`,
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := Parse(tt.args.line)
			if (err != nil) != tt.wantErr {
				t.Errorf("Parse() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Parse()\n  got: %#v\n want: %#v", got, tt.want)
			}
		})
	}
}
