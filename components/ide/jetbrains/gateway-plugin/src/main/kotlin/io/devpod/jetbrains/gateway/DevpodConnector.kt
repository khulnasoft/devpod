// Copyright (c) 2022 Gitpod GmbH. All rights reserved.
// Licensed under the GNU Affero General Public License (AGPL).
// See License.AGPL.txt in the project root for license information.

package io.devpod.jetbrains.gateway

import com.jetbrains.gateway.api.GatewayConnector
import com.jetbrains.gateway.api.GatewayConnectorDocumentationPage
import com.jetbrains.rd.util.lifetime.Lifetime
import io.devpod.jetbrains.icons.GitpodIcons
import java.awt.Component

class GitpodConnector : GatewayConnector {
    override val icon = GitpodIcons.Logo

    override fun createView(lifetime: Lifetime) = GitpodConnectorView(lifetime)

    override fun getActionText() = "Connect to Gitpod"

    override fun getDescription() = "Connect to Gitpod workspaces"

    override fun getDocumentationAction() = GatewayConnectorDocumentationPage("https://www.devpod.io/docs/ides-and-editors/jetbrains-gateway")

    override fun getConnectorId() = "devpod.connector"

    override fun getRecentConnections(setContentCallback: (Component) -> Unit) = GitpodRecentConnections()

    override fun getTitle() = "Gitpod"

    @Deprecated("Not used", ReplaceWith("null"))
    override fun getTitleAdornment() = null

    override fun initProcedure() {}
}
