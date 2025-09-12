package zenrpc_test

import (
	"bytes"
	"encoding/json"
	"log"
	"os"
	"testing"

	"github.com/vmkteam/zenrpc/v2"
	"github.com/vmkteam/zenrpc/v2/testdata"
)

var (
	testRPC     = zenrpc.NewServer(zenrpc.Options{BatchMaxLen: 5, AllowCORS: true})
	logRequests = false
)

func TestMain(m *testing.M) {
	testRPC.Register("arith", &testdata.ArithService{})
	testRPC.Register("", &testdata.ArithService{})

	if logRequests {
		testRPC.Use(zenrpc.Logger(log.New(os.Stderr, "", log.LstdFlags)))
	}

	result := m.Run()
	os.Exit(result)
}

func TestServer_SMD(t *testing.T) {
	r := testRPC.SMD()
	if b, err := json.Marshal(r); err != nil {
		t.Fatal(err)
	} else if !bytes.Contains(b, []byte("default")) {
		t.Error(string(b))
	}
}

func TestServer_SmdGenerate(t *testing.T) {
	rpc := zenrpc.NewServer(zenrpc.Options{})
	rpc.Register("phonebook", testdata.PhoneBook{DB: testdata.People})
	rpc.Register("arith", testdata.ArithService{})
	rpc.Register("printer", testdata.PrintService{})
	rpc.Register("", testdata.ArithService{})

	b, _ := json.MarshalIndent(rpc.SMD(), "", "  ")

	testData, err := os.ReadFile("./testdata/testdata/arithsrv-smd.json")
	if err != nil {
		t.Fatalf("open test data file")
	}

	if !bytes.Equal(b, testData) {
		t.Fatalf("bad zenrpc output")
	}
}
