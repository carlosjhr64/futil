package futil

import (
  "os"
  "io"
  "bytes"
  "strings"
  "io/ioutil"
)

const VERSION string = "0.2.0"

// http://stackoverflow.com/questions/24562942/golang-how-do-i-determine-the-number-of-lines-in-a-file-efficiently
func Wcl(fn string) int {
  r, e := os.Open(fn)
  defer r.Close()
  if e != nil { panic(e) }

  buf := make([]byte, 8196)
  count := 0
  lineSep := []byte{'\n'}

  for {
    c, err := r.Read(buf)
    if err != nil && err != io.EOF{ panic(err) }
    count += bytes.Count(buf[:c], lineSep)
    if err == io.EOF { break }
  }

  return count
}

func Read(fn string) string {
  b, e := ioutil.ReadFile(fn)
  if e != nil { panic(e) }
  return string(b)
}

func FileOK(fn string) bool {
  s, e := os.Stat(fn)
  if e != nil { return false }
  m := s.Mode()
  if !m.IsRegular() { return false }
  p := m.Perm()
  return (p &^ 0377) == 0400
}

func ReadTable(fn string, cols int) [][]string {
  lines := strings.Split(Read(fn), "\n")
  table := make([][]string, len(lines))
  i := 0
  for _, line := range(lines) {
    if line == "" { continue }
    fields := strings.Fields(line)
    if cols>0 && len(fields)!=cols { panic("Unexpected fields length.") }
    table[i] = fields
    i++
  }
  return table[:i]
}
