// Code generated by Kitex v0.11.3. DO NOT EDIT.

package re

import (
	"bytes"
	"fmt"
	"reflect"
	"strings"

	"github.com/cloudwego/gopkg/protocol/thrift"
)

// unused protection
var (
	_ = fmt.Formatter(nil)
	_ = (*bytes.Buffer)(nil)
	_ = (*strings.Builder)(nil)
	_ = reflect.Type(nil)
	_ = thrift.STOP
)

func (p *Image) FastRead(buf []byte) (int, error) {
	var err error
	var offset int
	var l int
	var fieldTypeId thrift.TType
	var fieldId int16
	for {
		fieldTypeId, fieldId, l, err = thrift.Binary.ReadFieldBegin(buf[offset:])
		offset += l
		if err != nil {
			goto ReadFieldBeginError
		}
		if fieldTypeId == thrift.STOP {
			break
		}
		switch fieldId {
		case 1:
			if fieldTypeId == thrift.STRING {
				l, err = p.FastReadField1(buf[offset:])
				offset += l
				if err != nil {
					goto ReadFieldError
				}
			} else {
				l, err = thrift.Binary.Skip(buf[offset:], fieldTypeId)
				offset += l
				if err != nil {
					goto SkipFieldError
				}
			}
		case 2:
			if fieldTypeId == thrift.I32 {
				l, err = p.FastReadField2(buf[offset:])
				offset += l
				if err != nil {
					goto ReadFieldError
				}
			} else {
				l, err = thrift.Binary.Skip(buf[offset:], fieldTypeId)
				offset += l
				if err != nil {
					goto SkipFieldError
				}
			}
		case 3:
			if fieldTypeId == thrift.I32 {
				l, err = p.FastReadField3(buf[offset:])
				offset += l
				if err != nil {
					goto ReadFieldError
				}
			} else {
				l, err = thrift.Binary.Skip(buf[offset:], fieldTypeId)
				offset += l
				if err != nil {
					goto SkipFieldError
				}
			}
		default:
			l, err = thrift.Binary.Skip(buf[offset:], fieldTypeId)
			offset += l
			if err != nil {
				goto SkipFieldError
			}
		}
	}

	return offset, nil
ReadFieldBeginError:
	return offset, thrift.PrependError(fmt.Sprintf("%T read field %d begin error: ", p, fieldId), err)
ReadFieldError:
	return offset, thrift.PrependError(fmt.Sprintf("%T read field %d '%s' error: ", p, fieldId, fieldIDToName_Image[fieldId]), err)
SkipFieldError:
	return offset, thrift.PrependError(fmt.Sprintf("%T field %d skip type %d error: ", p, fieldId, fieldTypeId), err)
}

func (p *Image) FastReadField1(buf []byte) (int, error) {
	offset := 0

	var _field []byte
	if v, l, err := thrift.Binary.ReadBinary(buf[offset:]); err != nil {
		return offset, err
	} else {
		offset += l

		_field = []byte(v)
	}
	p.Content = _field
	return offset, nil
}

func (p *Image) FastReadField2(buf []byte) (int, error) {
	offset := 0

	var _field *int32
	if v, l, err := thrift.Binary.ReadI32(buf[offset:]); err != nil {
		return offset, err
	} else {
		offset += l
		_field = &v
	}
	p.Width = _field
	return offset, nil
}

func (p *Image) FastReadField3(buf []byte) (int, error) {
	offset := 0

	var _field *int32
	if v, l, err := thrift.Binary.ReadI32(buf[offset:]); err != nil {
		return offset, err
	} else {
		offset += l
		_field = &v
	}
	p.Height = _field
	return offset, nil
}

// for compatibility
func (p *Image) FastWrite(buf []byte) int {
	return 0
}

func (p *Image) FastWriteNocopy(buf []byte, w thrift.NocopyWriter) int {
	offset := 0
	if p != nil {
		offset += p.fastWriteField2(buf[offset:], w)
		offset += p.fastWriteField3(buf[offset:], w)
		offset += p.fastWriteField1(buf[offset:], w)
	}
	offset += thrift.Binary.WriteFieldStop(buf[offset:])
	return offset
}

