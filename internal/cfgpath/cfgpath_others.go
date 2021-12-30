// Dinkur the task time tracking utility.
// <https://github.com/dinkur/dinkur>
//
// Copyright (C) 2021 Kalle Fagerberg
// SPDX-FileCopyrightText: 2021 Kalle Fagerberg
// SPDX-License-Identifier: GPL-3.0-or-later
//
// This program is free software: you can redistribute it and/or modify it under
// the terms of the GNU General Public License as published by the
// Free Software Foundation, either version 3 of the License, or (at your option)
// any later version.
//
// This program is distributed in the hope that it will be useful, but WITHOUT
// ANY WARRANTY; without even the implied warranty of MERCHANTABILITY or FITNESS
// FOR A PARTICULAR PURPOSE.  See the GNU General Public License for more
// details.
//
// You should have received a copy of the GNU General Public License along with
// this program.  If not, see <http://www.gnu.org/licenses/>.

//go:build !windows && !linux
// +build !windows,!linux

package cfgpath

import (
	"os"
	"path/filepath"
)

func getConfigPath() string {
	configDir, err := os.UserHomeDir()
	if err != nil {
		return "dinkur.yml"
	}
	return filepath.Join(configDir, ".dinkur.yml")
}

func getDataPath() string {
	configDir, err := os.UserHomeDir()
	if err != nil {
		return "dinkur.db"
	}
	return filepath.Join(configDir, ".dinkur.db")
}
