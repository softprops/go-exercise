package bufferedio

import (
	"bytes"
	"io"
	"reflect"
	"testing"
)

func TestNew(t *testing.T) {
	type args struct {
		cap int
	}
	tests := []struct {
		name        string
		args        args
		want        *BufWriter
		wantWrapped string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			wrapped := &bytes.Buffer{}
			if got := New(wrapped, tt.args.cap); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("New() = %v, want %v", got, tt.want)
			}
			if gotWrapped := wrapped.String(); gotWrapped != tt.wantWrapped {
				t.Errorf("New() = %v, want %v", gotWrapped, tt.wantWrapped)
			}
		})
	}
}

func TestBufWriter_Write(t *testing.T) {
	type fields struct {
		cap     int
		buf     []byte
		wrapped io.Writer
		pos     int
	}
	type args struct {
		data []byte
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			bw := &BufWriter{
				cap:     tt.fields.cap,
				buf:     tt.fields.buf,
				wrapped: tt.fields.wrapped,
				pos:     tt.fields.pos,
			}
			bw.Write(tt.args.data)
		})
	}
}

func TestBufWriter_Flush(t *testing.T) {
	type fields struct {
		cap     int
		buf     []byte
		wrapped io.Writer
		pos     int
	}
	tests := []struct {
		name   string
		fields fields
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			bw := &BufWriter{
				cap:     tt.fields.cap,
				buf:     tt.fields.buf,
				wrapped: tt.fields.wrapped,
				pos:     tt.fields.pos,
			}
			bw.Flush()
		})
	}
}

func TestFlow(t *testing.T) {
	writer := bytes.NewBufferString("")
	bw := New(writer, 5)
	bw.Write([]byte("hello world"))
	bw.Flush()
	got := writer.String()
	if !reflect.DeepEqual(got, "hello world") {
		t.Errorf("expected 'hello world' but got '%s'", got)
	}
}
