package futil

import "testing"

import "fmt"
var puts = fmt.Println

func TestUtil(test *testing.T){
  bad := test.Error
  fn := "./futil_test.go"

  // Obviously, this file exist, is regular, and is readable!
  if !FileOK(fn) { bad("OK") }

  // This file as 21 lines.
  lines := Wcl(fn)
  if lines != 21 { bad("Wcl") }

  s := Read(fn)
  if s[:13] != "package futil" { bad("Read") }
}
