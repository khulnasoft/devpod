# Copyright (c) 2020 Gitpod GmbH. All rights reserved.
# Licensed under the GNU Affero General Public License (AGPL).
# See License.AGPL.txt in the project root for license information.

FROM cgr.dev/chainguard/wolfi-base:latest@sha256:79884062be1fd5f698859ee4044c6d13729fabeb9eb1f34aacfc6120414f8a91
COPY components--all-docker/versions.yaml components--all-docker/provenance-bundle.jsonl /
