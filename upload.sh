#!/bin/bash

set -euo pipefail

cf delete-buildpack -f datadog-sidecar-test
cf create-buildpack datadog-sidecar-test datadog-cloudfoundry-buildpack.zip 99
