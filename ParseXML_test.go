package main

import(
  "testing"
  "io/ioutil"
  "fmt"
  "bytes"
)

func TestParseXML(t *testing.T) {
  xml, err := ioutil.ReadFile("./test_data/test.xml")
  if err != nil {
    t.Errorf(err.Error())
  }
  actual := []byte(fmt.Sprint(ParseXML(xml)))
  expected, err := ioutil.ReadFile("./test_data/test.txt")
  if err != nil {
    t.Errorf(err.Error())
  }
  if !bytes.Equal(actual,expected) {
    t.Errorf("got %s\nwant %s", actual, expected)
  }
}
