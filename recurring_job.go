package main

import (
	"time"

	"github.com/vpenando/piggy/piggy"
)

// RecurringJobFrequency is an enum that represents the execution
// time interval between each execution of a recurring job.
type RecurringJobFrequency int

// Available frequencies.
const (
	RecurringJobFrequencyNone RecurringJobFrequency = iota
	RecurringJobFrequencyWeekly
	RecurringJobFrequencyMonthly
)

// RecurringJob a
type RecurringJob interface {
	Check() (bool, error)
}

type RecurringOperationJob struct {
	Operation         piggy.Operation
	OperationID       int
	NextExecutionTime time.Time
	Frequency         RecurringJobFrequency
}

func NewRecurringOperationJob(operation piggy.Operation, nextTime time.Time, frequency RecurringJobFrequency) *RecurringOperationJob {
	return &RecurringOperationJob{
		Operation:         operation,
		Frequency:         frequency,
		NextExecutionTime: nextTime,
	}
}

func today() time.Time {
	now := time.Now()
	return time.Date(now.Year(), now.Month(), now.Day(), 0, 0, 0, 0, time.Local)
}

func (job *RecurringOperationJob) canExecute() bool {
	if job.Frequency == RecurringJobFrequencyNone {
		return false
	}
	return job.NextExecutionTime == today()
}

func (job *RecurringOperationJob) planNextExecution() {
	year, month, day := job.NextExecutionTime.Date()
	if job.Frequency == RecurringJobFrequencyMonthly {
		month++
	} else if job.Frequency == RecurringJobFrequencyWeekly {
		day += 7
	}
	job.NextExecutionTime = time.Date(year, month, day, 0, 0, 0, 0, time.Local)
}
