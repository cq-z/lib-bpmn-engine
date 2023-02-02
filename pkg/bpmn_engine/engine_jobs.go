package bpmn_engine

import (
	"encoding/xml"
	"github.com/nitram509/lib-bpmn-engine/pkg/bpmn_engine/var_holder"
	"github.com/nitram509/lib-bpmn-engine/pkg/spec/BPMN20/process_instance"
	"time"

	"github.com/nitram509/lib-bpmn-engine/pkg/spec/BPMN20"
	"github.com/nitram509/lib-bpmn-engine/pkg/spec/BPMN20/activity"
)

type job struct {
	ElementId          string
	ElementInstanceKey int64
	ProcessInstanceKey int64
	JobKey             int64
	State              activity.LifecycleState
	CreatedAt          time.Time
}

func (state *BpmnEngineState) handleServiceTask(process *ProcessInfo, instance *ProcessInstanceInfo, element *BPMN20.TaskElement) bool {
	id := (*element).GetId()
	job := findOrCreateJob(&state.jobs, id, instance, state.generateKey)

	handler := state.findTaskHandler(element)
	if handler != nil {
		job.State = activity.Active
		variableHolder := var_holder.New(&instance.variableHolder, nil)
		attributes := (*element).GetAttributes()
		if attributes == nil {
			attributes = make([]xml.Attr, 0)
		}
		activatedJob := &activatedJob{
			processInstanceInfo:      instance,
			failHandler:              func(reason string) { job.State = activity.Failed },
			completeHandler:          func() { job.State = activity.Completed },
			key:                      state.generateKey(),
			processInstanceKey:       instance.instanceKey,
			bpmnProcessId:            process.BpmnProcessId,
			processDefinitionVersion: process.Version,
			processDefinitionKey:     process.ProcessKey,
			elementId:                job.ElementId,
			createdAt:                job.CreatedAt,
			variableHolder:           variableHolder,
			attributes:               attributes,
		}
		if err := evaluateLocalVariables(variableHolder, (*element).GetInputMapping()); err != nil {
			job.State = activity.Failed
			instance.state = process_instance.FAILED
			return false
		}
		handler(activatedJob)
		if job.State == activity.Completed {
			if err := propagateProcessInstanceVariables(variableHolder, (*element).GetOutputMapping()); err != nil {
				job.State = activity.Failed
				instance.state = process_instance.FAILED
			}
		}
	}

	return job.State == activity.Completed
}

func findOrCreateJob(jobs *[]*job, id string, instance *ProcessInstanceInfo, generateKey func() int64) *job {
	for _, job := range *jobs {
		if job.ElementId == id {
			return job
		}
	}

	elementInstanceKey := generateKey()
	job := job{
		ElementId:          id,
		ElementInstanceKey: elementInstanceKey,
		ProcessInstanceKey: instance.GetInstanceKey(),
		JobKey:             elementInstanceKey + 1,
		State:              activity.Active,
		CreatedAt:          time.Now(),
	}

	*jobs = append(*jobs, &job)

	return &job
}
