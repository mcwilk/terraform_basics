resource "local_file" "example_resource1" {
  filename        = var.filename1
  content         = var.file_content1
  file_permission = "0640"
}

resource "local_file" "example_resource2" {
  filename = var.filename2
  content  = var.file_content2
}