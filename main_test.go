package main

import (
	"log"
	"os/exec"
	"strconv"
	"strings"
	"testing"
)

func TestNumOfTurnsSolution_ex00(t *testing.T) {
	// Construct the command string
	output := ExecuteCommand("go run . examples/example00.txt")
	inputFile := ExecuteCommand("cat examples/example00.txt")
	numOfTurns := (NumOfLines(output) - 1) - NumOfLines(inputFile) // (NumOfLines(output)-1) because of the new line after printing whats in the file
	maxNumOfTurns := 6

	// Compare the output with the solution content
	if numOfTurns > maxNumOfTurns {
		t.Errorf("Output does not match the solution.\nOutput:\n%s\nWith %s Number of turns\n\nMax Number of turns is: %s", LastLines(output, numOfTurns), strconv.Itoa(numOfTurns), strconv.Itoa(maxNumOfTurns))
	}
}

func TestNumOfTurnsSolution_ex01(t *testing.T) {
	// Construct the command string
	output := ExecuteCommand("go run . examples/example01.txt")
	inputFile := ExecuteCommand("cat examples/example01.txt")
	numOfTurns := (NumOfLines(output) - 1) - NumOfLines(inputFile) // (NumOfLines(output)-1) because of the new line after printing whats in the file
	maxNumOfTurns := 8

	// Compare the output with the solution content
	if numOfTurns > maxNumOfTurns {
		t.Errorf("Output does not match the solution.\nOutput:\n%s\nWith %s Number of turns\n\nMax Number of turns is: %s", LastLines(output, numOfTurns), strconv.Itoa(numOfTurns), strconv.Itoa(maxNumOfTurns))
	}
}

func TestNumOfTurnsSolution_ex02(t *testing.T) {
	// Construct the command string
	output := ExecuteCommand("go run . examples/example02.txt")
	inputFile := ExecuteCommand("cat examples/example02.txt")
	numOfTurns := (NumOfLines(output) - 1) - NumOfLines(inputFile) // (NumOfLines(output)-1) because of the new line after printing whats in the file
	maxNumOfTurns := 11

	// Compare the output with the solution content
	if numOfTurns > maxNumOfTurns {
		t.Errorf("Output does not match the solution.\nOutput:\n%s\nWith %s Number of turns\n\nMax Number of turns is: %s", LastLines(output, numOfTurns), strconv.Itoa(numOfTurns), strconv.Itoa(maxNumOfTurns))
	}
}

func TestNumOfTurnsSolution_ex03(t *testing.T) {
	// Construct the command string
	output := ExecuteCommand("go run . examples/example03.txt")
	inputFile := ExecuteCommand("cat examples/example03.txt")
	numOfTurns := (NumOfLines(output) - 1) - NumOfLines(inputFile) // (NumOfLines(output)-1) because of the new line after printing whats in the file
	maxNumOfTurns := 6

	// Compare the output with the solution content
	if numOfTurns > maxNumOfTurns {
		t.Errorf("Output does not match the solution.\nOutput:\n%s\nWith %s Number of turns\n\nMax Number of turns is: %s", LastLines(output, numOfTurns), strconv.Itoa(numOfTurns), strconv.Itoa(maxNumOfTurns))
	}
}

func TestNumOfTurnsSolution_ex04(t *testing.T) {
	// Construct the command string
	output := ExecuteCommand("go run . examples/example04.txt")
	inputFile := ExecuteCommand("cat examples/example04.txt")
	numOfTurns := (NumOfLines(output) - 1) - NumOfLines(inputFile) // (NumOfLines(output)-1) because of the new line after printing whats in the file
	maxNumOfTurns := 6

	// Compare the output with the solution content
	if numOfTurns > maxNumOfTurns {
		t.Errorf("Output does not match the solution.\nOutput:\n%s\nWith %s Number of turns\n\nMax Number of turns is: %s", LastLines(output, numOfTurns), strconv.Itoa(numOfTurns), strconv.Itoa(maxNumOfTurns))
	}
}

func TestNumOfTurnsSolution_ex05(t *testing.T) {
	// Construct the command string
	output := ExecuteCommand("go run . examples/example05.txt")
	inputFile := ExecuteCommand("cat examples/example05.txt")
	numOfTurns := (NumOfLines(output) - 1) - NumOfLines(inputFile) // (NumOfLines(output)-1) because of the new line after printing whats in the file
	maxNumOfTurns := 8

	// Compare the output with the solution content
	if numOfTurns > maxNumOfTurns {
		t.Errorf("Output does not match the solution.\nOutput:\n%s\nWith %s Number of turns\n\nMax Number of turns is: %s", LastLines(output, numOfTurns), strconv.Itoa(numOfTurns), strconv.Itoa(maxNumOfTurns))
	}
}

func ExecuteCommand(command string) string {
	// Run the command through the bash shell
	cmd := exec.Command("bash", "-c", command)

	// Capture the output
	output, err := cmd.Output()
	if err != nil {
		log.Fatalf("Error running command: %v", err)
	}

	return string(output)
}

func LastLines(output string, numLines int) string {
	// Split the output into individual lines
	lines := strings.Split(strings.TrimSpace(output), "\n")

	// Extract the last `numLines` lines
	startIndex := len(lines) - numLines
	if startIndex < 0 {
		startIndex = 0
	}
	lastLines := lines[startIndex:]

	// Join the lines back into a single string
	return strings.Join(lastLines, "\n")
}

func NumOfLines(output string) int {
	// Split the output into individual lines
	lines := strings.Split(strings.TrimSpace(output), "\n")

	return len(lines)
}

func TestBadExample_ex00(t *testing.T) {
	// Construct the command string for the bad example
	command := "go run . examples/badexample00.txt"

	// Execute the command and capture the error
	err := ExecuteCommandWithError(command)

	// Check if an error occurred
	if err == nil {
		t.Errorf("Expected an error, but no error occurred")
	}
}

func TestBadExample_ex01(t *testing.T) {
	// Construct the command string for the bad example
	command := "go run . examples/badexample01.txt"

	// Execute the command and capture the error
	err := ExecuteCommandWithError(command)

	// Check if an error occurred
	if err == nil {
		t.Errorf("Expected an error, but no error occurred")
	}
}

func ExecuteCommandWithError(command string) error {
	// Run the command through the bash shell
	cmd := exec.Command("bash", "-c", command)

	// Execute the command and capture the error
	err := cmd.Run()

	return err
}