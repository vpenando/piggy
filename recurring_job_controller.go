package main

import (
	"github.com/vpenando/piggy/piggy"
	"gorm.io/gorm"
)

type RecurringOperationJobController struct {
	db      *gorm.DB
	Verbose bool
}

func NewRecurringOperationJobController(db *gorm.DB) (*RecurringOperationJobController, error) {
	controller := &RecurringOperationJobController{
		db:      db,
		Verbose: false,
	}
	err := controller.db.AutoMigrate(&RecurringOperationJob{})
	return controller, err
}

func (rojc *RecurringOperationJobController) CheckPendingJobs() {

}

func (rojc *RecurringOperationJobController) CheckJob(job RecurringOperationJob) error {
	err := rojc.recoverMissingExecutions(job)
	if err != nil {
		return nil
	}
	if job.canExecute() {
		// Set next execution time to today() + job.Frequency
		job.planNextExecution()
		return rojc.execute(job)
	}
	return nil
}

func (rojc *RecurringOperationJobController) execute(job RecurringOperationJob) error {
	operationController, err := piggy.NewOperationController(rojc.db)
	if err != nil {
		return err
	}
	operation, err := operationController.CreateOne(job.Operation)
	if err != nil {
		return err
	}
	job.OperationID = operation.ID
	return rojc.db.Create(&job).Error
}

func (rojc *RecurringOperationJobController) recoverMissingExecutions(job RecurringOperationJob) error {
	missingJobs := []RecurringOperationJob{}
	for _, j := range missingJobs {
		j.Operation.Date = j.NextExecutionTime
		if err := rojc.execute(j); err != nil {
			return err
		}
	}
	return nil
}
