# Nuleculization

This document describes a proof of concept to generate most of a Nulecule file
with an automated process.

## Approach

We will enrich a RHEL7 Dockerfile, so that all environment variables that are
required or optional and all volumes that are required are present as LABELs.

Using these labels we will generate a Nulecule file by mapping the lables
to Nulecule objects. This will be done by Dockerfile2Nulecule.

As a result of this POC a

## Hacking on it

```
cd hack
go get github.com/docker/docker/builder/dockerfile/parser
go build
cd ../9.4/
../hack/hack
```
