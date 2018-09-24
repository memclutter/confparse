package confparse

import (
	"os"
	"reflect"
	"strconv"
	"testing"
)

type testParseContainer struct {
	Addr        string `name:"addr" value:":8000" usage:"Listen and serve address"`
	DatabaseUrl string `name:"databaseUrl" value:"mongodb://localhost:27017/db" usage:"Database connection url"`
	OptionalVal string `name:"optional"`
}

var testParseTable = []struct {
	container         *testParseContainer
	args              []string
	exceptedContainer *testParseContainer
}{
	{
		&testParseContainer{},
		[]string{"-addr", "localhost:8000"},
		&testParseContainer{Addr: "localhost:8000", DatabaseUrl: "mongodb://localhost:27017/db"},
	},
}

func TestParse(t *testing.T) {
	for i, tt := range testParseTable {
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			os.Args = []string{"test"}
			os.Args = append(os.Args, tt.args...)

			if err := Parse(tt.container); err != nil {
				t.Fatalf("Error parse valid config container: %s", err)
			}

			if !reflect.DeepEqual(tt.container, tt.exceptedContainer) {
				t.Errorf("Unequal container\nActual: %#v\nExcepted: %#v", tt.container, tt.exceptedContainer)
			}
		})
	}
}
