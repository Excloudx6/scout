package scan

import (
	"bytes"
	"time"

	"github.com/liamg/scout/internal/app/scout/data"
	"github.com/liamg/scout/pkg/wordlist"
)

type VHOSTOptions struct {
	BaseDomain          string           // target url
	Timeout             time.Duration    // http request timeout
	Parallelism         int              // parallel routines
	ResultChan          chan VHOSTResult // chan to return results on - otherwise will be returned in slice
	BusyChan            chan string      // chan to use to update current job
	Wordlist            wordlist.Wordlist
	SkipSSLVerification bool
	UseSSL              bool
	IP                  string
	Port                int
	ContentHashing      bool
}

type VHOSTResult struct {
	VHOST      string
	StatusCode int
}

var DefaultVHOSTOptions = VHOSTOptions{
	Timeout:     time.Second * 5,
	Parallelism: 10,
}

func (opt *VHOSTOptions) Inherit() {
	if opt.Timeout == 0 {
		opt.Timeout = DefaultVHOSTOptions.Timeout
	}
	if opt.Parallelism == 0 {
		opt.Parallelism = DefaultVHOSTOptions.Parallelism
	}
	if opt.Wordlist == nil {
		wordlistBytes, err := data.Asset("assets/vhost.txt")
		if err != nil {
			wordlistBytes = []byte{}
		}
		opt.Wordlist = wordlist.FromReader(bytes.NewReader(wordlistBytes))
	}
}
