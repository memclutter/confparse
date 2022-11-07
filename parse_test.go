package confparse

import (
	"flag"
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
	"time"
)

type testParseContainer struct {
	Addr        string        `name:"addr" value:":8000" usage:"Listen and serve address"`
	DatabaseUrl string        `name:"databaseUrl" value:"mongodb://localhost:27017/db" usage:"Database connection url"`
	Timeout     time.Duration `name:"timeout" value:"200ms" usage:"Timeout value"`
	OptionalVal string        `name:"optional"`
	ApiKey      string        `name:"apiKey" envVar:"API_KEY" usage:"API key"`
	BatchSize   int           `name:"batchSize" envVar:"BATCH_SIZE" usage:"Batch size for query"`
	MaxCount    int64         `name:"maxCount"`
	Debug       bool          `name:"debug" envVar:"DEBUG"`
	Profit      uint          `name:"profit" value:"500"`
	NonProfit   uint          `name:"nonProfit"`
	MegaProfit  uint64        `name:"megaProfit"`
}

func TestParse(t *testing.T) {
	var testingTable = []struct {
		title       string
		args        []string
		environment map[string]string
		excepted    *testParseContainer
	}{
		{
			"String parse success",
			[]string{"-addr", "localhost:8000"},
			map[string]string{"API_KEY": "test-key"},
			&testParseContainer{Addr: "localhost:8000", DatabaseUrl: "mongodb://localhost:27017/db", Timeout: 200 * time.Millisecond, ApiKey: "test-key", Profit: 500},
		},
		{
			"Duration parse success",
			[]string{"-timeout", "30s"},
			map[string]string{},
			&testParseContainer{Addr: ":8000", DatabaseUrl: "mongodb://localhost:27017/db", Timeout: 30 * time.Second, Profit: 500},
		},
		{
			"Int parse success",
			[]string{"-batchSize", "10000"},
			map[string]string{},
			&testParseContainer{Addr: ":8000", DatabaseUrl: "mongodb://localhost:27017/db", Timeout: 200 * time.Millisecond, BatchSize: 10000, Profit: 500},
		},
		{
			"Bool parse success",
			[]string{"-debug"},
			map[string]string{},
			&testParseContainer{Addr: ":8000", DatabaseUrl: "mongodb://localhost:27017/db", Timeout: 200 * time.Millisecond, Debug: true, Profit: 500},
		},
		{
			"Int64 parse success",
			[]string{"-maxCount", "850300"},
			map[string]string{},
			&testParseContainer{Addr: ":8000", DatabaseUrl: "mongodb://localhost:27017/db", Timeout: 200 * time.Millisecond, MaxCount: 850300, Profit: 500},
		},
		{
			"Uint and Uint64 parse success",
			[]string{"-profit", "100", "-megaProfit", "500000"},
			map[string]string{},
			&testParseContainer{Addr: ":8000", DatabaseUrl: "mongodb://localhost:27017/db", Timeout: 200 * time.Millisecond, Profit: 100, MegaProfit: 500000},
		},
	}

	for _, tt := range testingTable {
		t.Run(tt.title, func(t *testing.T) {

			// clear flags
			flag.CommandLine = flag.NewFlagSet(os.Args[0], flag.ExitOnError)

			// clear env vars
			os.Clearenv()

			// set environment vars
			for key, value := range tt.environment {
				assert.NoError(t, os.Setenv(key, value), "must be set environment")
			}

			// insert arguments
			os.Args = []string{"test"}
			os.Args = append(os.Args, tt.args...)

			// testing run
			actual := &testParseContainer{}
			assert.NoError(t, Parse(actual), "must be parse without errors")
			assert.Equal(t, tt.excepted, actual, "must be equal parse result")
		})
	}
}

type testParseErrorsBoolContainer struct {
	Debug bool `name:"debug" value:"foo"`
}

type testParseErrorsIntContainer struct {
	Size int `name:"size" value:"abc"`
}

type testParseErrorsInt64Container struct {
	Size64 int64 `name:"size64" value:"abc"`
}

type testParseErrorsUintContainer struct {
	Profit uint `name:"profit" value:"-100"`
}

type testParseErrorsUint64Container struct {
	Profit64 uint64 `name:"profit" value:"-1000"`
}

type testParseErrorsDurationContainer struct {
	Timeout time.Duration `name:"timeout" value:"bad"`
}

func TestParseErrors(t *testing.T) {
	var testingTable = []struct {
		title  string
		actual interface{}
		error  string
	}{
		{
			"Invalid bool default value throw error",
			&testParseErrorsBoolContainer{},
			`strconv.ParseBool: parsing "foo": invalid syntax`,
		},
		{
			"Invalid int default value throw error",
			&testParseErrorsIntContainer{},
			`strconv.Atoi: parsing "abc": invalid syntax`,
		},
		{
			"Invalid int64 default value throw error",
			&testParseErrorsInt64Container{},
			`strconv.ParseInt: parsing "abc": invalid syntax`,
		},
		{
			"Invalid uint default value throw error",
			&testParseErrorsUintContainer{},
			`strconv.ParseUint: parsing "-100": invalid syntax`,
		},
		{
			"Invalid uint64 default value throw error",
			&testParseErrorsUint64Container{},
			`strconv.ParseUint: parsing "-1000": invalid syntax`,
		},
		{
			"Invalid duration default value throw error",
			&testParseErrorsDurationContainer{},
			`time: invalid duration "bad"`,
		},
	}

	for _, tt := range testingTable {
		t.Run(tt.title, func(t *testing.T) {

			// clear flags
			flag.CommandLine = flag.NewFlagSet(os.Args[0], flag.ContinueOnError)

			// insert arguments
			os.Args = []string{"test"}

			// testing run
			assert.EqualError(t, Parse(tt.actual), tt.error, "must be parse without errors")
		})
	}
}
