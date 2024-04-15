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

package storagegit

import (
	"testing"

	"github.com/alis-exchange/buf/private/pkg/git/gittest"
	"github.com/alis-exchange/buf/private/pkg/storage/storagetesting"
	"github.com/stretchr/testify/require"
)

func TestNewBucketAtTreeHash(t *testing.T) {
	t.Parallel()

	repo := gittest.ScaffoldGitRepository(t)
	provider := NewProvider(repo.Objects())
	repo.Commit(t, "first commit", map[string]string{
		"randomBinary":                       "some executable",
		"proto/buf.yaml":                     "some buf.yaml",
		"proto/acme/petstore/v1/a.proto":     "cats",
		"proto/acme/petstore/v1/b.proto":     "animals",
		"proto/acme/petstore/v1/e.proto":     "loblaws",
		"proto/acme/petstore/v1/f.proto":     "merchant of venice",
		"proto/acme/grocerystore/v1/c.proto": "toysrus",
		"proto/acme/grocerystore/v1/d.proto": "petsrus",
		"proto/acme/grocerystore/v1/g.proto": "hamlet",
		"proto/acme/grocerystore/v1/h.proto": "bethoven",
	}, gittest.CommitWithExecutableFile("randomBinary"))
	headCommit, err := repo.HEADCommit()
	require.NoError(t, err)
	require.NotNil(t, headCommit)
	bucket, err := provider.NewReadBucket(headCommit.Tree())
	require.NoError(t, err)

	storagetesting.AssertPaths(
		t,
		bucket,
		"",
		"README.md",
		"proto/acme/grocerystore/v1/c.proto",
		"proto/acme/grocerystore/v1/d.proto",
		"proto/acme/grocerystore/v1/g.proto",
		"proto/acme/grocerystore/v1/h.proto",
		"proto/acme/petstore/v1/a.proto",
		"proto/acme/petstore/v1/b.proto",
		"proto/acme/petstore/v1/e.proto",
		"proto/acme/petstore/v1/f.proto",
		"proto/buf.yaml",
		"randomBinary",
	)
	storagetesting.AssertObjectInfo(
		t,
		bucket,
		"proto/acme/grocerystore/v1/c.proto",
		"proto/acme/grocerystore/v1/c.proto",
	)
	storagetesting.AssertNotExist(t, bucket, "random-path")
	storagetesting.AssertPathToContent(
		t,
		bucket,
		"",
		map[string]string{
			"README.md":                          "This is a scaffold repository.\n",
			"proto/acme/grocerystore/v1/c.proto": "toysrus",
			"proto/acme/grocerystore/v1/d.proto": "petsrus",
			"proto/acme/grocerystore/v1/g.proto": "hamlet",
			"proto/acme/grocerystore/v1/h.proto": "bethoven",
			"proto/acme/petstore/v1/a.proto":     "cats",
			"proto/acme/petstore/v1/b.proto":     "animals",
			"proto/acme/petstore/v1/e.proto":     "loblaws",
			"proto/acme/petstore/v1/f.proto":     "merchant of venice",
			"proto/buf.yaml":                     "some buf.yaml",
			"randomBinary":                       "some executable",
		},
	)
}
