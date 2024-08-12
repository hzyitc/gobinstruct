package gobinstruct

import (
	"bytes"
	"io"
)

func (s *Encoder) _WriteStream(w io.Writer, data any, info *Encoder_ExtraInfo) error {
	for _, Encoder := range *s {
		err := Encoder(w, data, info)
		if err == ErrMisType {
			continue
		}
		return err
	}

	return io.ErrUnexpectedEOF
}

func (s *Encoder) WriteStream(w io.Writer, data any) error {
	return s._WriteStream(w, data, &Encoder_ExtraInfo{
		Encoder: s,
	})
}

func (s *Encoder) Write(buf []byte, data any) (int, error) {
	w := new(bytes.Buffer)
	err := s.WriteStream(w, data)
	if err != nil {
		return 0, err
	}

	err = nil
	if len(buf) < w.Len() {
		err = io.ErrShortBuffer
	}
	return copy(buf, w.Bytes()), err
}

func (s *Encoder) WriteTo(w io.Writer, data any) (int, error) {
	buf := new(bytes.Buffer)
	err := s.WriteStream(buf, data)
	if err != nil {
		return 0, err
	}

	n, err := io.Copy(w, buf)
	return int(n), err
}

func (s *Encoder) WriteToOnce(w io.Writer, data any) (int, error) {
	buf := new(bytes.Buffer)
	err := s.WriteStream(buf, data)
	if err != nil {
		return 0, err
	}

	n, err := w.Write(buf.Bytes())
	if err != nil {
		return n, err
	}
	if n < buf.Len() {
		return n, io.ErrShortWrite
	}
	return n, nil
}
