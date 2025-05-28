# datadog-examples

This is a repository of various examples on setting up / configuring observability using Datadog, and also instrumenting cloud native applications.

A high overview of what to expect in each example category (these correspond to the directories) can be seen here:
- **ansible**: Configuration and script examples for setting up the Datadog Agent using Ansible automation.
- **containerization**: Examples of how to instrument containerized applications of different languages for APM and/or Logging.
- **ddagent-custom-checks**: Custom-developed Datadog Agent custom checks, for extending the agent to monitor non-standard objects.
- **docker/agent**: Datadog Agent deployment script and Docker compose deployment definition template for workloads running on Docker.
- **ecs-ec2**: Datadog Agent deployment defnition example for workloads running on AWS ECS EC2
- **ecs-fargate**:  Datadog Agent instrumentation example for applications running on AWS ECS Fargate.
- **eks-fargate**: Datadog Agent instrumentation example for applications running on AWS EKS Fargate.
- **integrations/nginx-ingress**: Custom check configuration to ensure the Datadog Agent collects NGINX Ingress metrics.
- **kubernetes**: Datadog Agent deployment defnition example for monitoring Kubernetes and applications running on Kubernetes.
- **utilities**: Personal scripts/tools to make life easier when using Datadog.

Some of these folders have their own README.md files, please refere to these where available.
