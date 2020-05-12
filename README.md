# CAKE

! NOTE: This project is currently under heavy early development and is a work in progress.  Things can and do change quickly and drastically.  For now, one should not expect the workflow or interfaces to be stable, nor should one expect the deployments to be fully functional.  This note will be updated soon, when things are stable.

## What is the cake binary

The cake binary is utility written in golang to automate bootstrapping a CAPv or RKE management cluster.

## Primary Focus

Provide a mechanism for easy, automated deployment of a fully functional and officially supported Rancher installation on both vSphere and Bare Metal Linux environments.

Cake is a utility, it's designed to be useful utility to make the task of deploying Rancher just a bit easier.  As a utility it's expected to be modular, have minimal dependencies and be able to run in multiple envionrments and models:

* Remote from Multiple OS Platforms (Mac OS, Linux and even Windows)
* Local as a bootstrapper on a linux machine that will eventually be part of the Rancher cluster

## Principals

* Easy Deploy
* Fast Deploy
* Fully functional Rancher deployment including Rancher UI and Kubernets API access
* Modular Building Block
* Upgradable

[![Test and Build](https://github.com/NetApp/cake/workflows/Test%20and%20Build/badge.svg)](https://github.com/NetApp/cake/actions?query=workflow%3A%22Test+and+Build%22)
[![Go Report](https://goreportcard.com/badge/github.com/netapp/cake)](https://goreportcard.com/report/github.com/netapp/cake)

Kubernetes bootstrap is a piece of cake!
Or "Cloud Adjacent Kubernetes Engine", is a simple tool used to deploy on-prem kubernetes and Rancher clusters.

## How to use it

### Install

Fetch the latest binary release for your platform from the projects Github Release page

### genconfig

`cake genconfig` 

Takes user input and builds a config.yaml file that includes your VSphere endpoint credentials, and options
for extra items to install.

### deploy

`cake capv-deploy --config myconfig.yaml`

Will deploy a management cluster on the specified VSphere cluster or if the `--config` option is omitted, then the
tool will interactively create a config and initiate the deployment.

### destroy

`cake destroy --cluster-id xxx` will destroy the cluster of the given id if it exists.
