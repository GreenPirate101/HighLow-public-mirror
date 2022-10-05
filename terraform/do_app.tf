resource "digitalocean_app" "highlow" {
  spec {
    name   = "highlow"
    region = local.do.region

    alert {
      rule = "DEPLOYMENT_FAILED"
    }
  }
}
