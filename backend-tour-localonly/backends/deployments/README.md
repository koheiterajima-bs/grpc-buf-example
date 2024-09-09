# `/deployments`

以下の作業には、az cliが必要です。`moc/patient/scripts/install-azure-cli.sh` にubuntu用のscriptがあります。

## ログイン

最初に、AzureとAzure Container Registry(ACR)にログインする。

```sh
$ az login --use-device-code --tenant devmykarte.onmicrosoft.com
$ az acr login -n crmhv25ua3ddectf7dg.azurecr.io
```

az loginの先はあまり変更されないが、az acr loginの部分は随時変わる。
ここでは、.envrcに、`export KO_DOCKER_REPO=crmhv25ua3ddectf7dg.azurecr.io` として定義しているので以下は、それを使う。

## ビルドとACRへの登録

Makefileに、koを使ったローカルビルド`make build`と、ACR向けのビルド`make publish`のターゲットがある。

### ローカルビルド

`make build`でローカルビルドする。`ko.local/patientapis`のようなイメージができる。

```sh
$ make build
env SOURCE_DATE_EPOCH=$(git log -1 --format="%ct") ko build --local ../cmd/patientapis
2022/08/25 09:48:49 Using base gcr.io/distroless/static:nonroot@sha256:1f580b0a1922c3e54ae15b0758b5747b260bd99d39d40c2edb3e7f6e2452298b for github.com/plusmedi/mhv2-backends/moc/patient/cmd/patientapis
2022/08/25 09:48:52 Building github.com/plusmedi/mhv2-backends/moc/patient/cmd/patientapis for linux/amd64
2022/08/25 09:48:53 Loading ko.local/patientapis-6d8bb96cb4919f922ff1c3d8cb1e1f1e:06b0cf003b0c401637cb5d490df0df3d49507529aece3eb58c864bee948d48ed
2022/08/25 09:48:55 Loaded ko.local/patientapis-6d8bb96cb4919f922ff1c3d8cb1e1f1e:06b0cf003b0c401637cb5d490df0df3d49507529aece3eb58c864bee948d48ed
2022/08/25 09:48:55 Adding tag latest
2022/08/25 09:48:56 Added tag latest
ko.local/patientapis-6d8bb96cb4919f922ff1c3d8cb1e1f1e:06b0cf003b0c401637cb5d490df0df3d49507529aece3eb58c864bee948d48ed

$ docker images | grep patientapis
ko.local/patientapis-6d8bb96cb4919f922ff1c3d8cb1e1f1e   06b0cf003b0c401637cb5d490df0df3d49507529aece3eb58c864bee948d48ed   3f5edbd50b03   59 minutes ago   20.1MB
ko.local/patientapis-6d8bb96cb4919f922ff1c3d8cb1e1f1e   latest                                                             3f5edbd50b03   59 minutes ago   20.1MB
ko.local/patientapis-6d8bb96cb4919f922ff1c3d8cb1e1f1e   18f1d73802a72aebc55d4d25577ab758b70f9a86bd8312a8bb6aa460cf917264   3f1731b22d74   52 years ago     20.1MB
takekazu@m900:~/ghq/github.com/plusmedi/mhv2-backends/moc/patient/deployments$
```

### ACRへの登録

`make publish`では、ko buildして、acrにpushする。

