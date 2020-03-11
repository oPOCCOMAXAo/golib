package binary

import (
	"encoding/binary"
	"math"
)

func WriteUint32BE(buffer []byte, position int, data uint32) {
	binary.BigEndian.PutUint32(buffer[position:position+4], data)
}

func ReadUint32BE(buffer []byte, position int) uint32 {
	return binary.BigEndian.Uint32(buffer[position : position+4])
}

func WriteUint8BE(buffer []byte, position int, data uint8) {
	buffer[position] = data
}

func ReadUint8BE(buffer []byte, position int) uint8 {
	return buffer[position]
}

func WriteBooleanBE(buffer []byte, position int, data bool) {
	if data {
		buffer[position] = 1
	} else {
		buffer[position] = 0
	}
}

func WriteFloat64BE(buffer []byte, position int, data float64) {
	binary.BigEndian.PutUint64(buffer[position:position+8], math.Float64bits(data))
}

func ReadFloat64BE(buffer []byte, position int) float64 {
	return math.Float64frombits(binary.BigEndian.Uint64(buffer[position : position+8]))
}
