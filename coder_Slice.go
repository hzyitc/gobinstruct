package gobinstruct

import (
	"errors"
	"io"
	"reflect"
	"strconv"
)

var ErrIncorrentLength = errors.New("incorrect length")

func Slice_Encoder(w io.Writer, data any, info *Encoder_ExtraInfo) error {
	v := reflect.Indirect(reflect.ValueOf(data))

	if v.Kind() != reflect.Slice {
		return ErrMisType
	}

	lenStr := info.StructTag.Get("gbs.len")
	if lenStr != "" {
		l, err := strconv.Atoi(lenStr)
		if err != nil {
			return err
		}
		if v.Len() != l {
			return ErrIncorrentLength
		}
	}

	for i := 0; i < v.Len(); i++ {
		err := info.Encoder.WriteStream(w, v.Index(i).Interface())
		if err != nil {
			return err
		}
	}
	return nil
}

func Slice_Decoder(r io.Reader, data any, info *Decoder_ExtraInfo) error {
	v := reflect.Indirect(reflect.ValueOf(data))

	if v.Kind() != reflect.Slice {
		return ErrMisType
	}

	l, err := strconv.Atoi(info.StructTag.Get("gbs.len"))
	if err != nil {
		return err
	}

	v.Set(reflect.MakeSlice(v.Type(), l, l))
	for i := 0; i < l; i++ {
		err := info.Decoder.ReadStream(r, v.Index(i).Addr().Interface())
		if err != nil {
			return err
		}
	}
	return nil
}
