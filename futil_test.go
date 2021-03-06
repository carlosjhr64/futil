package futil

import "testing"

func TestUtil(test *testing.T){
  bad := test.Error
  fn := "./futil_test.go"

  // Obviously, this file exist, is regular, and is readable!
  if !FileOK(fn) { bad("OK") }

  // This file has 18 lines.
  lines := Wcl(fn)
  if lines != 20 { bad("Wcl") }

  s := Read(fn)
  if s[:13] != "package futil" { bad("Read") }

  // TODO: ReadTable
}
