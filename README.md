# Kubeload


## What is it

This operator is a layer above kubernetes jobs that allows you to manage them in much easier and IAC-oriented way.
I came up with the idea when load-testing my site, because I manually had to increase the load every time.
this operator will let you configure your load-test initial load, max load, interval and hatch-rate, and most important - be able to reproduce the exact same load test over and over again
## Installation
Apply the CRD:
```console
kubectl apply -f https://raw.githubusercontent.com/Efrat19/kubeload/master/crd.yaml
```
## Usage
Customize your load manager:
```yaml
apiVersion: kubeload.kubeload.efrat19.io/v1
kind: LoadManager
metadata:
  name: loadmanager-sample
spec:
  loadSetup:
    # With how many pods the load should start:
    initialLoad: 2
    # How much time to wait before increasing the pod count:
    interval: 1m
    # How many pods should be added on each interval:
    hatchRate: 2
    # The manager will stop increasing the pods count once hitting maxLoad:
    maxLoad: 8
  selector:
    # Only jobs that has this set of labels will be controlled by this manager
    matchLabels:
      app: load-test
```
Once the above load manager will be applied, all jobs with the label `app: load-test` will be updated with pod count 2 - (the `initialRate`). than, every 1m (the `interval`), the loadmanager will increase the pod count by 2 more pods (the `hatchRate`), until reaching 8 pods (`maxLoad`).

## Annotations
Any time you can annotate controlled jobs with `kubeload.efrat19.io/freeze`. the load manager will not increase the pods count as long as this annotation is set to `true`
Example - to freeze a job:
```console
k annotate job load-test kubeload.efrat19.io/freeze="true"
```
To unfreeze-
```console
k annotate job load-test kubeload.efrat19.io/freeze-
```

## Examples
apply an example job and behold the results
```yaml
apiVersion: batch/v1
kind: Job
metadata:
  name: load-test
  namespace: kubeload
  labels:
    app: load-test
  # annotations:
  #  kubeload.efrat19.io/freeze: "true"
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
## Metrics
in addition to regular go metrics, the operator also exports a custom metric-
### kubeload_load_volume
- **Type:** Gauge
- **Description:** specifies pod count on a given time.
- **Labels:** job, namespace

## Roadmap (I welcome your contributions :star-struck: )

- [X] build CI
- [X] Documentation
- [X] Example
- [ ] Helm chart
- [X] Export Metrics
- [ ] Grafana Dashboard
- [ ] Tests

## Local Build
```console
make install
make run
``` 

## Built With
- [kubebuilder](https://book.kubebuilder.io/quick-start.html)









