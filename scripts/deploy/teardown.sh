# Copyright (c) 2021-2022. Marvin Hansen | marvin.hansen@gmail.com


# bin/bash
# set -o errexit
set -o nounset
set -o pipefail

echo "* Teardown OMX service"
echo "* Delete OMX Secrets"
command kubectl delete secret omx-creds

echo "* Delete OMX Service"
command kubectl delete "$(kubectl get svc -o name --namespace default | grep service/omx-*)"

echo "* Delete OMX Pods"
command kubectl delete "$(kubectl get pod -o name --namespace default | grep pod/omx-*)"

echo "* Delete OMX Deployment"
command kubectl delete "$(kubectl get deployment -o name --namespace default | grep deployment.apps/omx)"

echo "DONE!"
