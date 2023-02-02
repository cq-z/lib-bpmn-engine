package BPMN20

import (
	"encoding/xml"
	"github.com/nitram509/lib-bpmn-engine/pkg/spec/BPMN20/extensions"
)

type ElementType string

const (
	StartEvent             ElementType = "START_EVENT"
	EndEvent               ElementType = "END_EVENT"
	ServiceTask            ElementType = "SERVICE_TASK"
	UserTask               ElementType = "USER_TASK"
	ParallelGateway        ElementType = "PARALLEL_GATEWAY"
	ExclusiveGateway       ElementType = "EXCLUSIVE_GATEWAY"
	IntermediateCatchEvent ElementType = "INTERMEDIATE_CATCH_EVENT"
	EventBasedGateway      ElementType = "EVENT_BASED_GATEWAY"

	SequenceFlow ElementType = "SEQUENCE_FLOW"
)

type BaseElement interface {
	GetId() string
	GetName() string
	GetIncomingAssociation() []string
	GetOutgoingAssociation() []string
	GetType() ElementType
	GetAttributes() []xml.Attr
}

type TaskElement interface {
	BaseElement
	GetInputMapping() []extensions.TIoMapping
	GetOutputMapping() []extensions.TIoMapping
	GetTaskDefinitionType() string
	GetAssignmentAssignee() string
	GetAssignmentCandidateGroups() []string
}

func (startEvent TStartEvent) GetId() string {
	return startEvent.Id
}

func (startEvent TStartEvent) GetName() string {
	return startEvent.Name
}

func (startEvent TStartEvent) GetIncomingAssociation() []string {
	return startEvent.IncomingAssociation
}

func (startEvent TStartEvent) GetOutgoingAssociation() []string {
	return startEvent.OutgoingAssociation
}

func (startEvent TStartEvent) GetType() ElementType {
	return StartEvent
}

func (startEvent TStartEvent) GetAttributes() []xml.Attr {
	return startEvent.Attributes
}

func (endEvent TEndEvent) GetId() string {
	return endEvent.Id
}

func (endEvent TEndEvent) GetName() string {
	return endEvent.Name
}

func (endEvent TEndEvent) GetIncomingAssociation() []string {
	return endEvent.IncomingAssociation
}

func (endEvent TEndEvent) GetOutgoingAssociation() []string {
	return endEvent.OutgoingAssociation
}

func (endEvent TEndEvent) GetType() ElementType {
	return EndEvent
}

func (endEvent TEndEvent) GetAttributes() []xml.Attr {
	return endEvent.Attributes
}

func (serviceTask TServiceTask) GetId() string {
	return serviceTask.Id
}

func (serviceTask TServiceTask) GetName() string {
	return serviceTask.Name
}

func (serviceTask TServiceTask) GetIncomingAssociation() []string {
	return serviceTask.IncomingAssociation
}

func (serviceTask TServiceTask) GetOutgoingAssociation() []string {
	return serviceTask.OutgoingAssociation
}

func (serviceTask TServiceTask) GetType() ElementType {
	return ServiceTask
}

func (serviceTask TServiceTask) GetInputMapping() []extensions.TIoMapping {
	return serviceTask.Input
}

func (serviceTask TServiceTask) GetOutputMapping() []extensions.TIoMapping {
	return serviceTask.Output
}

func (serviceTask TServiceTask) GetTaskDefinitionType() string {
	return serviceTask.TaskDefinition.TypeName
}

func (serviceTask TServiceTask) GetAssignmentAssignee() string {
	return ""
}

func (serviceTask TServiceTask) GetAssignmentCandidateGroups() []string {
	return []string{}
}

func (serviceTask TServiceTask) GetAttributes() []xml.Attr {
	return serviceTask.Attributes
}

func (userTask TUserTask) GetId() string {
	return userTask.Id
}

func (userTask TUserTask) GetName() string {
	return userTask.Name
}

func (userTask TUserTask) GetIncomingAssociation() []string {
	return userTask.IncomingAssociation
}

