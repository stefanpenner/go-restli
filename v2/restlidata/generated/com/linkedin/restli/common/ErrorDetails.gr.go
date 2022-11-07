/*
Code generated by "github.com/PapaCharlie/go-restli/v2"; DO NOT EDIT.

Source file: /Users/pchesnai/code/personal/go-restli/v2/go-restli-spec-parser.jar
*/

package common

import (
	"github.com/PapaCharlie/go-restli/v2/fnv1a"
	"github.com/PapaCharlie/go-restli/v2/restli/patch"
	"github.com/PapaCharlie/go-restli/v2/restlicodec"
)

type ErrorDetails struct{}

func (e *ErrorDetails) Equals(other *ErrorDetails) bool {
	if e == other {
		return true
	}
	if e == nil || other == nil {
		return false
	}

	return true
}

func (e *ErrorDetails) ComputeHash() fnv1a.Hash {
	if e == nil {
		return fnv1a.ZeroHash()
	}
	hash := fnv1a.NewHash()

	return hash
}

func (e *ErrorDetails) MarshalFields(keyWriter func(string) restlicodec.Writer) (err error) {
	return nil
}

func (e *ErrorDetails) MarshalRestLi(writer restlicodec.Writer) (err error) {
	return writer.WriteMap(e.MarshalFields)
}

func (e *ErrorDetails) MarshalJSON() (data []byte, err error) {
	writer := restlicodec.NewCompactJsonWriter()
	err = e.MarshalRestLi(writer)
	if err != nil {
		return nil, err
	}
	return []byte(writer.Finalize()), nil
}

var ErrorDetailsRequiredFields = restlicodec.NewRequiredFields()

func (e *ErrorDetails) UnmarshalField(reader restlicodec.Reader, field string) (found bool, err error) {
	return false, nil
}

func (e *ErrorDetails) UnmarshalRestLi(reader restlicodec.Reader) (err error) {
	err = reader.ReadRecord(ErrorDetailsRequiredFields, func(reader restlicodec.Reader, field string) (err error) {
		found, err := e.UnmarshalField(reader, field)
		if err != nil {
			return err
		}
		if !found {
			err = reader.Skip()
		}
		return err
	})
	if err != nil {
		return err
	}

	return nil
}

func (e *ErrorDetails) UnmarshalJSON(data []byte) error {
	return restlicodec.UnmarshalJSON(data, e)
}

func (e *ErrorDetails) NewInstance() *ErrorDetails {
	return new(ErrorDetails)
}

/*
================================================================================
PARTIAL UPDATE STRUCTS
================================================================================
*/

type ErrorDetails_PartialUpdate_Set_Fields struct{}

func (e *ErrorDetails_PartialUpdate_Set_Fields) MarshalFields(keyWriter func(string) restlicodec.Writer) (err error) {
	return nil
}

func (e *ErrorDetails_PartialUpdate_Set_Fields) MarshalRestLi(writer restlicodec.Writer) (err error) {
	return writer.WriteMap(e.MarshalFields)
}

func (e *ErrorDetails_PartialUpdate_Set_Fields) MarshalJSON() (data []byte, err error) {
	writer := restlicodec.NewCompactJsonWriter()
	err = e.MarshalRestLi(writer)
	if err != nil {
		return nil, err
	}
	return []byte(writer.Finalize()), nil
}

var ErrorDetails_PartialUpdate_Set_FieldsRequiredFields = restlicodec.NewRequiredFields()

func (e *ErrorDetails_PartialUpdate_Set_Fields) UnmarshalField(reader restlicodec.Reader, field string) (found bool, err error) {
	return false, nil
}

func (e *ErrorDetails_PartialUpdate_Set_Fields) UnmarshalRestLi(reader restlicodec.Reader) (err error) {
	err = reader.ReadRecord(ErrorDetails_PartialUpdate_Set_FieldsRequiredFields, func(reader restlicodec.Reader, field string) (err error) {
		found, err := e.UnmarshalField(reader, field)
		if err != nil {
			return err
		}
		if !found {
			err = reader.Skip()
		}
		return err
	})
	if err != nil {
		return err
	}

	return nil
}

func (e *ErrorDetails_PartialUpdate_Set_Fields) UnmarshalJSON(data []byte) error {
	return restlicodec.UnmarshalJSON(data, e)
}

func (e *ErrorDetails_PartialUpdate_Set_Fields) NewInstance() *ErrorDetails_PartialUpdate_Set_Fields {
	return new(ErrorDetails_PartialUpdate_Set_Fields)
}

// ErrorDetails_PartialUpdate is used to represent a partial update on ErrorDetails. Toggling the value of a field
// in Delete represents selecting it for deletion in a partial update, while
// setting the value of a field in Update represents setting that field in the
// current struct. Other fields in this struct represent record fields that can
// themselves be partially updated.
type ErrorDetails_PartialUpdate struct{}

func (e *ErrorDetails_PartialUpdate) MarshalRestLiPatch(writer restlicodec.Writer) (err error) {
	return writer.WriteMap(func(keyWriter func(string) restlicodec.Writer) (err error) {
		return nil
	})
}

func (e *ErrorDetails_PartialUpdate) MarshalRestLi(writer restlicodec.Writer) (err error) {
	return writer.WriteMap(func(keyWriter func(string) restlicodec.Writer) (err error) {
		return e.MarshalRestLiPatch(keyWriter(patch.PatchField).SetScope())
	})
}

func (e *ErrorDetails_PartialUpdate) MarshalJSON() (data []byte, err error) {
	writer := restlicodec.NewCompactJsonWriter()
	err = e.MarshalRestLi(writer)
	if err != nil {
		return nil, err
	}
	return []byte(writer.Finalize()), nil
}

func (e *ErrorDetails_PartialUpdate) UnmarshalRestLiPatch(reader restlicodec.Reader) (err error) {
	return reader.ReadMap(func(reader restlicodec.Reader, key string) (err error) {
		return reader.Skip()
	})
}

func (e *ErrorDetails_PartialUpdate) UnmarshalRestLi(reader restlicodec.Reader) (err error) {
	return reader.ReadRecord(patch.RequiredPatchRecordFields, func(reader restlicodec.Reader, field string) (err error) {
		if field == patch.PatchField {
			return e.UnmarshalRestLiPatch(reader)
		} else {
			return reader.Skip()
		}
	})
}

func (e *ErrorDetails_PartialUpdate) UnmarshalJSON(data []byte) error {
	return restlicodec.UnmarshalJSON(data, e)
}

func (e *ErrorDetails_PartialUpdate) NewInstance() *ErrorDetails_PartialUpdate {
	return new(ErrorDetails_PartialUpdate)
}