```sh
$ make publish
env SOURCE_DATE_EPOCH=$(git log -1 --format="%ct") KO_DOCKER_REPO=crmhv25ua3ddectf7dg.azurecr.io ko build ../cmd/patientapis > .ko_name
2022/08/25 09:51:03 Using base gcr.io/distroless/static:nonroot@sha256:1f580b0a1922c3e54ae15b0758b5747b260bd99d39d40c2edb3e7f6e2452298b for github.com/plusmedi/mhv2-backends/moc/patient/cmd/patientapis
2022/08/25 09:51:07 Building github.com/plusmedi/mhv2-backends/moc/patient/cmd/patientapis for linux/amd64
2022/08/25 09:51:08 Publishing crmhv25ua3ddectf7dg.azurecr.io/patientapis-6d8bb96cb4919f922ff1c3d8cb1e1f1e:latest
2022/08/25 09:51:13 existing blob: sha256:33cf526771a470bc3e74f34ceff09b5834c2722b92a89874f4000efa14c3ec5b
2022/08/25 09:51:13 existing blob: sha256:9b6d3920e5ad3f7d07ec6f3ca4b6cad9de403fe402dad26cc54f290395fb245f
2022/08/25 09:51:14 crmhv25ua3ddectf7dg.azurecr.io/patientapis-6d8bb96cb4919f922ff1c3d8cb1e1f1e:sha256-06b0cf003b0c401637cb5d490df0df3d49507529aece3eb58c864bee948d48ed.sbom: digest: sha256:e0c50ac2f15e6b3ec6040c88a2baebbde35e817b37f12123c514f89b9a69ff51 size: 369
2022/08/25 09:51:14 Published SBOM crmhv25ua3ddectf7dg.azurecr.io/patientapis-6d8bb96cb4919f922ff1c3d8cb1e1f1e:sha256-06b0cf003b0c401637cb5d490df0df3d49507529aece3eb58c864bee948d48ed.sbom
2022/08/25 09:51:18 existing blob: sha256:3f5edbd50b03de130088bc3ce0a009e24cd35b8404977c35bdbeda51cb1c190d
2022/08/25 09:51:18 existing blob: sha256:db1366fb18b6eb8092748e2f88f178613489defab496f48c6a10a05436b60dd2
2022/08/25 09:51:18 existing blob: sha256:0a602d5f6ca3de9b0e0d4d64e8857e504ec7a8c47f1ec617d82a81f6c64b0fe8
2022/08/25 09:51:18 existing blob: sha256:250c06f7c38e52dc77e5c7586c3e40280dc7ff9bb9007c396e06d96736cf8542
2022/08/25 09:51:19 crmhv25ua3ddectf7dg.azurecr.io/patientapis-6d8bb96cb4919f922ff1c3d8cb1e1f1e:latest: digest: sha256:06b0cf003b0c401637cb5d490df0df3d49507529aece3eb58c864bee948d48ed size: 751
2022/08/25 09:51:19 Published crmhv25ua3ddectf7dg.azurecr.io/patientapis-6d8bb96cb4919f922ff1c3d8cb1e1f1e@sha256:06b0cf003b0c401637cb5d490df0df3d49507529aece3eb58c864bee948d48ed
```

実行すると、.ko_name にkoが作成した名前を入れたファイルを作る。
ACRへの登録内容は、`make show-acr`で確認できる。2つのイメージが登録される。アプリの実態は、`"digest": "sha256:06b0cf003b0c401637cb5d490df0df3d49507529aece3eb58c864bee948d48ed` の方。これは、`.ko_name`のダイジェストと一致する。

```sh
$ make show-acr
az acr repository show-tags -n crmhv25ua3ddectf7dg.azurecr.io --repository $(cat .ko_name | sed -r 's/.*\/(.*)@.*/\1/') --detail
The login server endpoint suffix '.azurecr.io' is automatically omitted.
[
  {
    "changeableAttributes": {
      "deleteEnabled": true,
      "listEnabled": true,
      "readEnabled": true,
      "writeEnabled": true
    },
    "createdTime": "2022-08-25T09:30:10.7264574Z",
    "digest": "sha256:06b0cf003b0c401637cb5d490df0df3d49507529aece3eb58c864bee948d48ed",
    "lastUpdateTime": "2022-08-25T09:30:10.7264574Z",
    "name": "latest",
    "signed": false
  },
  {
    "changeableAttributes": {
      "deleteEnabled": true,
      "listEnabled": true,
      "readEnabled": true,
      "writeEnabled": true
    },
    "createdTime": "2022-08-25T09:29:58.8058009Z",
    "digest": "sha256:e0c50ac2f15e6b3ec6040c88a2baebbde35e817b37f12123c514f89b9a69ff51",
    "lastUpdateTime": "2022-08-25T09:29:58.8058009Z",
    "name": "sha256-06b0cf003b0c401637cb5d490df0df3d49507529aece3eb58c864bee948d48ed.sbom",
    "signed": false
  }
]
```

## デプロイ

予め、ACAを作成し、MSIに権限をつけておく(Role Assignmentしておく)。

```sh
$ make deploy
az deployment sub create \
        --location japaneast \
        --name patientapis-v1alpha1 \
        --template-file ./main.bicep \
        --parameters \
                name=patientapis \
                acrName=crmhv25ua3ddectf7dg \
                acaRgName=rg-mhv2-playground-capps-backend \
                envName=caenv-mhv2-backend \
                managedIdName=id-mhv2-backend \
                imageName=patientapis-6d8bb96cb4919f922ff1c3d8cb1e1f1e@sha256:06b0cf003b0c401637cb5d490df0df3d49507529aece3eb58c864bee948d48ed
{
  "id": "/subscriptions/9b425b64-2564-4960-b281-13eecc5d6b77/providers/Microsoft.Resources/deployments/patientapis-v1alpha1",
  "location": "japaneast",
  "name": "patientapis-v1alpha1",
  "properties": {
    "correlationId": "3aacc169-67d5-4da3-96aa-37944225418d",
    "debugSetting": null,
    "dependencies": [],
    "duration": "PT33.3652324S",
    "error": null,
    "mode": "Incremental",
    "onErrorDeployment": null,
    "outputResources": [
      {
        "id": "/subscriptions/9b425b64-2564-4960-b281-13eecc5d6b77/resourceGroups/rg-mhv2-playground-capps-backend/providers/Microsoft.App/containerApps/patientapis",
        "resourceGroup": "rg-mhv2-playground-capps-backend"
      }
    ],
... 以下省略 ...
```

