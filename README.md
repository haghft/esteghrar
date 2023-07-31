# Esteghrar - Golang version   
[![GoDoc](https://pkg.go.dev/badge/github.com/GeniusesGroup/esteghrar)](https://pkg.go.dev/github.com/GeniusesGroup/esteghrar)
[![Go Report](https://goreportcard.com/badge/github.com/GeniusesGroup/esteghrar)](https://goreportcard.com/report/github.com/GeniusesGroup/esteghrar)

Esteghrar is a software orchestrator based on [Memar](https://github.com/GeniusesGroup/memar) framework. Because we chose Unikernels(Exokernels) instead of Containers style, It will be a VM Orchestration not precisely like other tools (e.g. K8s, ...). Unlike other tools that need to operate and maintain separately from organization software(External observer), Esteghrar makes the organization software self-awareness(Internal observer). This will add more intelligence to any software that uses(embeds) Estghrar modules, without the need to use any AI models.

## Some resources
- https://kubernetes.io/docs/reference/kubernetes-api/ - https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.27/#node-v1-core
- https://mesosphere.github.io/marathon/api-console/index.html
- https://github.com/Yelp/paasta/blob/master/comparison.md
- https://docs.digitalocean.com/reference/api/api-reference/#tag/Images
- https://github.com/erikh/spin
