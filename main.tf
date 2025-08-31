resource "local_file" "example_resource1" {
  filename        = var.filename1
  content         = var.file_content1
  file_permission = "0640"
}

data "local_file" "existing_data" {
  filename = "${path.module}/infra/resources/existing_file.txt"
}

resource "local_file" "example_resource2" {
  filename = var.filename2
  content  = data.local_file.existing_data.content
}