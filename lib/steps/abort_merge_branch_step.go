package steps

import "github.com/Originate/git-town/lib/script"

// AbortMergeBranchStep aborts the current merge conflict.
type AbortMergeBranchStep struct {
	NoAutomaticAbortOnError
	NoUndoStep
}

// CreateAbortStep returns the abort step for this step.
func (step AbortMergeBranchStep) CreateAbortStep() Step {
	return NoOpStep{}
}

// CreateContinueStep returns the continue step for this step.
func (step AbortMergeBranchStep) CreateContinueStep() Step {
	return NoOpStep{}
}

// Run executes this step.
func (step AbortMergeBranchStep) Run() error {
	return script.RunCommand("git", "merge", "--abort")
}
