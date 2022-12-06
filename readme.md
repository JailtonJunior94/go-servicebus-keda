# GOlang + Service Bus + AKS + Keda

## Instalando KEDA no cluster Kubernetes

1. Adicionado Helm Repo
   ```
   helm repo add kedacore https://kedacore.github.io/charts
   ```
2. Atualizando Helm Repo
   ```
   helm repo update
   ```
3. Instalando Keda Helm Chart
   ```
   kubectl create namespace keda
   helm install keda kedacore/keda --namespace keda
   ```