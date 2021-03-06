// Code generated by Kitex v0.1.0. DO NOT EDIT.

package base

import (
	"bytes"
	"fmt"
	"reflect"
	"strings"

	"github.com/apache/thrift/lib/go/thrift"

	"github.com/cloudwego/kitex/pkg/protocol/bthrift"
)

// unused protection
var (
	_ = fmt.Formatter(nil)
	_ = (*bytes.Buffer)(nil)
	_ = (*strings.Builder)(nil)
	_ = reflect.Type(nil)
	_ = thrift.TProtocol(nil)
	_ = bthrift.BinaryWriter(nil)
)

func (p *TrafficEnv) FastRead(buf []byte) (int, error) {
	var err error
	var offset int
	var l int
	var fieldTypeId thrift.TType
	var fieldId int16
	_, l, err = bthrift.Binary.ReadStructBegin(buf)
	offset += l
	if err != nil {
		goto ReadStructBeginError
	}

	for {
		_, fieldTypeId, fieldId, l, err = bthrift.Binary.ReadFieldBegin(buf[offset:])
		offset += l
		if err != nil {
			goto ReadFieldBeginError
		}
		if fieldTypeId == thrift.STOP {
			break
		}
		switch fieldId {
		case 1:
			if fieldTypeId == thrift.BOOL {
				l, err = p.FastReadField1(buf[offset:])
				offset += l
				if err != nil {
					goto ReadFieldError
				}
			} else {
				l, err = bthrift.Binary.Skip(buf[offset:], fieldTypeId)
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
				l, err = bthrift.Binary.Skip(buf[offset:], fieldTypeId)
				offset += l
				if err != nil {
					goto SkipFieldError
				}
			}
		default:
			l, err = bthrift.Binary.Skip(buf[offset:], fieldTypeId)
			offset += l
			if err != nil {
				goto SkipFieldError
			}
		}

		l, err = bthrift.Binary.ReadFieldEnd(buf[offset:])
		offset += l
		if err != nil {
			goto ReadFieldEndError
		}
	}
	l, err = bthrift.Binary.ReadStructEnd(buf[offset:])
	offset += l
	if err != nil {
		goto ReadStructEndError
	}

	return offset, nil
ReadStructBeginError:
	return offset, thrift.PrependError(fmt.Sprintf("%T read struct begin error: ", p), err)
ReadFieldBeginError:
	return offset, thrift.PrependError(fmt.Sprintf("%T read field %d begin error: ", p, fieldId), err)
ReadFieldError:
	return offset, thrift.PrependError(fmt.Sprintf("%T read field %d '%s' error: ", p, fieldId, fieldIDToName_TrafficEnv[fieldId]), err)
SkipFieldError:
	return offset, thrift.PrependError(fmt.Sprintf("%T field %d skip type %d error: ", p, fieldId, fieldTypeId), err)
ReadFieldEndError:
	return offset, thrift.PrependError(fmt.Sprintf("%T read field end error", p), err)
ReadStructEndError:
	return offset, thrift.PrependError(fmt.Sprintf("%T read struct end error: ", p), err)
}

func (p *TrafficEnv) FastReadField1(buf []byte) (int, error) {
	offset := 0

	if v, l, err := bthrift.Binary.ReadBool(buf[offset:]); err != nil {
		return offset, err
	} else {
		offset += l

		p.Open = v

	}
	return offset, nil
}

func (p *TrafficEnv) FastReadField2(buf []byte) (int, error) {
	offset := 0

	if v, l, err := bthrift.Binary.ReadString(buf[offset:]); err != nil {
		return offset, err
	} else {
		offset += l

		p.Env = v

	}
	return offset, nil
}

// for compatibility
func (p *TrafficEnv) FastWrite(buf []byte) int {
	return 0
}

func (p *TrafficEnv) FastWriteNocopy(buf []byte, binaryWriter bthrift.BinaryWriter) int {
	offset := 0
	offset += bthrift.Binary.WriteStructBegin(buf[offset:], "TrafficEnv")
	if p != nil {
		offset += p.fastWriteField1(buf[offset:], binaryWriter)
		offset += p.fastWriteField2(buf[offset:], binaryWriter)
	}
	offset += bthrift.Binary.WriteFieldStop(buf[offset:])
	offset += bthrift.Binary.WriteStructEnd(buf[offset:])
	return offset
}

func (p *TrafficEnv) BLength() int {
	l := 0
	l += bthrift.Binary.StructBeginLength("TrafficEnv")
	if p != nil {
		l += p.field1Length()
		l += p.field2Length()
	}
	l += bthrift.Binary.FieldStopLength()
	l += bthrift.Binary.StructEndLength()
	return l
}

func (p *TrafficEnv) fastWriteField1(buf []byte, binaryWriter bthrift.BinaryWriter) int {
	offset := 0
	offset += bthrift.Binary.WriteFieldBegin(buf[offset:], "Open", thrift.BOOL, 1)
	offset += bthrift.Binary.WriteBool(buf[offset:], p.Open)

	offset += bthrift.Binary.WriteFieldEnd(buf[offset:])
	return offset
}

func (p *TrafficEnv) fastWriteField2(buf []byte, binaryWriter bthrift.BinaryWriter) int {
	offset := 0
	offset += bthrift.Binary.WriteFieldBegin(buf[offset:], "Env", thrift.STRING, 2)
	offset += bthrift.Binary.WriteStringNocopy(buf[offset:], binaryWriter, p.Env)

	offset += bthrift.Binary.WriteFieldEnd(buf[offset:])
	return offset
}

func (p *TrafficEnv) field1Length() int {
	l := 0
	l += bthrift.Binary.FieldBeginLength("Open", thrift.BOOL, 1)
	l += bthrift.Binary.BoolLength(p.Open)

	l += bthrift.Binary.FieldEndLength()
	return l
}

