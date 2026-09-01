output "argocd_namespace" {
  description = "Namespace ArgoCD was installed into."
  value       = module.argocd.namespace
}
