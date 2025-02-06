package model

import "time"

type JobStatusEnum uint
type HelmJobStage uint

const (
	Submitting JobStatusEnum = iota
	Submitted
	Queued
	Running
	Completed
	Error
	Warning
	Cancelled
	Finished
	TimedOut
)

const (
	HelmSubmit HelmJobStage = iota
	HelmTemplating
	HelmLint
	HelmExecute
)

type Job struct {
	ID             string        `json:"id" binding:"required"`
	AccountId      string        `json:"accountId" binding:"required"`
	Type           string        `json:"type" binding:"required"`
	Status         JobStatusEnum `json:"status" binding:"required"`
	Subject        string        `json:"subject"`
	CreationTime   time.Time     `json:"creationTime"`
	LastUpdateTime time.Time     `json:"lastUpdateTime"`
	Payload        string        `json:"payload"`
	RetryJobType   string        `json:"retryJobType"`
}

type JobInfo struct {
	JobId   string       `json:"jobId" binding:"required"`
	Message string       `json:"message" binding:"required"`
	Type    string       `json:"type" binding:"required"`
	Stage   HelmJobStage `json:"stage"`
	Time    time.Time    `json:"time"`
}
