// Code generated by thriftgo (0.1.3). DO NOT EDIT.

package quota

import (
	"fmt"
	"github.com/apache/thrift/lib/go/thrift"
	"github.com/volcengine/vkectl/pkg/model/resource/kitex_gen/base"
	"strings"
)

type GetQuotaRequest struct {
	QuotaCode string         `thrift:"QuotaCode,1,required" json:"QuotaCode"`
	Top       *base.TopParam `thrift:"Top,254,required" json:"Top"`
	Base      *base.Base     `thrift:"Base,255" json:"Base,omitempty"`
}

func NewGetQuotaRequest() *GetQuotaRequest {
	return &GetQuotaRequest{}
}

func (p *GetQuotaRequest) GetQuotaCode() (v string) {
	return p.QuotaCode
}

var GetQuotaRequest_Top_DEFAULT *base.TopParam

func (p *GetQuotaRequest) GetTop() (v *base.TopParam) {
	if !p.IsSetTop() {
		return GetQuotaRequest_Top_DEFAULT
	}
	return p.Top
}

var GetQuotaRequest_Base_DEFAULT *base.Base

func (p *GetQuotaRequest) GetBase() (v *base.Base) {
	if !p.IsSetBase() {
		return GetQuotaRequest_Base_DEFAULT
	}
	return p.Base
}
func (p *GetQuotaRequest) SetQuotaCode(val string) {
	p.QuotaCode = val
}
func (p *GetQuotaRequest) SetTop(val *base.TopParam) {
	p.Top = val
}
func (p *GetQuotaRequest) SetBase(val *base.Base) {
	p.Base = val
}

var fieldIDToName_GetQuotaRequest = map[int16]string{
	1:   "QuotaCode",
	254: "Top",
	255: "Base",
}

func (p *GetQuotaRequest) IsSetTop() bool {
	return p.Top != nil
}

func (p *GetQuotaRequest) IsSetBase() bool {
	return p.Base != nil
}

func (p *GetQuotaRequest) Read(iprot thrift.TProtocol) (err error) {

	var fieldTypeId thrift.TType
	var fieldId int16
	var issetQuotaCode bool = false
	var issetTop bool = false

	if _, err = iprot.ReadStructBegin(); err != nil {
		goto ReadStructBeginError
	}

	for {
		_, fieldTypeId, fieldId, err = iprot.ReadFieldBegin()
		if err != nil {
			goto ReadFieldBeginError
		}
		if fieldTypeId == thrift.STOP {
			break
		}

		switch fieldId {
		case 1:
			if fieldTypeId == thrift.STRING {
				if err = p.ReadField1(iprot); err != nil {
					goto ReadFieldError
				}
				issetQuotaCode = true
			} else {
				if err = iprot.Skip(fieldTypeId); err != nil {
					goto SkipFieldError
				}
			}
		case 254:
			if fieldTypeId == thrift.STRUCT {
				if err = p.ReadField254(iprot); err != nil {
					goto ReadFieldError
				}
				issetTop = true
			} else {
				if err = iprot.Skip(fieldTypeId); err != nil {
					goto SkipFieldError
				}
			}
		case 255:
			if fieldTypeId == thrift.STRUCT {
				if err = p.ReadField255(iprot); err != nil {
					goto ReadFieldError
				}
			} else {
				if err = iprot.Skip(fieldTypeId); err != nil {
					goto SkipFieldError
				}
			}
		default:
			if err = iprot.Skip(fieldTypeId); err != nil {
				goto SkipFieldError
			}
		}

		if err = iprot.ReadFieldEnd(); err != nil {
			goto ReadFieldEndError
		}
	}
	if err = iprot.ReadStructEnd(); err != nil {
		goto ReadStructEndError
	}

	if !issetQuotaCode {
		fieldId = 1
		goto RequiredFieldNotSetError
	}

	if !issetTop {
		fieldId = 254
		goto RequiredFieldNotSetError
	}
	return nil
ReadStructBeginError:
	return thrift.PrependError(fmt.Sprintf("%T read struct begin error: ", p), err)
ReadFieldBeginError:
	return thrift.PrependError(fmt.Sprintf("%T read field %d begin error: ", p, fieldId), err)
ReadFieldError:
	return thrift.PrependError(fmt.Sprintf("%T read field %d '%s' error: ", p, fieldId, fieldIDToName_GetQuotaRequest[fieldId]), err)
SkipFieldError:
	return thrift.PrependError(fmt.Sprintf("%T field %d skip type %d error: ", p, fieldId, fieldTypeId), err)

ReadFieldEndError:
	return thrift.PrependError(fmt.Sprintf("%T read field end error", p), err)
ReadStructEndError:
	return thrift.PrependError(fmt.Sprintf("%T read struct end error: ", p), err)
RequiredFieldNotSetError:
	return thrift.NewTProtocolExceptionWithType(thrift.INVALID_DATA, fmt.Errorf("Required field %s is not set", fieldIDToName_GetQuotaRequest[fieldId]))
}

func (p *GetQuotaRequest) ReadField1(iprot thrift.TProtocol) error {
	if v, err := iprot.ReadString(); err != nil {
		return err
	} else {
		p.QuotaCode = v
	}
	return nil
}

func (p *GetQuotaRequest) ReadField254(iprot thrift.TProtocol) error {
	p.Top = base.NewTopParam()
	if err := p.Top.Read(iprot); err != nil {
		return err
	}
	return nil
}

