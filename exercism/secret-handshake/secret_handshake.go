package secret

func Handshake(in uint) []string {
	var res []string
	if in&(1<<0) > 0 {
		res = append(res, "wink")
	}
	if in&(1<<1) > 0 {
		res = append(res, "double blink")
	}
	if in&(1<<2) > 0 {
		res = append(res, "close your eyes")
	}
	if in&(1<<3) > 0 {
		res = append(res, "jump")
	}
	if in&(1<<4) > 0 {
		for i := len(res)/2 - 1; i >= 0; i-- {
			opp := len(res) - 1 - i
			res[i], res[opp] = res[opp], res[i]
		}
	}
	return res
}
