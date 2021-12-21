<h1 align="center">vkectl</h1>

## Introduction

vkectl is a tool to manage VKE(VolcanoEngine Kubernetes Engine) resources through a CLI(Command Line Interface). It is written in Go and built on the top of VKE OpenAPI.

VKECTL IS EXPERIMENTAL. THE BEHAVIOR MAY CHANGE. 

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
The supported modules are resource and security.

You can use the following commands to get the supported actions.

``` 
vkectl MODULE --help
```

Examples:
```
vkectl resource ListClusters
```