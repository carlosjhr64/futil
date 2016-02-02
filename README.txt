package futil // import "github.com/carlosjhr64/futil"

const VERSION string = "0.2.0"

func FileOK(fn string) bool
func Read(fn string) string
func Wcl(fn string) int
func ReadTable(fn string, cols int) [][]string
