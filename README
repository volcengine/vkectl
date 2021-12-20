<h1 align="center">VKE CLI</h1>

## Introduction

The VKE(VolcanoEngine Kubernetes Engine) CLI is a tool to manage VKE resources through a command line interface. It is written in Go and built on the top of VKE OpenAPI.

## Usage
### Set Environment
Before using vkectl, you need to set the values of the following environment variables
``` 
export AK=YOUR AK
export SK=YOUR SK
export HOST=YOUR HOST
export REGION=YOUR REGION
```

### Execute
``` 
vkectl [OPTIONS] MODULE ACTION
```
The supported modules are app, resource and security.

You can use the following commands to get the supported actions.

``` 
vkectl MODULE --help
```

Examples:
```
vkectl -d '{}' resource ListClusters
```