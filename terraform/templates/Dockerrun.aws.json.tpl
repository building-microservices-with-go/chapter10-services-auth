{
  "AWSEBDockerrunVersion": "1",
  "Image": {
    "Name": "${docker_image}:${docker_tag}",
    "Update": "true"
  },
  "Ports": [
    { "ContainerPort":"8080" }
  ]
}
