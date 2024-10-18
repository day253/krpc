// Code generated by thriftgo (0.2.11). DO NOT EDIT.

package re

import (
	"context"
	"fmt"
	"github.com/apache/thrift/lib/go/thrift"
	"strings"
)

type Event struct {
}

func NewEvent() *Event {
	return &Event{}
}

func (p *Event) InitDefault() {
	*p = Event{}
}

var fieldIDToName_Event = map[int16]string{}

func (p *Event) Read(iprot thrift.TProtocol) (err error) {

	var fieldTypeId thrift.TType
	var fieldId int16

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
		if err = iprot.Skip(fieldTypeId); err != nil {
			goto SkipFieldTypeError
		}

		if err = iprot.ReadFieldEnd(); err != nil {
			goto ReadFieldEndError
		}
	}
	if err = iprot.ReadStructEnd(); err != nil {
		goto ReadStructEndError
	}

	return nil
ReadStructBeginError:
	return thrift.PrependError(fmt.Sprintf("%T read struct begin error: ", p), err)
ReadFieldBeginError:
	return thrift.PrependError(fmt.Sprintf("%T read field %d begin error: ", p, fieldId), err)
SkipFieldTypeError:
	return thrift.PrependError(fmt.Sprintf("%T skip field type %d error", p, fieldTypeId), err)

ReadFieldEndError:
	return thrift.PrependError(fmt.Sprintf("%T read field end error", p), err)
ReadStructEndError:
	return thrift.PrependError(fmt.Sprintf("%T read struct end error: ", p), err)
}

