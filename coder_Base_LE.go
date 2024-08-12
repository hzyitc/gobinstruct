package gobinstruct

import (
	"encoding/binary"
	"io"
)

func Base_LE_Encoder(w io.Writer, data any, info *Encoder_ExtraInfo) error {
	switch data.(type) {
	case int8, uint8,
		int32, uint32:
		break
	default:
		return ErrMisType
	}

	return binary.Write(w, binary.LittleEndian, data)
}

func Base_LE_Decoder(r io.Reader, data any, info *Decoder_ExtraInfo) error {
	switch data.(type) {
	case *int8, *uint8,
		*int32, *uint32:
		break
	default:
		return ErrMisType
	}

	return binary.Read(r, binary.LittleEndian, data)
}
