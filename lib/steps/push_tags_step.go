package steps

import (
	"github.com/Originate/git-town/lib/script"
)

// PushTagsStep pushes newly created Git tags to the remote.
type PushTagsStep struct {
	NoAutomaticAbortOnError
	NoUndoStep
}

// CreateAbortStep returns the abort step for this step.
func (step PushTagsStep) CreateAbortStep() Step {
	return NoOpStep{}
}

// CreateContinueStep returns the continue step for this step.
func (step PushTagsStep) CreateContinueStep() Step {
	return NoOpStep{}
}

// Run executes this step.
func (step PushTagsStep) Run() error {
	return script.RunCommand("git", "push", "--tags")
}