func (p *Event) Write(oprot thrift.TProtocol) (err error) {
	if err = oprot.WriteStructBegin("Event"); err != nil {
		goto WriteStructBeginError
	}
	if p != nil {

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
WriteFieldStopError:
	return thrift.PrependError(fmt.Sprintf("%T write field stop error: ", p), err)
WriteStructEndError:
	return thrift.PrependError(fmt.Sprintf("%T write struct end error: ", p), err)
}

func (p *Event) String() string {
	if p == nil {
		return "<nil>"
	}
	return fmt.Sprintf("Event(%+v)", *p)
}

func (p *Event) DeepEqual(ano *Event) bool {
	if p == ano {
		return true
	} else if p == nil || ano == nil {
		return false
	}
	return true
}

type EventPredictRequest struct {
	RequestId    *string `thrift:"requestId,1,optional" frugal:"1,optional,string" json:"requestId,omitempty"`
	Organization *string `thrift:"organization,2,optional" frugal:"2,optional,string" json:"organization,omitempty"`
	Event        *Event  `thrift:"event,3,optional" frugal:"3,optional,Event" json:"event,omitempty"`
}

func NewEventPredictRequest() *EventPredictRequest {
	return &EventPredictRequest{}
}

func (p *EventPredictRequest) InitDefault() {
	*p = EventPredictRequest{}
}

var EventPredictRequest_RequestId_DEFAULT string

func (p *EventPredictRequest) GetRequestId() (v string) {
	if !p.IsSetRequestId() {
		return EventPredictRequest_RequestId_DEFAULT
	}
	return *p.RequestId
}

var EventPredictRequest_Organization_DEFAULT string

func (p *EventPredictRequest) GetOrganization() (v string) {
	if !p.IsSetOrganization() {
		return EventPredictRequest_Organization_DEFAULT
	}
	return *p.Organization
}

var EventPredictRequest_Event_DEFAULT *Event

func (p *EventPredictRequest) GetEvent() (v *Event) {
	if !p.IsSetEvent() {
		return EventPredictRequest_Event_DEFAULT
	}
	return p.Event
}
func (p *EventPredictRequest) SetRequestId(val *string) {
	p.RequestId = val
}
func (p *EventPredictRequest) SetOrganization(val *string) {
	p.Organization = val
}
func (p *EventPredictRequest) SetEvent(val *Event) {
	p.Event = val
}

var fieldIDToName_EventPredictRequest = map[int16]string{
	1: "requestId",
	2: "organization",
	3: "event",
}

func (p *EventPredictRequest) IsSetRequestId() bool {
	return p.RequestId != nil
}

func (p *EventPredictRequest) IsSetOrganization() bool {
	return p.Organization != nil
}

func (p *EventPredictRequest) IsSetEvent() bool {
	return p.Event != nil
}

func (p *EventPredictRequest) Read(iprot thrift.TProtocol) (err error) {

	var fieldTypeId thrift.TType
	var fieldId int16

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
			} else {
				if err = iprot.Skip(fieldTypeId); err != nil {
					goto SkipFieldError
				}
			}
		case 3:
			if fieldTypeId == thrift.STRUCT {
				if err = p.ReadField3(iprot); err != nil {
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

	return nil
ReadStructBeginError:
	return thrift.PrependError(fmt.Sprintf("%T read struct begin error: ", p), err)
ReadFieldBeginError:
	return thrift.PrependError(fmt.Sprintf("%T read field %d begin error: ", p, fieldId), err)
ReadFieldError:
	return thrift.PrependError(fmt.Sprintf("%T read field %d '%s' error: ", p, fieldId, fieldIDToName_EventPredictRequest[fieldId]), err)
SkipFieldError:
	return thrift.PrependError(fmt.Sprintf("%T field %d skip type %d error: ", p, fieldId, fieldTypeId), err)

ReadFieldEndError:
	return thrift.PrependError(fmt.Sprintf("%T read field end error", p), err)
ReadStructEndError:
	return thrift.PrependError(fmt.Sprintf("%T read struct end error: ", p), err)
}

func (p *EventPredictRequest) ReadField1(iprot thrift.TProtocol) error {
	if v, err := iprot.ReadString(); err != nil {
		return err
	} else {
		p.RequestId = &v
	}
	return nil
}

func (p *EventPredictRequest) ReadField2(iprot thrift.TProtocol) error {
	if v, err := iprot.ReadString(); err != nil {
		return err
	} else {
		p.Organization = &v
	}
	return nil
}

func (p *EventPredictRequest) ReadField3(iprot thrift.TProtocol) error {
	p.Event = NewEvent()
	if err := p.Event.Read(iprot); err != nil {
		return err
	}
	return nil
}

func (p *EventPredictRequest) Write(oprot thrift.TProtocol) (err error) {
	var fieldId int16
	if err = oprot.WriteStructBegin("EventPredictRequest"); err != nil {
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

func (p *EventPredictRequest) writeField1(oprot thrift.TProtocol) (err error) {
	if p.IsSetRequestId() {
		if err = oprot.WriteFieldBegin("requestId", thrift.STRING, 1); err != nil {
			goto WriteFieldBeginError
		}
		if err := oprot.WriteString(*p.RequestId); err != nil {
			return err
		}
		if err = oprot.WriteFieldEnd(); err != nil {
			goto WriteFieldEndError
		}
	}
	return nil
WriteFieldBeginError:
	return thrift.PrependError(fmt.Sprintf("%T write field 1 begin error: ", p), err)
WriteFieldEndError:
	return thrift.PrependError(fmt.Sprintf("%T write field 1 end error: ", p), err)
}

func (p *EventPredictRequest) writeField2(oprot thrift.TProtocol) (err error) {
	if p.IsSetOrganization() {
		if err = oprot.WriteFieldBegin("organization", thrift.STRING, 2); err != nil {
			goto WriteFieldBeginError
		}
		if err := oprot.WriteString(*p.Organization); err != nil {
			return err
		}
		if err = oprot.WriteFieldEnd(); err != nil {
			goto WriteFieldEndError
		}
	}
	return nil
WriteFieldBeginError:
	return thrift.PrependError(fmt.Sprintf("%T write field 2 begin error: ", p), err)
WriteFieldEndError:
	return thrift.PrependError(fmt.Sprintf("%T write field 2 end error: ", p), err)
}

func (p *EventPredictRequest) writeField3(oprot thrift.TProtocol) (err error) {
	if p.IsSetEvent() {
		if err = oprot.WriteFieldBegin("event", thrift.STRUCT, 3); err != nil {
			goto WriteFieldBeginError
		}
		if err := p.Event.Write(oprot); err != nil {
			return err
		}
		if err = oprot.WriteFieldEnd(); err != nil {
			goto WriteFieldEndError
		}
	}
	return nil
WriteFieldBeginError:
	return thrift.PrependError(fmt.Sprintf("%T write field 3 begin error: ", p), err)
WriteFieldEndError:
	return thrift.PrependError(fmt.Sprintf("%T write field 3 end error: ", p), err)
}

func (p *EventPredictRequest) String() string {
	if p == nil {
		return "<nil>"
	}
	return fmt.Sprintf("EventPredictRequest(%+v)", *p)
}

func (p *EventPredictRequest) DeepEqual(ano *EventPredictRequest) bool {
	if p == ano {
		return true
	} else if p == nil || ano == nil {
		return false
	}
	if !p.Field1DeepEqual(ano.RequestId) {
		return false
	}
	if !p.Field2DeepEqual(ano.Organization) {
		return false
	}
	if !p.Field3DeepEqual(ano.Event) {
		return false
	}
	return true
}

func (p *EventPredictRequest) Field1DeepEqual(src *string) bool {

	if p.RequestId == src {
		return true
	} else if p.RequestId == nil || src == nil {
		return false
	}
	if strings.Compare(*p.RequestId, *src) != 0 {
		return false
	}
	return true
}
func (p *EventPredictRequest) Field2DeepEqual(src *string) bool {

	if p.Organization == src {
		return true
	} else if p.Organization == nil || src == nil {
		return false
	}
	if strings.Compare(*p.Organization, *src) != 0 {
		return false
	}
	return true
}
func (p *EventPredictRequest) Field3DeepEqual(src *Event) bool {

	if !p.Event.DeepEqual(src) {
		return false
	}
	return true
}

type EventPredictResult_ struct {
	Result_ *string `thrift:"result,1,optional" frugal:"1,optional,string" json:"result,omitempty"`
}

func NewEventPredictResult_() *EventPredictResult_ {
	return &EventPredictResult_{}
}

func (p *EventPredictResult_) InitDefault() {
	*p = EventPredictResult_{}
}

var EventPredictResult__Result__DEFAULT string

func (p *EventPredictResult_) GetResult_() (v string) {
	if !p.IsSetResult_() {
		return EventPredictResult__Result__DEFAULT
	}
	return *p.Result_
}
func (p *EventPredictResult_) SetResult_(val *string) {
	p.Result_ = val
}

var fieldIDToName_EventPredictResult_ = map[int16]string{
	1: "result",
}

func (p *EventPredictResult_) IsSetResult_() bool {
	return p.Result_ != nil
}

func (p *EventPredictResult_) Read(iprot thrift.TProtocol) (err error) {

	var fieldTypeId thrift.TType
	var fieldId int16

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

	return nil
ReadStructBeginError:
	return thrift.PrependError(fmt.Sprintf("%T read struct begin error: ", p), err)
ReadFieldBeginError:
	return thrift.PrependError(fmt.Sprintf("%T read field %d begin error: ", p, fieldId), err)
ReadFieldError:
	return thrift.PrependError(fmt.Sprintf("%T read field %d '%s' error: ", p, fieldId, fieldIDToName_EventPredictResult_[fieldId]), err)
SkipFieldError:
	return thrift.PrependError(fmt.Sprintf("%T field %d skip type %d error: ", p, fieldId, fieldTypeId), err)

ReadFieldEndError:
	return thrift.PrependError(fmt.Sprintf("%T read field end error", p), err)
ReadStructEndError:
	return thrift.PrependError(fmt.Sprintf("%T read struct end error: ", p), err)
}

func (p *EventPredictResult_) ReadField1(iprot thrift.TProtocol) error {
	if v, err := iprot.ReadString(); err != nil {
		return err
	} else {
		p.Result_ = &v
	}
	return nil
}

func (p *EventPredictResult_) Write(oprot thrift.TProtocol) (err error) {
	var fieldId int16
	if err = oprot.WriteStructBegin("EventPredictResult"); err != nil {
		goto WriteStructBeginError
	}
	if p != nil {
		if err = p.writeField1(oprot); err != nil {
			fieldId = 1
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

func (p *EventPredictResult_) writeField1(oprot thrift.TProtocol) (err error) {
	if p.IsSetResult_() {
		if err = oprot.WriteFieldBegin("result", thrift.STRING, 1); err != nil {
			goto WriteFieldBeginError
		}
		if err := oprot.WriteString(*p.Result_); err != nil {
			return err
		}
		if err = oprot.WriteFieldEnd(); err != nil {
			goto WriteFieldEndError
		}
	}
	return nil
WriteFieldBeginError:
	return thrift.PrependError(fmt.Sprintf("%T write field 1 begin error: ", p), err)
WriteFieldEndError:
	return thrift.PrependError(fmt.Sprintf("%T write field 1 end error: ", p), err)
}

func (p *EventPredictResult_) String() string {
	if p == nil {
		return "<nil>"
	}
	return fmt.Sprintf("EventPredictResult_(%+v)", *p)
}

func (p *EventPredictResult_) DeepEqual(ano *EventPredictResult_) bool {
	if p == ano {
		return true
	} else if p == nil || ano == nil {
		return false
	}
	if !p.Field1DeepEqual(ano.Result_) {
		return false
	}
	return true
}

func (p *EventPredictResult_) Field1DeepEqual(src *string) bool {

	if p.Result_ == src {
		return true
	} else if p.Result_ == nil || src == nil {
		return false
	}
	if strings.Compare(*p.Result_, *src) != 0 {
		return false
	}
	return true
}

type EventPredictor interface {
	Predict(ctx context.Context, request *EventPredictRequest) (r *EventPredictResult_, err error)

	Health(ctx context.Context) (r bool, err error)
}

type EventPredictorClient struct {
	c thrift.TClient
}

func NewEventPredictorClientFactory(t thrift.TTransport, f thrift.TProtocolFactory) *EventPredictorClient {
	return &EventPredictorClient{
		c: thrift.NewTStandardClient(f.GetProtocol(t), f.GetProtocol(t)),
	}
}

func NewEventPredictorClientProtocol(t thrift.TTransport, iprot thrift.TProtocol, oprot thrift.TProtocol) *EventPredictorClient {
	return &EventPredictorClient{
		c: thrift.NewTStandardClient(iprot, oprot),
	}
}

func NewEventPredictorClient(c thrift.TClient) *EventPredictorClient {
	return &EventPredictorClient{
		c: c,
	}
}

func (p *EventPredictorClient) Client_() thrift.TClient {
	return p.c
}

func (p *EventPredictorClient) Predict(ctx context.Context, request *EventPredictRequest) (r *EventPredictResult_, err error) {
	var _args EventPredictorPredictArgs
	_args.Request = request
	var _result EventPredictorPredictResult
	if err = p.Client_().Call(ctx, "predict", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}
func (p *EventPredictorClient) Health(ctx context.Context) (r bool, err error) {
	var _args EventPredictorHealthArgs
	var _result EventPredictorHealthResult
	if err = p.Client_().Call(ctx, "health", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}

type EventPredictorProcessor struct {
	processorMap map[string]thrift.TProcessorFunction
	handler      EventPredictor
}

func (p *EventPredictorProcessor) AddToProcessorMap(key string, processor thrift.TProcessorFunction) {
	p.processorMap[key] = processor
}

func (p *EventPredictorProcessor) GetProcessorFunction(key string) (processor thrift.TProcessorFunction, ok bool) {
	processor, ok = p.processorMap[key]
	return processor, ok
}

func (p *EventPredictorProcessor) ProcessorMap() map[string]thrift.TProcessorFunction {
	return p.processorMap
}

func NewEventPredictorProcessor(handler EventPredictor) *EventPredictorProcessor {
	self := &EventPredictorProcessor{handler: handler, processorMap: make(map[string]thrift.TProcessorFunction)}
	self.AddToProcessorMap("predict", &eventPredictorProcessorPredict{handler: handler})
	self.AddToProcessorMap("health", &eventPredictorProcessorHealth{handler: handler})
	return self
}
func (p *EventPredictorProcessor) Process(ctx context.Context, iprot, oprot thrift.TProtocol) (success bool, err thrift.TException) {
	name, _, seqId, err := iprot.ReadMessageBegin()
	if err != nil {
		return false, err
	}
	if processor, ok := p.GetProcessorFunction(name); ok {
		return processor.Process(ctx, seqId, iprot, oprot)
	}
	iprot.Skip(thrift.STRUCT)
	iprot.ReadMessageEnd()
	x := thrift.NewTApplicationException(thrift.UNKNOWN_METHOD, "Unknown function "+name)
	oprot.WriteMessageBegin(name, thrift.EXCEPTION, seqId)
	x.Write(oprot)
	oprot.WriteMessageEnd()
	oprot.Flush(ctx)
	return false, x
}

type eventPredictorProcessorPredict struct {
	handler EventPredictor
}

func (p *eventPredictorProcessorPredict) Process(ctx context.Context, seqId int32, iprot, oprot thrift.TProtocol) (success bool, err thrift.TException) {
	args := EventPredictorPredictArgs{}
	if err = args.Read(iprot); err != nil {
		iprot.ReadMessageEnd()
		x := thrift.NewTApplicationException(thrift.PROTOCOL_ERROR, err.Error())
		oprot.WriteMessageBegin("predict", thrift.EXCEPTION, seqId)
		x.Write(oprot)
		oprot.WriteMessageEnd()
		oprot.Flush(ctx)
		return false, err
	}

	iprot.ReadMessageEnd()
	var err2 error
	result := EventPredictorPredictResult{}
	var retval *EventPredictResult_
	if retval, err2 = p.handler.Predict(ctx, args.Request); err2 != nil {
		x := thrift.NewTApplicationException(thrift.INTERNAL_ERROR, "Internal error processing predict: "+err2.Error())
		oprot.WriteMessageBegin("predict", thrift.EXCEPTION, seqId)
		x.Write(oprot)
		oprot.WriteMessageEnd()
		oprot.Flush(ctx)
		return true, err2
	} else {
		result.Success = retval
	}
	if err2 = oprot.WriteMessageBegin("predict", thrift.REPLY, seqId); err2 != nil {
		err = err2
	}
	if err2 = result.Write(oprot); err == nil && err2 != nil {
		err = err2
	}
	if err2 = oprot.WriteMessageEnd(); err == nil && err2 != nil {
		err = err2
	}
	if err2 = oprot.Flush(ctx); err == nil && err2 != nil {
		err = err2
	}
	if err != nil {
		return
	}
	return true, err
}

type eventPredictorProcessorHealth struct {
	handler EventPredictor
}

func (p *eventPredictorProcessorHealth) Process(ctx context.Context, seqId int32, iprot, oprot thrift.TProtocol) (success bool, err thrift.TException) {
	args := EventPredictorHealthArgs{}
	if err = args.Read(iprot); err != nil {
		iprot.ReadMessageEnd()
		x := thrift.NewTApplicationException(thrift.PROTOCOL_ERROR, err.Error())
		oprot.WriteMessageBegin("health", thrift.EXCEPTION, seqId)
		x.Write(oprot)
		oprot.WriteMessageEnd()
		oprot.Flush(ctx)
		return false, err
	}

	iprot.ReadMessageEnd()
	var err2 error
	result := EventPredictorHealthResult{}
	var retval bool
	if retval, err2 = p.handler.Health(ctx); err2 != nil {
		x := thrift.NewTApplicationException(thrift.INTERNAL_ERROR, "Internal error processing health: "+err2.Error())
		oprot.WriteMessageBegin("health", thrift.EXCEPTION, seqId)
		x.Write(oprot)
		oprot.WriteMessageEnd()
		oprot.Flush(ctx)
		return true, err2
	} else {
		result.Success = &retval
	}
	if err2 = oprot.WriteMessageBegin("health", thrift.REPLY, seqId); err2 != nil {
		err = err2
	}
	if err2 = result.Write(oprot); err == nil && err2 != nil {
		err = err2
	}
	if err2 = oprot.WriteMessageEnd(); err == nil && err2 != nil {
		err = err2
	}
	if err2 = oprot.Flush(ctx); err == nil && err2 != nil {
		err = err2
	}
	if err != nil {
		return
	}
	return true, err
}

type EventPredictorPredictArgs struct {
	Request *EventPredictRequest `thrift:"request,1" frugal:"1,default,EventPredictRequest" json:"request"`
}

func NewEventPredictorPredictArgs() *EventPredictorPredictArgs {
	return &EventPredictorPredictArgs{}
}

func (p *EventPredictorPredictArgs) InitDefault() {
	*p = EventPredictorPredictArgs{}
}

var EventPredictorPredictArgs_Request_DEFAULT *EventPredictRequest

func (p *EventPredictorPredictArgs) GetRequest() (v *EventPredictRequest) {
	if !p.IsSetRequest() {
		return EventPredictorPredictArgs_Request_DEFAULT
	}
	return p.Request
}
func (p *EventPredictorPredictArgs) SetRequest(val *EventPredictRequest) {
	p.Request = val
}

var fieldIDToName_EventPredictorPredictArgs = map[int16]string{
	1: "request",
}

func (p *EventPredictorPredictArgs) IsSetRequest() bool {
	return p.Request != nil
}

func (p *EventPredictorPredictArgs) Read(iprot thrift.TProtocol) (err error) {

	var fieldTypeId thrift.TType
	var fieldId int16

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
			if fieldTypeId == thrift.STRUCT {
				if err = p.ReadField1(iprot); err != nil {
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

	return nil
ReadStructBeginError:
	return thrift.PrependError(fmt.Sprintf("%T read struct begin error: ", p), err)
ReadFieldBeginError:
	return thrift.PrependError(fmt.Sprintf("%T read field %d begin error: ", p, fieldId), err)
ReadFieldError:
	return thrift.PrependError(fmt.Sprintf("%T read field %d '%s' error: ", p, fieldId, fieldIDToName_EventPredictorPredictArgs[fieldId]), err)
SkipFieldError:
	return thrift.PrependError(fmt.Sprintf("%T field %d skip type %d error: ", p, fieldId, fieldTypeId), err)

ReadFieldEndError:
	return thrift.PrependError(fmt.Sprintf("%T read field end error", p), err)
ReadStructEndError:
	return thrift.PrependError(fmt.Sprintf("%T read struct end error: ", p), err)
}

func (p *EventPredictorPredictArgs) ReadField1(iprot thrift.TProtocol) error {
	p.Request = NewEventPredictRequest()
	if err := p.Request.Read(iprot); err != nil {
		return err
	}
	return nil
}

func (p *EventPredictorPredictArgs) Write(oprot thrift.TProtocol) (err error) {
	var fieldId int16
	if err = oprot.WriteStructBegin("predict_args"); err != nil {
		goto WriteStructBeginError
	}
	if p != nil {
		if err = p.writeField1(oprot); err != nil {
			fieldId = 1
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

func (p *EventPredictorPredictArgs) writeField1(oprot thrift.TProtocol) (err error) {
	if err = oprot.WriteFieldBegin("request", thrift.STRUCT, 1); err != nil {
		goto WriteFieldBeginError
	}
	if err := p.Request.Write(oprot); err != nil {
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

func (p *EventPredictorPredictArgs) String() string {
	if p == nil {
		return "<nil>"
	}
	return fmt.Sprintf("EventPredictorPredictArgs(%+v)", *p)
}

func (p *EventPredictorPredictArgs) DeepEqual(ano *EventPredictorPredictArgs) bool {
	if p == ano {
		return true
	} else if p == nil || ano == nil {
		return false
	}
	if !p.Field1DeepEqual(ano.Request) {
		return false
	}
	return true
}

func (p *EventPredictorPredictArgs) Field1DeepEqual(src *EventPredictRequest) bool {

	if !p.Request.DeepEqual(src) {
		return false
	}
	return true
}

type EventPredictorPredictResult struct {
	Success *EventPredictResult_ `thrift:"success,0,optional" frugal:"0,optional,EventPredictResult_" json:"success,omitempty"`
}

func NewEventPredictorPredictResult() *EventPredictorPredictResult {
	return &EventPredictorPredictResult{}
}

func (p *EventPredictorPredictResult) InitDefault() {
	*p = EventPredictorPredictResult{}
}

var EventPredictorPredictResult_Success_DEFAULT *EventPredictResult_

func (p *EventPredictorPredictResult) GetSuccess() (v *EventPredictResult_) {
	if !p.IsSetSuccess() {
		return EventPredictorPredictResult_Success_DEFAULT
	}
	return p.Success
}
func (p *EventPredictorPredictResult) SetSuccess(x interface{}) {
	p.Success = x.(*EventPredictResult_)
}

var fieldIDToName_EventPredictorPredictResult = map[int16]string{
	0: "success",
}

func (p *EventPredictorPredictResult) IsSetSuccess() bool {
	return p.Success != nil
}

func (p *EventPredictorPredictResult) Read(iprot thrift.TProtocol) (err error) {

	var fieldTypeId thrift.TType
	var fieldId int16

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
		case 0:
			if fieldTypeId == thrift.STRUCT {
				if err = p.ReadField0(iprot); err != nil {
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

	return nil
ReadStructBeginError:
	return thrift.PrependError(fmt.Sprintf("%T read struct begin error: ", p), err)
ReadFieldBeginError:
	return thrift.PrependError(fmt.Sprintf("%T read field %d begin error: ", p, fieldId), err)
ReadFieldError:
	return thrift.PrependError(fmt.Sprintf("%T read field %d '%s' error: ", p, fieldId, fieldIDToName_EventPredictorPredictResult[fieldId]), err)
SkipFieldError:
	return thrift.PrependError(fmt.Sprintf("%T field %d skip type %d error: ", p, fieldId, fieldTypeId), err)

ReadFieldEndError:
	return thrift.PrependError(fmt.Sprintf("%T read field end error", p), err)
ReadStructEndError:
	return thrift.PrependError(fmt.Sprintf("%T read struct end error: ", p), err)
}

func (p *EventPredictorPredictResult) ReadField0(iprot thrift.TProtocol) error {
	p.Success = NewEventPredictResult_()
	if err := p.Success.Read(iprot); err != nil {
		return err
	}
	return nil
}

func (p *EventPredictorPredictResult) Write(oprot thrift.TProtocol) (err error) {
	var fieldId int16
	if err = oprot.WriteStructBegin("predict_result"); err != nil {
		goto WriteStructBeginError
	}
	if p != nil {
		if err = p.writeField0(oprot); err != nil {
			fieldId = 0
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

func (p *EventPredictorPredictResult) writeField0(oprot thrift.TProtocol) (err error) {
	if p.IsSetSuccess() {
		if err = oprot.WriteFieldBegin("success", thrift.STRUCT, 0); err != nil {
			goto WriteFieldBeginError
		}
		if err := p.Success.Write(oprot); err != nil {
			return err
		}
		if err = oprot.WriteFieldEnd(); err != nil {
			goto WriteFieldEndError
		}
	}
	return nil
WriteFieldBeginError:
	return thrift.PrependError(fmt.Sprintf("%T write field 0 begin error: ", p), err)
WriteFieldEndError:
	return thrift.PrependError(fmt.Sprintf("%T write field 0 end error: ", p), err)
}

func (p *EventPredictorPredictResult) String() string {
	if p == nil {
		return "<nil>"
	}
	return fmt.Sprintf("EventPredictorPredictResult(%+v)", *p)
}

func (p *EventPredictorPredictResult) DeepEqual(ano *EventPredictorPredictResult) bool {
	if p == ano {
		return true
	} else if p == nil || ano == nil {
		return false
	}
	if !p.Field0DeepEqual(ano.Success) {
		return false
	}
	return true
}

func (p *EventPredictorPredictResult) Field0DeepEqual(src *EventPredictResult_) bool {

	if !p.Success.DeepEqual(src) {
		return false
	}
	return true
}

type EventPredictorHealthArgs struct {
}

func NewEventPredictorHealthArgs() *EventPredictorHealthArgs {
	return &EventPredictorHealthArgs{}
}

func (p *EventPredictorHealthArgs) InitDefault() {
	*p = EventPredictorHealthArgs{}
}

var fieldIDToName_EventPredictorHealthArgs = map[int16]string{}

func (p *EventPredictorHealthArgs) Read(iprot thrift.TProtocol) (err error) {

	var fieldTypeId thrift.TType
	var fieldId int16

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
		if err = iprot.Skip(fieldTypeId); err != nil {
			goto SkipFieldTypeError
		}

		if err = iprot.ReadFieldEnd(); err != nil {
			goto ReadFieldEndError
		}
	}
	if err = iprot.ReadStructEnd(); err != nil {
		goto ReadStructEndError
	}

	return nil
ReadStructBeginError:
	return thrift.PrependError(fmt.Sprintf("%T read struct begin error: ", p), err)
ReadFieldBeginError:
	return thrift.PrependError(fmt.Sprintf("%T read field %d begin error: ", p, fieldId), err)
SkipFieldTypeError:
	return thrift.PrependError(fmt.Sprintf("%T skip field type %d error", p, fieldTypeId), err)

ReadFieldEndError:
	return thrift.PrependError(fmt.Sprintf("%T read field end error", p), err)
ReadStructEndError:
	return thrift.PrependError(fmt.Sprintf("%T read struct end error: ", p), err)
}

func (p *EventPredictorHealthArgs) Write(oprot thrift.TProtocol) (err error) {
	if err = oprot.WriteStructBegin("health_args"); err != nil {
		goto WriteStructBeginError
	}
	if p != nil {

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
WriteFieldStopError:
	return thrift.PrependError(fmt.Sprintf("%T write field stop error: ", p), err)
WriteStructEndError:
	return thrift.PrependError(fmt.Sprintf("%T write struct end error: ", p), err)
}

func (p *EventPredictorHealthArgs) String() string {
	if p == nil {
		return "<nil>"
	}
	return fmt.Sprintf("EventPredictorHealthArgs(%+v)", *p)
}

func (p *EventPredictorHealthArgs) DeepEqual(ano *EventPredictorHealthArgs) bool {
	if p == ano {
		return true
	} else if p == nil || ano == nil {
		return false
	}
	return true
}

type EventPredictorHealthResult struct {
	Success *bool `thrift:"success,0,optional" frugal:"0,optional,bool" json:"success,omitempty"`
}

func NewEventPredictorHealthResult() *EventPredictorHealthResult {
	return &EventPredictorHealthResult{}
}

func (p *EventPredictorHealthResult) InitDefault() {
	*p = EventPredictorHealthResult{}
}

var EventPredictorHealthResult_Success_DEFAULT bool

func (p *EventPredictorHealthResult) GetSuccess() (v bool) {
	if !p.IsSetSuccess() {
		return EventPredictorHealthResult_Success_DEFAULT
	}
	return *p.Success
}
func (p *EventPredictorHealthResult) SetSuccess(x interface{}) {
	p.Success = x.(*bool)
}

var fieldIDToName_EventPredictorHealthResult = map[int16]string{
	0: "success",
}

func (p *EventPredictorHealthResult) IsSetSuccess() bool {
	return p.Success != nil
}

func (p *EventPredictorHealthResult) Read(iprot thrift.TProtocol) (err error) {

	var fieldTypeId thrift.TType
	var fieldId int16

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
		case 0:
			if fieldTypeId == thrift.BOOL {
				if err = p.ReadField0(iprot); err != nil {
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

	return nil
ReadStructBeginError:
	return thrift.PrependError(fmt.Sprintf("%T read struct begin error: ", p), err)
ReadFieldBeginError:
	return thrift.PrependError(fmt.Sprintf("%T read field %d begin error: ", p, fieldId), err)
ReadFieldError:
	return thrift.PrependError(fmt.Sprintf("%T read field %d '%s' error: ", p, fieldId, fieldIDToName_EventPredictorHealthResult[fieldId]), err)
SkipFieldError:
	return thrift.PrependError(fmt.Sprintf("%T field %d skip type %d error: ", p, fieldId, fieldTypeId), err)

ReadFieldEndError:
	return thrift.PrependError(fmt.Sprintf("%T read field end error", p), err)
ReadStructEndError:
	return thrift.PrependError(fmt.Sprintf("%T read struct end error: ", p), err)
}

func (p *EventPredictorHealthResult) ReadField0(iprot thrift.TProtocol) error {
	if v, err := iprot.ReadBool(); err != nil {
		return err
	} else {
		p.Success = &v
	}
	return nil
}

func (p *EventPredictorHealthResult) Write(oprot thrift.TProtocol) (err error) {
	var fieldId int16
	if err = oprot.WriteStructBegin("health_result"); err != nil {
		goto WriteStructBeginError
	}
	if p != nil {
		if err = p.writeField0(oprot); err != nil {
			fieldId = 0
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

func (p *EventPredictorHealthResult) writeField0(oprot thrift.TProtocol) (err error) {
	if p.IsSetSuccess() {
		if err = oprot.WriteFieldBegin("success", thrift.BOOL, 0); err != nil {
			goto WriteFieldBeginError
		}
		if err := oprot.WriteBool(*p.Success); err != nil {
			return err
		}
		if err = oprot.WriteFieldEnd(); err != nil {
			goto WriteFieldEndError
		}
	}
	return nil
WriteFieldBeginError:
	return thrift.PrependError(fmt.Sprintf("%T write field 0 begin error: ", p), err)
WriteFieldEndError:
	return thrift.PrependError(fmt.Sprintf("%T write field 0 end error: ", p), err)
}

func (p *EventPredictorHealthResult) String() string {
	if p == nil {
		return "<nil>"
	}
	return fmt.Sprintf("EventPredictorHealthResult(%+v)", *p)
}

func (p *EventPredictorHealthResult) DeepEqual(ano *EventPredictorHealthResult) bool {
	if p == ano {
		return true
	} else if p == nil || ano == nil {
		return false
	}
	if !p.Field0DeepEqual(ano.Success) {
		return false
	}
	return true
}

func (p *EventPredictorHealthResult) Field0DeepEqual(src *bool) bool {

	if p.Success == src {
		return true
	} else if p.Success == nil || src == nil {
		return false
	}
	if *p.Success != *src {
		return false
	}
	return true
}
