# Provider Vultr

`provider-jet-vultr` is a [Crossplane](https://crossplane.io/) provider that
is built using [Upjet](https://github.com/upbound/upjet) code
generation tools and exposes XRM-conformant managed resources for the
Vultr API.

## Getting Started

Very quick overview on how to get up and running.

### Install Crossplane

Get crossplane installed by doing something like the following:

```
kubectl create namespace crossplane-system
helm install crossplane --namespace crossplane-system crossplane-stable/crossplane
```
### Install the provider crds

Install the provider crds:

```
kubectl apply -f provider/crds
```

### Run the provider

Just run the provider by issuing:

````
make run
````

### Install provider config

First, create a `secret.yaml` file (using the `secret.yaml.tmpl` as a template)

Then just apply the whole directory:

```
kubectl apply -f examples/providerconfig/
```

### Now create a resource

Example is within the instance directory:

```
kubectl apply -f examples/instance/instance.yaml
```
