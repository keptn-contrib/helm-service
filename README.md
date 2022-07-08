# Helm Service

![GitHub release (latest by date)](https://img.shields.io/github/v/release/keptn-contrib/helm-service)
[![Go Report Card](https://goreportcard.com/badge/github.com/keptn-contrib/helm-service)](https://goreportcard.com/report/github.com/keptn-contrib/helm-service)


The *helm-service* allows deploying services to a Kubernetes cluster and releasing them to user traffic using Istio.
Therefore, these services have to be packed as [Helm charts](https://helm.sh/docs/topics/charts/).
For details about the Helm chart and how to onboard a service, please checkout the [docs](https://keptn.sh/docs/0.15.x/manage/service/#onboard-a-service).

In order to deploy and release services to user-traffic, the *helm-service* implements two tasks:
1. Deployment task: Here, the `helm-service` executes a
Helm upgrade on the Helm chart provided by the user. Furthermore, the `helm-service` routes traffic to this new version.
1. Release task: Here, the `helm-service`
either promotes or rolls back the new version depending on the (evaluation) result.

## Installation

The *helm-service* is part of the *Execution Plane for Continuous Delivery*.

To install it next to your Keptn control-plane installation, you can use the following command:

```console
HELM_SERVICE_VERSION=1.0.0
HELM_SERVICE_NAMESPACE=keptn

helm upgrade --install -n ${HELM_SERVICE_NAMESPACE} helm-service \
  https://github.com/keptn-contrib/helm-service/releases/download/${HELM_SERVICE_VERSION}/helm-service-${HELM_SERVICE_VERSION}.tgz
```

### Istio

In order for blue-green and direct deployments to work properly, helm-service requires Istio installed on your Kubernetes cluster.

Please follow the [Official Istio Getting Started Guide](https://istio.io/latest/docs/setup/getting-started/), or alternatively execute the following commands:

```bash
curl -L https://istio.io/downloadIstio | ISTIO_VERSION=1.14.1 sh -
./istio-1.14.1/bin/istioctl install
```

## Usage

Please read up about multi-stage delivery within the Keptn docs, e.g.:

* [Multi stage delivery](https://keptn.sh/docs/0.17.x/continuous_delivery/multi_stage/)
* [Deployment using helm](https://keptn.sh/docs/0.17.x/continuous_delivery/deployment_helm/)


## Development

You can use `skaffold run --tail` to build and deploy from this directory.

## Handled events
The *helm-service* handles a set of events. The following sequence diagrams describe the respectively executed actions
and the involved components.

### Handling of `sh.keptn.event.service.delete.finished` events
The `sh.keptn.event.service.delete.finished` event states that a Keptn service has been deleted by the `shipyard-controller`.
In case this service was deployed by the `helm-service`, the `helm-service` uninstalls all releases of this Keptn service.

![](sequence_diagrams/service-deleted.png)

### Handling of `sh.keptn.event.deployment.triggered` events
The `sh.keptn.event.deployment.triggered` event states that a new deployment has been triggered e.g. by the user.
The `helm-service` executes a Helm upgrade on the Helm chart provided by the user, i.e. the `user-chart`
and routes traffic to this new version.

![](./sequence_diagrams/deployment-triggered.png)

### Handling of `sh.keptn.event.release.triggered` events
The `sh.keptn.event.release.triggered` event states that a release has been triggered.

For a direct deployment, the `helm-service` does not have to apply anything.

![](./sequence_diagrams/release-triggered-direct.png)

For a b/g deployment with an (evaluation) result equals pass or warning, the `helm-service` promotes the new version
to be stable.

![](./sequence_diagrams/release-triggered-bg-promote.png)

For a b/g deployment with an (evaluation) result equals fail, the `helm-service` rolls back the new version.

![](./sequence_diagrams/release-triggered-bg-rollback.png)


### Handling of `sh.keptn.event.action.triggered` events
The `sh.keptn.event.action.triggered` event stats that a remediation action has been triggered.
The `helm-service` provides a replica scaling remediation action.
![](./sequence_diagrams/action-triggered.png)

## License

Please find more information in the [LICENSE](LICENSE) file.