func (userTask TUserTask) GetOutgoingAssociation() []string {
	return userTask.OutgoingAssociation
}

func (userTask TUserTask) GetType() ElementType {
	return UserTask
}

func (userTask TUserTask) GetInputMapping() []extensions.TIoMapping {
	return userTask.Input
}

func (userTask TUserTask) GetOutputMapping() []extensions.TIoMapping {
	return userTask.Output
}

func (userTask TUserTask) GetTaskDefinitionType() string {
	return ""
}

func (userTask TUserTask) GetAssignmentAssignee() string {
	return userTask.AssignmentDefinition.Assignee
}

func (userTask TUserTask) GetAssignmentCandidateGroups() []string {
	return userTask.AssignmentDefinition.GetCandidateGroups()
}

func (userTask TUserTask) GetAttributes() []xml.Attr {
	return userTask.Attributes
}

func (parallelGateway TParallelGateway) GetId() string {
	return parallelGateway.Id
}

func (parallelGateway TParallelGateway) GetName() string {
	return parallelGateway.Name
}

func (parallelGateway TParallelGateway) GetIncomingAssociation() []string {
	return parallelGateway.IncomingAssociation
}

func (parallelGateway TParallelGateway) GetOutgoingAssociation() []string {
	return parallelGateway.OutgoingAssociation
}

func (parallelGateway TParallelGateway) GetType() ElementType {
	return ParallelGateway
}

func (parallelGateway TParallelGateway) GetAttributes() []xml.Attr {
	return parallelGateway.Attributes
}

func (exclusiveGateway TExclusiveGateway) GetId() string {
	return exclusiveGateway.Id
}

func (exclusiveGateway TExclusiveGateway) GetName() string {
	return exclusiveGateway.Name
}

func (exclusiveGateway TExclusiveGateway) GetIncomingAssociation() []string {
	return exclusiveGateway.IncomingAssociation
}

func (exclusiveGateway TExclusiveGateway) GetOutgoingAssociation() []string {
	return exclusiveGateway.OutgoingAssociation
}

func (exclusiveGateway TExclusiveGateway) GetType() ElementType {
	return ExclusiveGateway
}

func (exclusiveGateway TExclusiveGateway) GetAttributes() []xml.Attr {
	return exclusiveGateway.Attributes
}

func (intermediateCatchEvent TIntermediateCatchEvent) GetId() string {
	return intermediateCatchEvent.Id
}

func (intermediateCatchEvent TIntermediateCatchEvent) GetName() string {
	return intermediateCatchEvent.Name
}

func (intermediateCatchEvent TIntermediateCatchEvent) GetIncomingAssociation() []string {
	return intermediateCatchEvent.IncomingAssociation
}

func (intermediateCatchEvent TIntermediateCatchEvent) GetOutgoingAssociation() []string {
	return intermediateCatchEvent.OutgoingAssociation
}

func (intermediateCatchEvent TIntermediateCatchEvent) GetType() ElementType {
	return IntermediateCatchEvent
}

func (intermediateCatchEvent TIntermediateCatchEvent) GetAttributes() []xml.Attr {
	return intermediateCatchEvent.Attributes
}

func (eventBasedGateway TEventBasedGateway) GetId() string {
	return eventBasedGateway.Id
}

func (eventBasedGateway TEventBasedGateway) GetName() string {
	return eventBasedGateway.Name
}

func (eventBasedGateway TEventBasedGateway) GetIncomingAssociation() []string {
	return eventBasedGateway.IncomingAssociation
}

func (eventBasedGateway TEventBasedGateway) GetOutgoingAssociation() []string {
	return eventBasedGateway.OutgoingAssociation
}

func (eventBasedGateway TEventBasedGateway) GetType() ElementType {
	return EventBasedGateway
}

func (eventBasedGateway TEventBasedGateway) GetAttributes() []xml.Attr {
	return eventBasedGateway.Attributes
}
