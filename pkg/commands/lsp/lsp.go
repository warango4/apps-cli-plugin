/*
Copyright 2023 VMware, Inc.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package lsp

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"

	"github.com/vmware-tanzu/apps-cli-plugin/pkg/apis/lsp"
	"github.com/vmware-tanzu/apps-cli-plugin/pkg/cli-runtime"
	"github.com/vmware-tanzu/apps-cli-plugin/pkg/source"
)

type lspResponse struct {
	Message    string `json:"message"`
	StatusCode string `json:"statuscode"`
}

func GetStatus(ctx context.Context, c *cli.Config) (lsp.LSPStatus, error) {
	r := &lspResponse{}
	var localTransport *source.Wrapper
	var resp *http.Response
	var err error

	if localTransport, err = source.LocalRegistryTransport(ctx, c.KubeRestConfig(), c.GetClientSet().CoreV1().RESTClient(), "health"); err != nil {
		return lsp.LSPStatus{}, err
	}
	if req, err := http.NewRequest(http.MethodGet, localTransport.URL.Path, nil); err != nil {
		return lsp.LSPStatus{}, err
	} else if resp, err = localTransport.RoundTrip(req); err != nil {
		return lsp.LSPStatus{}, err
	}

	if b, err := ioutil.ReadAll(resp.Body); err != nil {
		return lsp.LSPStatus{}, err
	} else if err = json.Unmarshal(b, r); err != nil {
		return lsp.LSPStatus{}, err
	}

	if r.StatusCode == "" {
		return lsp.LSPStatus{}, fmt.Errorf("unable to read local source proxy response: %+v", r)
	}

	s, err := strconv.Atoi(r.StatusCode)
	if err != nil {
		return lsp.LSPStatus{}, err
	}

	switch s {
	case http.StatusOK:
		return lsp.LSPStatus{
			UserHasPermission:     true,
			Reachable:             true,
			UpstreamAuthenticated: true,
			OverallHealth:         true,
		}, nil
	case http.StatusNotFound:
	default:
		return lsp.LSPStatus{
			Reachable: true,
			Message:   r.Message,
		}, nil
	}

	return lsp.LSPStatus{}, nil
}
