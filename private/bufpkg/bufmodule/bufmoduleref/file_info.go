// Copyright 2020-2024 Buf Technologies, Inc.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package bufmoduleref

import (
	"github.com/alis-exchange/buf/private/pkg/protodescriptor"
)

var _ FileInfo = &fileInfo{}

type fileInfo struct {
	path           string
	externalPath   string
	moduleIdentity ModuleIdentity
	commit         string
}

func newFileInfo(
	path string,
	externalPath string,
	moduleIdentity ModuleIdentity,
	commit string,
) (*fileInfo, error) {
	if err := protodescriptor.ValidateProtoPath("root relative file path", path); err != nil {
		return nil, err
	}
	if externalPath == "" {
		externalPath = path
	}
	return newFileInfoNoValidate(
		path,
		externalPath,
		moduleIdentity,
		commit,
	), nil
}

func newFileInfoNoValidate(
	path string,
	externalPath string,
	moduleIdentity ModuleIdentity,
	commit string,
) *fileInfo {
	return &fileInfo{
		path:           path,
		externalPath:   externalPath,
		moduleIdentity: moduleIdentity,
		commit:         commit,
	}
}

func (f *fileInfo) Path() string {
	return f.path
}

func (f *fileInfo) ExternalPath() string {
	return f.externalPath
}

func (f *fileInfo) ModuleIdentity() ModuleIdentity {
	return f.moduleIdentity
}

func (f *fileInfo) Commit() string {
	return f.commit
}

func (*fileInfo) isFileInfo() {}