func (p *TrafficEnv) field2Length() int {
	l := 0
	l += bthrift.Binary.FieldBeginLength("Env", thrift.STRING, 2)
	l += bthrift.Binary.StringLengthNocopy(p.Env)

	l += bthrift.Binary.FieldEndLength()
	return l
}

func (p *Base) FastRead(buf []byte) (int, error) {
	var err error
	var offset int
	var l int
	var fieldTypeId thrift.TType
	var fieldId int16
	_, l, err = bthrift.Binary.ReadStructBegin(buf)
	offset += l
	if err != nil {
		goto ReadStructBeginError
	}

	for {
		_, fieldTypeId, fieldId, l, err = bthrift.Binary.ReadFieldBegin(buf[offset:])
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
				l, err = bthrift.Binary.Skip(buf[offset:], fieldTypeId)
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
				l, err = bthrift.Binary.Skip(buf[offset:], fieldTypeId)
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
				l, err = bthrift.Binary.Skip(buf[offset:], fieldTypeId)
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
				l, err = bthrift.Binary.Skip(buf[offset:], fieldTypeId)
				offset += l
				if err != nil {
					goto SkipFieldError
				}
			}
		case 5:
			if fieldTypeId == thrift.STRUCT {
				l, err = p.FastReadField5(buf[offset:])
				offset += l
				if err != nil {
					goto ReadFieldError
				}
			} else {
				l, err = bthrift.Binary.Skip(buf[offset:], fieldTypeId)
				offset += l
				if err != nil {
					goto SkipFieldError
				}
			}
		case 6:
			if fieldTypeId == thrift.MAP {
				l, err = p.FastReadField6(buf[offset:])
				offset += l
				if err != nil {
					goto ReadFieldError
				}
			} else {
				l, err = bthrift.Binary.Skip(buf[offset:], fieldTypeId)
				offset += l
				if err != nil {
					goto SkipFieldError
				}
			}
		default:
			l, err = bthrift.Binary.Skip(buf[offset:], fieldTypeId)
			offset += l
			if err != nil {
				goto SkipFieldError
			}
		}

		l, err = bthrift.Binary.ReadFieldEnd(buf[offset:])
		offset += l
		if err != nil {
			goto ReadFieldEndError
		}
	}
	l, err = bthrift.Binary.ReadStructEnd(buf[offset:])
	offset += l
	if err != nil {
		goto ReadStructEndError
	}

	return offset, nil
ReadStructBeginError:
	return offset, thrift.PrependError(fmt.Sprintf("%T read struct begin error: ", p), err)
ReadFieldBeginError:
	return offset, thrift.PrependError(fmt.Sprintf("%T read field %d begin error: ", p, fieldId), err)
ReadFieldError:
	return offset, thrift.PrependError(fmt.Sprintf("%T read field %d '%s' error: ", p, fieldId, fieldIDToName_Base[fieldId]), err)
SkipFieldError:
	return offset, thrift.PrependError(fmt.Sprintf("%T field %d skip type %d error: ", p, fieldId, fieldTypeId), err)
ReadFieldEndError:
	return offset, thrift.PrependError(fmt.Sprintf("%T read field end error", p), err)
ReadStructEndError:
	return offset, thrift.PrependError(fmt.Sprintf("%T read struct end error: ", p), err)
}

func (p *Base) FastReadField1(buf []byte) (int, error) {
	offset := 0

	if v, l, err := bthrift.Binary.ReadString(buf[offset:]); err != nil {
		return offset, err
	} else {
		offset += l

		p.LogID = v

	}
	return offset, nil
}

func (p *Base) FastReadField2(buf []byte) (int, error) {
	offset := 0

	if v, l, err := bthrift.Binary.ReadString(buf[offset:]); err != nil {
		return offset, err
	} else {
		offset += l

		p.Caller = v

	}
	return offset, nil
}

func (p *Base) FastReadField3(buf []byte) (int, error) {
	offset := 0

	if v, l, err := bthrift.Binary.ReadString(buf[offset:]); err != nil {
		return offset, err
	} else {
		offset += l

		p.Addr = v

	}
	return offset, nil
}

func (p *Base) FastReadField4(buf []byte) (int, error) {
	offset := 0

	if v, l, err := bthrift.Binary.ReadString(buf[offset:]); err != nil {
		return offset, err
	} else {
		offset += l

		p.Client = v

	}
	return offset, nil
}

func (p *Base) FastReadField5(buf []byte) (int, error) {
	offset := 0
	p.TrafficEnv = NewTrafficEnv()
	if l, err := p.TrafficEnv.FastRead(buf[offset:]); err != nil {
		return offset, err
	} else {
		offset += l
	}
	return offset, nil
}

func (p *Base) FastReadField6(buf []byte) (int, error) {
	offset := 0

	_, _, size, l, err := bthrift.Binary.ReadMapBegin(buf[offset:])
	offset += l
	if err != nil {
		return offset, err
	}
	p.Extra = make(map[string]string, size)
	for i := 0; i < size; i++ {
		var _key string
		if v, l, err := bthrift.Binary.ReadString(buf[offset:]); err != nil {
			return offset, err
		} else {
			offset += l

			_key = v

		}

		var _val string
		if v, l, err := bthrift.Binary.ReadString(buf[offset:]); err != nil {
			return offset, err
		} else {
			offset += l

			_val = v

		}

		p.Extra[_key] = _val
	}
	if l, err := bthrift.Binary.ReadMapEnd(buf[offset:]); err != nil {
		return offset, err
	} else {
		offset += l
	}
	return offset, nil
}

