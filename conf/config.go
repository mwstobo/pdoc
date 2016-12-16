package conf

import (
	"flag"
)

var (
	BaseDirectory = flag.String("dir", "~/docs", "directory containing your documentation")
)

func LoadConfig() {
	flag.Parse()
}