func (p *Image) BLength() int {
	l := 0
	if p != nil {
		l += p.field1Length()
		l += p.field2Length()
		l += p.field3Length()
	}
	l += thrift.Binary.FieldStopLength()
	return l
}

func (p *Image) fastWriteField1(buf []byte, w thrift.NocopyWriter) int {
	offset := 0
	if p.IsSetContent() {
		offset += thrift.Binary.WriteFieldBegin(buf[offset:], thrift.STRING, 1)
		offset += thrift.Binary.WriteBinaryNocopy(buf[offset:], w, []byte(p.Content))
	}
	return offset
}

func (p *Image) fastWriteField2(buf []byte, w thrift.NocopyWriter) int {
	offset := 0
	if p.IsSetWidth() {
		offset += thrift.Binary.WriteFieldBegin(buf[offset:], thrift.I32, 2)
		offset += thrift.Binary.WriteI32(buf[offset:], *p.Width)
	}
	return offset
}

func (p *Image) fastWriteField3(buf []byte, w thrift.NocopyWriter) int {
	offset := 0
	if p.IsSetHeight() {
		offset += thrift.Binary.WriteFieldBegin(buf[offset:], thrift.I32, 3)
		offset += thrift.Binary.WriteI32(buf[offset:], *p.Height)
	}
	return offset
}

func (p *Image) field1Length() int {
	l := 0
	if p.IsSetContent() {
		l += thrift.Binary.FieldBeginLength()
		l += thrift.Binary.BinaryLengthNocopy([]byte(p.Content))
	}
	return l
}

func (p *Image) field2Length() int {
	l := 0
	if p.IsSetWidth() {
		l += thrift.Binary.FieldBeginLength()
		l += thrift.Binary.I32Length()
	}
	return l
}

func (p *Image) field3Length() int {
	l := 0
	if p.IsSetHeight() {
		l += thrift.Binary.FieldBeginLength()
		l += thrift.Binary.I32Length()
	}
	return l
}

func (p *ImagePredictRequest) FastRead(buf []byte) (int, error) {
	var err error
	var offset int
	var l int
	var fieldTypeId thrift.TType
	var fieldId int16
	for {
		fieldTypeId, fieldId, l, err = thrift.Binary.ReadFieldBegin(buf[offset:])
		offset += l
		if err != nil {
			goto ReadFieldBeginError
		}
		if fieldTypeId == thrift.STOP {
			break
		}
		switch fieldId {
		case 1:
			if fieldTypeId == thrift.STRING {
				l, err = p.FastReadField1(buf[offset:])
				offset += l
				if err != nil {
					goto ReadFieldError
				}
			} else {
				l, err = thrift.Binary.Skip(buf[offset:], fieldTypeId)
				offset += l
				if err != nil {
					goto SkipFieldError
				}
			}
		case 2:
			if fieldTypeId == thrift.STRING {
				l, err = p.FastReadField2(buf[offset:])
				offset += l
				if err != nil {
					goto ReadFieldError
				}
			} else {
				l, err = thrift.Binary.Skip(buf[offset:], fieldTypeId)
				offset += l
				if err != nil {
					goto SkipFieldError
				}
			}
		case 3:
			if fieldTypeId == thrift.STRING {
				l, err = p.FastReadField3(buf[offset:])
				offset += l
				if err != nil {
					goto ReadFieldError
				}
			} else {
				l, err = thrift.Binary.Skip(buf[offset:], fieldTypeId)
				offset += l
				if err != nil {
					goto SkipFieldError
				}
			}
		case 4:
			if fieldTypeId == thrift.STRING {
				l, err = p.FastReadField4(buf[offset:])
				offset += l
				if err != nil {
					goto ReadFieldError
				}
			} else {
				l, err = thrift.Binary.Skip(buf[offset:], fieldTypeId)
				offset += l
				if err != nil {
					goto SkipFieldError
				}
			}
		case 5:
			if fieldTypeId == thrift.STRING {
				l, err = p.FastReadField5(buf[offset:])
				offset += l
				if err != nil {
					goto ReadFieldError
				}
			} else {
				l, err = thrift.Binary.Skip(buf[offset:], fieldTypeId)
				offset += l
				if err != nil {
					goto SkipFieldError
				}
			}
		case 6:
			if fieldTypeId == thrift.STRING {
				l, err = p.FastReadField6(buf[offset:])
				offset += l
				if err != nil {
					goto ReadFieldError
				}
			} else {
				l, err = thrift.Binary.Skip(buf[offset:], fieldTypeId)
				offset += l
				if err != nil {
					goto SkipFieldError
				}
			}
		case 7:
			if fieldTypeId == thrift.STRING {
				l, err = p.FastReadField7(buf[offset:])
				offset += l
				if err != nil {
					goto ReadFieldError
				}
			} else {
				l, err = thrift.Binary.Skip(buf[offset:], fieldTypeId)
				offset += l
				if err != nil {
					goto SkipFieldError
				}
			}
		case 8:
			if fieldTypeId == thrift.STRUCT {
				l, err = p.FastReadField8(buf[offset:])
				offset += l
				if err != nil {
					goto ReadFieldError
				}
			} else {
				l, err = thrift.Binary.Skip(buf[offset:], fieldTypeId)
				offset += l
				if err != nil {
					goto SkipFieldError
				}
			}
		default:
			l, err = thrift.Binary.Skip(buf[offset:], fieldTypeId)
			offset += l
			if err != nil {
				goto SkipFieldError
			}
		}
	}

	return offset, nil
ReadFieldBeginError:
	return offset, thrift.PrependError(fmt.Sprintf("%T read field %d begin error: ", p, fieldId), err)
ReadFieldError:
	return offset, thrift.PrependError(fmt.Sprintf("%T read field %d '%s' error: ", p, fieldId, fieldIDToName_ImagePredictRequest[fieldId]), err)
SkipFieldError:
	return offset, thrift.PrependError(fmt.Sprintf("%T field %d skip type %d error: ", p, fieldId, fieldTypeId), err)
}

