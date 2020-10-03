package spammasker

import (
	"fmt"
	"os"
)

// SpamMasker testing
func SpamMasker() {
	var (
		text = os.Args[1:][0]
		size = len(text)
		// buf  = make([]byte, 0, size)
		buf = []byte(text)
		in  = false
	)

	const (
		link  = "https://"
		lLink = len(link)
		mask  = '*'
	)

	for i := 0; i < size; i++ {

		if len(text[i:]) >= lLink && text[i:i+lLink] == link {
			i += lLink
			in = true
		}

		switch text[i] {
		case '\n', '\t', ' ':
			in = false
		}

		if in {
			buf[i] = mask
		}
	}

	fmt.Println(string(buf))

}
