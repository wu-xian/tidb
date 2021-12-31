// Copyright 2021 PingCAP, Inc.
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

package state

import "go.uber.org/atomic"

// Default Top-SQL state values.
const (
	DefTiDBTopSQLEnable                = false
	DefTiDBTopSQLPrecisionSeconds      = 1
	DefTiDBTopSQLMaxStatementCount     = 100
	DefTiDBTopSQLMaxCollect            = 5000
	DefTiDBTopSQLReportIntervalSeconds = 60
)

// GlobalState is the global Top-SQL state.
var GlobalState = State{
	Enable:                atomic.NewBool(DefTiDBTopSQLEnable),
	PrecisionSeconds:      atomic.NewInt64(DefTiDBTopSQLPrecisionSeconds),
	MaxStatementCount:     atomic.NewInt64(DefTiDBTopSQLMaxStatementCount),
	MaxCollect:            atomic.NewInt64(DefTiDBTopSQLMaxCollect),
	ReportIntervalSeconds: atomic.NewInt64(DefTiDBTopSQLReportIntervalSeconds),
}

// State is the state for control top sql feature.
type State struct {
	// Enable top-sql or not.
	Enable *atomic.Bool
	// The refresh interval of top-sql.
	PrecisionSeconds *atomic.Int64
	// The maximum number of statements kept in memory.
	MaxStatementCount *atomic.Int64
	// The maximum capacity of the collect map.
	MaxCollect *atomic.Int64
	// The report data interval of top-sql.
	ReportIntervalSeconds *atomic.Int64
}

// EnableTopSQL enables the top SQL feature.
func EnableTopSQL() {
	GlobalState.Enable.Store(true)
}

// DisableTopSQL disables the top SQL feature.
func DisableTopSQL() {
	GlobalState.Enable.Store(false)
}

// TopSQLEnabled uses to check whether enabled the top SQL feature.
func TopSQLEnabled() bool {
	return GlobalState.Enable.Load()
}