func (p *GetQuotaRequest) ReadField255(iprot thrift.TProtocol) error {
	p.Base = base.NewBase()
	if err := p.Base.Read(iprot); err != nil {
		return err
	}
	return nil
}

func (p *GetQuotaRequest) Write(oprot thrift.TProtocol) (err error) {
	var fieldId int16
	if err = oprot.WriteStructBegin("GetQuotaRequest"); err != nil {
		goto WriteStructBeginError
	}
	if p != nil {
		if err = p.writeField1(oprot); err != nil {
			fieldId = 1
			goto WriteFieldError
		}
		if err = p.writeField254(oprot); err != nil {
			fieldId = 254
			goto WriteFieldError
		}
		if err = p.writeField255(oprot); err != nil {
			fieldId = 255
			goto WriteFieldError
		}

	}
	if err = oprot.WriteFieldStop(); err != nil {
		goto WriteFieldStopError
	}
	if err = oprot.WriteStructEnd(); err != nil {
		goto WriteStructEndError
	}
	return nil
WriteStructBeginError:
	return thrift.PrependError(fmt.Sprintf("%T write struct begin error: ", p), err)
WriteFieldError:
	return thrift.PrependError(fmt.Sprintf("%T write field %d error: ", p, fieldId), err)
WriteFieldStopError:
	return thrift.PrependError(fmt.Sprintf("%T write field stop error: ", p), err)
WriteStructEndError:
	return thrift.PrependError(fmt.Sprintf("%T write struct end error: ", p), err)
}

func (p *GetQuotaRequest) writeField1(oprot thrift.TProtocol) (err error) {
	if err = oprot.WriteFieldBegin("QuotaCode", thrift.STRING, 1); err != nil {
		goto WriteFieldBeginError
	}
	if err := oprot.WriteString(p.QuotaCode); err != nil {
		return err
	}
	if err = oprot.WriteFieldEnd(); err != nil {
		goto WriteFieldEndError
	}
	return nil
WriteFieldBeginError:
	return thrift.PrependError(fmt.Sprintf("%T write field 1 begin error: ", p), err)
WriteFieldEndError:
	return thrift.PrependError(fmt.Sprintf("%T write field 1 end error: ", p), err)
}

func (p *GetQuotaRequest) writeField254(oprot thrift.TProtocol) (err error) {
	if err = oprot.WriteFieldBegin("Top", thrift.STRUCT, 254); err != nil {
		goto WriteFieldBeginError
	}
	if err := p.Top.Write(oprot); err != nil {
		return err
	}
	if err = oprot.WriteFieldEnd(); err != nil {
		goto WriteFieldEndError
	}
	return nil
WriteFieldBeginError:
	return thrift.PrependError(fmt.Sprintf("%T write field 254 begin error: ", p), err)
WriteFieldEndError:
	return thrift.PrependError(fmt.Sprintf("%T write field 254 end error: ", p), err)
}

func (p *GetQuotaRequest) writeField255(oprot thrift.TProtocol) (err error) {
	if p.IsSetBase() {
		if err = oprot.WriteFieldBegin("Base", thrift.STRUCT, 255); err != nil {
			goto WriteFieldBeginError
		}
		if err := p.Base.Write(oprot); err != nil {
			return err
		}
		if err = oprot.WriteFieldEnd(); err != nil {
			goto WriteFieldEndError
		}
	}
	return nil
WriteFieldBeginError:
	return thrift.PrependError(fmt.Sprintf("%T write field 255 begin error: ", p), err)
WriteFieldEndError:
	return thrift.PrependError(fmt.Sprintf("%T write field 255 end error: ", p), err)
}

func (p *GetQuotaRequest) String() string {
	if p == nil {
		return "<nil>"
	}
	return fmt.Sprintf("GetQuotaRequest(%+v)", *p)
}

func (p *GetQuotaRequest) DeepEqual(ano *GetQuotaRequest) bool {
	if p == ano {
		return true
	} else if p == nil || ano == nil {
		return false
	}
	if !p.Field1DeepEqual(ano.QuotaCode) {
		return false
	}
	if !p.Field254DeepEqual(ano.Top) {
		return false
	}
	if !p.Field255DeepEqual(ano.Base) {
		return false
	}
	return true
}

func (p *GetQuotaRequest) Field1DeepEqual(src string) bool {

	if strings.Compare(p.QuotaCode, src) != 0 {
		return false
	}
	return true
}
func (p *GetQuotaRequest) Field254DeepEqual(src *base.TopParam) bool {

	if !p.Top.DeepEqual(src) {
		return false
	}
	return true
}
func (p *GetQuotaRequest) Field255DeepEqual(src *base.Base) bool {

	if !p.Base.DeepEqual(src) {
		return false
	}
	return true
}

type GetQuotaResponse struct {
	QuotaCode  string     `thrift:"QuotaCode,1,required" json:"QuotaCode"`
	QuotaValue float64    `thrift:"QuotaValue,2,required" json:"QuotaValue"`
	Base       *base.Base `thrift:"Base,255" json:"Base,omitempty"`
}

func NewGetQuotaResponse() *GetQuotaResponse {
	return &GetQuotaResponse{}
}

func (p *GetQuotaResponse) GetQuotaCode() (v string) {
	return p.QuotaCode
}

func (p *GetQuotaResponse) GetQuotaValue() (v float64) {
	return p.QuotaValue
}

var GetQuotaResponse_Base_DEFAULT *base.Base

