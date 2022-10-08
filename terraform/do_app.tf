resource "digitalocean_app" "highlow" {
  spec {
    name   = "highlow"
    region = local.do.region

    alert {
      rule = "DEPLOYMENT_FAILED"
    }


    function {
      name = "highlow-functions"
      github {
        repo           = "CyberCitizen01/HighLow"
        branch         = "main"
        deploy_on_push = false
      }
    }
  }
}

output "do_app" {
  value = {
    id              = digitalocean_app.highlow.id
    default_ingress = digitalocean_app.highlow.default_ingress
    live_url        = digitalocean_app.highlow.live_url
  }
}