func (p *ImagePredictRequest) FastReadField1(buf []byte) (int, error) {
	offset := 0

	var _field *string
	if v, l, err := thrift.Binary.ReadString(buf[offset:]); err != nil {
		return offset, err
	} else {
		offset += l
		_field = &v
	}
	p.RequestId = _field
	return offset, nil
}

func (p *ImagePredictRequest) FastReadField2(buf []byte) (int, error) {
	offset := 0

	var _field *string
	if v, l, err := thrift.Binary.ReadString(buf[offset:]); err != nil {
		return offset, err
	} else {
		offset += l
		_field = &v
	}
	p.Organization = _field
	return offset, nil
}

func (p *ImagePredictRequest) FastReadField3(buf []byte) (int, error) {
	offset := 0

	var _field *string
	if v, l, err := thrift.Binary.ReadString(buf[offset:]); err != nil {
		return offset, err
	} else {
		offset += l
		_field = &v
	}
	p.Product = _field
	return offset, nil
}

func (p *ImagePredictRequest) FastReadField4(buf []byte) (int, error) {
	offset := 0

	var _field *string
	if v, l, err := thrift.Binary.ReadString(buf[offset:]); err != nil {
		return offset, err
	} else {
		offset += l
		_field = &v
	}
	p.SceneId = _field
	return offset, nil
}

func (p *ImagePredictRequest) FastReadField5(buf []byte) (int, error) {
	offset := 0

	var _field *string
	if v, l, err := thrift.Binary.ReadString(buf[offset:]); err != nil {
		return offset, err
	} else {
		offset += l
		_field = &v
	}
	p.AppId = _field
	return offset, nil
}

func (p *ImagePredictRequest) FastReadField6(buf []byte) (int, error) {
	offset := 0

	var _field *string
	if v, l, err := thrift.Binary.ReadString(buf[offset:]); err != nil {
		return offset, err
	} else {
		offset += l
		_field = &v
	}
	p.Channel = _field
	return offset, nil
}

func (p *ImagePredictRequest) FastReadField7(buf []byte) (int, error) {
	offset := 0

	var _field *string
	if v, l, err := thrift.Binary.ReadString(buf[offset:]); err != nil {
		return offset, err
	} else {
		offset += l
		_field = &v
	}
	p.Data = _field
	return offset, nil
}

