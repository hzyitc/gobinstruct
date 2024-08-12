package gobinstruct

import (
	"bytes"
	"io"
)

func (s *Decoder) _ReadStream(r io.Reader, data any, info *Decoder_ExtraInfo) error {
	for _, Decoder := range *s {
		err := Decoder(r, data, info)
		if err == ErrMisType {
			continue
		}
		return err
	}

	return io.ErrUnexpectedEOF
}

func (s *Decoder) ReadStream(r io.Reader, data any) error {
	return s._ReadStream(r, data, &Decoder_ExtraInfo{
		Decoder: s,
	})
}

func (s *Decoder) Read(buf []byte, data any) (int, error) {
	r := bytes.NewReader(buf)
	err := s.ReadStream(r, data)
	if err != nil {
		return 0, err
	}
	return int(r.Size() - int64(r.Len())), nil
}
