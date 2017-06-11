output "auth_alb" {
  value = "${aws_elastic_beanstalk_environment.default.cname}"
}
