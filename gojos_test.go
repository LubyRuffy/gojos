package gojos

import (
	"os"
	"testing"
	"time"

	"fmt"

	"path/filepath"

	"log"

	"github.com/m-pavel/gojos/lib/javaos"
)

const (
	fldr = "./target/tdata/"
)

func doTest(fname string, t *testing.T) (res *javaos.JavaModel) {
	log.Printf("Testing %s\n", fname)
	file, err := os.OpenFile(fname, os.O_RDONLY, 0644)
	if err != nil {
		t.Fatal(err)
	}
	res, err = Deserialize(file)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println(res)
	return res
}

func TestAllParse(t *testing.T) {
	filepath.Walk(fldr, func(path string, f os.FileInfo, _ error) error {
		if !f.IsDir() {
			doTest(path, t)
		}
		return nil
	})
}

type Tkn struct {
	Gen        time.Time `java:"generatedDateTime"`
	ValidUntil time.Time
	User       string `java:"userId"`
	Metadata   map[string]string
}

func Test55(t *testing.T) {
	file, err := os.OpenFile("./target/tdata/test55.bin", os.O_RDONLY, 0644)
	if err != nil {
		t.Fatal(err)
	}
	t1 := Tkn{}
	err = UnmarshallStream(file, &t1)
	if err != nil {
		t.Fatal(err)
	}
	log.Printf("Go model %s", t1)
}