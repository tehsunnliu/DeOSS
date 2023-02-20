/*
   Copyright 2022 CESS scheduler authors

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

package configs

import (
	"time"
)

// system
const (
	// name
	Name = "cess-oss"
	// Name space
	NameSpace = Name
	// version
	Version = Name + " " + "v0.1.2 230220.0954 dev"
	// description
	Description = "Implementation of object storage service based on cess platform"
)

const (
	// base dir
	BaseDir = Name
	// log file dir
	Log = "log"
	// database dir
	Cache = "cache"
	// file dir
	File = "file"
	// tracked files
	Track = "track"
)

const (
	// BlockInterval is the time interval for generating blocks, in seconds
	BlockInterval = time.Second * time.Duration(6)
)
