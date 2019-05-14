package atbash

import "bytes"

func Atbash(plain string) string {
	cipher := make([]byte, 0, len(plain)+len(plain)/5)
	for _, r := range []byte(plain) {
		switch {
		case r >= 'a' && r <= 'z':
			cipher = append(cipher, 'z'-r+'a')
		case r >= 'A' && r <= 'Z':
			cipher = append(cipher, 'z'-r+'A')
		case r >= '0' && r <= '9':
			cipher = append(cipher, r)
		default:
			continue
		}
		if len(cipher)%6 == 5 {
			cipher = append(cipher, ' ')
		}
	}
	return string(bytes.TrimSuffix(cipher, []byte{' '}))
}
