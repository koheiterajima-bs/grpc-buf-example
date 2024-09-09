//
// MyHospital Azure infrastructure deploy script
//
// Copyright(C)2022 plus-medi-corp. All Rights Reserved.
//

targetScope = 'subscription'

// must parameter(s)
@description('name for azure container registry')
param acrName string

//@description('resource group name for azure container registry')
//param acrRgName string

@description('resource group name for azure container apps')
param acaRgName string

@description('name for azure container apps environment')
param envName string

@description('name for managed id')
param managedIdName string

@description('container image name (without server name)')
param imageName string

@description('container apps name')
param name string

// optional parameter(s)
param location string = deployment().location

// variables
var containerImage = '${acrName}.azurecr.io/${imageName}'
var containerPort = 51051
var minReplicas = 1
var maxReplicas = 1
var cpu = '0.25'
var memory = '0.5Gi'
var transport = 'http2'
var tags = {
  description: 'patientapi'
}

var identity = {
  type: 'UserAssigned'
  userAssignedIdentities: {
    '${userManagedId.id}': {}
  }
}

var registries = [
  {
    server: '${acrName}.azurecr.io'
    identity: userManagedId.id
  }
]

var secrets = [
]

var env = json(loadTextContent('env.json'))

var rules = [
  {
    name: 'http-scale'
    http: {
      metadata: {
        concurrentRequests: '100'
      }
    }
  }
]

module acaAppContainerModule 'modules/aca.bicep' = {
  scope: resourceGroup(acaRgName)
  name: 'acaAppContainerModule'
  params: {
    location: location
    name: name
    containerImage: containerImage
    containerPort: containerPort
    environmentId: acaEnvironment.id
    isExternalIngress: true
    minReplicas: minReplicas
    maxReplicas: maxReplicas
    transport: transport
    allowInsecure: false
    env: env
    secrets: secrets
    rules: rules
    registries: registries
    cpu: cpu
    memory: memory
    identity: identity
    tags: tags
  }
}

// get existing resources
resource acaEnvironment 'Microsoft.App/managedEnvironments@2022-01-01-preview' existing = {
  name: envName
  scope: resourceGroup(acaRgName)
}

resource userManagedId 'Microsoft.ManagedIdentity/userAssignedIdentities@2021-09-30-preview' existing = {
  name: managedIdName
  scope: resourceGroup(acaRgName)
}