// for compatibility
func (p *Base) FastWrite(buf []byte) int {
	return 0
}

func (p *Base) FastWriteNocopy(buf []byte, binaryWriter bthrift.BinaryWriter) int {
	offset := 0
	offset += bthrift.Binary.WriteStructBegin(buf[offset:], "Base")
	if p != nil {
		offset += p.fastWriteField1(buf[offset:], binaryWriter)
		offset += p.fastWriteField2(buf[offset:], binaryWriter)
		offset += p.fastWriteField3(buf[offset:], binaryWriter)
		offset += p.fastWriteField4(buf[offset:], binaryWriter)
		offset += p.fastWriteField5(buf[offset:], binaryWriter)
		offset += p.fastWriteField6(buf[offset:], binaryWriter)
	}
	offset += bthrift.Binary.WriteFieldStop(buf[offset:])
	offset += bthrift.Binary.WriteStructEnd(buf[offset:])
	return offset
}

func (p *Base) BLength() int {
	l := 0
	l += bthrift.Binary.StructBeginLength("Base")
	if p != nil {
		l += p.field1Length()
		l += p.field2Length()
		l += p.field3Length()
		l += p.field4Length()
		l += p.field5Length()
		l += p.field6Length()
	}
	l += bthrift.Binary.FieldStopLength()
	l += bthrift.Binary.StructEndLength()
	return l
}

func (p *Base) fastWriteField1(buf []byte, binaryWriter bthrift.BinaryWriter) int {
	offset := 0
	offset += bthrift.Binary.WriteFieldBegin(buf[offset:], "LogID", thrift.STRING, 1)
	offset += bthrift.Binary.WriteStringNocopy(buf[offset:], binaryWriter, p.LogID)

	offset += bthrift.Binary.WriteFieldEnd(buf[offset:])
	return offset
}

func (p *Base) fastWriteField2(buf []byte, binaryWriter bthrift.BinaryWriter) int {
	offset := 0
	offset += bthrift.Binary.WriteFieldBegin(buf[offset:], "Caller", thrift.STRING, 2)
	offset += bthrift.Binary.WriteStringNocopy(buf[offset:], binaryWriter, p.Caller)

	offset += bthrift.Binary.WriteFieldEnd(buf[offset:])
	return offset
}

func (p *Base) fastWriteField3(buf []byte, binaryWriter bthrift.BinaryWriter) int {
	offset := 0
	offset += bthrift.Binary.WriteFieldBegin(buf[offset:], "Addr", thrift.STRING, 3)
	offset += bthrift.Binary.WriteStringNocopy(buf[offset:], binaryWriter, p.Addr)

	offset += bthrift.Binary.WriteFieldEnd(buf[offset:])
	return offset
}

func (p *Base) fastWriteField4(buf []byte, binaryWriter bthrift.BinaryWriter) int {
	offset := 0
	offset += bthrift.Binary.WriteFieldBegin(buf[offset:], "Client", thrift.STRING, 4)
	offset += bthrift.Binary.WriteStringNocopy(buf[offset:], binaryWriter, p.Client)

	offset += bthrift.Binary.WriteFieldEnd(buf[offset:])
	return offset
}

func (p *Base) fastWriteField5(buf []byte, binaryWriter bthrift.BinaryWriter) int {
	offset := 0
	if p.IsSetTrafficEnv() {
		offset += bthrift.Binary.WriteFieldBegin(buf[offset:], "TrafficEnv", thrift.STRUCT, 5)
		offset += p.TrafficEnv.FastWriteNocopy(buf[offset:], binaryWriter)
		offset += bthrift.Binary.WriteFieldEnd(buf[offset:])
	}
	return offset
}

func (p *Base) fastWriteField6(buf []byte, binaryWriter bthrift.BinaryWriter) int {
	offset := 0
	if p.IsSetExtra() {
		offset += bthrift.Binary.WriteFieldBegin(buf[offset:], "Extra", thrift.MAP, 6)
		mapBeginOffset := offset
		offset += bthrift.Binary.MapBeginLength(thrift.STRING, thrift.STRING, 0)
		var length int
		for k, v := range p.Extra {
			length++

			offset += bthrift.Binary.WriteStringNocopy(buf[offset:], binaryWriter, k)

			offset += bthrift.Binary.WriteStringNocopy(buf[offset:], binaryWriter, v)

		}
		bthrift.Binary.WriteMapBegin(buf[mapBeginOffset:], thrift.STRING, thrift.STRING, length)
		offset += bthrift.Binary.WriteMapEnd(buf[offset:])
		offset += bthrift.Binary.WriteFieldEnd(buf[offset:])
	}
	return offset
}

func (p *Base) field1Length() int {
	l := 0
	l += bthrift.Binary.FieldBeginLength("LogID", thrift.STRING, 1)
	l += bthrift.Binary.StringLengthNocopy(p.LogID)

	l += bthrift.Binary.FieldEndLength()
	return l
}

func (p *Base) field2Length() int {
	l := 0
	l += bthrift.Binary.FieldBeginLength("Caller", thrift.STRING, 2)
	l += bthrift.Binary.StringLengthNocopy(p.Caller)

	l += bthrift.Binary.FieldEndLength()
	return l
}

func (p *Base) field3Length() int {
	l := 0
	l += bthrift.Binary.FieldBeginLength("Addr", thrift.STRING, 3)
	l += bthrift.Binary.StringLengthNocopy(p.Addr)

	l += bthrift.Binary.FieldEndLength()
	return l
}

func (p *Base) field4Length() int {
	l := 0
	l += bthrift.Binary.FieldBeginLength("Client", thrift.STRING, 4)
	l += bthrift.Binary.StringLengthNocopy(p.Client)

	l += bthrift.Binary.FieldEndLength()
	return l
}

