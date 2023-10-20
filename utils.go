package ilog10

func CountUint32(n uint32) uint {
	if n == 0 {
		return 1
	}
	return 1 + FastUint32(n)
}

func CountUint64(n uint64) uint {
	if n == 0 {
		return 1
	}
	return 1 + FastUint64(n)
}