func (p *GetQuotaResponse) GetBase() (v *base.Base) {
	if !p.IsSetBase() {
		return GetQuotaResponse_Base_DEFAULT
	}
	return p.Base
}
func (p *GetQuotaResponse) SetQuotaCode(val string) {
	p.QuotaCode = val
}
func (p *GetQuotaResponse) SetQuotaValue(val float64) {
	p.QuotaValue = val
}
func (p *GetQuotaResponse) SetBase(val *base.Base) {
	p.Base = val
}

var fieldIDToName_GetQuotaResponse = map[int16]string{
	1:   "QuotaCode",
	2:   "QuotaValue",
	255: "Base",
}

func (p *GetQuotaResponse) IsSetBase() bool {
	return p.Base != nil
}

func (p *GetQuotaResponse) Read(iprot thrift.TProtocol) (err error) {

	var fieldTypeId thrift.TType
	var fieldId int16
	var issetQuotaCode bool = false
	var issetQuotaValue bool = false

	if _, err = iprot.ReadStructBegin(); err != nil {
		goto ReadStructBeginError
	}

	for {
		_, fieldTypeId, fieldId, err = iprot.ReadFieldBegin()
		if err != nil {
			goto ReadFieldBeginError
		}
		if fieldTypeId == thrift.STOP {
			break
		}

		switch fieldId {
		case 1:
			if fieldTypeId == thrift.STRING {
				if err = p.ReadField1(iprot); err != nil {
					goto ReadFieldError
				}
				issetQuotaCode = true
			} else {
				if err = iprot.Skip(fieldTypeId); err != nil {
					goto SkipFieldError
				}
			}
		case 2:
			if fieldTypeId == thrift.DOUBLE {
				if err = p.ReadField2(iprot); err != nil {
					goto ReadFieldError
				}
				issetQuotaValue = true
			} else {
				if err = iprot.Skip(fieldTypeId); err != nil {
					goto SkipFieldError
				}
			}
		case 255:
			if fieldTypeId == thrift.STRUCT {
				if err = p.ReadField255(iprot); err != nil {
					goto ReadFieldError
				}
			} else {
				if err = iprot.Skip(fieldTypeId); err != nil {
					goto SkipFieldError
				}
			}
		default:
			if err = iprot.Skip(fieldTypeId); err != nil {
				goto SkipFieldError
			}
		}

		if err = iprot.ReadFieldEnd(); err != nil {
			goto ReadFieldEndError
		}
	}
	if err = iprot.ReadStructEnd(); err != nil {
		goto ReadStructEndError
	}

	if !issetQuotaCode {
		fieldId = 1
		goto RequiredFieldNotSetError
	}

	if !issetQuotaValue {
		fieldId = 2
		goto RequiredFieldNotSetError
	}
	return nil
ReadStructBeginError:
	return thrift.PrependError(fmt.Sprintf("%T read struct begin error: ", p), err)
ReadFieldBeginError:
	return thrift.PrependError(fmt.Sprintf("%T read field %d begin error: ", p, fieldId), err)
ReadFieldError:
	return thrift.PrependError(fmt.Sprintf("%T read field %d '%s' error: ", p, fieldId, fieldIDToName_GetQuotaResponse[fieldId]), err)
SkipFieldError:
	return thrift.PrependError(fmt.Sprintf("%T field %d skip type %d error: ", p, fieldId, fieldTypeId), err)

ReadFieldEndError:
	return thrift.PrependError(fmt.Sprintf("%T read field end error", p), err)
ReadStructEndError:
	return thrift.PrependError(fmt.Sprintf("%T read struct end error: ", p), err)
RequiredFieldNotSetError:
	return thrift.NewTProtocolExceptionWithType(thrift.INVALID_DATA, fmt.Errorf("Required field %s is not set", fieldIDToName_GetQuotaResponse[fieldId]))
}

func (p *GetQuotaResponse) ReadField1(iprot thrift.TProtocol) error {
	if v, err := iprot.ReadString(); err != nil {
		return err
	} else {
		p.QuotaCode = v
	}
	return nil
}

func (p *GetQuotaResponse) ReadField2(iprot thrift.TProtocol) error {
	if v, err := iprot.ReadDouble(); err != nil {
		return err
	} else {
		p.QuotaValue = v
	}
	return nil
}

func (p *GetQuotaResponse) ReadField255(iprot thrift.TProtocol) error {
	p.Base = base.NewBase()
	if err := p.Base.Read(iprot); err != nil {
		return err
	}
	return nil
}

func (p *GetQuotaResponse) Write(oprot thrift.TProtocol) (err error) {
	var fieldId int16
	if err = oprot.WriteStructBegin("GetQuotaResponse"); err != nil {
		goto WriteStructBeginError
	}
	if p != nil {
		if err = p.writeField1(oprot); err != nil {
			fieldId = 1
			goto WriteFieldError
		}
		if err = p.writeField2(oprot); err != nil {
			fieldId = 2
			goto WriteFieldError
		}
		if err = p.writeField255(oprot); err != nil {
			fieldId = 255
			goto WriteFieldError
		}

	}
	if err = oprot.WriteFieldStop(); err != nil {
		goto WriteFieldStopError
	}
	if err = oprot.WriteStructEnd(); err != nil {
		goto WriteStructEndError
	}
	return nil
WriteStructBeginError:
	return thrift.PrependError(fmt.Sprintf("%T write struct begin error: ", p), err)
WriteFieldError:
	return thrift.PrependError(fmt.Sprintf("%T write field %d error: ", p, fieldId), err)
WriteFieldStopError:
	return thrift.PrependError(fmt.Sprintf("%T write field stop error: ", p), err)
WriteStructEndError:
	return thrift.PrependError(fmt.Sprintf("%T write struct end error: ", p), err)
}