func (p *Base) field5Length() int {
	l := 0
	if p.IsSetTrafficEnv() {
		l += bthrift.Binary.FieldBeginLength("TrafficEnv", thrift.STRUCT, 5)
		l += p.TrafficEnv.BLength()
		l += bthrift.Binary.FieldEndLength()
	}
	return l
}

func (p *Base) field6Length() int {
	l := 0
	if p.IsSetExtra() {
		l += bthrift.Binary.FieldBeginLength("Extra", thrift.MAP, 6)
		l += bthrift.Binary.MapBeginLength(thrift.STRING, thrift.STRING, len(p.Extra))
		for k, v := range p.Extra {

			l += bthrift.Binary.StringLengthNocopy(k)

			l += bthrift.Binary.StringLengthNocopy(v)

		}
		l += bthrift.Binary.MapEndLength()
		l += bthrift.Binary.FieldEndLength()
	}
	return l
}

func (p *BaseResp) FastRead(buf []byte) (int, error) {
	var err error
	var offset int
	var l int
	var fieldTypeId thrift.TType
	var fieldId int16
	_, l, err = bthrift.Binary.ReadStructBegin(buf)
	offset += l
	if err != nil {
		goto ReadStructBeginError
	}

	for {
		_, fieldTypeId, fieldId, l, err = bthrift.Binary.ReadFieldBegin(buf[offset:])
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
				l, err = bthrift.Binary.Skip(buf[offset:], fieldTypeId)
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
				l, err = bthrift.Binary.Skip(buf[offset:], fieldTypeId)
				offset += l
				if err != nil {
					goto SkipFieldError
				}
			}
		case 3:
			if fieldTypeId == thrift.MAP {
				l, err = p.FastReadField3(buf[offset:])
				offset += l
				if err != nil {
					goto ReadFieldError
				}
			} else {
				l, err = bthrift.Binary.Skip(buf[offset:], fieldTypeId)
				offset += l
				if err != nil {
					goto SkipFieldError
				}
			}
		default:
			l, err = bthrift.Binary.Skip(buf[offset:], fieldTypeId)
			offset += l
			if err != nil {
				goto SkipFieldError
			}
		}

		l, err = bthrift.Binary.ReadFieldEnd(buf[offset:])
		offset += l
		if err != nil {
			goto ReadFieldEndError
		}
	}
	l, err = bthrift.Binary.ReadStructEnd(buf[offset:])
	offset += l
	if err != nil {
		goto ReadStructEndError
	}

	return offset, nil
ReadStructBeginError:
	return offset, thrift.PrependError(fmt.Sprintf("%T read struct begin error: ", p), err)
ReadFieldBeginError:
	return offset, thrift.PrependError(fmt.Sprintf("%T read field %d begin error: ", p, fieldId), err)
ReadFieldError:
	return offset, thrift.PrependError(fmt.Sprintf("%T read field %d '%s' error: ", p, fieldId, fieldIDToName_BaseResp[fieldId]), err)
SkipFieldError:
	return offset, thrift.PrependError(fmt.Sprintf("%T field %d skip type %d error: ", p, fieldId, fieldTypeId), err)
ReadFieldEndError:
	return offset, thrift.PrependError(fmt.Sprintf("%T read field end error", p), err)
ReadStructEndError:
	return offset, thrift.PrependError(fmt.Sprintf("%T read struct end error: ", p), err)
}

func (p *BaseResp) FastReadField1(buf []byte) (int, error) {
	offset := 0

	if v, l, err := bthrift.Binary.ReadString(buf[offset:]); err != nil {
		return offset, err
	} else {
		offset += l

		p.StatusMessage = v

	}
	return offset, nil
}

func (p *BaseResp) FastReadField2(buf []byte) (int, error) {
	offset := 0

	if v, l, err := bthrift.Binary.ReadI32(buf[offset:]); err != nil {
		return offset, err
	} else {
		offset += l

		p.StatusCode = v

	}
	return offset, nil
}

func (p *BaseResp) FastReadField3(buf []byte) (int, error) {
	offset := 0

	_, _, size, l, err := bthrift.Binary.ReadMapBegin(buf[offset:])
	offset += l
	if err != nil {
		return offset, err
	}
	p.Extra = make(map[string]string, size)
	for i := 0; i < size; i++ {
		var _key string
		if v, l, err := bthrift.Binary.ReadString(buf[offset:]); err != nil {
			return offset, err
		} else {
			offset += l

			_key = v

		}

		var _val string
		if v, l, err := bthrift.Binary.ReadString(buf[offset:]); err != nil {
			return offset, err
		} else {
			offset += l

			_val = v

		}

		p.Extra[_key] = _val
	}
	if l, err := bthrift.Binary.ReadMapEnd(buf[offset:]); err != nil {
		return offset, err
	} else {
		offset += l
	}
	return offset, nil
}

// for compatibility
func (p *BaseResp) FastWrite(buf []byte) int {
	return 0
}

func (p *BaseResp) FastWriteNocopy(buf []byte, binaryWriter bthrift.BinaryWriter) int {
	offset := 0
	offset += bthrift.Binary.WriteStructBegin(buf[offset:], "BaseResp")
	if p != nil {
		offset += p.fastWriteField2(buf[offset:], binaryWriter)
		offset += p.fastWriteField1(buf[offset:], binaryWriter)
		offset += p.fastWriteField3(buf[offset:], binaryWriter)
	}
	offset += bthrift.Binary.WriteFieldStop(buf[offset:])
	offset += bthrift.Binary.WriteStructEnd(buf[offset:])
	return offset
}

