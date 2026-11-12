package sendtablescs2

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func fieldOf(varType, encoder string) *field {
	return &field{varType: varType, encoder: encoder, fieldType: newFieldType(varType)}
}

func TestFindDecoder_Fixed8(t *testing.T) {
	t.Parallel()

	wants := map[string][]interface{}{
		"uint8":                 {uint64(0xFF), uint64(0x01), uint64(0x80)},
		"uint8[8]":              {uint64(0xFF), uint64(0x01), uint64(0x80)},
		"MoveType_t":            {uint64(0xFF), uint64(0x01), uint64(0x80)},
		"int8":                  {int32(-1), int32(1), int32(-128)},
		"EntityPlatformTypes_t": {uint32(0xFF), uint32(0x01), uint32(0x80)},
	}

	for varType, want := range wants {
		dec := findDecoder(fieldOf(varType, "fixed8"))
		r := newReader([]byte{0xFF, 0x01, 0x80})

		for _, w := range want {
			assert.Equal(t, w, dec(r), varType)
		}

		assert.Zero(t, r.remBytes(), varType)
	}
}

func TestFindDecoderByBaseType_Fixed8(t *testing.T) {
	t.Parallel()

	dec := findDecoderByBaseType(fieldOf("CNetworkUtlVectorBase< uint8 >", "fixed8"))
	r := newReader([]byte{0xFF, 0x80})

	assert.Equal(t, uint64(0xFF), dec(r))
	assert.Equal(t, uint64(0x80), dec(r))
	assert.Zero(t, r.remBytes())
}

func TestFindDecoder_Fixed16And32(t *testing.T) {
	t.Parallel()

	assert.Equal(t, uint64(0xBEEF), findDecoder(fieldOf("uint16", "fixed16"))(newReader([]byte{0xEF, 0xBE})))
	assert.Equal(t, int32(-2), findDecoder(fieldOf("int16", "fixed16"))(newReader([]byte{0xFE, 0xFF})))
	assert.Equal(t, uint64(0xDEADBEEF), findDecoder(fieldOf("uint32", "fixed32"))(newReader([]byte{0xEF, 0xBE, 0xAD, 0xDE})))
	assert.Equal(t, int32(-1), findDecoder(fieldOf("int32", "fixed32"))(newReader([]byte{0xFF, 0xFF, 0xFF, 0xFF})))
}

func TestFindDecoder_WithoutEncoderStaysVarint(t *testing.T) {
	t.Parallel()

	assert.Equal(t, uint64(300), findDecoder(fieldOf("uint8", ""))(newReader([]byte{0xAC, 0x02})))
	assert.Equal(t, int32(-1), findDecoder(fieldOf("int8", ""))(newReader([]byte{0x01})))
	assert.Equal(t, uint32(300), findDecoder(fieldOf("EntityPlatformTypes_t", ""))(newReader([]byte{0xAC, 0x02})))
}
