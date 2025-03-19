// Copyright (c) 2021 Gitpod GmbH. All rights reserved.
// Licensed under the GNU Affero General Public License (AGPL).
// See License.AGPL.txt in the project root for license information.

package io.devpod.devpodprotocol.testclient;

import io.devpod.devpodprotocol.api.GitpodClient;
import io.devpod.devpodprotocol.api.GitpodServer;
import io.devpod.devpodprotocol.api.GitpodServerLauncher;
import io.devpod.devpodprotocol.api.entities.SendHeartBeatOptions;
import io.devpod.devpodprotocol.api.entities.User;

import java.util.Collections;

public class TestClient {
    public static void main(String[] args) throws Exception {
        String uri = "wss://devpod.io/api/v1";
        String token = "CHANGE-ME";
        String origin = "https://CHANGE-ME.devpod.io/";

        GitpodClient client = new GitpodClient();
        GitpodServerLauncher.create(client).listen(uri, origin, token, "Test", "Test", Collections.emptyList(), null);
        GitpodServer devpodServer = client.getServer();
        User user = devpodServer.getLoggedInUser().join();
        System.out.println("logged in user:" + user);

        Void result = devpodServer
                .sendHeartBeat(new SendHeartBeatOptions("CHANGE-ME", false)).join();
        System.out.println("send heart beat:" + result);
    }
}
