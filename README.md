# Kubeload

UNDER DEVELOPMENT

## What is it

This operator is a layer above kubernetes jobs that allows you to manage them in much easier and IAC-oriented way.
I came up with the idea when load-testing my site, because I manually had to increase the load every time
this operator will let you configure your loadtest's initial load, max load, interval and hatch-rate, and you will also be able to freeze or reset it any time.

## Install

- Apply the CRD:
```console
kubectl apply -f 
```

## Local Build

```console
make install
make run
```
## TODO

- [X] build CI
- [ ] Documentation
- [ ] Example
- [ ] Helm chart
- [X] Export Metrics
- [ ] Grafana Dashboard
- [ ] Tests
 
## Built With
- [kubebuilder](https://book.kubebuilder.io/quick-start.html)
