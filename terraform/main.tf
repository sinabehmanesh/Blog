provider "helm" {
    kubernetes {
        config_path = var.config_path
    }
}

resource "helm_release" {
    name = "blog"
    repositroy = "oci://registry-1.docker.io/sinabehmanesh"
    chart = "blog-chart
}