func (p *ImagePredictRequest) FastReadField8(buf []byte) (int, error) {
	offset := 0
	_field := NewImage()
	if l, err := _field.FastRead(buf[offset:]); err != nil {
		return offset, err
	} else {
		offset += l
	}
	p.Image = _field
	return offset, nil
}

// for compatibility
func (p *ImagePredictRequest) FastWrite(buf []byte) int {
	return 0
}

func (p *ImagePredictRequest) FastWriteNocopy(buf []byte, w thrift.NocopyWriter) int {
	offset := 0
	if p != nil {
		offset += p.fastWriteField1(buf[offset:], w)
		offset += p.fastWriteField2(buf[offset:], w)
		offset += p.fastWriteField3(buf[offset:], w)
		offset += p.fastWriteField4(buf[offset:], w)
		offset += p.fastWriteField5(buf[offset:], w)
		offset += p.fastWriteField6(buf[offset:], w)
		offset += p.fastWriteField7(buf[offset:], w)
		offset += p.fastWriteField8(buf[offset:], w)
	}
	offset += thrift.Binary.WriteFieldStop(buf[offset:])
	return offset
}

func (p *ImagePredictRequest) BLength() int {
	l := 0
	if p != nil {
		l += p.field1Length()
		l += p.field2Length()
		l += p.field3Length()
		l += p.field4Length()
		l += p.field5Length()
		l += p.field6Length()
		l += p.field7Length()
		l += p.field8Length()
	}
	l += thrift.Binary.FieldStopLength()
	return l
}

func (p *ImagePredictRequest) fastWriteField1(buf []byte, w thrift.NocopyWriter) int {
	offset := 0
	if p.IsSetRequestId() {
		offset += thrift.Binary.WriteFieldBegin(buf[offset:], thrift.STRING, 1)
		offset += thrift.Binary.WriteStringNocopy(buf[offset:], w, *p.RequestId)
	}
	return offset
}

func (p *ImagePredictRequest) fastWriteField2(buf []byte, w thrift.NocopyWriter) int {
	offset := 0
	if p.IsSetOrganization() {
		offset += thrift.Binary.WriteFieldBegin(buf[offset:], thrift.STRING, 2)
		offset += thrift.Binary.WriteStringNocopy(buf[offset:], w, *p.Organization)
	}
	return offset
}

func (p *ImagePredictRequest) fastWriteField3(buf []byte, w thrift.NocopyWriter) int {
	offset := 0
	if p.IsSetProduct() {
		offset += thrift.Binary.WriteFieldBegin(buf[offset:], thrift.STRING, 3)
		offset += thrift.Binary.WriteStringNocopy(buf[offset:], w, *p.Product)
	}
	return offset
}

func (p *ImagePredictRequest) fastWriteField4(buf []byte, w thrift.NocopyWriter) int {
	offset := 0
	if p.IsSetSceneId() {
		offset += thrift.Binary.WriteFieldBegin(buf[offset:], thrift.STRING, 4)
		offset += thrift.Binary.WriteStringNocopy(buf[offset:], w, *p.SceneId)
	}
	return offset
}

func (p *ImagePredictRequest) fastWriteField5(buf []byte, w thrift.NocopyWriter) int {
	offset := 0
	if p.IsSetAppId() {
		offset += thrift.Binary.WriteFieldBegin(buf[offset:], thrift.STRING, 5)
		offset += thrift.Binary.WriteStringNocopy(buf[offset:], w, *p.AppId)
	}
	return offset
}

func (p *ImagePredictRequest) fastWriteField6(buf []byte, w thrift.NocopyWriter) int {
	offset := 0
	if p.IsSetChannel() {
		offset += thrift.Binary.WriteFieldBegin(buf[offset:], thrift.STRING, 6)
		offset += thrift.Binary.WriteStringNocopy(buf[offset:], w, *p.Channel)
	}
	return offset
}

func (p *ImagePredictRequest) fastWriteField7(buf []byte, w thrift.NocopyWriter) int {
	offset := 0
	if p.IsSetData() {
		offset += thrift.Binary.WriteFieldBegin(buf[offset:], thrift.STRING, 7)
		offset += thrift.Binary.WriteStringNocopy(buf[offset:], w, *p.Data)
	}
	return offset
}

