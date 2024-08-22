//
// MyHospital Azure infrastructure deploy script
//
// Copyright(C)2022 plus-medi-corp. All Rights Reserved.
//

// must parameter(s)
param environmentId string

param name string

@allowed([
  'multiple'
  'single'
])
param revisionMode string = 'single'

@allowed([
  'auto'
  'http'
  'http2'
])
param transport string = 'auto'

param allowInsecure bool = false
param isExternalIngress bool = false

#disable-next-line secure-secrets-in-params
param secrets array = []
param registries array = []
param env array = []
param rules array = []
param containerImage string
param containerPort int
param minReplicas int = 1
param maxReplicas int = 1
param cpu string = '0.25'
param memory string = '0.5Gi'

@description('target azure region')
param location string = resourceGroup().location

param identity object = {}

param tags object = {}

resource acaApp 'Microsoft.App/containerApps@2022-03-01' = {
  name: name
  location: location
  tags: tags
  identity:identity
  properties: {
    managedEnvironmentId: environmentId
    configuration: {
      activeRevisionsMode: revisionMode
      secrets: secrets
      registries: registries

      ingress: {
        external: isExternalIngress
        targetPort: containerPort
        transport: transport
        allowInsecure: allowInsecure
      }
    }
    template: {
      containers: [
        {
          image: containerImage
          name: name
          env: env
          resources: {
            cpu: any(cpu)
            memory: memory
          }
        }
      ]
      scale: {
        minReplicas: minReplicas
        maxReplicas: maxReplicas
        rules: rules
      }
    }
  }
}

