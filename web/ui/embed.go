// Copyright The Prometheus Authors
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
// http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

//go:build builtinassets

package ui

import "embed"

//go:embed static/mantine-ui/assets/codicon-B_nZgZYP.ttf.gz static/mantine-ui/assets/index-DVD_1QSo.js.gz static/mantine-ui/assets/index-Dli9JPrB.css.gz static/mantine-ui/favicon.svg.gz static/mantine-ui/index.html.gz
var EmbedFS embed.FS
