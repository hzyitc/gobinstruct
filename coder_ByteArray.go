package gobinstruct

import (
	"io"
)

func ByteArray_Encoder(w io.Writer, data any, info *Encoder_ExtraInfo) error {
	d, ok := data.([]byte)
	if !ok {
		return ErrMisType
	}

	_, err := w.Write(d)
	return err
}

func ByteArray_Decoder(r io.Reader, data any, info *Decoder_ExtraInfo) error {
	return ErrMisType
}