func (p *GetQuotaResponse) writeField1(oprot thrift.TProtocol) (err error) {
	if err = oprot.WriteFieldBegin("QuotaCode", thrift.STRING, 1); err != nil {
		goto WriteFieldBeginError
	}
	if err := oprot.WriteString(p.QuotaCode); err != nil {
		return err
	}
	if err = oprot.WriteFieldEnd(); err != nil {
		goto WriteFieldEndError
	}
	return nil
WriteFieldBeginError:
	return thrift.PrependError(fmt.Sprintf("%T write field 1 begin error: ", p), err)
WriteFieldEndError:
	return thrift.PrependError(fmt.Sprintf("%T write field 1 end error: ", p), err)
}

func (p *GetQuotaResponse) writeField2(oprot thrift.TProtocol) (err error) {
	if err = oprot.WriteFieldBegin("QuotaValue", thrift.DOUBLE, 2); err != nil {
		goto WriteFieldBeginError
	}
	if err := oprot.WriteDouble(p.QuotaValue); err != nil {
		return err
	}
	if err = oprot.WriteFieldEnd(); err != nil {
		goto WriteFieldEndError
	}
	return nil
WriteFieldBeginError:
	return thrift.PrependError(fmt.Sprintf("%T write field 2 begin error: ", p), err)
WriteFieldEndError:
	return thrift.PrependError(fmt.Sprintf("%T write field 2 end error: ", p), err)
}

func (p *GetQuotaResponse) writeField255(oprot thrift.TProtocol) (err error) {
	if p.IsSetBase() {
		if err = oprot.WriteFieldBegin("Base", thrift.STRUCT, 255); err != nil {
			goto WriteFieldBeginError
		}
		if err := p.Base.Write(oprot); err != nil {
			return err
		}
		if err = oprot.WriteFieldEnd(); err != nil {
			goto WriteFieldEndError
		}
	}
	return nil
WriteFieldBeginError:
	return thrift.PrependError(fmt.Sprintf("%T write field 255 begin error: ", p), err)
WriteFieldEndError:
	return thrift.PrependError(fmt.Sprintf("%T write field 255 end error: ", p), err)
}

func (p *GetQuotaResponse) String() string {
	if p == nil {
		return "<nil>"
	}
	return fmt.Sprintf("GetQuotaResponse(%+v)", *p)
}

func (p *GetQuotaResponse) DeepEqual(ano *GetQuotaResponse) bool {
	if p == ano {
		return true
	} else if p == nil || ano == nil {
		return false
	}
	if !p.Field1DeepEqual(ano.QuotaCode) {
		return false
	}
	if !p.Field2DeepEqual(ano.QuotaValue) {
		return false
	}
	if !p.Field255DeepEqual(ano.Base) {
		return false
	}
	return true
}

func (p *GetQuotaResponse) Field1DeepEqual(src string) bool {

	if strings.Compare(p.QuotaCode, src) != 0 {
		return false
	}
	return true
}
func (p *GetQuotaResponse) Field2DeepEqual(src float64) bool {

	if p.QuotaValue != src {
		return false
	}
	return true
}
func (p *GetQuotaResponse) Field255DeepEqual(src *base.Base) bool {

	if !p.Base.DeepEqual(src) {
		return false
	}
	return true
}

type ListQuotasRequest struct {
	Top  *base.TopParam `thrift:"Top,254,required" json:"Top"`
	Base *base.Base     `thrift:"Base,255" json:"Base,omitempty"`
}

func NewListQuotasRequest() *ListQuotasRequest {
	return &ListQuotasRequest{}
}

var ListQuotasRequest_Top_DEFAULT *base.TopParam

func (p *ListQuotasRequest) GetTop() (v *base.TopParam) {
	if !p.IsSetTop() {
		return ListQuotasRequest_Top_DEFAULT
	}
	return p.Top
}

var ListQuotasRequest_Base_DEFAULT *base.Base

func (p *ListQuotasRequest) GetBase() (v *base.Base) {
	if !p.IsSetBase() {
		return ListQuotasRequest_Base_DEFAULT
	}
	return p.Base
}
func (p *ListQuotasRequest) SetTop(val *base.TopParam) {
	p.Top = val
}
func (p *ListQuotasRequest) SetBase(val *base.Base) {
	p.Base = val
}

var fieldIDToName_ListQuotasRequest = map[int16]string{
	254: "Top",
	255: "Base",
}

func (p *ListQuotasRequest) IsSetTop() bool {
	return p.Top != nil
}

func (p *ListQuotasRequest) IsSetBase() bool {
	return p.Base != nil
}

