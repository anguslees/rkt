// Copyright 2014 The rkt Authors
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

//+build linux

package main

import (
	"path/filepath"

	"github.com/coreos/rkt/Godeps/_workspace/src/github.com/appc/spec/schema/types"
	"github.com/coreos/rkt/Godeps/_workspace/src/github.com/coreos/go-systemd/unit"
	"github.com/coreos/rkt/common"
)

const (
	envDir          = "/rkt/env" // TODO(vc): perhaps this doesn't belong in /rkt?
	unitsDir        = "/usr/lib/systemd/system"
	defaultWantsDir = unitsDir + "/default.target.wants"
	socketsWantsDir = unitsDir + "/sockets.target.wants"
)

// ServiceUnitName returns a systemd service unit name for the given app name.
func ServiceUnitName(appName types.ACName) string {
	return appName.String() + ".service"
}

// ServiceUnitPath returns the path to the systemd service file for the given
// app name.
func ServiceUnitPath(root string, appName types.ACName) string {
	return filepath.Join(common.Stage1RootfsPath(root), unitsDir, ServiceUnitName(appName))
}

// RelEnvFilePath returns the path to the environment file for the given app name
// relative to the pod's root.
func RelEnvFilePath(appName types.ACName) string {
	return filepath.Join(envDir, appName.String())
}

// EnvFilePath returns the path to the environment file for the given app name.
func EnvFilePath(root string, appName types.ACName) string {
	return filepath.Join(common.Stage1RootfsPath(root), RelEnvFilePath(appName))
}

// ServiceWantPath returns the systemd default.target want symlink path for the
// given app name.
func ServiceWantPath(root string, appName types.ACName) string {
	return filepath.Join(common.Stage1RootfsPath(root), defaultWantsDir, ServiceUnitName(appName))
}

// InstantiatedPrepareAppUnitName returns the systemd service unit name for prepare-app
// instantiated for the given root.
func InstantiatedPrepareAppUnitName(appName types.ACName) string {
	// Naming respecting escaping rules, see systemd.unit(5) and systemd-escape(1)
	escaped_root := unit.UnitNamePathEscape(common.RelAppRootfsPath(appName))
	return "prepare-app@" + escaped_root + ".service"
}

// SocketUnitName returns a systemd socket unit name for the given app name.
func SocketUnitName(appName types.ACName) string {
	return appName.String() + ".socket"
}

// SocketUnitPath returns the path to the systemd socket file for the given app name.
func SocketUnitPath(root string, appName types.ACName) string {
	return filepath.Join(common.Stage1RootfsPath(root), unitsDir, SocketUnitName(appName))
}

// SocketWantPath returns the systemd sockets.target.wants symlink path for the
// given app name.
func SocketWantPath(root string, appName types.ACName) string {
	return filepath.Join(common.Stage1RootfsPath(root), socketsWantsDir, SocketUnitName(appName))
}
