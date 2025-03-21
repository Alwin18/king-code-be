package services

import (
	"bytes"
	"fmt"
	"os"

	"os/exec"
	"strings"

	"github.com/Alwin18/king-code/models"
	"github.com/Alwin18/king-code/repositories"
)

type SubmissionService struct {
	submissionRepo *repositories.SubmissionRepository
	challengeRepo  *repositories.ChallengeRepository
}

func NewSubmissionService(subRepo *repositories.SubmissionRepository, chRepo *repositories.ChallengeRepository) *SubmissionService {
	return &SubmissionService{subRepo, chRepo}
}

// Evaluasi jawaban user
func (s *SubmissionService) EvaluateSubmission(userID, challengeID, code, language string) (*models.UserSubmission, error) {
	challenge, err := s.challengeRepo.GetChallengeByID(challengeID)
	if err != nil {
		return nil, fmt.Errorf("Challenge tidak ditemukan")
	}

	// Eksekusi kode dalam sandbox
	correct := true
	for _, testCase := range challenge.TestCases {
		output, err := executeCode(code, testCase.Input, language)
		if err != nil || strings.TrimSpace(output) != strings.TrimSpace(testCase.ExpectedOutput) {
			correct = false
			break
		}
	}

	// Simpan hasil submission
	status := "correct"
	score := 100
	if !correct {
		status = "incorrect"
		score = 0
	}

	submission := &models.UserSubmission{
		UserID:      userID,
		ChallengeID: challengeID,
		Code:        code,
		Language:    language,
		Status:      status,
		Score:       score,
	}
	err = s.submissionRepo.CreateSubmission(submission)
	if err != nil {
		return nil, err
	}
	return submission, nil
}

// Fungsi untuk mengeksekusi kode di dalam sistem secara aman
func executeCode(code, input, language string) (string, error) {
	var cmd *exec.Cmd

	switch language {
	case "python":
		cmd = exec.Command("python3", "-c", code)

	case "javascript":
		cmd = exec.Command("node", "-e", code)

	case "go":
		tmpFile := "/tmp/main.go"
		err := os.WriteFile(tmpFile, []byte(code), 0644)
		if err != nil {
			return "", err
		}
		cmd = exec.Command("go", "run", tmpFile)
	case "java":
		tmpFile := "/tmp/Main.java"
		err := os.WriteFile(tmpFile, []byte(code), 0644)
		if err != nil {
			return "", err
		}
		compileCmd := exec.Command("javac", tmpFile)
		if err := compileCmd.Run(); err != nil {
			return "", err
		}
		cmd = exec.Command("java", "-cp", "/tmp", "Main")

	default:
		return "", fmt.Errorf("Bahasa tidak didukung")
	}

	// Masukkan input ke dalam program
	cmd.Stdin = strings.NewReader(input)

	// Tangkap output
	var out bytes.Buffer
	var stderr bytes.Buffer
	cmd.Stdout = &out
	cmd.Stderr = &stderr

	err := cmd.Run()
	if err != nil {
		return "", fmt.Errorf("execution error: %s, stderr: %s", err.Error(), stderr.String())
	}

	// Trim hasil output agar tidak ada karakter tambahan seperti newline
	result := strings.TrimSpace(out.String())

	return result, nil
}
