# Our Life Before Corona

[![Docker Repository on Quay](https://quay.io/repository/marcelmue/our-life-before-corona/status "Docker Repository on Quay")](https://quay.io/repository/marcelmue/our-life-before-corona)
[![Actions Status](https://github.com/traaction/our-life-before-corona/workflows/Go/badge.svg)](https://github.com/traaction/our-life-before-corona/actions)

## Setup with Frontend
- Download [KIND](https://github.com/kubernetes-sigs/kind#installation-and-usage) (easiest install with `curl`)
- Create cluster with `kind create cluster`
- Apply all manifests `kubectl apply -f kubernetes`
- Start port forwarding api to localhost `kubectl port-forward service/our-life-before-corona 8080:8080`
- Hit `http://localhost:8080/dev/init` **once** from your browser to load in seed data
- Point your frontend to `localhost:8080` for the api!
- To remove everything from the cluster simply run `kubectl delete -f kubernetes`

## Coding Setup

This project should be cloned in the go path like this:
- My go path `/home/marcel/golang`
- Full folder path to this readme `/home/marcel/golang/src/gitlab/wirvsvirus/our-life-before-corona`

### General interaction:
- Run this app in your cmd with `go run main.go`
- Build a binary of this app with `go build`
- Add a new third party dependency with `go get -u dependencyname` follow by `go mod tidy`

### Kubernetes Database:
- Download [KIND](https://github.com/kubernetes-sigs/kind#installation-and-usage) (easiest install with `curl`)
- Create cluster with `kind create cluster`
- Apply database manifest `kubectl apply -f kubernetes/database.yaml`
- Start port forwarding db to localhost `kubectl port-forward service/postgres 5432:5432`
- Build locally `go build`
- Test locally with `./our-life-beforecorona`
