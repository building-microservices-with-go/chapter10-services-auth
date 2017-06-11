terraform {
  backend "s3" {
    bucket = "nicjackson-terraform-state"
    key    = "chapter11-auth.tfstate"
    region = "eu-west-1"
  }
}

provider "aws" {}

provider "datadog" {}
