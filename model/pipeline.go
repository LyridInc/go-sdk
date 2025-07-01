package model

import "time"

const (
	PipelineRunningStageStatus  = "RUNNING"
	PipelineCompleteStageStatus = "COMPLETE"
	PipelineErrorStageStatus    = "ERROR"
)

const (
	PipelinePackageStartedStatus       = "PACKAGE_STARTED"
	PipelinePackageCompleteStatus      = "PACKAGE_COMPLETE"
	PipelinePackageErrorStatus         = "PACKAGE_ERROR"
	PipelineSubmitStartedStatus        = "SUBMIT_STARTED"
	PipelineSubmitCompleteStatus       = "SUBMIT_COMPLETE"
	PipelineSubmitErrorStatus          = "SUBMIT_ERROR"
	PipelineBuildStartedStatus         = "BUILD_STARTED"
	PipelineBuildCompleteStatus        = "BUILD_COMPLETE"
	PipelineBuildErrorStatus           = "BUILD_ERROR"
	PipelinePushingImageStartedStatus  = "PUSHING_IMAGE_STARTED"
	PipelinePushingImageCompleteStatus = "PUSHING_IMAGE_COMPLETE"
	PipelinePushingImageErrorStatus    = "PUSHING_IMAGE_ERROR"
	PipelineDeployStartedStatus        = "DEPLOY_STARTED"
	PipelineDeployCompleteStatus       = "DEPLOY_COMPLETE"
	PipelineDeployErrorStatus          = "DEPLOY_ERROR"
	PipelinePodSpawningStartedStatus   = "POD_SPAWNING_STARTED"
	PipelinePodSpawningCompleteStatus  = "POD_SPAWNING_COMPLETE"
	PipelinePodSpawningErrorStatus     = "POD_SPAWNING_ERROR"
)

const (
	PipelinePackageStage      = "PACKAGE"
	PipelineSubmitStage       = "SUBMIT"
	PipelineBuildStage        = "BUILD"
	PipelinePushingImageStage = "PUSHING_IMAGE"
	PipelineDeployStage       = "DEPLOY"
	PipelinePodSpawningStage  = "POD_SPAWNING"
)

const (
	MessageType            string = "MESSAGE"
	ErrorType              string = "ERROR"
	ProgressPercentageType string = "PROGRESSPERCENTAGE"
)

type StageDefinition struct {
	Status string                `json:"status"`
	Stages map[string]*StageLogs `json:"stages"`
}

type StageLogs struct {
	Logs []StageLog `json:"logs"`
}

type StageLog struct {
	TargetPlatform string          `json:"targetPlatform"`
	TargetRegion   string          `json:"targetRegion"`
	Status         string          `json:"status"`
	JobID          string          `json:"jobId"`
	Messages       []*StageMessage `json:"messages"`
}

type StageMessageDetail struct {
	Subject            string `json:"subject"`
	Message            string `json:"message"`
	ProgressPercentage string `json:"progressPercentage"`
}

type StageMessage struct {
	Time     time.Time          `json:"time"`
	Severity string             `json:"severity"`
	Message  string             `json:"message"`
	Detail   StageMessageDetail `json:"detail"`
}

type RevisionLogContext struct {
	RevisionId string
	JobId      string
	Stage      string
}

func NewStageDefinition() *StageDefinition {
	definition := &StageDefinition{}
	definition.Stages = make(map[string]*StageLogs, 0)
	return definition
}

func (definition *StageDefinition) CreateStage(stage string) {
	definition.Stages[stage] = &StageLogs{Logs: make([]StageLog, 0)}
}

func (definition *StageDefinition) GetStage(stage string) *StageLogs {
	if definition.Stages[stage] == nil {
		definition.CreateStage(stage)
	}

	return definition.Stages[stage]
}

func (logs *StageLogs) CreateStageLog(platform string, region string) {
	if logs.GetStageLog(platform, region) == nil {
		logs.Logs = append(logs.Logs, StageLog{
			TargetPlatform: platform,
			TargetRegion:   region,
		})
	}
}

func (logs *StageLogs) GetStageLog(platform string, region string) *StageLog {
	for _, log := range logs.Logs {
		if log.TargetRegion == region && log.TargetPlatform == platform {
			return &log
		}
	}
	return nil
}

func (log *StageLog) SetStageStatus(status string) {
	log.Status = status
}

func (log *StageLog) AppendMessage(message *StageMessage) {
	log.Messages = append(log.Messages, message)
}
