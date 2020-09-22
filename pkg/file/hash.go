package file

import (
	"crypto/sha256"
	"os"
)

// fp: file path
func FileHash(fp string) ([]byte, error) {
	f, err := os.Open(fp)
	defer f.Close()
	if err != nil {
		return nil, err
	}
	hash := sha256.New()
	split := make([]byte, MaxReadSize)
	for {
		switch n, err := f.Read(split); {
		case n > 0:
			hash.Write(split)
		// EOF
		case n == 0:
			return hash.Sum(nil), nil
		case n < 0:
			return nil, err
		}
	}
}