func (p *BaseResp) BLength() int {
	l := 0
	l += bthrift.Binary.StructBeginLength("BaseResp")
	if p != nil {
		l += p.field1Length()
		l += p.field2Length()
		l += p.field3Length()
	}
	l += bthrift.Binary.FieldStopLength()
	l += bthrift.Binary.StructEndLength()
	return l
}

func (p *BaseResp) fastWriteField1(buf []byte, binaryWriter bthrift.BinaryWriter) int {
	offset := 0
	offset += bthrift.Binary.WriteFieldBegin(buf[offset:], "StatusMessage", thrift.STRING, 1)
	offset += bthrift.Binary.WriteStringNocopy(buf[offset:], binaryWriter, p.StatusMessage)

	offset += bthrift.Binary.WriteFieldEnd(buf[offset:])
	return offset
}

func (p *BaseResp) fastWriteField2(buf []byte, binaryWriter bthrift.BinaryWriter) int {
	offset := 0
	offset += bthrift.Binary.WriteFieldBegin(buf[offset:], "StatusCode", thrift.I32, 2)
	offset += bthrift.Binary.WriteI32(buf[offset:], p.StatusCode)

	offset += bthrift.Binary.WriteFieldEnd(buf[offset:])
	return offset
}

func (p *BaseResp) fastWriteField3(buf []byte, binaryWriter bthrift.BinaryWriter) int {
	offset := 0
	if p.IsSetExtra() {
		offset += bthrift.Binary.WriteFieldBegin(buf[offset:], "Extra", thrift.MAP, 3)
		mapBeginOffset := offset
		offset += bthrift.Binary.MapBeginLength(thrift.STRING, thrift.STRING, 0)
		var length int
		for k, v := range p.Extra {
			length++

			offset += bthrift.Binary.WriteStringNocopy(buf[offset:], binaryWriter, k)

			offset += bthrift.Binary.WriteStringNocopy(buf[offset:], binaryWriter, v)

		}
		bthrift.Binary.WriteMapBegin(buf[mapBeginOffset:], thrift.STRING, thrift.STRING, length)
		offset += bthrift.Binary.WriteMapEnd(buf[offset:])
		offset += bthrift.Binary.WriteFieldEnd(buf[offset:])
	}
	return offset
}

func (p *BaseResp) field1Length() int {
	l := 0
	l += bthrift.Binary.FieldBeginLength("StatusMessage", thrift.STRING, 1)
	l += bthrift.Binary.StringLengthNocopy(p.StatusMessage)

	l += bthrift.Binary.FieldEndLength()
	return l
}

func (p *BaseResp) field2Length() int {
	l := 0
	l += bthrift.Binary.FieldBeginLength("StatusCode", thrift.I32, 2)
	l += bthrift.Binary.I32Length(p.StatusCode)

	l += bthrift.Binary.FieldEndLength()
	return l
}

func (p *BaseResp) field3Length() int {
	l := 0
	if p.IsSetExtra() {
		l += bthrift.Binary.FieldBeginLength("Extra", thrift.MAP, 3)
		l += bthrift.Binary.MapBeginLength(thrift.STRING, thrift.STRING, len(p.Extra))
		for k, v := range p.Extra {

			l += bthrift.Binary.StringLengthNocopy(k)

			l += bthrift.Binary.StringLengthNocopy(v)

		}
		l += bthrift.Binary.MapEndLength()
		l += bthrift.Binary.FieldEndLength()
	}
	return l
}

func (p *TopParam) FastRead(buf []byte) (int, error) {
	var err error
	var offset int
	var l int
	var fieldTypeId thrift.TType
	var fieldId int16
	var issetRequestId bool = false
	var issetAccountId bool = false
	var issetDestService bool = false
	var issetSourceService bool = false
	_, l, err = bthrift.Binary.ReadStructBegin(buf)
	offset += l
	if err != nil {
		goto ReadStructBeginError
	}

	for {
		_, fieldTypeId, fieldId, l, err = bthrift.Binary.ReadFieldBegin(buf[offset:])
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
				issetRequestId = true
			} else {
				l, err = bthrift.Binary.Skip(buf[offset:], fieldTypeId)
				offset += l
				if err != nil {
					goto SkipFieldError
				}
			}
		case 2:
			if fieldTypeId == thrift.I64 {
				l, err = p.FastReadField2(buf[offset:])
				offset += l
				if err != nil {
					goto ReadFieldError
				}
				issetAccountId = true
			} else {
				l, err = bthrift.Binary.Skip(buf[offset:], fieldTypeId)
				offset += l
				if err != nil {
					goto SkipFieldError
				}
			}
		case 3:
			if fieldTypeId == thrift.I64 {
				l, err = p.FastReadField3(buf[offset:])
				offset += l
				if err != nil {
					goto ReadFieldError
				}
			} else {
				l, err = bthrift.Binary.Skip(buf[offset:], fieldTypeId)
				offset += l
				if err != nil {
					goto SkipFieldError
				}
			}
		case 4:
			if fieldTypeId == thrift.I64 {
				l, err = p.FastReadField4(buf[offset:])
				offset += l
				if err != nil {
					goto ReadFieldError
				}
			} else {
				l, err = bthrift.Binary.Skip(buf[offset:], fieldTypeId)
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
				issetDestService = true
			} else {
				l, err = bthrift.Binary.Skip(buf[offset:], fieldTypeId)
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
				issetSourceService = true
			} else {
				l, err = bthrift.Binary.Skip(buf[offset:], fieldTypeId)
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
				l, err = bthrift.Binary.Skip(buf[offset:], fieldTypeId)
				offset += l
				if err != nil {
					goto SkipFieldError
				}
			}
		case 8:
			if fieldTypeId == thrift.STRING {
				l, err = p.FastReadField8(buf[offset:])
				offset += l
				if err != nil {
					goto ReadFieldError
				}
			} else {
				l, err = bthrift.Binary.Skip(buf[offset:], fieldTypeId)
				offset += l
				if err != nil {
					goto SkipFieldError
				}
			}
		case 9:
			if fieldTypeId == thrift.I32 {
				l, err = p.FastReadField9(buf[offset:])
				offset += l
				if err != nil {
					goto ReadFieldError
				}
			} else {
				l, err = bthrift.Binary.Skip(buf[offset:], fieldTypeId)
				offset += l
				if err != nil {
					goto SkipFieldError
				}
			}
		case 10:
			if fieldTypeId == thrift.STRING {
				l, err = p.FastReadField10(buf[offset:])
				offset += l
				if err != nil {
					goto ReadFieldError
				}
			} else {
				l, err = bthrift.Binary.Skip(buf[offset:], fieldTypeId)
				offset += l
				if err != nil {
					goto SkipFieldError
				}
			}
		case 11:
			if fieldTypeId == thrift.STRING {
				l, err = p.FastReadField11(buf[offset:])
				offset += l
				if err != nil {
					goto ReadFieldError
				}
			} else {
				l, err = bthrift.Binary.Skip(buf[offset:], fieldTypeId)
				offset += l
				if err != nil {
					goto SkipFieldError
				}
			}
		default:
			l, err = bthrift.Binary.Skip(buf[offset:], fieldTypeId)
			offset += l
			if err != nil {
				goto SkipFieldError
			}
		}

		l, err = bthrift.Binary.ReadFieldEnd(buf[offset:])
		offset += l
		if err != nil {
			goto ReadFieldEndError
		}
	}
	l, err = bthrift.Binary.ReadStructEnd(buf[offset:])
	offset += l
	if err != nil {
		goto ReadStructEndError
	}

	if !issetRequestId {
		fieldId = 1
		goto RequiredFieldNotSetError
	}

	if !issetAccountId {
		fieldId = 2
		goto RequiredFieldNotSetError
	}

	if !issetDestService {
		fieldId = 5
		goto RequiredFieldNotSetError
	}

	if !issetSourceService {
		fieldId = 6
		goto RequiredFieldNotSetError
	}
	return offset, nil