### 設定

|      環境変数名      |            内容            |                例                |
+----------------------+----------------------------+----------------------------------+
| ACA_NAME             | ACAの名前                  | patientapis                      |
| ACR_NAME             | ACRの名前                  | crmhv25ua3ddectf7dg              |
| ACA_BACKEND_RGNAME   | ACAのリソースグループ名    | rg-mhv2-playground-capps-backend |
| ACA_BACKEND_ENVNAME  | ACEの名前                  | caenv-mhv2-backend               |
| ACA_BACKEND_MSIDNAME | ACRにアクセスするMSIの名前 | id-mhv2-backend                  |

ACA_NAMEは、アプリ毎に変わる。その他は、Azure環境毎に決まる。ここでは、playgroundを記載した。

## HOW TO

az az deployment のエラーが読みづらいので、少しは良くなる方法をメモする。

例えば、こんなエラーになったとする。

```sh
$ make deploy
az deployment sub create \
        --location japaneast \
        --name patientapis-v1alpha1 \
        --template-file ./main.bicep \
        --parameters \
                name=patientapis \
                acrName=crmhv25ua3ddectf7dg.azurecr.io \
                acaRgName=rg-mhv2-playground-capps-backend \
                envName=caenv-mhv2-backend \
                managedIdName=id-mhv2-backend \
                imageName=crmhv25ua3ddectf7dg.azurecr.io/patientapis-6d8bb96cb4919f922ff1c3d8cb1e1f1e@sha256:06b0cf003b0c401637cb5d490df0df3d49507529aece3eb58c864bee948d48ed
{"status":"Failed","error":{"code":"DeploymentFailed","message":"At least one resource deployment operation failed. Please list deployment operations for details. Please see https://aka.ms/DeployOperations for usage details.","details":[{"code":"Conflict","message":"{\r\n  \"status\": \"Failed\",\r\n  \"error\": {\r\n    \"code\": \"ResourceDeploymentFailure\",\r\n    \"message\": \"The resource operation completed with terminal provisioning state 'Failed'.\",\r\n    \"details\": [\r\n      {\r\n        \"code\": \"DeploymentFailed\",\r\n        \"message\": \"At least one resource deployment operation failed. Please list deployment operations for details. Please see https://aka.ms/DeployOperations for usage details.\",\r\n        \"details\": [\r\n          {\r\n            \"code\": \"BadRequest\",\r\n            \"message\": \"{\\r\\n  \\\"code\\\": \\\"WebhookInvalidParameterValue\\\",\\r\\n  \\\"message\\\": \\\"The following field(s) are either invalid or missing. Invalid value: \\\\\\\"crmhv25ua3ddectf7dg.azurecr.io.azurecr.io/crmhv25ua3ddectf7dg.azurecr.io/patientapis-6d8bb96cb4919f922ff1c3d8cb1e1f1e@sha256:06b0cf003b0c401637cb5d490df0df3d49507529aece3eb58c864bee948d48ed\\\\\\\": Unable to pull image using Managed identity /subscriptions/9b425b64-2564-4960-b281-13eecc5d6b77/resourceGroups/rg-mhv2-playground-capps-backend/providers/Microsoft.ManagedIdentity/userAssignedIdentities/id-mhv2-backend for registry crmhv25ua3ddectf7dg.azurecr.io.azurecr.io: template.containers.patientapis.image.\\\"\\r\\n}\"\r\n          }\r\n        ]\r\n      }\r\n    ]\r\n  }\r\n}"}]}}
make: *** [Makefile:24: deploy] Error 1
```

json 文字列の中にjsonが埋め込まれて読みづらいので、json部分をファイルに書き出してjqで展開する。3重構造になっている。

