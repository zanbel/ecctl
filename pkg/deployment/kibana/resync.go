// Licensed to Elasticsearch B.V. under one or more contributor
// license agreements. See the NOTICE file distributed with
// this work for additional information regarding copyright
// ownership. Elasticsearch B.V. licenses this file to you under
// the Apache License, Version 2.0 (the "License"); you may
// not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing,
// software distributed under the License is distributed on an
// "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY
// KIND, either express or implied.  See the License for the
// specific language governing permissions and limitations
// under the License.

package kibana

import (
	"github.com/elastic/cloud-sdk-go/pkg/client/clusters_kibana"

	"github.com/elastic/ecctl/pkg/util"
)

// Resync forces indexer to immediately resynchronize the search index
// and cache for a given Kibana instance.
func Resync(params DeploymentParams) error {
	if err := params.Validate(); err != nil {
		return err
	}

	return util.ReturnErrOnly(
		params.API.V1API.ClustersKibana.ResyncKibanaCluster(
			clusters_kibana.NewResyncKibanaClusterParams().
				WithClusterID(params.ID),
			params.API.AuthWriter,
		),
	)
}
