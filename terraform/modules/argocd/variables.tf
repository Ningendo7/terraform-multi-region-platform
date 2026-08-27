variable "namespace" {
  description = "Kubernetes namespace to install ArgoCD into."
  type        = string
  default     = "argocd"
}

variable "chart_version" {
  description = "ArgoCD Helm chart version — verify against the current release before bumping blindly."
  type        = string
  default     = "7.7.11"
}