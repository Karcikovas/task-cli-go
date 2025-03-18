package task

import "errors"

var ErrMissingStatusFlag = errors.New("missing status flag")
var ErrEmptyStatus = errors.New("empty status")
var ErrFailedToExtractStatus = errors.New("failed to extract status")
var ErrWrongStatusTypePassed = errors.New("wrong status type passed")
var ErrWrongArgumentPassed = errors.New("wrong argument passed")
var ErrFailedAddTask = errors.New("failed add task")
var ErrUnableDeleteTask = errors.New("unable delete task")