```json
$ jq '.' tmp/x.json
{
  "status": "Failed",
  "error": {
    "code": "DeploymentFailed",
    "message": "At least one resource deployment operation failed. Please list deployment operations for details. Please see https://aka.ms/DeployOperations for usage details.",
    "details": [
      {
        "code": "Conflict",
        "message": "{\r\n  \"status\": \"Failed\",\r\n  \"error\": {\r\n    \"code\": \"ResourceDeploymentFailure\",\r\n    \"message\": \"The resource operation completed with terminal provisioning state 'Failed'.\",\r\n    \"details\": [\r\n      {\r\n        \"code\": \"DeploymentFailed\",\r\n        \"message\": \"At least one resource deployment operation failed. Please list deployment operations for details. Please see https://aka.ms/DeployOperations for usage details.\",\r\n        \"details\": [\r\n          {\r\n            \"code\": \"BadRequest\",\r\n            \"message\": \"{\\r\\n  \\\"code\\\": \\\"WebhookInvalidParameterValue\\\",\\r\\n  \\\"message\\\": \\\"The following field(s) are either invalid or missing. Invalid value: \\\\\\\"crmhv25ua3ddectf7dg.azurecr.io.azurecr.io/crmhv25ua3ddectf7dg.azurecr.io/patientapis-6d8bb96cb4919f922ff1c3d8cb1e1f1e@sha256:06b0cf003b0c401637cb5d490df0df3d49507529aece3eb58c864bee948d48ed\\\\\\\": Unable to pull image using Managed identity /subscriptions/9b425b64-2564-4960-b281-13eecc5d6b77/resourceGroups/rg-mhv2-playground-capps-backend/providers/Microsoft.ManagedIdentity/userAssignedIdentities/id-mhv2-backend for registry crmhv25ua3ddectf7dg.azurecr.io.azurecr.io: template.containers.patientapis.image.\\\"\\r\\n}\"\r\n          }\r\n        ]\r\n      }\r\n    ]\r\n  }\r\n}"
      }
    ]
  }
}

$ jq -r '.error.details[0].message | fromjson' tmp/x.json
{
  "status": "Failed",
  "error": {
    "code": "ResourceDeploymentFailure",
    "message": "The resource operation completed with terminal provisioning state 'Failed'.",
    "details": [
      {
        "code": "DeploymentFailed",
        "message": "At least one resource deployment operation failed. Please list deployment operations for details. Please see https://aka.ms/DeployOperations for usage details.",
        "details": [
          {
            "code": "BadRequest",
            "message": "{\r\n  \"code\": \"WebhookInvalidParameterValue\",\r\n  \"message\": \"The following field(s) are either invalid or missing. Invalid value: \\\"crmhv25ua3ddectf7dg.azurecr.io.azurecr.io/crmhv25ua3ddectf7dg.azurecr.io/patientapis-6d8bb96cb4919f922ff1c3d8cb1e1f1e@sha256:06b0cf003b0c401637cb5d490df0df3d49507529aece3eb58c864bee948d48ed\\\": Unable to pull image using Managed identity /subscriptions/9b425b64-2564-4960-b281-13eecc5d6b77/resourceGroups/rg-mhv2-playground-capps-backend/providers/Microsoft.ManagedIdentity/userAssignedIdentities/id-mhv2-backend for registry crmhv25ua3ddectf7dg.azurecr.io.azurecr.io: template.containers.patientapis.image.\"\r\n}"
          }
        ]
      }
    ]
  }
}
$ jq -r '.error.details[0].message | fromjson' tmp/x.json | jq -r '.error.details[0].details[0].message | fromjson'
{
  "code": "WebhookInvalidParameterValue",
  "message": "The following field(s) are either invalid or missing. Invalid value: \"crmhv25ua3ddectf7dg.azurecr.io.azurecr.io/crmhv25ua3ddectf7dg.azurecr.io/patientapis-6d8bb96cb4919f922ff1c3d8cb1e1f1e@sha256:06b0cf003b0c401637cb5d490df0df3d49507529aece3eb58c864bee948d48ed\": Unable to pull image using Managed identity /subscriptions/9b425b64-2564-4960-b281-13eecc5d6b77/resourceGroups/rg-mhv2-playground-capps-backend/providers/Microsoft.ManagedIdentity/userAssignedIdentities/id-mhv2-backend for registry crmhv25ua3ddectf7dg.azurecr.io.azurecr.io: template.containers.patientapis.image."
}
```

### まとめ

`jq -r '.error.details[0].message | fromjson' tmp/x.json | jq -r '.error.details[0].details[0].message | fromjson'` のようにすると、参考となるエラーメッセージを取り出せる。

## TODO

- mhv2-infra 対応にする
- tagが無い場合に、イメージ名が変わらない
