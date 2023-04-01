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
	"io"
	"net/http"
	"strconv"

	"github.com/vmware-tanzu/apps-cli-plugin/pkg/apis/lsp"
	"github.com/vmware-tanzu/apps-cli-plugin/pkg/cli-runtime"
	"github.com/vmware-tanzu/apps-cli-plugin/pkg/source"
)

const errFormat = "%s\nErrors:\n- %s"

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

	if b, err := io.ReadAll(resp.Body); err != nil {
		return lsp.LSPStatus{}, err
	} else if err := json.Unmarshal(b, r); err != nil {
		r = &lspResponse{Message: string(b)}
	}

	if s := checkRequestResponseCode(resp, r.Message); s != nil {
		return *s, nil
	}

	return getStatusFromLSPResponse(*r)
}

func checkRequestResponseCode(resp *http.Response, msg string) *lsp.LSPStatus {
	if resp.StatusCode == http.StatusForbidden || resp.StatusCode == http.StatusUnauthorized {
		return &lsp.LSPStatus{
			Message: fmt.Sprintf(errFormat, "The current user does not have permission to access the local source proxy.", msg),
		}
	}

	if resp.StatusCode >= http.StatusInternalServerError && resp.StatusCode != http.StatusServiceUnavailable {
		return &lsp.LSPStatus{
			UserHasPermission: true,
			Reachable:         true,
			Message:           fmt.Sprintf(errFormat, "Local source proxy is not healthy.", msg),
		}
	}

	if resp.StatusCode == http.StatusServiceUnavailable {
		return &lsp.LSPStatus{
			UserHasPermission: true,
			Message:           fmt.Sprintf(errFormat, "Local source proxy is not healthy.", msg),
		}
	}

	if resp.StatusCode == http.StatusNotFound {
		return &lsp.LSPStatus{
			UserHasPermission: true,
			Message:           fmt.Sprintf(errFormat, "Local source proxy is not installed on the cluster.", msg),
		}
	}
	return nil
}

func getStatusFromLSPResponse(r lspResponse) (lsp.LSPStatus, error) {
	if r.StatusCode == "" {
		return lsp.LSPStatus{}, fmt.Errorf("unable to read local source proxy response: %+v", r)
	}

	if s, err := strconv.Atoi(r.StatusCode); err == nil {
		switch s {
		case http.StatusOK:
			return lsp.LSPStatus{
				UserHasPermission:     true,
				Reachable:             true,
				UpstreamAuthenticated: true,
				OverallHealth:         true,
			}, nil
		default:
			return lsp.LSPStatus{
				Reachable: true,
				Message:   r.Message,
			}, nil
		}
	} else {
		return lsp.LSPStatus{}, err
	}
}
