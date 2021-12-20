// Code generated by thriftgo (0.1.3). DO NOT EDIT.

package appserver

import (
	"fmt"
	"github.com/apache/thrift/lib/go/thrift"
	"github.com/volcengine/vkectl/pkg/model/app/kitex_gen/base"
	"strings"
)

type ForwardKubernetesApiRequest struct {
	ClusterId string              `thrift:"ClusterId,1,required" json:"ClusterId"`
	Path      string              `thrift:"Path,2,required" json:"Path"`
	Method    string              `thrift:"Method,3,required" validate:"required,oneof=GET POST PUT PATCH DELETE"`
	Body      *string             `thrift:"Body,4" json:"Body,omitempty"`
	Headers   map[string][]string `thrift:"Headers,5" json:"Headers,omitempty"`
	Top       *base.TopParam      `thrift:"Top,255,required" json:"Top"`
}

func NewForwardKubernetesApiRequest() *ForwardKubernetesApiRequest {
	return &ForwardKubernetesApiRequest{}
}

func (p *ForwardKubernetesApiRequest) GetClusterId() (v string) {
	return p.ClusterId
}

func (p *ForwardKubernetesApiRequest) GetPath() (v string) {
	return p.Path
}

func (p *ForwardKubernetesApiRequest) GetMethod() (v string) {
	return p.Method
}

var ForwardKubernetesApiRequest_Body_DEFAULT string

func (p *ForwardKubernetesApiRequest) GetBody() (v string) {
	if !p.IsSetBody() {
		return ForwardKubernetesApiRequest_Body_DEFAULT
	}
	return *p.Body
}

var ForwardKubernetesApiRequest_Headers_DEFAULT map[string][]string

func (p *ForwardKubernetesApiRequest) GetHeaders() (v map[string][]string) {
	if !p.IsSetHeaders() {
		return ForwardKubernetesApiRequest_Headers_DEFAULT
	}
	return p.Headers
}

var ForwardKubernetesApiRequest_Top_DEFAULT *base.TopParam

func (p *ForwardKubernetesApiRequest) GetTop() (v *base.TopParam) {
	if !p.IsSetTop() {
		return ForwardKubernetesApiRequest_Top_DEFAULT
	}
	return p.Top
}
func (p *ForwardKubernetesApiRequest) SetClusterId(val string) {
	p.ClusterId = val
}
func (p *ForwardKubernetesApiRequest) SetPath(val string) {
	p.Path = val
}
func (p *ForwardKubernetesApiRequest) SetMethod(val string) {
	p.Method = val
}
func (p *ForwardKubernetesApiRequest) SetBody(val *string) {
	p.Body = val
}
func (p *ForwardKubernetesApiRequest) SetHeaders(val map[string][]string) {
	p.Headers = val
}
func (p *ForwardKubernetesApiRequest) SetTop(val *base.TopParam) {
	p.Top = val
}

var fieldIDToName_ForwardKubernetesApiRequest = map[int16]string{
	1:   "ClusterId",
	2:   "Path",
	3:   "Method",
	4:   "Body",
	5:   "Headers",
	255: "Top",
}

func (p *ForwardKubernetesApiRequest) IsSetBody() bool {
	return p.Body != nil
}

func (p *ForwardKubernetesApiRequest) IsSetHeaders() bool {
	return p.Headers != nil
}

func (p *ForwardKubernetesApiRequest) IsSetTop() bool {
	return p.Top != nil
}

