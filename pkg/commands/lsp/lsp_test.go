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
	"net/http"
	"strconv"
	"testing"

	"github.com/google/go-cmp/cmp"

	"github.com/vmware-tanzu/apps-cli-plugin/pkg/apis/lsp"
)

func Test_checkRequestResponseCode(t *testing.T) {
	msg := "my cool message"
	type args struct {
		resp *http.Response
		msg  string
	}
	tests := []struct {
		name string
		args args
		want *lsp.LSPStatus
	}{
		{
			name: "200 OK",
			args: args{
				resp: &http.Response{
					StatusCode: http.StatusOK,
				},
			},
			want: nil,
		},
		{
			name: "403 Forbidden",
			args: args{
				resp: &http.Response{
					StatusCode: http.StatusForbidden,
				},
				msg: msg,
			},
			want: &lsp.LSPStatus{
				Message: `The current user does not have permission to access the local source proxy.
Errors:
- my cool message`,
			},
		},
		{
			name: "403 Unauthorized",
			args: args{
				resp: &http.Response{
					StatusCode: http.StatusUnauthorized,
				},
				msg: msg,
			},
			want: &lsp.LSPStatus{
				Message: `The current user does not have permission to access the local source proxy.
Errors:
- my cool message`,
			},
		},
		{
			name: "500 Internal Server Error",
			args: args{
				resp: &http.Response{
					StatusCode: http.StatusInternalServerError,
				},
				msg: msg,
			},
			want: &lsp.LSPStatus{
				UserHasPermission: true,
				Reachable:         true,
				Message: `Local source proxy is not healthy.
Errors:
- my cool message`,
			},
		},
		{
			name: "501 Not Implemented",
			args: args{
				resp: &http.Response{
					StatusCode: http.StatusNotImplemented,
				},
				msg: msg,
			},
			want: &lsp.LSPStatus{
				UserHasPermission: true,
				Reachable:         true,
				Message: `Local source proxy is not healthy.
Errors:
- my cool message`,
			},
		},
		{
			name: "504 Gateway Timeout",
			args: args{
				resp: &http.Response{
					StatusCode: http.StatusGatewayTimeout,
				},
				msg: msg,
			},
			want: &lsp.LSPStatus{
				UserHasPermission: true,
				Reachable:         true,
				Message: `Local source proxy is not healthy.
Errors:
- my cool message`,
			},
		},
		{
			name: "503 Service Unavailable",
			args: args{
				resp: &http.Response{
					StatusCode: http.StatusServiceUnavailable,
				},
				msg: msg,
			},
			want: &lsp.LSPStatus{
				UserHasPermission: true,
				Message: `Local source proxy is not healthy.
Errors:
- my cool message`,
			},
		},
		{
			name: "404 Not Found",
			args: args{
				resp: &http.Response{
					StatusCode: http.StatusNotFound,
				},
				msg: msg,
			},
			want: &lsp.LSPStatus{
				UserHasPermission: true,
				Message: `Local source proxy is not installed on the cluster.
Errors:
- my cool message`,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := checkRequestResponseCode(tt.args.resp, tt.args.msg)
			if diff := cmp.Diff(tt.want, got); diff != "" {
				t.Errorf("checkRequestResponseCode(): Unexpected output (-expected, +actual): %s", diff)
			}
		})
	}
}

func Test_getStatusFromLSPResponse(t *testing.T) {
	msg := "my cool message"
	type args struct {
		r lspResponse
	}
	tests := []struct {
		name    string
		args    args
		want    lsp.LSPStatus
		wantErr bool
	}{
		{
			name: "No status code",
			args: args{
				r: lspResponse{Message: msg},
			},
			wantErr: true,
		},
		{
			name: "200 OK",
			args: args{
				r: lspResponse{
					StatusCode: strconv.Itoa(http.StatusOK),
					Message:    msg,
				},
			},
			want: lsp.LSPStatus{
				UserHasPermission:     true,
				Reachable:             true,
				UpstreamAuthenticated: true,
				OverallHealth:         true,
				Message:               "All health checks passed",
			},
		},
		{
			name: "400 Bad Request",
			args: args{
				r: lspResponse{
					StatusCode: strconv.Itoa(http.StatusBadRequest),
					Message:    msg,
				},
			},
			want: lsp.LSPStatus{
				Reachable: true,
				Message:   "Local source proxy was unable to authenticate against the target registry.\nErrors:\n- my cool message",
			},
		},
		{
			name: "401 Unauthorized",
			args: args{
				r: lspResponse{
					StatusCode: strconv.Itoa(http.StatusUnauthorized),
					Message:    msg,
				},
			},
			want: lsp.LSPStatus{
				Reachable: true,
				Message:   "Local source proxy was unable to authenticate against the target registry.\nErrors:\n- my cool message",
			},
		},
		{
			name: "403 Forbidden",
			args: args{
				r: lspResponse{
					StatusCode: strconv.Itoa(http.StatusForbidden),
					Message:    msg,
				},
			},
			want: lsp.LSPStatus{
				Reachable: true,
				Message:   "Local source proxy was unable to authenticate against the target registry.\nErrors:\n- my cool message",
			},
		},
		{
			name: "404 Not Found",
			args: args{
				r: lspResponse{
					StatusCode: strconv.Itoa(http.StatusNotFound),
					Message:    msg,
				},
			},
			want: lsp.LSPStatus{
				Reachable: true,
				Message:   "Local source proxy was unable to authenticate against the target registry.\nErrors:\n- my cool message",
			},
		},
		{
			name: "500 Internal Server Error",
			args: args{
				r: lspResponse{
					StatusCode: strconv.Itoa(http.StatusInternalServerError),
					Message:    msg,
				},
			},
			want: lsp.LSPStatus{
				Reachable: true,
				Message:   "Local source proxy was unable to authenticate against the target registry.\nErrors:\n- my cool message",
			},
		},
		{
			name: "501 Not Implemented",
			args: args{
				r: lspResponse{
					StatusCode: strconv.Itoa(http.StatusNotImplemented),
					Message:    msg,
				},
			},
			want: lsp.LSPStatus{
				Reachable: true,
				Message:   "Local source proxy was unable to authenticate against the target registry.\nErrors:\n- my cool message",
			},
		},
		{
			name: "502 Bad Gateway",
			args: args{
				r: lspResponse{
					StatusCode: strconv.Itoa(http.StatusBadGateway),
					Message:    msg,
				},
			},
			want: lsp.LSPStatus{
				Reachable: true,
				Message:   "Local source proxy was unable to authenticate against the target registry.\nErrors:\n- my cool message",
			},
		},
		{
			name: "503 Service Unavailable",
			args: args{
				r: lspResponse{
					StatusCode: strconv.Itoa(http.StatusServiceUnavailable),
					Message:    msg,
				},
			},
			want: lsp.LSPStatus{
				Reachable: true,
				Message:   "Local source proxy was unable to authenticate against the target registry.\nErrors:\n- my cool message",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := getStatusFromLSPResponse(tt.args.r)
			if (err != nil) != tt.wantErr {
				t.Errorf("getStatusFromLSPResponse() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if diff := cmp.Diff(tt.want, got); diff != "" {
				t.Errorf("getStatusFromLSPResponse(): Unexpected output (-expected, +actual): %s", diff)
			}
		})
	}
}
