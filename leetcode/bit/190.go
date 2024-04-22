package bit

func ReverseBits(num uint32) uint32 {
	var res uint32 = 0
	for i := 0; i < 32; i++ {
		if num&1 == 1 {
			res |= 1 << (31 - i)
		}
		num >>= 1
	}
	return res
}
