package tests_terratest

import {
	"testing"
	"github.com/gruntwork-io/terratest/modules/terraform"
	"github.com/stretchr/testify/assert"
}

func TestSum(t *testing.T) {
	result := 2 + 2

	if result != 4 {
		t.Error("Expected result: 4. Actual result: %d", result)
	}
}

func TestLocalFileCreation(t *testing.T) {
	
	t.Parallel()

	// Set directory where .tf files are stored
	terraformOptions := &terraform.Options {
		TerraformDir: "../",
	}

	// Tear-down
	defer terraform.Destroy(t, terraformOptions)

	// init and apply
	terraform.InitAndApply(t, terraformOptions)

	// Get outputs
	filePath1 := terraform.outputs(t, terraformOptions, "file_path1")
	fileContent1 := terraform.outputs(t, terraformOptions, "file_content1")

	// Assertions
	assert.Contains(t, file_path1, "file_path1")
	assert.Equal(t, "Content of the file created by Terraform (overridden default value)", fileContent1)
}