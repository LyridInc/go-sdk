package model

import "time"

type JobStatusEnum uint

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

type Job struct {
	ID             string        `json:"id" binding:"required"`
	AccountId      string        `json:"accountId" binding:"required"`
	Type           string        `json:"type" binding:"required"`
	Status         JobStatusEnum `json:"status" binding:"required"`
	CreationTime   time.Time     `json:"creationTime"`
	LastUpdateTime time.Time     `json:"lastUpdateTime"`
	Payload        string        `json:"payload"`
}

type JobInfo struct {
	JobId   string    `json:"jobId" binding:"required"`
	Message string    `json:"message" binding:"required"`
	Type    string    `json:"type" binding:"required"`
	Time    time.Time `json:"time"`
}
