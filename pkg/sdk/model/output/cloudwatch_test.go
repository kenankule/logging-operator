// Copyright © 2019 Banzai Cloud
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

package output_test

import (
	"testing"

	"github.com/banzaicloud/logging-operator/pkg/sdk/model/output"
	"github.com/banzaicloud/logging-operator/pkg/sdk/model/render"
	"github.com/ghodss/yaml"
	"github.com/stretchr/testify/require"
)

func TesCloudWatch(t *testing.T) {
	CONFIG := []byte(`
log_group_name: operator-log-group
log_stream_name: operator-log-stream
region: us-east-1
auto_create_stream: true
buffer:
  timekey: 1m
  timekey_wait: 30s
  timekey_use_utc: true
`)
	expected := `
  <match **>
    @type cloudwatch_logs
    @id test
    log_group_name operator-log-group
    log_stream_name operator-log-stream
    region us-east-1
    <buffer tag,time>
      @type file
	  chunk_limit_size 8MB
      path /buffers/test_loki.*.buffer
      retry_forever true
      timekey 1m
      timekey_use_utc true
      timekey_wait 30s
    </buffer>
  </match>
`
	cw := &output.CloudWatchOutput{}
	require.NoError(t, yaml.Unmarshal(CONFIG, cw))
	test := render.NewOutputPluginTest(t, cw)
	test.DiffResult(expected)
}
