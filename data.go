package otr3

import (
	"crypto/sha256"
	"hash"
	"math/big"
)

func appendWord(l []byte, r uint32) []byte {
	return append(l, byte(r>>24), byte(r>>16), byte(r>>8), byte(r))
}

func appendBytes(l, r []byte) []byte {
	return append(appendWord(l, uint32(len(r))), r...)
}

func appendMPI(l []byte, r *big.Int) []byte {
	mpiBytes := r.Bytes()
	return append(appendWord(l, uint32(len(mpiBytes))), mpiBytes...)
}

func appendMPIs(l []byte, r ...*big.Int) []byte {
	for _, mpi := range r {
		l = appendMPI(l, mpi)
	}
	return l
}

func hashMPIs(h hash.Hash, magic byte, mpis ...*big.Int) []byte {
	if h != nil {
		h.Reset()
	} else {
		h = sha256.New()
	}

	h.Write([]byte{magic})
	for _, mpi := range mpis {
		h.Write(appendMPI(nil, mpi))
	}
	return h.Sum(nil)
}