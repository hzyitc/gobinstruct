package gobinstruct

import (
	"io"
	"reflect"
)

func Struct_Encoder(w io.Writer, data any, info *Encoder_ExtraInfo) error {
	v := reflect.Indirect(reflect.ValueOf(data))

	if v.Kind() != reflect.Struct {
		return ErrMisType
	}
	for i := 0; i < v.NumField(); i++ {
		field := v.Field(i)
		err := info.Encoder._WriteStream(w, field.Interface(), &Encoder_ExtraInfo{
			Encoder:   info.Encoder,
			StructTag: v.Type().Field(i).Tag,
		})
		if err != nil {
			return err
		}
	}
	return nil
}

func Struct_Decoder(r io.Reader, data any, info *Decoder_ExtraInfo) error {
	v := reflect.Indirect(reflect.ValueOf(data))

	if v.Kind() != reflect.Struct {
		return ErrMisType
	}

	for i := 0; i < v.NumField(); i++ {
		err := info.Decoder._ReadStream(r, v.Field(i).Addr().Interface(), &Decoder_ExtraInfo{
			Decoder:   info.Decoder,
			StructTag: v.Type().Field(i).Tag,
		})
		if err != nil {
			return err
		}
	}
	return nil
}
