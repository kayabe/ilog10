package ilog10

var factors = [...]uint64{
	1e19,
	1e18,
	1e17,
	1e16,
	1e15,
	1e14,
	1e13,
	1e12,
	1e11,
	1e10,
	1e9,
	1e8,
	1e7,
	1e6,
	1e5,
	1e4,
	1e3,
	1e2,
	1e1,
}

const factorsCount = len(factors)

const maxUint64 = ^uint64(0)
const maxUint32 = ^uint32(0)

var lookup32 [32]uint32
var lookup64 [64]uint64

func init() {
	buildLookup32()
	buildLookup64()
}

func buildLookup32() {
	/*
		lookup32 table

		           4294967295 4294967295
		1000000000 4294967295 4294967295
		100000000  4294967295 4294967295
		10000000   4294967295 4294967295 4294967295
		1000000    4294967295 4294967295
		100000     4294967295 4294967295
		10000      4294967295 4294967295 4294967295
		1000       4294967295 4294967295
		100        4294967295 4294967295
		10         4294967295 4294967295 4294967295
	*/
	for i, j := 10, 0; i < factorsCount; i++ {
		if i == 10 {
			lookup32[j] = maxUint32
			lookup32[j+1] = maxUint32
			j += 2
		}

		lookup32[j] = uint32(factors[i])
		lookup32[j+1] = maxUint32
		lookup32[j+2] = maxUint32

		j += 3

		if i%3 == 0 {
			lookup32[j] = maxUint32
			j++
		}
	}
}

func buildLookup64() {
	/*
		lookup64 table

		10000000000000000000 18446744073709551615 18446744073709551615 18446744073709551615
		1000000000000000000  18446744073709551615 18446744073709551615
		100000000000000000   18446744073709551615 18446744073709551615
		10000000000000000    18446744073709551615 18446744073709551615 18446744073709551615
		1000000000000000     18446744073709551615 18446744073709551615
		100000000000000      18446744073709551615 18446744073709551615
		10000000000000       18446744073709551615 18446744073709551615 18446744073709551615
		1000000000000        18446744073709551615 18446744073709551615
		100000000000         18446744073709551615 18446744073709551615
		10000000000          18446744073709551615 18446744073709551615 18446744073709551615
		1000000000           18446744073709551615 18446744073709551615
		100000000            18446744073709551615 18446744073709551615
		10000000             18446744073709551615 18446744073709551615 18446744073709551615
		1000000              18446744073709551615 18446744073709551615
		100000               18446744073709551615 18446744073709551615
		10000                18446744073709551615 18446744073709551615 18446744073709551615
		1000                 18446744073709551615 18446744073709551615
		100                  18446744073709551615 18446744073709551615
		10                   18446744073709551615 18446744073709551615 18446744073709551615
	*/

	for i, j := 0, 0; i < factorsCount; i++ {
		lookup64[j] = factors[i]
		lookup64[j+1] = maxUint64
		lookup64[j+2] = maxUint64

		j += 3

		if i%3 == 0 {
			lookup64[j] = maxUint64
			j++
		}
	}
}
