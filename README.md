# Argo CD gitops demo with Platform9 K8s

A CI/CD pipeline utilizing DevOps best practices with the following features:

- Continuous Integration checks on new pull requests
- Docker image builds only on new application code changes
- Automatic application [Semantic Version](https://semver.org/) management for repository and Docker image tagging
- [Continuous Delivery](https://martinfowler.com/bliki/ContinuousDelivery.html) of new application versions into the Development environment
- Deterministic control over application version promotion through Staging and Production environments
- Environment customization of application configuration via [kustomize](https://kustomize.io/)
- Horizontal POD Autoscaling

# Trigger HPA

Export your `KUBECONFIG` environment variable to the correct cluster kubeconfig.yaml

```
make load-test
```

