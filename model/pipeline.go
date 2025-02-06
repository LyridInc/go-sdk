package model

import "time"

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

type StageMessage struct {
	Time     time.Time `json:"time"`
	Severity string    `json:"severity"`
	Message  string    `json:"message"`
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
