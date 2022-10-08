terraform {
  required_providers {
    digitalocean = {
      source  = "digitalocean/digitalocean"
      version = "~> 2.0"
    }
  }

  cloud {
    organization = "highlow"

    workspaces {
      name = "do-prod-blr1"
    }
  }
}