func (p *ImagePredictRequest) fastWriteField8(buf []byte, w thrift.NocopyWriter) int {
	offset := 0
	if p.IsSetImage() {
		offset += thrift.Binary.WriteFieldBegin(buf[offset:], thrift.STRUCT, 8)
		offset += p.Image.FastWriteNocopy(buf[offset:], w)
	}
	return offset
}

func (p *ImagePredictRequest) field1Length() int {
	l := 0
	if p.IsSetRequestId() {
		l += thrift.Binary.FieldBeginLength()
		l += thrift.Binary.StringLengthNocopy(*p.RequestId)
	}
	return l
}

func (p *ImagePredictRequest) field2Length() int {
	l := 0
	if p.IsSetOrganization() {
		l += thrift.Binary.FieldBeginLength()
		l += thrift.Binary.StringLengthNocopy(*p.Organization)
	}
	return l
}

func (p *ImagePredictRequest) field3Length() int {
	l := 0
	if p.IsSetProduct() {
		l += thrift.Binary.FieldBeginLength()
		l += thrift.Binary.StringLengthNocopy(*p.Product)
	}
	return l
}

func (p *ImagePredictRequest) field4Length() int {
	l := 0
	if p.IsSetSceneId() {
		l += thrift.Binary.FieldBeginLength()
		l += thrift.Binary.StringLengthNocopy(*p.SceneId)
	}
	return l
}

func (p *ImagePredictRequest) field5Length() int {
	l := 0
	if p.IsSetAppId() {
		l += thrift.Binary.FieldBeginLength()
		l += thrift.Binary.StringLengthNocopy(*p.AppId)
	}
	return l
}

func (p *ImagePredictRequest) field6Length() int {
	l := 0
	if p.IsSetChannel() {
		l += thrift.Binary.FieldBeginLength()
		l += thrift.Binary.StringLengthNocopy(*p.Channel)
	}
	return l
}

func (p *ImagePredictRequest) field7Length() int {
	l := 0
	if p.IsSetData() {
		l += thrift.Binary.FieldBeginLength()
		l += thrift.Binary.StringLengthNocopy(*p.Data)
	}
	return l
}

func (p *ImagePredictRequest) field8Length() int {
	l := 0
	if p.IsSetImage() {
		l += thrift.Binary.FieldBeginLength()
		l += p.Image.BLength()
	}
	return l
}

func (p *ImagePredictResult_) FastRead(buf []byte) (int, error) {
	var err error
	var offset int
	var l int
	var fieldTypeId thrift.TType
	var fieldId int16
	for {
		fieldTypeId, fieldId, l, err = thrift.Binary.ReadFieldBegin(buf[offset:])
		offset += l
		if err != nil {
			goto ReadFieldBeginError
		}
		if fieldTypeId == thrift.STOP {
			break
		}
		switch fieldId {
		case 1:
			if fieldTypeId == thrift.STRING {
				l, err = p.FastReadField1(buf[offset:])
				offset += l
				if err != nil {
					goto ReadFieldError
				}
			} else {
				l, err = thrift.Binary.Skip(buf[offset:], fieldTypeId)
				offset += l
				if err != nil {
					goto SkipFieldError
				}
			}
		default:
			l, err = thrift.Binary.Skip(buf[offset:], fieldTypeId)
			offset += l
			if err != nil {
				goto SkipFieldError
			}
		}
	}

	return offset, nil
ReadFieldBeginError:
	return offset, thrift.PrependError(fmt.Sprintf("%T read field %d begin error: ", p, fieldId), err)
ReadFieldError:
	return offset, thrift.PrependError(fmt.Sprintf("%T read field %d '%s' error: ", p, fieldId, fieldIDToName_ImagePredictResult_[fieldId]), err)
SkipFieldError:
	return offset, thrift.PrependError(fmt.Sprintf("%T field %d skip type %d error: ", p, fieldId, fieldTypeId), err)
}

func (p *ImagePredictResult_) FastReadField1(buf []byte) (int, error) {
	offset := 0

	var _field *string
	if v, l, err := thrift.Binary.ReadString(buf[offset:]); err != nil {
		return offset, err
	} else {
		offset += l
		_field = &v
	}
	p.Result_ = _field
	return offset, nil
}

