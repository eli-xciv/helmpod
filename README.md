# helmpod

A utility for using `podman kube play`  with helm charts. This work is inspired by [this issue](https://github.com/containers/podman/issues/15098). The goal of this project is for `helmpod` to be a viable drop-in replacement (if you so desire) for your helm commands when running against podman,  effectively allowing you `helm install` onto podman.

# How it works

At it's core, this is a wrapper for helm and podman. `podman` has the capability to `run` k8s (some) objects. `helmpod` will essentially run a `helm template`, and then iterate through the objects and try to create objects that fall in line with what podman supports and then creates a `podman` ready set of k8s objects to then execute `podman kube-play` on.

# Required Dependecies
- helm
- podman

# Supported k8s Objects
The following kind, or k8s object list is supported by `podman` per [their docs](https://docs.podman.io/en/latest/markdown/podman-kube-play.1.html)

- Pod
- Deployment
- PersistentVolumeClaim
- ConfigMap
- Secret

# Installation

TODO

# How to use

Simplest way to get going is to run your `helm` commands as you would, but with `helmpod` instead.

e.g. 
```
# Typical helm install command (Deploys to k8s)
helm install my-fancy-deployment-name bitnami/kubeapps

# helmpod (Deploys to podman)
helmpod install my-fancy-deployment-nname bitnami/kubeapps

```

# Contributing
Contributions are welcome. Please fork this project, create your changes and then submit a PR with your changes.