ReadStructBeginError:
	return offset, thrift.PrependError(fmt.Sprintf("%T read struct begin error: ", p), err)
ReadFieldBeginError:
	return offset, thrift.PrependError(fmt.Sprintf("%T read field %d begin error: ", p, fieldId), err)
ReadFieldError:
	return offset, thrift.PrependError(fmt.Sprintf("%T read field %d '%s' error: ", p, fieldId, fieldIDToName_TopParam[fieldId]), err)
SkipFieldError:
	return offset, thrift.PrependError(fmt.Sprintf("%T field %d skip type %d error: ", p, fieldId, fieldTypeId), err)
ReadFieldEndError:
	return offset, thrift.PrependError(fmt.Sprintf("%T read field end error", p), err)
ReadStructEndError:
	return offset, thrift.PrependError(fmt.Sprintf("%T read struct end error: ", p), err)
RequiredFieldNotSetError:
	return offset, thrift.NewTProtocolExceptionWithType(thrift.INVALID_DATA, fmt.Errorf("required field %s is not set", fieldIDToName_TopParam[fieldId]))
}

func (p *TopParam) FastReadField1(buf []byte) (int, error) {
	offset := 0

	if v, l, err := bthrift.Binary.ReadString(buf[offset:]); err != nil {
		return offset, err
	} else {
		offset += l

		p.RequestId = v

	}
	return offset, nil
}

func (p *TopParam) FastReadField2(buf []byte) (int, error) {
	offset := 0

	if v, l, err := bthrift.Binary.ReadI64(buf[offset:]); err != nil {
		return offset, err
	} else {
		offset += l

		p.AccountId = v

	}
	return offset, nil
}

func (p *TopParam) FastReadField3(buf []byte) (int, error) {
	offset := 0

	if v, l, err := bthrift.Binary.ReadI64(buf[offset:]); err != nil {
		return offset, err
	} else {
		offset += l
		p.UserId = &v

	}
	return offset, nil
}

func (p *TopParam) FastReadField4(buf []byte) (int, error) {
	offset := 0

	if v, l, err := bthrift.Binary.ReadI64(buf[offset:]); err != nil {
		return offset, err
	} else {
		offset += l
		p.RoleId = &v

	}
	return offset, nil
}

func (p *TopParam) FastReadField5(buf []byte) (int, error) {
	offset := 0

	if v, l, err := bthrift.Binary.ReadString(buf[offset:]); err != nil {
		return offset, err
	} else {
		offset += l

		p.DestService = v

	}
	return offset, nil
}

func (p *TopParam) FastReadField6(buf []byte) (int, error) {
	offset := 0

	if v, l, err := bthrift.Binary.ReadString(buf[offset:]); err != nil {
		return offset, err
	} else {
		offset += l

		p.SourceService = v

	}
	return offset, nil
}

func (p *TopParam) FastReadField7(buf []byte) (int, error) {
	offset := 0

	if v, l, err := bthrift.Binary.ReadString(buf[offset:]); err != nil {
		return offset, err
	} else {
		offset += l
		p.Region = &v

	}
	return offset, nil
}

func (p *TopParam) FastReadField8(buf []byte) (int, error) {
	offset := 0

	if v, l, err := bthrift.Binary.ReadString(buf[offset:]); err != nil {
		return offset, err
	} else {
		offset += l
		p.RealIp = &v

	}
	return offset, nil
}

