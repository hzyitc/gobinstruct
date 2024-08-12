package gobinstruct

import (
	"bytes"
	"encoding/binary"
	"io"
)

func NulTermString_Encoder(w io.Writer, data any, info *Encoder_ExtraInfo) error {
	d, ok := data.(string)
	if !ok {
		return ErrMisType
	}

	return binary.Write(w, binary.LittleEndian, []byte(d+"\000"))
}

func NulTermString_Decoder(r io.Reader, data any, info *Decoder_ExtraInfo) error {
	d, ok := data.(*string)
	if !ok {
		return ErrMisType
	}

	buf := new(bytes.Buffer)
	err := copyUntil(buf, r, 0)
	if err != nil {
		return err
	}
	*d = buf.String()
	return nil
}

func copyUntil(w io.Writer, r io.Reader, delim byte) error {
	buf := make([]byte, 1)
	for {
		_, err := r.Read(buf)
		if err != nil {
			return err
		}

		if buf[0] == delim {
			return nil
		}

		_, err = w.Write(buf)
		if err != nil {
			return err
		}
	}
}
