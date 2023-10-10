required_providers "helm" {
    kubernetes {
        config_path = var.config_path
    }
}