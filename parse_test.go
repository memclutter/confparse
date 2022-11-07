package confparse

import (
	"flag"
	"os"
	"reflect"
	"strconv"
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
	Profit      uint          `name:"profit"`
	MegaProfit  uint64        `name:"megaProfit"`
}

var testParseTable = []struct {
	container         *testParseContainer
	args              []string
	environment       map[string]string
	exceptedContainer *testParseContainer
}{
	{
		&testParseContainer{},
		[]string{"-addr", "localhost:8000"},
		map[string]string{"API_KEY": "test-key"},
		&testParseContainer{Addr: "localhost:8000", DatabaseUrl: "mongodb://localhost:27017/db", Timeout: 200 * time.Millisecond, ApiKey: "test-key"},
	},
	{
		&testParseContainer{},
		[]string{"-timeout", "30s"},
		map[string]string{},
		&testParseContainer{Addr: ":8000", DatabaseUrl: "mongodb://localhost:27017/db", Timeout: 30 * time.Second},
	},
	{
		&testParseContainer{},
		[]string{"-batchSize", "10000"},
		map[string]string{},
		&testParseContainer{Addr: ":8000", DatabaseUrl: "mongodb://localhost:27017/db", Timeout: 200 * time.Millisecond, BatchSize: 10000},
	},
	{
		&testParseContainer{},
		[]string{"-debug"},
		map[string]string{},
		&testParseContainer{Addr: ":8000", DatabaseUrl: "mongodb://localhost:27017/db", Timeout: 200 * time.Millisecond, Debug: true},
	},
	{
		&testParseContainer{},
		[]string{"-maxCount", "850300"},
		map[string]string{},
		&testParseContainer{Addr: ":8000", DatabaseUrl: "mongodb://localhost:27017/db", Timeout: 200 * time.Millisecond, MaxCount: 850300},
	},
	{
		&testParseContainer{},
		[]string{"-profit", "100", "-megaProfit", "500000"},
		map[string]string{},
		&testParseContainer{Addr: ":8000", DatabaseUrl: "mongodb://localhost:27017/db", Timeout: 200 * time.Millisecond, Profit: 100, MegaProfit: 500000},
	},
}

func TestParse(t *testing.T) {
	for i, tt := range testParseTable {
		t.Run(strconv.Itoa(i), func(t *testing.T) {

			// clear flags
			flag.CommandLine = flag.NewFlagSet(os.Args[0], flag.ExitOnError)

			// clear env vars
			os.Clearenv()

			// set environment vars
			for key, value := range tt.environment {
				os.Setenv(key, value)
			}

			// insert arguments
			os.Args = []string{"test"}
			os.Args = append(os.Args, tt.args...)

			// check no errors
			if err := Parse(tt.container); err != nil {
				t.Fatalf("Error parse valid config container: %s", err)
			}

			// compare containers
			if !reflect.DeepEqual(tt.container, tt.exceptedContainer) {
				t.Errorf("Unequal container\nActual: %#v\nExcepted: %#v", tt.container, tt.exceptedContainer)
			}
		})
	}
}
