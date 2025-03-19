/**
 * Copyright (c) 2020 Gitpod GmbH. All rights reserved.
 * Licensed under the GNU Affero General Public License (AGPL).
 * See License.AGPL.txt in the project root for license information.
 */

import { GitpodHostUrl } from "@devpod/devpod-protocol/lib/util/devpod-host-url";

export const workspaceUrl = new GitpodHostUrl(window.location.href);

export const serverUrl = workspaceUrl.withoutWorkspacePrefix();

export const startUrl = serverUrl.asStart(workspaceUrl.workspaceId);