func (p *ListQuotasRequest) Read(iprot thrift.TProtocol) (err error) {

	var fieldTypeId thrift.TType
	var fieldId int16
	var issetTop bool = false

	if _, err = iprot.ReadStructBegin(); err != nil {
		goto ReadStructBeginError
	}

	for {
		_, fieldTypeId, fieldId, err = iprot.ReadFieldBegin()
		if err != nil {
			goto ReadFieldBeginError
		}
		if fieldTypeId == thrift.STOP {
			break
		}

		switch fieldId {
		case 254:
			if fieldTypeId == thrift.STRUCT {
				if err = p.ReadField254(iprot); err != nil {
					goto ReadFieldError
				}
				issetTop = true
			} else {
				if err = iprot.Skip(fieldTypeId); err != nil {
					goto SkipFieldError
				}
			}
		case 255:
			if fieldTypeId == thrift.STRUCT {
				if err = p.ReadField255(iprot); err != nil {
					goto ReadFieldError
				}
			} else {
				if err = iprot.Skip(fieldTypeId); err != nil {
					goto SkipFieldError
				}
			}
		default:
			if err = iprot.Skip(fieldTypeId); err != nil {
				goto SkipFieldError
			}
		}

		if err = iprot.ReadFieldEnd(); err != nil {
			goto ReadFieldEndError
		}
	}
	if err = iprot.ReadStructEnd(); err != nil {
		goto ReadStructEndError
	}

	if !issetTop {
		fieldId = 254
		goto RequiredFieldNotSetError
	}
	return nil
ReadStructBeginError:
	return thrift.PrependError(fmt.Sprintf("%T read struct begin error: ", p), err)
ReadFieldBeginError:
	return thrift.PrependError(fmt.Sprintf("%T read field %d begin error: ", p, fieldId), err)
ReadFieldError:
	return thrift.PrependError(fmt.Sprintf("%T read field %d '%s' error: ", p, fieldId, fieldIDToName_ListQuotasRequest[fieldId]), err)
SkipFieldError:
	return thrift.PrependError(fmt.Sprintf("%T field %d skip type %d error: ", p, fieldId, fieldTypeId), err)

ReadFieldEndError:
	return thrift.PrependError(fmt.Sprintf("%T read field end error", p), err)
ReadStructEndError:
	return thrift.PrependError(fmt.Sprintf("%T read struct end error: ", p), err)
RequiredFieldNotSetError:
	return thrift.NewTProtocolExceptionWithType(thrift.INVALID_DATA, fmt.Errorf("Required field %s is not set", fieldIDToName_ListQuotasRequest[fieldId]))
}

func (p *ListQuotasRequest) ReadField254(iprot thrift.TProtocol) error {
	p.Top = base.NewTopParam()
	if err := p.Top.Read(iprot); err != nil {
		return err
	}
	return nil
}

func (p *ListQuotasRequest) ReadField255(iprot thrift.TProtocol) error {
	p.Base = base.NewBase()
	if err := p.Base.Read(iprot); err != nil {
		return err
	}
	return nil
}

func (p *ListQuotasRequest) Write(oprot thrift.TProtocol) (err error) {
	var fieldId int16
	if err = oprot.WriteStructBegin("ListQuotasRequest"); err != nil {
		goto WriteStructBeginError
	}
	if p != nil {
		if err = p.writeField254(oprot); err != nil {
			fieldId = 254
			goto WriteFieldError
		}
		if err = p.writeField255(oprot); err != nil {
			fieldId = 255
			goto WriteFieldError
		}

	}
	if err = oprot.WriteFieldStop(); err != nil {
		goto WriteFieldStopError
	}
	if err = oprot.WriteStructEnd(); err != nil {
		goto WriteStructEndError
	}
	return nil
WriteStructBeginError:
	return thrift.PrependError(fmt.Sprintf("%T write struct begin error: ", p), err)
WriteFieldError:
	return thrift.PrependError(fmt.Sprintf("%T write field %d error: ", p, fieldId), err)
WriteFieldStopError:
	return thrift.PrependError(fmt.Sprintf("%T write field stop error: ", p), err)
WriteStructEndError:
	return thrift.PrependError(fmt.Sprintf("%T write struct end error: ", p), err)
}

func (p *ListQuotasRequest) writeField254(oprot thrift.TProtocol) (err error) {
	if err = oprot.WriteFieldBegin("Top", thrift.STRUCT, 254); err != nil {
		goto WriteFieldBeginError
	}
	if err := p.Top.Write(oprot); err != nil {
		return err
	}
	if err = oprot.WriteFieldEnd(); err != nil {
		goto WriteFieldEndError
	}
	return nil
WriteFieldBeginError:
	return thrift.PrependError(fmt.Sprintf("%T write field 254 begin error: ", p), err)
WriteFieldEndError:
	return thrift.PrependError(fmt.Sprintf("%T write field 254 end error: ", p), err)
}

func (p *ListQuotasRequest) writeField255(oprot thrift.TProtocol) (err error) {
	if p.IsSetBase() {
		if err = oprot.WriteFieldBegin("Base", thrift.STRUCT, 255); err != nil {
			goto WriteFieldBeginError
		}
		if err := p.Base.Write(oprot); err != nil {
			return err
		}
		if err = oprot.WriteFieldEnd(); err != nil {
			goto WriteFieldEndError
		}
	}
	return nil
WriteFieldBeginError:
	return thrift.PrependError(fmt.Sprintf("%T write field 255 begin error: ", p), err)
WriteFieldEndError:
	return thrift.PrependError(fmt.Sprintf("%T write field 255 end error: ", p), err)
}

func (p *ListQuotasRequest) String() string {
	if p == nil {
		return "<nil>"
	}
	return fmt.Sprintf("ListQuotasRequest(%+v)", *p)
}

