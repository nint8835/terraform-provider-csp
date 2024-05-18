data "csp_example" "example" {
  directive {
    name     = "default-src"
    keywords = ["self"]
  }

  directive {
    name     = "img-src"
    keywords = ["self"]
    hosts    = ["cdn.example.com"]
  }
}