// for compatibility
func (p *ImagePredictResult_) FastWrite(buf []byte) int {
	return 0
}

func (p *ImagePredictResult_) FastWriteNocopy(buf []byte, w thrift.NocopyWriter) int {
	offset := 0
	if p != nil {
		offset += p.fastWriteField1(buf[offset:], w)
	}
	offset += thrift.Binary.WriteFieldStop(buf[offset:])
	return offset
}

func (p *ImagePredictResult_) BLength() int {
	l := 0
	if p != nil {
		l += p.field1Length()
	}
	l += thrift.Binary.FieldStopLength()
	return l
}

func (p *ImagePredictResult_) fastWriteField1(buf []byte, w thrift.NocopyWriter) int {
	offset := 0
	if p.IsSetResult_() {
		offset += thrift.Binary.WriteFieldBegin(buf[offset:], thrift.STRING, 1)
		offset += thrift.Binary.WriteStringNocopy(buf[offset:], w, *p.Result_)
	}
	return offset
}

func (p *ImagePredictResult_) field1Length() int {
	l := 0
	if p.IsSetResult_() {
		l += thrift.Binary.FieldBeginLength()
		l += thrift.Binary.StringLengthNocopy(*p.Result_)
	}
	return l
}

func (p *ImagePredictorPredictArgs) FastRead(buf []byte) (int, error) {
	var err error
	var offset int
	var l int
	var fieldTypeId thrift.TType
	var fieldId int16
	for {
		fieldTypeId, fieldId, l, err = thrift.Binary.ReadFieldBegin(buf[offset:])
		offset += l
		if err != nil {
			goto ReadFieldBeginError
		}
		if fieldTypeId == thrift.STOP {
			break
		}
		switch fieldId {
		case 1:
			if fieldTypeId == thrift.STRUCT {
				l, err = p.FastReadField1(buf[offset:])
				offset += l
				if err != nil {
					goto ReadFieldError
				}
			} else {
				l, err = thrift.Binary.Skip(buf[offset:], fieldTypeId)
				offset += l
				if err != nil {
					goto SkipFieldError
				}
			}
		default:
			l, err = thrift.Binary.Skip(buf[offset:], fieldTypeId)
			offset += l
			if err != nil {
				goto SkipFieldError
			}
		}
	}

	return offset, nil
ReadFieldBeginError:
	return offset, thrift.PrependError(fmt.Sprintf("%T read field %d begin error: ", p, fieldId), err)
ReadFieldError:
	return offset, thrift.PrependError(fmt.Sprintf("%T read field %d '%s' error: ", p, fieldId, fieldIDToName_ImagePredictorPredictArgs[fieldId]), err)
SkipFieldError:
	return offset, thrift.PrependError(fmt.Sprintf("%T field %d skip type %d error: ", p, fieldId, fieldTypeId), err)
}

func (p *ImagePredictorPredictArgs) FastReadField1(buf []byte) (int, error) {
	offset := 0
	_field := NewImagePredictRequest()
	if l, err := _field.FastRead(buf[offset:]); err != nil {
		return offset, err
	} else {
		offset += l
	}
	p.Request = _field
	return offset, nil
}

// for compatibility
func (p *ImagePredictorPredictArgs) FastWrite(buf []byte) int {
	return 0
}

func (p *ImagePredictorPredictArgs) FastWriteNocopy(buf []byte, w thrift.NocopyWriter) int {
	offset := 0
	if p != nil {
		offset += p.fastWriteField1(buf[offset:], w)
	}
	offset += thrift.Binary.WriteFieldStop(buf[offset:])
	return offset
}

func (p *ImagePredictorPredictArgs) BLength() int {
	l := 0
	if p != nil {
		l += p.field1Length()
	}
	l += thrift.Binary.FieldStopLength()
	return l
}

func (p *ImagePredictorPredictArgs) fastWriteField1(buf []byte, w thrift.NocopyWriter) int {
	offset := 0
	offset += thrift.Binary.WriteFieldBegin(buf[offset:], thrift.STRUCT, 1)
	offset += p.Request.FastWriteNocopy(buf[offset:], w)
	return offset
}

