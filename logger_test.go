package Logger

import (
	"bytes"
	"fmt"
	"io"
	"log"
	"os"
	"regexp"
	"strings"
	"testing"
)

func init() {

}

func TestWARN(t *testing.T) {
	type args struct {
		str  string
		args []interface{}
	}
	tests := []struct {
		name     string
		args     args
		expected string
	}{
		{
			name: "Multiple placeholders",
			args: args{
				str:  "This is a test: {}",
				args: []interface{}{"first"},
			},
			expected: "^.*WARN.*" + "This is a test: first" + "$",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var buf bytes.Buffer
			log.SetOutput(&buf)

			reader, writer, _ := os.Pipe()
			originalStdout := os.Stdout
			os.Stdout = writer

			buf.Reset()

			WARN(tt.args.str, tt.args.args...)

			err := writer.Close()
			if err != nil {
				panic(err.Error())
			}
			os.Stdout = originalStdout

			_, err = io.Copy(&buf, reader)
			if err != nil {
				panic(err.Error())
			}
			err = reader.Close()
			if err != nil {
				panic(err.Error())
			}

			logged := buf.String()
			logged = strings.TrimSuffix(logged, "\n")

			isPresent, _ := regexp.MatchString(tt.expected, logged)
			if !isPresent {
				t.Errorf("[%s] Expected log to contain %q, but got %q", tt.name, tt.expected, logged)
			} else {
				fmt.Printf("[%s] has passed\n", tt.name)
			}
		})
	}
}

func TestINFO(t *testing.T) {
	type args struct {
		str  string
		args []interface{}
	}
	tests := []struct {
		name     string
		args     args
		expected string
	}{
		{
			name: "Multiple placeholders",
			args: args{
				str:  "This is a test: {} and {}",
				args: []interface{}{"first", "second"},
			},
			expected: "^.*INFO.*" + "This is a test: first and second" + "$",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var buf bytes.Buffer
			log.SetOutput(&buf)

			reader, writer, _ := os.Pipe()
			originalStdout := os.Stdout
			os.Stdout = writer

			buf.Reset()

			INFO(tt.args.str, tt.args.args...)

			err := writer.Close()
			if err != nil {
				panic(err.Error())
			}
			os.Stdout = originalStdout

			_, err = io.Copy(&buf, reader)
			if err != nil {
				panic(err.Error())
			}
			err = reader.Close()
			if err != nil {
				panic(err.Error())
			}

			logged := buf.String()
			logged = strings.TrimSuffix(logged, "\n")

			isPresent, _ := regexp.MatchString(tt.expected, logged)
			if !isPresent {
				t.Errorf("[%s] Expected log to contain %q, but got %q", tt.name, tt.expected, logged)
			} else {
				fmt.Printf("[%s] has passed\n", tt.name)
			}
		})
	}
}

func TestERROR(t *testing.T) {
	type args struct {
		str  string
		args []interface{}
	}
	tests := []struct {
		name     string
		args     args
		expected string
	}{
		{
			name: "Multiple placeholders",
			args: args{
				str:  "This is a test: {} and {}",
				args: []interface{}{"first", "second"},
			},
			expected: "^.*ERROR.*" + "This is a test: first and second" + "$",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var buf bytes.Buffer
			log.SetOutput(&buf)

			reader, writer, _ := os.Pipe()
			originalStdout := os.Stdout
			os.Stdout = writer

			buf.Reset()

			ERROR(tt.args.str, tt.args.args...)

			err := writer.Close()
			if err != nil {
				panic(err.Error())
			}
			os.Stdout = originalStdout

			_, err = io.Copy(&buf, reader)
			if err != nil {
				panic(err.Error())
			}
			err = reader.Close()
			if err != nil {
				panic(err.Error())
			}

			logged := buf.String()
			logged = strings.TrimSuffix(logged, "\n")

			isPresent, _ := regexp.MatchString(tt.expected, logged)
			if !isPresent {
				t.Errorf("[%s] Expected log to contain %q, but got %q", tt.name, tt.expected, logged)
			}
		})
	}
}
