/**
 * Copyright (c) 2020 Gitpod GmbH. All rights reserved.
 * Licensed under the GNU Affero General Public License (AGPL).
 * See License.AGPL.txt in the project root for license information.
 */

import { User } from "@devpod/devpod-protocol";
import { injectable } from "inversify";

@injectable()
export class RepositoryService {
    async isGitpodWebhookEnabled(user: User, cloneUrl: string): Promise<boolean> {
        throw new Error("unsupported");
    }
}
