# Copyright (c) 2022. Marvin Hansen | marvin.hansen@gmail.com

# bin/bash
set -o errexit
set -o nounset
set -o pipefail

# "---------------------------------------------------------"
# "-                                                       -"
# "-  Installs all dependencies required by make           -"
# "-                                                       -"
# "---------------------------------------------------------"

echo ""
echo "==============================="
echo "Check for cloud  dependencies:"
echo "==============================="
echo ""

command -v gcloud >/dev/null 2>&1 || {
  # https://cloud.google.com/sdk/docs/install#deb
  echo "Download GCloud"
  command echo "deb https://packages.cloud.google.com/apt cloud-sdk main" | sudo tee -a /etc/apt/sources.list.d/google-cloud-sdk.list
  command curl https://packages.cloud.google.com/apt/doc/apt-key.gpg | sudo apt-key add -
  echo "Install GCloud"
  command sudo apt-get -qqq update && sudo apt-get -qqq -y install google-cloud-sdk
}
echo "* GCloud installed"

command -v kubectl >/dev/null 2>&1 || {
  # https://kubernetes.io/docs/tasks/tools/install-kubectl/
  echo "Download kubectl"
  command sudo curl -LO "https://storage.googleapis.com/kubernetes-release/release/$(curl -s https://storage.googleapis.com/kubernetes-release/release/stable.txt)/bin/linux/amd64/kubectl"
  echo "Install kubectl"
  command sudo chmod +x ./kubectl
  command sudo mv ./kubectl /usr/local/bin/kubectl
  echo "Done"
  kubectl version --client
}
echo "* Kubectl installed"

command helm version >/dev/null 2>&1 || {
  echo "Download Helm"
  command curl -fsSL -o get_helm.sh https://raw.githubusercontent.com/helm/helm/master/scripts/get-helm-3
  echo "Install Helm"
  command sudo chmod +x get_helm.sh
  command ./get_helm.sh
  command rm -f get_helm.sh
  echo "Done"
  helm version
}
echo "* Helm installed"

echo ""
echo "==============================="
echo "All dependencies installed."
echo "==============================="
echo ""
