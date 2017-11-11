package comm

import (
	"flag"
	"fmt"

	"os"

	"path/filepath"

	"github.com/xtfly/gokits/gcrypto"
	"github.com/xtfly/gokits/gfile"
	"github.com/xtfly/gokits/grand"
	"github.com/xtfly/log4g"
)

var (
	Version   = "0.0.1"
	BuildDate = ""
	CommitID  = ""

	v  = flag.Bool("version", false, "show the version")
	gf = flag.Bool("gfactor", false, "generate a encryption factor")
	f  = flag.String("factor", "", "encryption factor")
	p  = flag.String("password", "", "the password to be encrypted")

	gAppcfg = ""
	logfile = ""
)

// ParseCmd parse command flags
func ParseCmd() {
	flag.Parse()
	if *v {
		fmt.Printf(`Version: %s
Commit: %s
Build Date: %s
`, Version, CommitID, BuildDate)
		os.Exit(0)
	}

	if *gf {
		fmt.Printf("Factor: %s\n", grand.NewRand(16))
		os.Exit(0)
	}

	if *f != "" && *p != "" {
		crypto, _ := gcrypto.NewCrypto(*f)
		stxt, _ := crypto.EncryptStr(*p)
		fmt.Printf("Encrypted Password: %s\n", stxt)
		os.Exit(0)
	}

	pwd, _ := os.Getwd()
	gAppcfg = filepath.Join(pwd, "conf/app.yaml")
	if !gfile.FileExist(gAppcfg) {
		fmt.Printf("not find app.yaml in config dir")
		os.Exit(1)
	}

	logfile = filepath.Join(pwd, "conf/log4g.yaml")
	if gfile.FileExist(logfile) {
		log.GetManager().LoadConfigFile(logfile)
	} else {
		cfgdata := `
formats:
  - name: f1
    type: text
    layout: "%{time}|%{level}|%{module}|%{pid:6d}>> %{msg} (%{longfile}:%{line}) \n"
outputs:
  - name: c1
    type: console
    format: f1
    threshold: debug
loggers:
  - name: root
    level: debug
    outputs: ["c1"]
`
		log.GetManager().LoadConfig([]byte(cfgdata), "yaml")
	}
}
