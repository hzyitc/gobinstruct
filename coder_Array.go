package gobinstruct

import (
	"io"
	"reflect"
)

func Array_Encoder(w io.Writer, data any, info *Encoder_ExtraInfo) error {
	v := reflect.Indirect(reflect.ValueOf(data))

	if v.Kind() != reflect.Array {
		return ErrMisType
	}

	for i := 0; i < v.Len(); i++ {
		err := info.Encoder.WriteStream(w, v.Index(i).Interface())
		if err != nil {
			return err
		}
	}
	return nil
}

func Array_Decoder(r io.Reader, data any, info *Decoder_ExtraInfo) error {
	v := reflect.Indirect(reflect.ValueOf(data))

	if v.Kind() != reflect.Array {
		return ErrMisType
	}

	for i := 0; i < v.Len(); i++ {
		err := info.Decoder.ReadStream(r, v.Index(i).Addr().Interface())
		if err != nil {
			return err
		}
	}
	return nil
}