func (p *ForwardKubernetesApiRequest) Read(iprot thrift.TProtocol) (err error) {

	var fieldTypeId thrift.TType
	var fieldId int16
	var issetClusterId bool = false
	var issetPath bool = false
	var issetMethod bool = false
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
				issetClusterId = true
			} else {
				if err = iprot.Skip(fieldTypeId); err != nil {
					goto SkipFieldError
				}
			}
		case 2:
			if fieldTypeId == thrift.STRING {
				if err = p.ReadField2(iprot); err != nil {
					goto ReadFieldError
				}
				issetPath = true
			} else {
				if err = iprot.Skip(fieldTypeId); err != nil {
					goto SkipFieldError
				}
			}
		case 3:
			if fieldTypeId == thrift.STRING {
				if err = p.ReadField3(iprot); err != nil {
					goto ReadFieldError
				}
				issetMethod = true
			} else {
				if err = iprot.Skip(fieldTypeId); err != nil {
					goto SkipFieldError
				}
			}
		case 4:
			if fieldTypeId == thrift.STRING {
				if err = p.ReadField4(iprot); err != nil {
					goto ReadFieldError
				}
			} else {
				if err = iprot.Skip(fieldTypeId); err != nil {
					goto SkipFieldError
				}
			}
		case 5:
			if fieldTypeId == thrift.MAP {
				if err = p.ReadField5(iprot); err != nil {
					goto ReadFieldError
				}
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
				issetTop = true
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

	if !issetClusterId {
		fieldId = 1
		goto RequiredFieldNotSetError
	}

	if !issetPath {
		fieldId = 2
		goto RequiredFieldNotSetError
	}

	if !issetMethod {
		fieldId = 3
		goto RequiredFieldNotSetError
	}

	if !issetTop {
		fieldId = 255
		goto RequiredFieldNotSetError
	}
	return nil
ReadStructBeginError:
	return thrift.PrependError(fmt.Sprintf("%T read struct begin error: ", p), err)
ReadFieldBeginError:
	return thrift.PrependError(fmt.Sprintf("%T read field %d begin error: ", p, fieldId), err)
ReadFieldError:
	return thrift.PrependError(fmt.Sprintf("%T read field %d '%s' error: ", p, fieldId, fieldIDToName_ForwardKubernetesApiRequest[fieldId]), err)
SkipFieldError:
	return thrift.PrependError(fmt.Sprintf("%T field %d skip type %d error: ", p, fieldId, fieldTypeId), err)

ReadFieldEndError:
	return thrift.PrependError(fmt.Sprintf("%T read field end error", p), err)
ReadStructEndError:
	return thrift.PrependError(fmt.Sprintf("%T read struct end error: ", p), err)
RequiredFieldNotSetError:
	return thrift.NewTProtocolExceptionWithType(thrift.INVALID_DATA, fmt.Errorf("Required field %s is not set", fieldIDToName_ForwardKubernetesApiRequest[fieldId]))
}

func (p *ForwardKubernetesApiRequest) ReadField1(iprot thrift.TProtocol) error {
	if v, err := iprot.ReadString(); err != nil {
		return err
	} else {
		p.ClusterId = v
	}
	return nil
}

func (p *ForwardKubernetesApiRequest) ReadField2(iprot thrift.TProtocol) error {
	if v, err := iprot.ReadString(); err != nil {
		return err
	} else {
		p.Path = v
	}
	return nil
}

func (p *ForwardKubernetesApiRequest) ReadField3(iprot thrift.TProtocol) error {
	if v, err := iprot.ReadString(); err != nil {
		return err
	} else {
		p.Method = v
	}
	return nil
}

func (p *ForwardKubernetesApiRequest) ReadField4(iprot thrift.TProtocol) error {
	if v, err := iprot.ReadString(); err != nil {
		return err
	} else {
		p.Body = &v
	}
	return nil
}

func (p *ForwardKubernetesApiRequest) ReadField5(iprot thrift.TProtocol) error {
	_, _, size, err := iprot.ReadMapBegin()
	if err != nil {
		return err
	}
	p.Headers = make(map[string][]string, size)
	for i := 0; i < size; i++ {
		var _key string
		if v, err := iprot.ReadString(); err != nil {
			return err
		} else {
			_key = v
		}

		_, size, err := iprot.ReadListBegin()
		if err != nil {
			return err
		}
		_val := make([]string, 0, size)
		for i := 0; i < size; i++ {
			var _elem string
			if v, err := iprot.ReadString(); err != nil {
				return err
			} else {
				_elem = v
			}

			_val = append(_val, _elem)
		}
		if err := iprot.ReadListEnd(); err != nil {
			return err
		}

		p.Headers[_key] = _val
	}
	if err := iprot.ReadMapEnd(); err != nil {
		return err
	}
	return nil
}

