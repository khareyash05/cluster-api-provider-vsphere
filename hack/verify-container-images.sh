#!/bin/bash

# Copyright 2023 The Kubernetes Authors.
#
# Licensed under the Apache License, Version 2.0 (the "License");
# you may not use this file except in compliance with the License.
# You may obtain a copy of the License at
#
#     http://www.apache.org/licenses/LICENSE-2.0
#
# Unless required by applicable law or agreed to in writing, software
# distributed under the License is distributed on an "AS IS" BASIS,
# WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
# See the License for the specific language governing permissions and
# limitations under the License.

set -o errexit
set -o nounset
set -o pipefail

if [[ "${TRACE-0}" == "1" ]]; then
    set -o xtrace
fi

TRIVY_VERSION=0.34.0

GO_OS="$(go env GOOS)"
if [[ "${GO_OS}" == "linux" ]]; then
  TRIVY_OS="Linux"
elif [[ "${GO_OS}" == "darwin"* ]]; then
  TRIVY_OS="macOS"
fi

GO_ARCH="$(go env GOARCH)"
if [[ "${GO_ARCH}" == "amd" ]]; then
  TRIVY_ARCH="32bit"
elif [[ "${GO_ARCH}" == "amd64"* ]]; then
  TRIVY_ARCH="64bit"
elif [[ "${GO_ARCH}" == "arm" ]]; then
  TRIVY_ARCH="ARM"
elif [[ "${GO_ARCH}" == "arm64" ]]; then
  TRIVY_ARCH="ARM64"
fi

TOOL_BIN=hack/tools/bin
mkdir -p ${TOOL_BIN}

# Downloads trivy scanner
curl -L -o ${TOOL_BIN}/trivy.tar.gz "https://github.com/aquasecurity/trivy/releases/download/v${TRIVY_VERSION}/trivy_${TRIVY_VERSION}_${TRIVY_OS}-${TRIVY_ARCH}.tar.gz"

tar -xf "${TOOL_BIN}/trivy.tar.gz" -C "${TOOL_BIN}" trivy
chmod +x ${TOOL_BIN}/trivy
rm ${TOOL_BIN}/trivy.tar.gz

# Builds all the container images to be scanned
make DEV_REGISTRY="gcr.io/cluster-api-provider-vsphere" PULL_POLICY=IfNotPresent DEV_TAG="dev" docker-build

# Scan the images
${TOOL_BIN}/trivy image -q --exit-code 1 --ignore-unfixed --severity MEDIUM,HIGH,CRITICAL gcr.io/cluster-api-provider-vsphere/vsphere-manager:dev && R1=$? || R1=$?

echo ""
BRed='\033[1;31m'
BGreen='\033[1;32m'
NC='\033[0m' # No

if [ "$R1" -ne "0" ]
then
  echo -e "${BRed}Check container images failed! There are vulnerability to be fixed${NC}"
  exit 1
fi

echo -e "${BGreen}Check container images passed! No vulnerability found${NC}"
