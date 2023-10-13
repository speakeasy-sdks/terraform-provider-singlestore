terraform {
  required_providers {
    singlestore = {
      source  = "Singlestore/singlestore"
      version = "0.1.3"
    }
  }
}

provider "singlestore" {
  # Configuration options
}