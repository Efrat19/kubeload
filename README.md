# Kubeload

UNDER DEVELOPMENT

## What is it

This operator is a layer above kubernetes jobs that allows you to manage them in much easier and IAC-oriented way.
I came up with the idea when load-testing my site, because I manually had to increase the load every time
this operator will let you configure your loadtest's initial load, max load, interval and hatch-rate, and you will also be able to freeze or reset it any time.

## Install

1. Apply the CRD:
```console
kubectl apply -f https://raw.githubusercontent.com/Efrat19/kubeload/master/crd.yaml
```
2. Customize your load manager:
```yaml
apiVersion: kubeload.kubeload.efrat19.io/v1
kind: LoadManager
metadata:
  name: loadmanager-sample
spec:
  loadSetup:
    hatchRate: 1
    initialLoad: 2
    interval: 1m
    maxLoad: 8
  selector:
    matchLabels:
      app: load-test
```
3. OPTIONAL- Apply an example job: 
```yaml
apiVersion: batch/v1
kind: Job
metadata:
  name: load-test
  namespace: kubeload
  labels:
    app: load-test
  annotations:
    kubeload.efrat19.io/freeze: "true"
spec:
  parallelism: 1
  template:
    spec:
      containers:
      - name: load-test
        image: efrat19/locust-test:latest
        command: [ "locust","--host=https://www.zdnet.com","--no-web", "-c 1", "-r 1"]
      restartPolicy: Never
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