func (p *TopParam) FastReadField9(buf []byte) (int, error) {
	offset := 0

	if v, l, err := bthrift.Binary.ReadI32(buf[offset:]); err != nil {
		return offset, err
	} else {
		offset += l
		p.IsInternal = &v

	}
	return offset, nil
}

func (p *TopParam) FastReadField10(buf []byte) (int, error) {
	offset := 0

	if v, l, err := bthrift.Binary.ReadString(buf[offset:]); err != nil {
		return offset, err
	} else {
		offset += l
		p.Psm = &v

	}
	return offset, nil
}

func (p *TopParam) FastReadField11(buf []byte) (int, error) {
	offset := 0

	if v, l, err := bthrift.Binary.ReadString(buf[offset:]); err != nil {
		return offset, err
	} else {
		offset += l
		p.Site = &v

	}
	return offset, nil
}

// for compatibility
func (p *TopParam) FastWrite(buf []byte) int {
	return 0
}

func (p *TopParam) FastWriteNocopy(buf []byte, binaryWriter bthrift.BinaryWriter) int {
	offset := 0
	offset += bthrift.Binary.WriteStructBegin(buf[offset:], "TopParam")
	if p != nil {
		offset += p.fastWriteField2(buf[offset:], binaryWriter)
		offset += p.fastWriteField3(buf[offset:], binaryWriter)
		offset += p.fastWriteField4(buf[offset:], binaryWriter)
		offset += p.fastWriteField9(buf[offset:], binaryWriter)
		offset += p.fastWriteField1(buf[offset:], binaryWriter)
		offset += p.fastWriteField5(buf[offset:], binaryWriter)
		offset += p.fastWriteField6(buf[offset:], binaryWriter)
		offset += p.fastWriteField7(buf[offset:], binaryWriter)
		offset += p.fastWriteField8(buf[offset:], binaryWriter)
		offset += p.fastWriteField10(buf[offset:], binaryWriter)
		offset += p.fastWriteField11(buf[offset:], binaryWriter)
	}
	offset += bthrift.Binary.WriteFieldStop(buf[offset:])
	offset += bthrift.Binary.WriteStructEnd(buf[offset:])
	return offset
}

func (p *TopParam) BLength() int {
	l := 0
	l += bthrift.Binary.StructBeginLength("TopParam")
	if p != nil {
		l += p.field1Length()
		l += p.field2Length()
		l += p.field3Length()
		l += p.field4Length()
		l += p.field5Length()
		l += p.field6Length()
		l += p.field7Length()
		l += p.field8Length()
		l += p.field9Length()
		l += p.field10Length()
		l += p.field11Length()
	}
	l += bthrift.Binary.FieldStopLength()
	l += bthrift.Binary.StructEndLength()
	return l
}

func (p *TopParam) fastWriteField1(buf []byte, binaryWriter bthrift.BinaryWriter) int {
	offset := 0
	offset += bthrift.Binary.WriteFieldBegin(buf[offset:], "RequestId", thrift.STRING, 1)
	offset += bthrift.Binary.WriteStringNocopy(buf[offset:], binaryWriter, p.RequestId)

	offset += bthrift.Binary.WriteFieldEnd(buf[offset:])
	return offset
}

func (p *TopParam) fastWriteField2(buf []byte, binaryWriter bthrift.BinaryWriter) int {
	offset := 0
	offset += bthrift.Binary.WriteFieldBegin(buf[offset:], "AccountId", thrift.I64, 2)
	offset += bthrift.Binary.WriteI64(buf[offset:], p.AccountId)

	offset += bthrift.Binary.WriteFieldEnd(buf[offset:])
	return offset
}

func (p *TopParam) fastWriteField3(buf []byte, binaryWriter bthrift.BinaryWriter) int {
	offset := 0
	if p.IsSetUserId() {
		offset += bthrift.Binary.WriteFieldBegin(buf[offset:], "UserId", thrift.I64, 3)
		offset += bthrift.Binary.WriteI64(buf[offset:], *p.UserId)

		offset += bthrift.Binary.WriteFieldEnd(buf[offset:])
	}
	return offset
}

func (p *TopParam) fastWriteField4(buf []byte, binaryWriter bthrift.BinaryWriter) int {
	offset := 0
	if p.IsSetRoleId() {
		offset += bthrift.Binary.WriteFieldBegin(buf[offset:], "RoleId", thrift.I64, 4)
		offset += bthrift.Binary.WriteI64(buf[offset:], *p.RoleId)

		offset += bthrift.Binary.WriteFieldEnd(buf[offset:])
	}
	return offset
}

func (p *TopParam) fastWriteField5(buf []byte, binaryWriter bthrift.BinaryWriter) int {
	offset := 0
	offset += bthrift.Binary.WriteFieldBegin(buf[offset:], "DestService", thrift.STRING, 5)
	offset += bthrift.Binary.WriteStringNocopy(buf[offset:], binaryWriter, p.DestService)

	offset += bthrift.Binary.WriteFieldEnd(buf[offset:])
	return offset
}

func (p *TopParam) fastWriteField6(buf []byte, binaryWriter bthrift.BinaryWriter) int {
	offset := 0
	offset += bthrift.Binary.WriteFieldBegin(buf[offset:], "SourceService", thrift.STRING, 6)
	offset += bthrift.Binary.WriteStringNocopy(buf[offset:], binaryWriter, p.SourceService)

	offset += bthrift.Binary.WriteFieldEnd(buf[offset:])
	return offset
}

func (p *TopParam) fastWriteField7(buf []byte, binaryWriter bthrift.BinaryWriter) int {
	offset := 0
	if p.IsSetRegion() {
		offset += bthrift.Binary.WriteFieldBegin(buf[offset:], "Region", thrift.STRING, 7)
		offset += bthrift.Binary.WriteStringNocopy(buf[offset:], binaryWriter, *p.Region)

		offset += bthrift.Binary.WriteFieldEnd(buf[offset:])
	}
	return offset
}

