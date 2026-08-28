resource "helm_release" "argocd" {
  name             = "argocd"
  repository       = "https://argoproj.github.io/argo-helm"
  chart            = "argo-cd"
  version          = var.chart_version
  namespace        = var.namespace
  create_namespace = true

  # Deliberately minimal otherwise — default chart values. HA, ingress,
  # and SSO are hardening/scale concerns for later, not needed to prove
  # GitOps actually works end-to-end.
  values = [
    yamlencode({
      controller = {
        # The chart sets no resources by default. On Fargate that's not
        # "unbounded" — the Fargate scheduler sizes the pod itself, and
        # with nothing requested it lands on a small default that's fine
        # for reconciling a handful of small apps but OOMKills the
        # controller once an app's rendered manifest gets large (e.g.
        # kube-prometheus-stack, ~140 objects incl. its CRDs). Sizing
        # this explicitly is what actually avoids that, not anything
        # ArgoCD-side.
        resources = {
          requests = {
            cpu    = "500m"
            memory = "1Gi"
          }
          limits = {
            cpu    = "1"
            memory = "2Gi"
          }
        }
      }
    })
  ]
}