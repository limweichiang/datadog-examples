# NGINX Ingress Integration

## Applying Datadog Autodiscovery(v2) annotations for NGINX Ingress Open Source

At its simplest, you can run the following command to deploy NGINX Ingress with the required annotations, which configures the Datadog Agent to collect Prometheus metrics from the ingress service.
```
% helm install <RELEASE_NAME> oci://ghcr.io/nginxinc/charts/nginx-ingress --version 1.0.2 -f nginx-ingress-opensource_autodiscovery-v2_helm-values.yml
```

In a real world scenario, you'd really want to merge the configs from `nginx-ingress-opensource_autodiscovery-v2_helm-values.yml` into your own NGINX helm `values.yaml` file. Take note that for security reasons, you absolutely must scope down `controller.nginxStatus.allowCidrs` to an allowable CIDR within your VPC or Kubernetes cluster.

## Applying Datadog Autodiscovery(v2) annotations for NGINX Ingress Plus

This is work in progress; Let me know if there is interest, and I might need some help to test, as I do not have access to a NGINX Ingress Plus license.
