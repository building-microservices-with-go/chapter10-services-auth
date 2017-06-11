# Create a new Datadog timeboard
resource "datadog_timeboard" "auth" {
  title       = "Auth service Timeboard (created via Terraform)"
  description = "created using the Datadog provider in Terraform"
  read_only   = true

  graph {
    title = "Authentication"
    viz   = "timeseries"

    request {
      q    = "sum:chapter11.auth.jwt.badrequest{*}"
      type = "bars"

      style {
        palette = "warm"
      }
    }

    request {
      q    = "sum:chapter11.auth.jwt.success{*}"
      type = "bars"
    }
  }

  graph {
    title = "Health Check"
    viz   = "timeseries"

    request {
      q    = "sum:chapter11.auth.health.success{*}"
      type = "bars"
    }
  }
}
