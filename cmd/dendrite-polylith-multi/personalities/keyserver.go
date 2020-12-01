// Copyright 2020 The Matrix.org Foundation C.I.C.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package personalities

import (
	"github.com/matrix-org/dendrite/internal/config"
	"github.com/matrix-org/dendrite/internal/setup"
	"github.com/matrix-org/dendrite/keyserver"
)

func KeyServer(base *setup.BaseDendrite, cfg *config.Dendrite) {
	fsAPI := base.FederationSenderHTTPClient()
	intAPI := keyserver.NewInternalAPI(&base.Cfg.KeyServer, fsAPI)
	intAPI.SetUserAPI(base.UserAPIClient())

	keyserver.AddInternalRoutes(base.InternalAPIMux, intAPI)

	base.SetupAndServeHTTP(
		base.Cfg.KeyServer.InternalAPI.Listen, // internal listener
		setup.NoListener,                      // external listener
		nil, nil,
	)
}
