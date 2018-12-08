package main

import (
	"fmt"
	"github.com/pastushenkoy/adventOfCode/utils"
	"regexp"
	"sort"
	"strings"
)

type StepCondition struct {
	name      string
	dependsOn string
}

func main() {
	data := getConditions(utils.ReadFileOfStrings("input.txt"))

	fmt.Println("The step sequence is '", getStepSequence(data, false), "'")
	fmt.Println("Work will take ", getWorkTime(data, 5, 60), "seconds")
}

type Step struct {
	name        rune
	nameStr		string
	dependentOn []string
}

type Worker struct {
	currentStep string
	secondsLeft int
}

func getWorkTime(conditions []*StepCondition, workerCount int, workingTimeModifier int) int {
	sequence := []rune(getStepSequence(conditions, true))

	steps := make([]Step, 0)
	for _, stepName := range sequence {
		step := Step{name: stepName, nameStr: string(stepName)}
		for _, cond := range conditions {
			if cond.name == string(stepName) {
				step.dependentOn = append(step.dependentOn, cond.dependsOn)
			}
		}
		steps = append(steps, step)
	}

	workers := make([]*Worker, workerCount)
	for i := 0; i < workerCount; i++ {
		workers[i] = &Worker{}
	}

	doneSteps := ""
	i := 0
	for {
		jobIsDone := true
		doneInIteration := ""
		for _, worker := range workers {
			if worker.secondsLeft == 0 {
				stepIndex, canBeTaken := canTakeNextStep(steps, doneSteps)
				if canBeTaken {
					stepName := steps[stepIndex].name
					worker.secondsLeft = getTaskTime(stepName, workingTimeModifier)
					worker.currentStep = string(steps[stepIndex].name)
					steps = append(steps[:stepIndex], steps[stepIndex+1:]...)
					jobIsDone = false
				} else {
					worker.currentStep = ""
				}
			}

			if worker.secondsLeft > 0 {
				worker.secondsLeft--
				jobIsDone = false
			}

			if worker.secondsLeft == 0 {
				doneInIteration += string(worker.currentStep)

			}

		}


		//fmt.Printf("%v\t%s\t%s\t%s\n", i, workers[0].currentStep, workers[1].currentStep, doneSteps)
		fmt.Printf("%v\t%s\t%s\t%s\t%s\t%s\t%s\n", i, workers[0].currentStep, workers[1].currentStep, workers[2].currentStep, workers[3].currentStep, workers[4].currentStep, doneSteps)

		doneSteps += doneInIteration

		if jobIsDone {
			break
		}

		i++
	}

	return i
}

func canTakeNextStep(steps []Step, doneSteps string) (stepIndex int, canBeTaken bool) {
	if len(steps) == 0{
		return -1, false
	}

	for ind, step := range steps {
		dependenciesDone := true
		for _, dependentOn := range step.dependentOn {
			if !strings.Contains(doneSteps, dependentOn) {
				dependenciesDone = false
				break
			}
		}
		if dependenciesDone{
			return ind, true
		}
	}
	return -1, false
}

func getTaskTime(r rune, modifier int) int {
	return modifier + int(r-64)
}

func getStepSequence(conditions []*StepCondition, forWorkTimeCount bool) string {
	steps := getSteps(conditions)

	sequence := ""
	for {
		availableSteps := make([]string, 0)

		for stepName, stepConditions := range steps {
			if strings.Contains(sequence, stepName) {
				continue;
			}
			if allConditionsFullfiled(sequence, stepConditions) {
				availableSteps = append(availableSteps, stepName)
			}
		}

		if len(availableSteps) == 0 {
			break
		}

		sort.Strings(availableSteps)
		if forWorkTimeCount {
			for _, avSt := range availableSteps {
				sequence += avSt
			}
		}else {
			sequence += availableSteps[0]
		}
	}

	return sequence

}
func allConditionsFullfiled(sequence string, stepConditions []string) bool {
	for _, cond := range stepConditions {
		if !strings.Contains(sequence, cond) {
			return false
		}
	}
	return true
}

func getSteps(conditions []*StepCondition) map[string][]string {
	steps := make(map[string][]string, 0)
	for _, condition := range conditions {
		stepConds := createStepIfNeeded(steps, condition.name)
		steps[condition.name] = append(stepConds, condition.dependsOn)

		createStepIfNeeded(steps, condition.dependsOn)
	}
	return steps
}

func createStepIfNeeded(steps map[string][]string, stepName string) []string {
	stepConds, hasValue := steps[stepName]
	if !hasValue {
		steps[stepName] = make([]string, 0)
	}
	return stepConds
}

func getConditions(data []string) []*StepCondition {
	re := regexp.MustCompile("Step (\\w) must be finished before step (\\w) can begin.")

	steps := make([]*StepCondition, 0)
	for _, row := range data {
		submatch := re.FindStringSubmatch(row)

		steps = append(steps, &StepCondition{
			name:      submatch[2],
			dependsOn: submatch[1],
		})
	}

	return steps
}
