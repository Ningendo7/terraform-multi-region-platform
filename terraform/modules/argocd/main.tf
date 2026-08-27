resource "helm_release" "argocd" {
  name             = "argocd"
  repository       = "https://argoproj.github.io/argo-helm"
  chart            = "argo-cd"
  version          = var.chart_version
  namespace        = var.namespace
  create_namespace = true

  # Deliberately minimal — default chart values. HA, ingress, and SSO
  # are hardening/scale concerns for later, not needed to prove GitOps
  # actually works end-to-end.
}