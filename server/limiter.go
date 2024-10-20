// Copyright 2024 CloudWeGo Authors
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

package server

import (
	"github.com/kitex-contrib/config-consul/consul"
	"github.com/kitex-contrib/config-consul/utils"

	configserver "github.com/cloudwego-contrib/cwgo-pkg/config/consul/server"
	"github.com/cloudwego/kitex/server"
)

// WithLimiter sets the limiter config from consul configuration center.
func WithLimiter(dest string, consulClient consul.Client, uniqueID int64, opts utils.Options) server.Option {
	return configserver.WithLimiter(dest, consulClient, uniqueID, opts)
}
