variable "filename1" {
  type        = string
  default     = "sample_file1.txt"
  description = "1st file created by Terraform"
}

variable "filename2" {
  type        = string
  default     = "./infra/resources/sample_file2.txt"
  description = "2nd file created by Terraform & based on existing file"
}

variable "file_content1" {
  type        = string
  default     = "Default file content 1"
  description = "File content"
}

variable "file_content2" {
  type        = string
  default     = "Default file content 2"
  description = "File content"
}