func (p *ListQuotasRequest) DeepEqual(ano *ListQuotasRequest) bool {
	if p == ano {
		return true
	} else if p == nil || ano == nil {
		return false
	}
	if !p.Field254DeepEqual(ano.Top) {
		return false
	}
	if !p.Field255DeepEqual(ano.Base) {
		return false
	}
	return true
}

func (p *ListQuotasRequest) Field254DeepEqual(src *base.TopParam) bool {

	if !p.Top.DeepEqual(src) {
		return false
	}
	return true
}
func (p *ListQuotasRequest) Field255DeepEqual(src *base.Base) bool {

	if !p.Base.DeepEqual(src) {
		return false
	}
	return true
}

type QuotaItem struct {
	QuotaCode  string  `thrift:"QuotaCode,1,required" json:"QuotaCode"`
	QuotaValue float64 `thrift:"QuotaValue,2,required" json:"QuotaValue"`
}

func NewQuotaItem() *QuotaItem {
	return &QuotaItem{}
}

func (p *QuotaItem) GetQuotaCode() (v string) {
	return p.QuotaCode
}

func (p *QuotaItem) GetQuotaValue() (v float64) {
	return p.QuotaValue
}
func (p *QuotaItem) SetQuotaCode(val string) {
	p.QuotaCode = val
}
func (p *QuotaItem) SetQuotaValue(val float64) {
	p.QuotaValue = val
}

var fieldIDToName_QuotaItem = map[int16]string{
	1: "QuotaCode",
	2: "QuotaValue",
}

func (p *QuotaItem) Read(iprot thrift.TProtocol) (err error) {

	var fieldTypeId thrift.TType
	var fieldId int16
	var issetQuotaCode bool = false
	var issetQuotaValue bool = false

	if _, err = iprot.ReadStructBegin(); err != nil {
		goto ReadStructBeginError
	}

	for {
		_, fieldTypeId, fieldId, err = iprot.ReadFieldBegin()
		if err != nil {
			goto ReadFieldBeginError
		}
		if fieldTypeId == thrift.STOP {
			break
		}

		switch fieldId {
		case 1:
			if fieldTypeId == thrift.STRING {
				if err = p.ReadField1(iprot); err != nil {
					goto ReadFieldError
				}
				issetQuotaCode = true
			} else {
				if err = iprot.Skip(fieldTypeId); err != nil {
					goto SkipFieldError
				}
			}
		case 2:
			if fieldTypeId == thrift.DOUBLE {
				if err = p.ReadField2(iprot); err != nil {
					goto ReadFieldError
				}
				issetQuotaValue = true
			} else {
				if err = iprot.Skip(fieldTypeId); err != nil {
					goto SkipFieldError
				}
			}
		default:
			if err = iprot.Skip(fieldTypeId); err != nil {
				goto SkipFieldError
			}
		}

		if err = iprot.ReadFieldEnd(); err != nil {
			goto ReadFieldEndError
		}
	}
	if err = iprot.ReadStructEnd(); err != nil {
		goto ReadStructEndError
	}

	if !issetQuotaCode {
		fieldId = 1
		goto RequiredFieldNotSetError
	}

	if !issetQuotaValue {
		fieldId = 2
		goto RequiredFieldNotSetError
	}
	return nil
ReadStructBeginError:
	return thrift.PrependError(fmt.Sprintf("%T read struct begin error: ", p), err)
ReadFieldBeginError:
	return thrift.PrependError(fmt.Sprintf("%T read field %d begin error: ", p, fieldId), err)
ReadFieldError:
	return thrift.PrependError(fmt.Sprintf("%T read field %d '%s' error: ", p, fieldId, fieldIDToName_QuotaItem[fieldId]), err)
SkipFieldError:
	return thrift.PrependError(fmt.Sprintf("%T field %d skip type %d error: ", p, fieldId, fieldTypeId), err)

ReadFieldEndError:
	return thrift.PrependError(fmt.Sprintf("%T read field end error", p), err)
ReadStructEndError:
	return thrift.PrependError(fmt.Sprintf("%T read struct end error: ", p), err)
RequiredFieldNotSetError:
	return thrift.NewTProtocolExceptionWithType(thrift.INVALID_DATA, fmt.Errorf("Required field %s is not set", fieldIDToName_QuotaItem[fieldId]))
}

func (p *QuotaItem) ReadField1(iprot thrift.TProtocol) error {
	if v, err := iprot.ReadString(); err != nil {
		return err
	} else {
		p.QuotaCode = v
	}
	return nil
}

func (p *QuotaItem) ReadField2(iprot thrift.TProtocol) error {
	if v, err := iprot.ReadDouble(); err != nil {
		return err
	} else {
		p.QuotaValue = v
	}
	return nil
}

func (p *QuotaItem) Write(oprot thrift.TProtocol) (err error) {
	var fieldId int16
	if err = oprot.WriteStructBegin("QuotaItem"); err != nil {
		goto WriteStructBeginError
	}
	if p != nil {
		if err = p.writeField1(oprot); err != nil {
			fieldId = 1
			goto WriteFieldError
		}
		if err = p.writeField2(oprot); err != nil {
			fieldId = 2
			goto WriteFieldError
		}

	}
	if err = oprot.WriteFieldStop(); err != nil {
		goto WriteFieldStopError
	}
	if err = oprot.WriteStructEnd(); err != nil {
		goto WriteStructEndError
	}
	return nil
WriteStructBeginError:
	return thrift.PrependError(fmt.Sprintf("%T write struct begin error: ", p), err)
WriteFieldError:
	return thrift.PrependError(fmt.Sprintf("%T write field %d error: ", p, fieldId), err)
WriteFieldStopError:
	return thrift.PrependError(fmt.Sprintf("%T write field stop error: ", p), err)
WriteStructEndError:
	return thrift.PrependError(fmt.Sprintf("%T write struct end error: ", p), err)
}