func (p *ImagePredictorPredictArgs) field1Length() int {
	l := 0
	l += thrift.Binary.FieldBeginLength()
	l += p.Request.BLength()
	return l
}

func (p *ImagePredictorPredictResult) FastRead(buf []byte) (int, error) {
	var err error
	var offset int
	var l int
	var fieldTypeId thrift.TType
	var fieldId int16
	for {
		fieldTypeId, fieldId, l, err = thrift.Binary.ReadFieldBegin(buf[offset:])
		offset += l
		if err != nil {
			goto ReadFieldBeginError
		}
		if fieldTypeId == thrift.STOP {
			break
		}
		switch fieldId {
		case 0:
			if fieldTypeId == thrift.STRUCT {
				l, err = p.FastReadField0(buf[offset:])
				offset += l
				if err != nil {
					goto ReadFieldError
				}
			} else {
				l, err = thrift.Binary.Skip(buf[offset:], fieldTypeId)
				offset += l
				if err != nil {
					goto SkipFieldError
				}
			}
		default:
			l, err = thrift.Binary.Skip(buf[offset:], fieldTypeId)
			offset += l
			if err != nil {
				goto SkipFieldError
			}
		}
	}

	return offset, nil
ReadFieldBeginError:
	return offset, thrift.PrependError(fmt.Sprintf("%T read field %d begin error: ", p, fieldId), err)
ReadFieldError:
	return offset, thrift.PrependError(fmt.Sprintf("%T read field %d '%s' error: ", p, fieldId, fieldIDToName_ImagePredictorPredictResult[fieldId]), err)
SkipFieldError:
	return offset, thrift.PrependError(fmt.Sprintf("%T field %d skip type %d error: ", p, fieldId, fieldTypeId), err)
}

func (p *ImagePredictorPredictResult) FastReadField0(buf []byte) (int, error) {
	offset := 0
	_field := NewImagePredictResult_()
	if l, err := _field.FastRead(buf[offset:]); err != nil {
		return offset, err
	} else {
		offset += l
	}
	p.Success = _field
	return offset, nil
}

// for compatibility
func (p *ImagePredictorPredictResult) FastWrite(buf []byte) int {
	return 0
}

func (p *ImagePredictorPredictResult) FastWriteNocopy(buf []byte, w thrift.NocopyWriter) int {
	offset := 0
	if p != nil {
		offset += p.fastWriteField0(buf[offset:], w)
	}
	offset += thrift.Binary.WriteFieldStop(buf[offset:])
	return offset
}

func (p *ImagePredictorPredictResult) BLength() int {
	l := 0
	if p != nil {
		l += p.field0Length()
	}
	l += thrift.Binary.FieldStopLength()
	return l
}

func (p *ImagePredictorPredictResult) fastWriteField0(buf []byte, w thrift.NocopyWriter) int {
	offset := 0
	if p.IsSetSuccess() {
		offset += thrift.Binary.WriteFieldBegin(buf[offset:], thrift.STRUCT, 0)
		offset += p.Success.FastWriteNocopy(buf[offset:], w)
	}
	return offset
}

func (p *ImagePredictorPredictResult) field0Length() int {
	l := 0
	if p.IsSetSuccess() {
		l += thrift.Binary.FieldBeginLength()
		l += p.Success.BLength()
	}
	return l
}

func (p *ImagePredictorHealthArgs) FastRead(buf []byte) (int, error) {
	var err error
	var offset int
	var l int
	var fieldTypeId thrift.TType
	var fieldId int16
	for {
		fieldTypeId, fieldId, l, err = thrift.Binary.ReadFieldBegin(buf[offset:])
		offset += l
		if err != nil {
			goto ReadFieldBeginError
		}
		if fieldTypeId == thrift.STOP {
			break
		}
		l, err = thrift.Binary.Skip(buf[offset:], fieldTypeId)
		offset += l
		if err != nil {
			goto SkipFieldError
		}
	}

	return offset, nil
ReadFieldBeginError:
	return offset, thrift.PrependError(fmt.Sprintf("%T read field %d begin error: ", p, fieldId), err)
SkipFieldError:
	return offset, thrift.PrependError(fmt.Sprintf("%T field %d skip type %d error: ", p, fieldId, fieldTypeId), err)
}

