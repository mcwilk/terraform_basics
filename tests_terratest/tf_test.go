package tests_terratest

import (
	"testing"

	"github.com/gruntwork-io/terratest/modules/terraform"
	"github.com/stretchr/testify/assert"
)

func TestExampleFail(t *testing.T) {
	result := 2 + 2

	if result != 5 {
		t.Error("Expected result: 5. Actual result:", result)
	}
}

func TestLocalFileCreation(t *testing.T) {
	// Make it possible to run tests in parallel
	t.Parallel()

	// Set directory where .tf files are stored
	terraformOptions := &terraform.Options {
		TerraformDir: "../",
	}

	// Tear-down (after tests it will destroy all the resources)
	defer terraform.Destroy(t, terraformOptions)

	// terraform init & terraform apply
	terraform.InitAndApply(t, terraformOptions)

	// Get outputs
	filePath1 := terraform.Output(t, terraformOptions, "file_path1")
	fileContent1 := terraform.Output(t, terraformOptions, "file_content1")

	// Assertions
	t.Log("This assertion will PASS")
	assert.Contains(t, filePath1, "sample_file1.txt")
	t.Log("This assertion will FAIL")
	assert.Equal(t, "Content of the file created by Terraform", fileContent1)
}