func (p *QuotaItem) writeField1(oprot thrift.TProtocol) (err error) {
	if err = oprot.WriteFieldBegin("QuotaCode", thrift.STRING, 1); err != nil {
		goto WriteFieldBeginError
	}
	if err := oprot.WriteString(p.QuotaCode); err != nil {
		return err
	}
	if err = oprot.WriteFieldEnd(); err != nil {
		goto WriteFieldEndError
	}
	return nil
WriteFieldBeginError:
	return thrift.PrependError(fmt.Sprintf("%T write field 1 begin error: ", p), err)
WriteFieldEndError:
	return thrift.PrependError(fmt.Sprintf("%T write field 1 end error: ", p), err)
}

func (p *QuotaItem) writeField2(oprot thrift.TProtocol) (err error) {
	if err = oprot.WriteFieldBegin("QuotaValue", thrift.DOUBLE, 2); err != nil {
		goto WriteFieldBeginError
	}
	if err := oprot.WriteDouble(p.QuotaValue); err != nil {
		return err
	}
	if err = oprot.WriteFieldEnd(); err != nil {
		goto WriteFieldEndError
	}
	return nil
WriteFieldBeginError:
	return thrift.PrependError(fmt.Sprintf("%T write field 2 begin error: ", p), err)
WriteFieldEndError:
	return thrift.PrependError(fmt.Sprintf("%T write field 2 end error: ", p), err)
}

func (p *QuotaItem) String() string {
	if p == nil {
		return "<nil>"
	}
	return fmt.Sprintf("QuotaItem(%+v)", *p)
}

func (p *QuotaItem) DeepEqual(ano *QuotaItem) bool {
	if p == ano {
		return true
	} else if p == nil || ano == nil {
		return false
	}
	if !p.Field1DeepEqual(ano.QuotaCode) {
		return false
	}
	if !p.Field2DeepEqual(ano.QuotaValue) {
		return false
	}
	return true
}

func (p *QuotaItem) Field1DeepEqual(src string) bool {

	if strings.Compare(p.QuotaCode, src) != 0 {
		return false
	}
	return true
}
func (p *QuotaItem) Field2DeepEqual(src float64) bool {

	if p.QuotaValue != src {
		return false
	}
	return true
}

type ListQuotasResponse struct {
	Items []*QuotaItem `thrift:"Items,1,required" json:"Items"`
	Base  *base.Base   `thrift:"Base,255" json:"Base,omitempty"`
}

func NewListQuotasResponse() *ListQuotasResponse {
	return &ListQuotasResponse{}
}

func (p *ListQuotasResponse) GetItems() (v []*QuotaItem) {
	return p.Items
}

var ListQuotasResponse_Base_DEFAULT *base.Base

func (p *ListQuotasResponse) GetBase() (v *base.Base) {
	if !p.IsSetBase() {
		return ListQuotasResponse_Base_DEFAULT
	}
	return p.Base
}
func (p *ListQuotasResponse) SetItems(val []*QuotaItem) {
	p.Items = val
}
func (p *ListQuotasResponse) SetBase(val *base.Base) {
	p.Base = val
}

var fieldIDToName_ListQuotasResponse = map[int16]string{
	1:   "Items",
	255: "Base",
}

func (p *ListQuotasResponse) IsSetBase() bool {
	return p.Base != nil
}

func (p *ListQuotasResponse) Read(iprot thrift.TProtocol) (err error) {

	var fieldTypeId thrift.TType
	var fieldId int16
	var issetItems bool = false

	if _, err = iprot.ReadStructBegin(); err != nil {
		goto ReadStructBeginError
	}

	for {
		_, fieldTypeId, fieldId, err = iprot.ReadFieldBegin()
		if err != nil {
			goto ReadFieldBeginError
		}
		if fieldTypeId == thrift.STOP {
			break
		}

		switch fieldId {
		case 1:
			if fieldTypeId == thrift.LIST {
				if err = p.ReadField1(iprot); err != nil {
					goto ReadFieldError
				}
				issetItems = true
			} else {
				if err = iprot.Skip(fieldTypeId); err != nil {
					goto SkipFieldError
				}
			}
		case 255:
			if fieldTypeId == thrift.STRUCT {
				if err = p.ReadField255(iprot); err != nil {
					goto ReadFieldError
				}
			} else {
				if err = iprot.Skip(fieldTypeId); err != nil {
					goto SkipFieldError
				}
			}
		default:
			if err = iprot.Skip(fieldTypeId); err != nil {
				goto SkipFieldError
			}
		}

		if err = iprot.ReadFieldEnd(); err != nil {
			goto ReadFieldEndError
		}
	}
	if err = iprot.ReadStructEnd(); err != nil {
		goto ReadStructEndError
	}

	if !issetItems {
		fieldId = 1
		goto RequiredFieldNotSetError
	}
	return nil
ReadStructBeginError:
	return thrift.PrependError(fmt.Sprintf("%T read struct begin error: ", p), err)
ReadFieldBeginError:
	return thrift.PrependError(fmt.Sprintf("%T read field %d begin error: ", p, fieldId), err)
ReadFieldError:
	return thrift.PrependError(fmt.Sprintf("%T read field %d '%s' error: ", p, fieldId, fieldIDToName_ListQuotasResponse[fieldId]), err)
SkipFieldError:
	return thrift.PrependError(fmt.Sprintf("%T field %d skip type %d error: ", p, fieldId, fieldTypeId), err)

ReadFieldEndError:
	return thrift.PrependError(fmt.Sprintf("%T read field end error", p), err)
ReadStructEndError:
	return thrift.PrependError(fmt.Sprintf("%T read struct end error: ", p), err)
RequiredFieldNotSetError:
	return thrift.NewTProtocolExceptionWithType(thrift.INVALID_DATA, fmt.Errorf("Required field %s is not set", fieldIDToName_ListQuotasResponse[fieldId]))
}