func (p *ForwardKubernetesApiRequest) ReadField255(iprot thrift.TProtocol) error {
	p.Top = base.NewTopParam()
	if err := p.Top.Read(iprot); err != nil {
		return err
	}
	return nil
}

func (p *ForwardKubernetesApiRequest) Write(oprot thrift.TProtocol) (err error) {
	var fieldId int16
	if err = oprot.WriteStructBegin("ForwardKubernetesApiRequest"); err != nil {
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
		if err = p.writeField3(oprot); err != nil {
			fieldId = 3
			goto WriteFieldError
		}
		if err = p.writeField4(oprot); err != nil {
			fieldId = 4
			goto WriteFieldError
		}
		if err = p.writeField5(oprot); err != nil {
			fieldId = 5
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

func (p *ForwardKubernetesApiRequest) writeField1(oprot thrift.TProtocol) (err error) {
	if err = oprot.WriteFieldBegin("ClusterId", thrift.STRING, 1); err != nil {
		goto WriteFieldBeginError
	}
	if err := oprot.WriteString(p.ClusterId); err != nil {
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

func (p *ForwardKubernetesApiRequest) writeField2(oprot thrift.TProtocol) (err error) {
	if err = oprot.WriteFieldBegin("Path", thrift.STRING, 2); err != nil {
		goto WriteFieldBeginError
	}
	if err := oprot.WriteString(p.Path); err != nil {
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

func (p *ForwardKubernetesApiRequest) writeField3(oprot thrift.TProtocol) (err error) {
	if err = oprot.WriteFieldBegin("Method", thrift.STRING, 3); err != nil {
		goto WriteFieldBeginError
	}
	if err := oprot.WriteString(p.Method); err != nil {
		return err
	}
	if err = oprot.WriteFieldEnd(); err != nil {
		goto WriteFieldEndError
	}
	return nil
WriteFieldBeginError:
	return thrift.PrependError(fmt.Sprintf("%T write field 3 begin error: ", p), err)
WriteFieldEndError:
	return thrift.PrependError(fmt.Sprintf("%T write field 3 end error: ", p), err)
}

func (p *ForwardKubernetesApiRequest) writeField4(oprot thrift.TProtocol) (err error) {
	if p.IsSetBody() {
		if err = oprot.WriteFieldBegin("Body", thrift.STRING, 4); err != nil {
			goto WriteFieldBeginError
		}
		if err := oprot.WriteString(*p.Body); err != nil {
			return err
		}
		if err = oprot.WriteFieldEnd(); err != nil {
			goto WriteFieldEndError
		}
	}
	return nil
WriteFieldBeginError:
	return thrift.PrependError(fmt.Sprintf("%T write field 4 begin error: ", p), err)
WriteFieldEndError:
	return thrift.PrependError(fmt.Sprintf("%T write field 4 end error: ", p), err)
}

func (p *ForwardKubernetesApiRequest) writeField5(oprot thrift.TProtocol) (err error) {
	if p.IsSetHeaders() {
		if err = oprot.WriteFieldBegin("Headers", thrift.MAP, 5); err != nil {
			goto WriteFieldBeginError
		}
		if err := oprot.WriteMapBegin(thrift.STRING, thrift.LIST, len(p.Headers)); err != nil {
			return err
		}
		for k, v := range p.Headers {

			if err := oprot.WriteString(k); err != nil {
				return err
			}

			if err := oprot.WriteListBegin(thrift.STRING, len(v)); err != nil {
				return err
			}
			for _, v := range v {
				if err := oprot.WriteString(v); err != nil {
					return err
				}
			}
			if err := oprot.WriteListEnd(); err != nil {
				return err
			}
		}
		if err := oprot.WriteMapEnd(); err != nil {
			return err
		}
		if err = oprot.WriteFieldEnd(); err != nil {
			goto WriteFieldEndError
		}
	}
	return nil
WriteFieldBeginError:
	return thrift.PrependError(fmt.Sprintf("%T write field 5 begin error: ", p), err)
WriteFieldEndError:
	return thrift.PrependError(fmt.Sprintf("%T write field 5 end error: ", p), err)
}

func (p *ForwardKubernetesApiRequest) writeField255(oprot thrift.TProtocol) (err error) {
	if err = oprot.WriteFieldBegin("Top", thrift.STRUCT, 255); err != nil {
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
	return thrift.PrependError(fmt.Sprintf("%T write field 255 begin error: ", p), err)
WriteFieldEndError:
	return thrift.PrependError(fmt.Sprintf("%T write field 255 end error: ", p), err)
}

func (p *ForwardKubernetesApiRequest) String() string {
	if p == nil {
		return "<nil>"
	}
	return fmt.Sprintf("ForwardKubernetesApiRequest(%+v)", *p)
}

func (p *ForwardKubernetesApiRequest) DeepEqual(ano *ForwardKubernetesApiRequest) bool {
	if p == ano {
		return true
	} else if p == nil || ano == nil {
		return false
	}
	if !p.Field1DeepEqual(ano.ClusterId) {
		return false
	}
	if !p.Field2DeepEqual(ano.Path) {
		return false
	}
	if !p.Field3DeepEqual(ano.Method) {
		return false
	}
	if !p.Field4DeepEqual(ano.Body) {
		return false
	}
	if !p.Field5DeepEqual(ano.Headers) {
		return false
	}
	if !p.Field255DeepEqual(ano.Top) {
		return false
	}
	return true
}

func (p *ForwardKubernetesApiRequest) Field1DeepEqual(src string) bool {

	if strings.Compare(p.ClusterId, src) != 0 {
		return false
	}
	return true
}
func (p *ForwardKubernetesApiRequest) Field2DeepEqual(src string) bool {

	if strings.Compare(p.Path, src) != 0 {
		return false
	}
	return true
}
func (p *ForwardKubernetesApiRequest) Field3DeepEqual(src string) bool {

	if strings.Compare(p.Method, src) != 0 {
		return false
	}
	return true
}
func (p *ForwardKubernetesApiRequest) Field4DeepEqual(src *string) bool {

	if p.Body == src {
		return true
	} else if p.Body == nil || src == nil {
		return false
	}
	if strings.Compare(*p.Body, *src) != 0 {
		return false
	}
	return true
}
func (p *ForwardKubernetesApiRequest) Field5DeepEqual(src map[string][]string) bool {

	if len(p.Headers) != len(src) {
		return false
	}
	for k, v := range p.Headers {
		_src := src[k]
		if len(v) != len(_src) {
			return false
		}
		for i, v := range v {
			_src1 := _src[i]
			if strings.Compare(v, _src1) != 0 {
				return false
			}
		}
	}
	return true
}
func (p *ForwardKubernetesApiRequest) Field255DeepEqual(src *base.TopParam) bool {

	if !p.Top.DeepEqual(src) {
		return false
	}
	return true
}

type ForwardKubernetesApiResponse struct {
	Code int32  `thrift:"Code,1,required" json:"Code"`
	Body string `thrift:"Body,2,required" json:"Body"`
}

func NewForwardKubernetesApiResponse() *ForwardKubernetesApiResponse {
	return &ForwardKubernetesApiResponse{}
}

func (p *ForwardKubernetesApiResponse) GetCode() (v int32) {
	return p.Code
}

func (p *ForwardKubernetesApiResponse) GetBody() (v string) {
	return p.Body
}
func (p *ForwardKubernetesApiResponse) SetCode(val int32) {
	p.Code = val
}
func (p *ForwardKubernetesApiResponse) SetBody(val string) {
	p.Body = val
}

var fieldIDToName_ForwardKubernetesApiResponse = map[int16]string{
	1: "Code",
	2: "Body",
}

func (p *ForwardKubernetesApiResponse) Read(iprot thrift.TProtocol) (err error) {

	var fieldTypeId thrift.TType
	var fieldId int16
	var issetCode bool = false
	var issetBody bool = false

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
			if fieldTypeId == thrift.I32 {
				if err = p.ReadField1(iprot); err != nil {
					goto ReadFieldError
				}
				issetCode = true
			} else {
				if err = iprot.Skip(fieldTypeId); err != nil {
					goto SkipFieldError
				}
			}
		case 2:
			if fieldTypeId == thrift.STRING {
				if err = p.ReadField2(iprot); err != nil {
					goto ReadFieldError
				}
				issetBody = true
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

	if !issetCode {
		fieldId = 1
		goto RequiredFieldNotSetError
	}

	if !issetBody {
		fieldId = 2
		goto RequiredFieldNotSetError
	}
	return nil
ReadStructBeginError:
	return thrift.PrependError(fmt.Sprintf("%T read struct begin error: ", p), err)
ReadFieldBeginError:
	return thrift.PrependError(fmt.Sprintf("%T read field %d begin error: ", p, fieldId), err)
ReadFieldError:
	return thrift.PrependError(fmt.Sprintf("%T read field %d '%s' error: ", p, fieldId, fieldIDToName_ForwardKubernetesApiResponse[fieldId]), err)
SkipFieldError:
	return thrift.PrependError(fmt.Sprintf("%T field %d skip type %d error: ", p, fieldId, fieldTypeId), err)

ReadFieldEndError:
	return thrift.PrependError(fmt.Sprintf("%T read field end error", p), err)
ReadStructEndError:
	return thrift.PrependError(fmt.Sprintf("%T read struct end error: ", p), err)
RequiredFieldNotSetError:
	return thrift.NewTProtocolExceptionWithType(thrift.INVALID_DATA, fmt.Errorf("Required field %s is not set", fieldIDToName_ForwardKubernetesApiResponse[fieldId]))
}

func (p *ForwardKubernetesApiResponse) ReadField1(iprot thrift.TProtocol) error {
	if v, err := iprot.ReadI32(); err != nil {
		return err
	} else {
		p.Code = v
	}
	return nil
}

func (p *ForwardKubernetesApiResponse) ReadField2(iprot thrift.TProtocol) error {
	if v, err := iprot.ReadString(); err != nil {
		return err
	} else {
		p.Body = v
	}
	return nil
}

func (p *ForwardKubernetesApiResponse) Write(oprot thrift.TProtocol) (err error) {
	var fieldId int16
	if err = oprot.WriteStructBegin("ForwardKubernetesApiResponse"); err != nil {
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

func (p *ForwardKubernetesApiResponse) writeField1(oprot thrift.TProtocol) (err error) {
	if err = oprot.WriteFieldBegin("Code", thrift.I32, 1); err != nil {
		goto WriteFieldBeginError
	}
	if err := oprot.WriteI32(p.Code); err != nil {
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

func (p *ForwardKubernetesApiResponse) writeField2(oprot thrift.TProtocol) (err error) {
	if err = oprot.WriteFieldBegin("Body", thrift.STRING, 2); err != nil {
		goto WriteFieldBeginError
	}
	if err := oprot.WriteString(p.Body); err != nil {
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

func (p *ForwardKubernetesApiResponse) String() string {
	if p == nil {
		return "<nil>"
	}
	return fmt.Sprintf("ForwardKubernetesApiResponse(%+v)", *p)
}

func (p *ForwardKubernetesApiResponse) DeepEqual(ano *ForwardKubernetesApiResponse) bool {
	if p == ano {
		return true
	} else if p == nil || ano == nil {
		return false
	}
	if !p.Field1DeepEqual(ano.Code) {
		return false
	}
	if !p.Field2DeepEqual(ano.Body) {
		return false
	}
	return true
}

func (p *ForwardKubernetesApiResponse) Field1DeepEqual(src int32) bool {

	if p.Code != src {
		return false
	}
	return true
}
func (p *ForwardKubernetesApiResponse) Field2DeepEqual(src string) bool {

	if strings.Compare(p.Body, src) != 0 {
		return false
	}
	return true
}