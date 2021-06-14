/*
 * Copyright (c) 2020 Devtron Labs
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *    http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 *
 */

package pubsub

import (
	"github.com/caarlos0/env"
	"github.com/posthog/posthog-go"
	"go.uber.org/zap"
)

type PosthogClient struct {
	Client posthog.Client
}

type PosthogConfig struct {
	ApiKey string `env:"API_KEY" envDefault:""`
}

func NewPosthogClient(logger *zap.SugaredLogger) (*PosthogClient, error) {
	cfg := &PosthogConfig{}
	err := env.Parse(cfg)
	if err != nil {
		logger.Error("err", err)
		return &PosthogClient{}, err
	}
	client, _ := posthog.NewWithConfig(cfg.ApiKey, posthog.Config{Endpoint: "https://app.posthog.com"})
	//defer client.Close()
	pgClient := &PosthogClient{
		Client: client,
	}
	/*pgClient.Client.Enqueue(posthog.Capture{
		DistinctId: "localhost-user",
		Event:      fmt.Sprintf("event sent from orchestrator on startup userid=%d,time=%s", 1, time.Now()),
	})*/
	return pgClient, nil
}
