package main

import (
	"os"
	"testing"
)

func TestGenerateJobManifestFile(t *testing.T) {
	tests := []struct {
		name       string
		job        JobManifest
		filePath   string
		expectFail bool
	}{
		{   // Test case 1: vaild inputs
			name: "Valid inputs",
			job: JobManifest{
				Name:      "client-monitor-job",
				Namespace: "default",
				Action:    "processes",
				PodName:   "deploy-dotnet-monitor-77c45ff464-dsgr4",
				PID:       "NO_PID",
				UID:       "NO_UID",
				NameOverride: "NO_NAME",
				Duration:  "NO_DURATION",
				Egress:    "NO_EGRESS_PROVIDER",
				Tags:      "NO_TAG",
			},
			filePath:   "test-job1.yml",
			expectFail: false,
		},
		{   // Test case 2: invalid inputs with empty string
			name: "Invalid inputs with empty name",
			job: JobManifest{
				Name:      "",
				Namespace: "default",
				Action:    "processes",
				PodName:   "deploy-dotnet-monitor-77c45ff464-dsgr4",
				PID:       "NO_PID",
				UID:       "NO_UID",
				NameOverride: "NO_NAME",
				Duration:  "NO_DURATION",
				Egress:    "NO_EGRESS_PROVIDER",
				Tags:      "NO_TAG",
			},
			filePath:   "test-job2.yml",
			expectFail: true,
		},
		{   // Test case 3: invalid input
			name: "Invalid inputs with empty pod name",
			job: JobManifest{
				Name:      "client-monitor-job",
				Namespace: "default",
				Action:    "processes",
				PodName:   "",
				PID:       "NO_PID",
				UID:       "NO_UID",
				NameOverride: "NO_NAME",
				Duration:  "NO_DURATION",
				Egress:    "NO_EGRESS_PROVIDER",
				Tags:      "NO_TAG",
			},
			filePath:   "test-job3.yml",
			expectFail: true,
		},
		{   // Test case 4: invalid file path
			name: "invalid file path",
			job: JobManifest{
				Name:      "client-monitor-job",
				Namespace: "default",
				Action:    "processes",
				PodName:   "deploy-dotnet-monitor-77c45ff464-dsgr4",
				PID:       "NO_PID",
				UID:       "NO_UID",
				NameOverride: "NO_NAME",
				Duration:  "NO_DURATION",
				Egress:    "NO_EGRESS_PROVIDER",
				Tags:      "NO_TAG",
			},
			filePath:   "/test-folder/test-job4.yml",
			expectFail: true,
		},
	}

	for _, tc := range tests {
		err := generateJobManifestFile(tc.job, tc.filePath)
		defer os.Remove(tc.filePath)

		if tc.expectFail && err == nil {
			t.Errorf("%s: Expected error but got none", tc.name)
		} else if !tc.expectFail && err != nil {
			t.Errorf("%s: Unexpected error: %v", tc.name, err)
		}
	}
}

func TestMain(m *testing.M) {

	// Run the tests
	exitCode := m.Run()

	// Exit with the proper exit code
	os.Exit(exitCode)
}