func (p *TopParam) fastWriteField8(buf []byte, binaryWriter bthrift.BinaryWriter) int {
	offset := 0
	if p.IsSetRealIp() {
		offset += bthrift.Binary.WriteFieldBegin(buf[offset:], "RealIp", thrift.STRING, 8)
		offset += bthrift.Binary.WriteStringNocopy(buf[offset:], binaryWriter, *p.RealIp)

		offset += bthrift.Binary.WriteFieldEnd(buf[offset:])
	}
	return offset
}

func (p *TopParam) fastWriteField9(buf []byte, binaryWriter bthrift.BinaryWriter) int {
	offset := 0
	if p.IsSetIsInternal() {
		offset += bthrift.Binary.WriteFieldBegin(buf[offset:], "IsInternal", thrift.I32, 9)
		offset += bthrift.Binary.WriteI32(buf[offset:], *p.IsInternal)

		offset += bthrift.Binary.WriteFieldEnd(buf[offset:])
	}
	return offset
}

func (p *TopParam) fastWriteField10(buf []byte, binaryWriter bthrift.BinaryWriter) int {
	offset := 0
	if p.IsSetPsm() {
		offset += bthrift.Binary.WriteFieldBegin(buf[offset:], "Psm", thrift.STRING, 10)
		offset += bthrift.Binary.WriteStringNocopy(buf[offset:], binaryWriter, *p.Psm)

		offset += bthrift.Binary.WriteFieldEnd(buf[offset:])
	}
	return offset
}

func (p *TopParam) fastWriteField11(buf []byte, binaryWriter bthrift.BinaryWriter) int {
	offset := 0
	if p.IsSetSite() {
		offset += bthrift.Binary.WriteFieldBegin(buf[offset:], "Site", thrift.STRING, 11)
		offset += bthrift.Binary.WriteStringNocopy(buf[offset:], binaryWriter, *p.Site)

		offset += bthrift.Binary.WriteFieldEnd(buf[offset:])
	}
	return offset
}

func (p *TopParam) field1Length() int {
	l := 0
	l += bthrift.Binary.FieldBeginLength("RequestId", thrift.STRING, 1)
	l += bthrift.Binary.StringLengthNocopy(p.RequestId)

	l += bthrift.Binary.FieldEndLength()
	return l
}

func (p *TopParam) field2Length() int {
	l := 0
	l += bthrift.Binary.FieldBeginLength("AccountId", thrift.I64, 2)
	l += bthrift.Binary.I64Length(p.AccountId)

	l += bthrift.Binary.FieldEndLength()
	return l
}

func (p *TopParam) field3Length() int {
	l := 0
	if p.IsSetUserId() {
		l += bthrift.Binary.FieldBeginLength("UserId", thrift.I64, 3)
		l += bthrift.Binary.I64Length(*p.UserId)

		l += bthrift.Binary.FieldEndLength()
	}
	return l
}

func (p *TopParam) field4Length() int {
	l := 0
	if p.IsSetRoleId() {
		l += bthrift.Binary.FieldBeginLength("RoleId", thrift.I64, 4)
		l += bthrift.Binary.I64Length(*p.RoleId)

		l += bthrift.Binary.FieldEndLength()
	}
	return l
}

func (p *TopParam) field5Length() int {
	l := 0
	l += bthrift.Binary.FieldBeginLength("DestService", thrift.STRING, 5)
	l += bthrift.Binary.StringLengthNocopy(p.DestService)

	l += bthrift.Binary.FieldEndLength()
	return l
}

func (p *TopParam) field6Length() int {
	l := 0
	l += bthrift.Binary.FieldBeginLength("SourceService", thrift.STRING, 6)
	l += bthrift.Binary.StringLengthNocopy(p.SourceService)

	l += bthrift.Binary.FieldEndLength()
	return l
}

func (p *TopParam) field7Length() int {
	l := 0
	if p.IsSetRegion() {
		l += bthrift.Binary.FieldBeginLength("Region", thrift.STRING, 7)
		l += bthrift.Binary.StringLengthNocopy(*p.Region)

		l += bthrift.Binary.FieldEndLength()
	}
	return l
}

func (p *TopParam) field8Length() int {
	l := 0
	if p.IsSetRealIp() {
		l += bthrift.Binary.FieldBeginLength("RealIp", thrift.STRING, 8)
		l += bthrift.Binary.StringLengthNocopy(*p.RealIp)

		l += bthrift.Binary.FieldEndLength()
	}
	return l
}

func (p *TopParam) field9Length() int {
	l := 0
	if p.IsSetIsInternal() {
		l += bthrift.Binary.FieldBeginLength("IsInternal", thrift.I32, 9)
		l += bthrift.Binary.I32Length(*p.IsInternal)

		l += bthrift.Binary.FieldEndLength()
	}
	return l
}

func (p *TopParam) field10Length() int {
	l := 0
	if p.IsSetPsm() {
		l += bthrift.Binary.FieldBeginLength("Psm", thrift.STRING, 10)
		l += bthrift.Binary.StringLengthNocopy(*p.Psm)

		l += bthrift.Binary.FieldEndLength()
	}
	return l
}

func (p *TopParam) field11Length() int {
	l := 0
	if p.IsSetSite() {
		l += bthrift.Binary.FieldBeginLength("Site", thrift.STRING, 11)
		l += bthrift.Binary.StringLengthNocopy(*p.Site)

		l += bthrift.Binary.FieldEndLength()
	}
	return l
}
