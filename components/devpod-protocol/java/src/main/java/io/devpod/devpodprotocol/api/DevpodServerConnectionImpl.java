// Copyright (c) 2022 Gitpod GmbH. All rights reserved.
// Licensed under the GNU Affero General Public License (AGPL).
// See License.AGPL.txt in the project root for license information.

package io.devpod.devpodprotocol.api;

import javax.websocket.CloseReason;
import javax.websocket.Session;
import java.io.IOException;
import java.util.concurrent.CompletableFuture;
import java.util.logging.Level;
import java.util.logging.Logger;

public class GitpodServerConnectionImpl extends CompletableFuture<CloseReason> implements GitpodServerConnection {

    public static final Logger LOG = Logger.getLogger(GitpodServerConnectionImpl.class.getName());

    private final String devpodHost;

    private Session session;

    public GitpodServerConnectionImpl(String devpodHost) {
        this.devpodHost = devpodHost;
    }

    public void setSession(Session session) {
        this.session = session;
    }

    @Override
    public boolean cancel(boolean mayInterruptIfRunning) {
        Session session = this.session;
        this.session = null;
        if (session != null) {
            try {
                session.close();
            } catch (IOException e) {
                LOG.log(Level.WARNING, devpodHost + ": failed to close connection:", e);
            }
        }
        return super.cancel(mayInterruptIfRunning);
    }
}
