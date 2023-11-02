provider "helm" {
    kubernetes {
        config_path = var.config_path
    }
}

resource "helm_release" "blog" {
    name = "blog"
#    repository = "oci://registry-1.docker.io/sinabehmanesh"
    chart = "../blog-chart"
    version           = "0.1.1"
}