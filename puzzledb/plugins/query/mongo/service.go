// Copyright (C) 2022 The PuzzleDB Authors.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//    http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package mongo

import (
	"github.com/cybergarage/go-mongo/mongo"
	"github.com/cybergarage/puzzledb-go/puzzledb/context"
	"github.com/cybergarage/puzzledb-go/puzzledb/plugins/query"
	"github.com/cybergarage/puzzledb-go/puzzledb/store"
)

type Service struct {
	mongo.Server
	*query.BaseService
	*Coder
}

// NewService returns a MongoDB service instance.
func NewService() *Service {
	server := &Service{
		Server:      mongo.NewServer(),
		BaseService: query.NewBaseService(),
		Coder:       NewCoder(),
	}

	server.SetMessageListener(server)
	server.SetUserCommandExecutor(server)

	return server
}

// ServiceName returns the plug-in service name.
func (service *Service) ServiceName() string {
	return "mongo"
}

// MessageReceived passes a request message from MongoDB client.
func (service *Service) MessageReceived(msg mongo.OpMessage) {
	// fmt.Printf("-> %s\n", msg.String())
	// log.Hexdump(log.LevelInfo, msg.Bytes())
}

// MessageRespond passes a response message from mongo.Server.
func (service *Service) MessageRespond(msg mongo.OpMessage) {
	// fmt.Printf("<- %s\n", msg.String())
	// log.Hexdump(log.LevelInfo, msg.Bytes())
}

// GetDatabase returns the database with the specified name.
func (service *Service) GetDatabase(ctx context.Context, name string) (store.Database, error) {
	store := service.Store()
	db, err := store.GetDatabase(ctx, name)
	if err == nil {
		return db, nil
	}
	err = store.CreateDatabase(ctx, name)
	if err != nil {
		return nil, err
	}
	return store.GetDatabase(ctx, name)
}

// Start starts the service.
func (service *Service) Start() error {
	port, err := service.GetServiceConfigPort(service)
	if err == nil {
		service.SetPort(port)
	}

	tlsConfig, ok := service.BaseService.TLSConfig()
	if ok {
		service.Server.SetTLSConfig(tlsConfig)
	} else {
		service.Server.SetTLSConfig(nil)
	}

	if err := service.Server.Start(); err != nil {
		return err
	}
	return nil
}

// Stop stops the service.
func (service *Service) Stop() error {
	if err := service.Server.Stop(); err != nil {
		return err
	}
	return nil
}
