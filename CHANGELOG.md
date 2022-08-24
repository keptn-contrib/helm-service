# Changelog

All notable changes to this project will be documented in this file. See [standard-version](https://github.com/conventional-changelog/standard-version) for commit guidelines.

## [0.18.0](https://github.com/keptn-contrib/helm-service/compare/0.17.0...0.18.0) (2022-08-24)


### ⚠ BREAKING CHANGES

* All Keptn core now depends on resource-service. From this moment on resource-service installation is mandatory.

### Features

* **installer:** Add options for setting image repository and tag globally ([#8152](https://github.com/keptn-contrib/helm-service/issues/8152)) ([deeec1d](https://github.com/keptn-contrib/helm-service/commit/deeec1d132772dd13923de2a1ac5ac15f3e2cf52))
* Repository migration ([#2](https://github.com/keptn-contrib/helm-service/issues/2)) ([5c7abf9](https://github.com/keptn-contrib/helm-service/commit/5c7abf98936417eae4b1060db54948c809fc9629))


### Bug Fixes

* log.Fatal will call os.Exit, use log.Println instead ([#8492](https://github.com/keptn-contrib/helm-service/issues/8492)) ([981cdc0](https://github.com/keptn-contrib/helm-service/commit/981cdc09c7a0b815a9444fb032bb76f292398f6c))


### Refactoring

* Refactor all services to use resource-service ([#8271](https://github.com/keptn-contrib/helm-service/issues/8271)) ([819963c](https://github.com/keptn-contrib/helm-service/commit/819963c2118e2a06dcaf132af5840d95d8a8630a))


### Other

* Add helm dependencies directly to repository charts ([#8472](https://github.com/keptn-contrib/helm-service/issues/8472)) ([d859d6a](https://github.com/keptn-contrib/helm-service/commit/d859d6a99e1db8056d94113637f26878e548f17f))
* **installer:** Improve NATS configuration ([#8475](https://github.com/keptn-contrib/helm-service/issues/8475)) ([57a949d](https://github.com/keptn-contrib/helm-service/commit/57a949d83a3f07fae18c0ee9dc71c210b21640b3))
