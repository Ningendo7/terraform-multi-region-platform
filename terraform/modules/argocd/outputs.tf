output "namespace" {
  description = "Namespace ArgoCD was installed into."
  value       = helm_release.argocd.namespace
}

output "release_name" {
  description = "Helm release name."
  value       = helm_release.argocd.name
}
