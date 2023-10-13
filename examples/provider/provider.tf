terraform {
  required_providers {
    singlestore = {
      source  = "Singlestore/singlestore"
      version = "0.1.1"
    }
  }
}

provider "singlestore" {
  # Configuration options
}