func (p *ListQuotasResponse) ReadField1(iprot thrift.TProtocol) error {
	_, size, err := iprot.ReadListBegin()
	if err != nil {
		return err
	}
	p.Items = make([]*QuotaItem, 0, size)
	for i := 0; i < size; i++ {
		_elem := NewQuotaItem()
		if err := _elem.Read(iprot); err != nil {
			return err
		}

		p.Items = append(p.Items, _elem)
	}
	if err := iprot.ReadListEnd(); err != nil {
		return err
	}
	return nil
}

func (p *ListQuotasResponse) ReadField255(iprot thrift.TProtocol) error {
	p.Base = base.NewBase()
	if err := p.Base.Read(iprot); err != nil {
		return err
	}
	return nil
}

func (p *ListQuotasResponse) Write(oprot thrift.TProtocol) (err error) {
	var fieldId int16
	if err = oprot.WriteStructBegin("ListQuotasResponse"); err != nil {
		goto WriteStructBeginError
	}
	if p != nil {
		if err = p.writeField1(oprot); err != nil {
			fieldId = 1
			goto WriteFieldError
		}
		if err = p.writeField255(oprot); err != nil {
			fieldId = 255
			goto WriteFieldError
		}

	}
	if err = oprot.WriteFieldStop(); err != nil {
		goto WriteFieldStopError
	}
	if err = oprot.WriteStructEnd(); err != nil {
		goto WriteStructEndError
	}
	return nil
WriteStructBeginError:
	return thrift.PrependError(fmt.Sprintf("%T write struct begin error: ", p), err)
WriteFieldError:
	return thrift.PrependError(fmt.Sprintf("%T write field %d error: ", p, fieldId), err)
WriteFieldStopError:
	return thrift.PrependError(fmt.Sprintf("%T write field stop error: ", p), err)
WriteStructEndError:
	return thrift.PrependError(fmt.Sprintf("%T write struct end error: ", p), err)
}

func (p *ListQuotasResponse) writeField1(oprot thrift.TProtocol) (err error) {
	if err = oprot.WriteFieldBegin("Items", thrift.LIST, 1); err != nil {
		goto WriteFieldBeginError
	}
	if err := oprot.WriteListBegin(thrift.STRUCT, len(p.Items)); err != nil {
		return err
	}
	for _, v := range p.Items {
		if err := v.Write(oprot); err != nil {
			return err
		}
	}
	if err := oprot.WriteListEnd(); err != nil {
		return err
	}
	if err = oprot.WriteFieldEnd(); err != nil {
		goto WriteFieldEndError
	}
	return nil
WriteFieldBeginError:
	return thrift.PrependError(fmt.Sprintf("%T write field 1 begin error: ", p), err)
WriteFieldEndError:
	return thrift.PrependError(fmt.Sprintf("%T write field 1 end error: ", p), err)
}

func (p *ListQuotasResponse) writeField255(oprot thrift.TProtocol) (err error) {
	if p.IsSetBase() {
		if err = oprot.WriteFieldBegin("Base", thrift.STRUCT, 255); err != nil {
			goto WriteFieldBeginError
		}
		if err := p.Base.Write(oprot); err != nil {
			return err
		}
		if err = oprot.WriteFieldEnd(); err != nil {
			goto WriteFieldEndError
		}
	}
	return nil
WriteFieldBeginError:
	return thrift.PrependError(fmt.Sprintf("%T write field 255 begin error: ", p), err)
WriteFieldEndError:
	return thrift.PrependError(fmt.Sprintf("%T write field 255 end error: ", p), err)
}

func (p *ListQuotasResponse) String() string {
	if p == nil {
		return "<nil>"
	}
	return fmt.Sprintf("ListQuotasResponse(%+v)", *p)
}

func (p *ListQuotasResponse) DeepEqual(ano *ListQuotasResponse) bool {
	if p == ano {
		return true
	} else if p == nil || ano == nil {
		return false
	}
	if !p.Field1DeepEqual(ano.Items) {
		return false
	}
	if !p.Field255DeepEqual(ano.Base) {
		return false
	}
	return true
}

func (p *ListQuotasResponse) Field1DeepEqual(src []*QuotaItem) bool {

	if len(p.Items) != len(src) {
		return false
	}
	for i, v := range p.Items {
		_src := src[i]
		if !v.DeepEqual(_src) {
			return false
		}
	}
	return true
}
func (p *ListQuotasResponse) Field255DeepEqual(src *base.Base) bool {

	if !p.Base.DeepEqual(src) {
		return false
	}
	return true
}
