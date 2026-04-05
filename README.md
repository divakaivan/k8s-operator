# EC2Instance Operator

A Kubernetes operator that manages the lifecycle of AWS EC2 instances as native Kubernetes resources. Define, create, and delete EC2 instances using standard `kubectl` commands — no AWS console or CLI required.

The operator introduces a custom resource called **EC2Instance**. When you apply one, the operator launches the EC2 instance via the AWS API, tracks its state, and writes the instance ID, IP addresses, and DNS names back to the resource's `status`. When you delete the resource, the operator terminates the EC2 instance automatically.

```yaml
apiVersion: compute.cloud.com/v1
kind: EC2Instance
metadata:
  name: my-instance
  namespace: default
spec:
  instanceType: t3.micro
  amiId: ami-0c02fb55956c7d316
  region: us-east-1
  keyPair: my-key-pair
  subnet: subnet-0123456789abcdef0
```

## Installation

The operator requires AWS credentials (`AWS_ACCESS_KEY_ID` and `AWS_SECRET_ACCESS_KEY`) with EC2 permissions. Choose one of the two methods below.

### Option 1 — kubectl

```sh
kubectl apply -f https://raw.githubusercontent.com/divakaivan/ec2instance-k8s-operator/main/dist/install.yaml
```

### Option 2 — Helm

```sh
helm install ec2instance-k8s-operator ./dist/chart
```

## Usage

Apply an `EC2Instance` manifest and watch the status update with the instance details:

```sh
kubectl apply -f my-instance.yaml
kubectl get ec2instances -w
```

Delete the resource to terminate the EC2 instance:

```sh
kubectl delete ec2instance my-instance
```

## License

Copyright 2026.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
