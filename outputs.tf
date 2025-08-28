output "file_path1" {
  value = local_file.example_resource1.filename
}

output "file_path2" {
  value = local_file.example_resource2.filename
}

output "file_content1" {
  value = local_file.example_resource1.content
}

output "file_content2" {
  value = local_file.example_resource2.content
}