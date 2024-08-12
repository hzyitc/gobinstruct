package gobinstruct

import (
	"encoding/binary"
	"io"
)

func Base_BE_Encoder(w io.Writer, data any, info *Encoder_ExtraInfo) error {
	switch data.(type) {
	case int8, uint8,
		int32, uint32:
		break
	default:
		return ErrMisType
	}

	return binary.Write(w, binary.BigEndian, data)
}

func Base_BE_Decoder(r io.Reader, data any, info *Decoder_ExtraInfo) error {
	switch data.(type) {
	case *int8, *uint8,
		*int32, *uint32:
		break
	default:
		return ErrMisType
	}

	return binary.Read(r, binary.BigEndian, data)
}
