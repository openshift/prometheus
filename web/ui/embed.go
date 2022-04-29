// Copyright 2022 The Prometheus Authors
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
// +build builtinassets

package ui

import "embed"

//go:embed static/css/prom_console.css.gz static/js/prom_console.js.gz static/react/favicon.ico.gz static/react/index.html.gz static/react/asset-manifest.json.gz static/react/static/css/main.faad45b4.css.gz static/react/static/css/main.faad45b4.css.map.gz static/react/static/js/main.b0a7c7cf.js.gz static/react/static/js/main.b0a7c7cf.js.LICENSE.txt.gz static/react/static/js/main.b0a7c7cf.js.map.gz static/react/static/media/prometheus_logo_grey.3cf697e5443028ca5e5255b93c7906c5.svg.gz static/react/static/media/codicon.b3726f0165bf67ac6849.ttf.gz static/react/static/media/index.cd351d7c31d0d3fccf96.cjs.gz static/react/manifest.json.gz static/vendor/rickshaw/rickshaw.min.js.gz static/vendor/rickshaw/rickshaw.min.css.gz static/vendor/rickshaw/vendor/d3.v3.js.gz static/vendor/rickshaw/vendor/d3.layout.min.js.gz static/vendor/js/jquery-3.5.1.min.js.gz static/vendor/js/jquery.selection.js.gz static/vendor/js/jquery.hotkeys.js.gz static/vendor/js/popper.min.js.gz static/vendor/bootstrap4-glyphicons/maps/glyphicons-fontawesome.min.css.gz static/vendor/bootstrap4-glyphicons/maps/glyphicons-fontawesome.less.gz static/vendor/bootstrap4-glyphicons/maps/glyphicons-fontawesome.css.gz static/vendor/bootstrap4-glyphicons/css/bootstrap-glyphicons.css.gz static/vendor/bootstrap4-glyphicons/css/bootstrap-glyphicons.min.css.gz static/vendor/bootstrap4-glyphicons/fonts/fontawesome/fa-solid-900.ttf.gz static/vendor/bootstrap4-glyphicons/fonts/fontawesome/fa-regular-400.svg.gz static/vendor/bootstrap4-glyphicons/fonts/fontawesome/fa-regular-400.woff2.gz static/vendor/bootstrap4-glyphicons/fonts/fontawesome/fa-solid-900.eot.gz static/vendor/bootstrap4-glyphicons/fonts/fontawesome/fa-brands-400.svg.gz static/vendor/bootstrap4-glyphicons/fonts/fontawesome/fa-regular-400.woff.gz static/vendor/bootstrap4-glyphicons/fonts/fontawesome/fa-brands-400.eot.gz static/vendor/bootstrap4-glyphicons/fonts/fontawesome/fa-solid-900.svg.gz static/vendor/bootstrap4-glyphicons/fonts/fontawesome/fa-solid-900.woff.gz static/vendor/bootstrap4-glyphicons/fonts/fontawesome/fa-regular-400.ttf.gz static/vendor/bootstrap4-glyphicons/fonts/fontawesome/fa-solid-900.woff2.gz static/vendor/bootstrap4-glyphicons/fonts/fontawesome/fa-brands-400.woff2.gz static/vendor/bootstrap4-glyphicons/fonts/fontawesome/fa-brands-400.woff.gz static/vendor/bootstrap4-glyphicons/fonts/fontawesome/fa-brands-400.ttf.gz static/vendor/bootstrap4-glyphicons/fonts/fontawesome/fa-regular-400.eot.gz static/vendor/bootstrap4-glyphicons/fonts/glyphicons/glyphicons-halflings-regular.woff.gz static/vendor/bootstrap4-glyphicons/fonts/glyphicons/glyphicons-halflings-regular.eot.gz static/vendor/bootstrap4-glyphicons/fonts/glyphicons/glyphicons-halflings-regular.woff2.gz static/vendor/bootstrap4-glyphicons/fonts/glyphicons/glyphicons-halflings-regular.ttf.gz static/vendor/bootstrap4-glyphicons/fonts/glyphicons/glyphicons-halflings-regular.svg.gz static/vendor/bootstrap-4.5.2/css/bootstrap.min.css.gz static/vendor/bootstrap-4.5.2/css/bootstrap-reboot.min.css.map.gz static/vendor/bootstrap-4.5.2/css/bootstrap.css.gz static/vendor/bootstrap-4.5.2/css/bootstrap-grid.css.map.gz static/vendor/bootstrap-4.5.2/css/bootstrap-grid.min.css.gz static/vendor/bootstrap-4.5.2/css/bootstrap.css.map.gz static/vendor/bootstrap-4.5.2/css/bootstrap.min.css.map.gz static/vendor/bootstrap-4.5.2/css/bootstrap-reboot.min.css.gz static/vendor/bootstrap-4.5.2/css/bootstrap-reboot.css.gz static/vendor/bootstrap-4.5.2/css/bootstrap-grid.css.gz static/vendor/bootstrap-4.5.2/css/bootstrap-grid.min.css.map.gz static/vendor/bootstrap-4.5.2/css/bootstrap-reboot.css.map.gz static/vendor/bootstrap-4.5.2/js/bootstrap.bundle.js.gz static/vendor/bootstrap-4.5.2/js/bootstrap.bundle.min.js.map.gz static/vendor/bootstrap-4.5.2/js/bootstrap.bundle.js.map.gz static/vendor/bootstrap-4.5.2/js/bootstrap.js.gz static/vendor/bootstrap-4.5.2/js/bootstrap.bundle.min.js.gz static/vendor/bootstrap-4.5.2/js/bootstrap.min.js.gz static/vendor/bootstrap-4.5.2/js/bootstrap.js.map.gz static/vendor/bootstrap-4.5.2/js/bootstrap.min.js.map.gz
var EmbedFS embed.FS