// for compatibility
func (p *ImagePredictorHealthArgs) FastWrite(buf []byte) int {
	return 0
}

func (p *ImagePredictorHealthArgs) FastWriteNocopy(buf []byte, w thrift.NocopyWriter) int {
	offset := 0
	if p != nil {
	}
	offset += thrift.Binary.WriteFieldStop(buf[offset:])
	return offset
}

func (p *ImagePredictorHealthArgs) BLength() int {
	l := 0
	if p != nil {
	}
	l += thrift.Binary.FieldStopLength()
	return l
}

func (p *ImagePredictorHealthResult) FastRead(buf []byte) (int, error) {
	var err error
	var offset int
	var l int
	var fieldTypeId thrift.TType
	var fieldId int16
	for {
		fieldTypeId, fieldId, l, err = thrift.Binary.ReadFieldBegin(buf[offset:])
		offset += l
		if err != nil {
			goto ReadFieldBeginError
		}
		if fieldTypeId == thrift.STOP {
			break
		}
		switch fieldId {
		case 0:
			if fieldTypeId == thrift.BOOL {
				l, err = p.FastReadField0(buf[offset:])
				offset += l
				if err != nil {
					goto ReadFieldError
				}
			} else {
				l, err = thrift.Binary.Skip(buf[offset:], fieldTypeId)
				offset += l
				if err != nil {
					goto SkipFieldError
				}
			}
		default:
			l, err = thrift.Binary.Skip(buf[offset:], fieldTypeId)
			offset += l
			if err != nil {
				goto SkipFieldError
			}
		}
	}

	return offset, nil
ReadFieldBeginError:
	return offset, thrift.PrependError(fmt.Sprintf("%T read field %d begin error: ", p, fieldId), err)
ReadFieldError:
	return offset, thrift.PrependError(fmt.Sprintf("%T read field %d '%s' error: ", p, fieldId, fieldIDToName_ImagePredictorHealthResult[fieldId]), err)
SkipFieldError:
	return offset, thrift.PrependError(fmt.Sprintf("%T field %d skip type %d error: ", p, fieldId, fieldTypeId), err)
}

func (p *ImagePredictorHealthResult) FastReadField0(buf []byte) (int, error) {
	offset := 0

	var _field *bool
	if v, l, err := thrift.Binary.ReadBool(buf[offset:]); err != nil {
		return offset, err
	} else {
		offset += l
		_field = &v
	}
	p.Success = _field
	return offset, nil
}

// for compatibility
func (p *ImagePredictorHealthResult) FastWrite(buf []byte) int {
	return 0
}

func (p *ImagePredictorHealthResult) FastWriteNocopy(buf []byte, w thrift.NocopyWriter) int {
	offset := 0
	if p != nil {
		offset += p.fastWriteField0(buf[offset:], w)
	}
	offset += thrift.Binary.WriteFieldStop(buf[offset:])
	return offset
}

func (p *ImagePredictorHealthResult) BLength() int {
	l := 0
	if p != nil {
		l += p.field0Length()
	}
	l += thrift.Binary.FieldStopLength()
	return l
}

func (p *ImagePredictorHealthResult) fastWriteField0(buf []byte, w thrift.NocopyWriter) int {
	offset := 0
	if p.IsSetSuccess() {
		offset += thrift.Binary.WriteFieldBegin(buf[offset:], thrift.BOOL, 0)
		offset += thrift.Binary.WriteBool(buf[offset:], *p.Success)
	}
	return offset
}

func (p *ImagePredictorHealthResult) field0Length() int {
	l := 0
	if p.IsSetSuccess() {
		l += thrift.Binary.FieldBeginLength()
		l += thrift.Binary.BoolLength()
	}
	return l
}

func (p *ImagePredictorPredictArgs) GetFirstArgument() interface{} {
	return p.Request
}

func (p *ImagePredictorPredictResult) GetResult() interface{} {
	return p.Success
}

func (p *ImagePredictorHealthArgs) GetFirstArgument() interface{} {
	return nil
}

func (p *ImagePredictorHealthResult) GetResult() interface{} {
	return p.Success
}
