// Forked from:
// https://github.com/redis/go-redis/blob/91dddc2e1108c779e8c5b85fd667029873c95172/result.go
//
// License for this file: BSD-2-Clause license

package valkey

import (
	"time"

	"github.com/valkey-io/valkey-go/valkeycompat"
)

// NewCmdResult returns a Cmd initialised with val and err for testing.
func NewCmdResult(val interface{}, err error) *valkeycompat.Cmd {
	var cmd valkeycompat.Cmd
	cmd.SetVal(val)
	cmd.SetErr(err)
	return &cmd
}

// NewSliceResult returns a SliceCmd initialised with val and err for testing.
func NewSliceResult(val []interface{}, err error) *valkeycompat.SliceCmd {
	var cmd valkeycompat.SliceCmd
	cmd.SetVal(val)
	cmd.SetErr(err)
	return &cmd
}

// NewStatusResult returns a StatusCmd initialised with val and err for testing.
func NewStatusResult(val string, err error) *valkeycompat.StatusCmd {
	var cmd valkeycompat.StatusCmd
	cmd.SetVal(val)
	cmd.SetErr(err)
	return &cmd
}

// NewIntResult returns an IntCmd initialised with val and err for testing.
func NewIntResult(val int64, err error) *valkeycompat.IntCmd {
	var cmd valkeycompat.IntCmd
	cmd.SetVal(val)
	cmd.SetErr(err)
	return &cmd
}

// NewDurationResult returns a DurationCmd initialised with val and err for testing.
func NewDurationResult(val time.Duration, err error) *valkeycompat.DurationCmd {
	var cmd valkeycompat.DurationCmd
	cmd.SetVal(val)
	cmd.SetErr(err)
	return &cmd
}

// NewBoolResult returns a BoolCmd initialised with val and err for testing.
func NewBoolResult(val bool, err error) *valkeycompat.BoolCmd {
	var cmd valkeycompat.BoolCmd
	cmd.SetVal(val)
	cmd.SetErr(err)
	return &cmd
}

// NewStringResult returns a StringCmd initialised with val and err for testing.
func NewStringResult(val string, err error) *valkeycompat.StringCmd {
	var cmd valkeycompat.StringCmd
	cmd.SetVal(val)
	cmd.SetErr(err)
	return &cmd
}

// NewFloatResult returns a FloatCmd initialised with val and err for testing.
func NewFloatResult(val float64, err error) *valkeycompat.FloatCmd {
	var cmd valkeycompat.FloatCmd
	cmd.SetVal(val)
	cmd.SetErr(err)
	return &cmd
}

// NewStringSliceResult returns a StringSliceCmd initialised with val and err for testing.
func NewStringSliceResult(val []string, err error) *valkeycompat.StringSliceCmd {
	var cmd valkeycompat.StringSliceCmd
	cmd.SetVal(val)
	cmd.SetErr(err)
	return &cmd
}

// NewBoolSliceResult returns a BoolSliceCmd initialised with val and err for testing.
func NewBoolSliceResult(val []bool, err error) *valkeycompat.BoolSliceCmd {
	var cmd valkeycompat.BoolSliceCmd
	cmd.SetVal(val)
	cmd.SetErr(err)
	return &cmd
}

// NewMapStringStringResult returns a MapStringStringCmd initialised with val and err for testing.
// func NewMapStringStringResult(val map[string]string, err error) *valkeycompat.MapStringStringCmd {
// 	var cmd valkeycompat.MapStringStringCmd
// 	cmd.SetVal(val)
// 	cmd.SetErr(err)
// 	return &cmd
// }

// NewMapStringIntCmdResult returns a MapStringIntCmd initialised with val and err for testing.
func NewMapStringIntCmdResult(val map[string]int64, err error) *valkeycompat.MapStringIntCmd {
	var cmd valkeycompat.MapStringIntCmd
	cmd.SetVal(val)
	cmd.SetErr(err)
	return &cmd
}

// NewTimeCmdResult returns a TimeCmd initialised with val and err for testing.
func NewTimeCmdResult(val time.Time, err error) *valkeycompat.TimeCmd {
	var cmd valkeycompat.TimeCmd
	cmd.SetVal(val)
	cmd.SetErr(err)
	return &cmd
}

// NewZSliceCmdResult returns a ZSliceCmd initialised with val and err for testing.
func NewZSliceCmdResult(val []valkeycompat.Z, err error) *valkeycompat.ZSliceCmd {
	var cmd valkeycompat.ZSliceCmd
	cmd.SetVal(val)
	cmd.SetErr(err)
	return &cmd
}

// NewZWithKeyCmdResult returns a ZWithKeyCmd initialised with val and err for testing.
func NewZWithKeyCmdResult(val *valkeycompat.ZWithKey, err error) *valkeycompat.ZWithKeyCmd {
	var cmd valkeycompat.ZWithKeyCmd
	cmd.SetVal(*val)
	cmd.SetErr(err)
	return &cmd
}

// NewScanCmdResult returns a ScanCmd initialised with val and err for testing.
func NewScanCmdResult(keys []string, cursor uint64, err error) *valkeycompat.ScanCmd {
	var cmd valkeycompat.ScanCmd
	cmd.SetVal(keys, cursor)
	cmd.SetErr(err)
	return &cmd
}

// NewClusterSlotsCmdResult returns a ClusterSlotsCmd initialised with val and err for testing.
func NewClusterSlotsCmdResult(val []valkeycompat.ClusterSlot, err error) *valkeycompat.ClusterSlotsCmd {
	var cmd valkeycompat.ClusterSlotsCmd
	cmd.SetVal(val)
	cmd.SetErr(err)
	return &cmd
}

// NewGeoLocationCmdResult returns a GeoLocationCmd initialised with val and err for testing.
func NewGeoLocationCmdResult(val []valkeycompat.GeoLocation, err error) *valkeycompat.GeoLocationCmd {
	var cmd valkeycompat.GeoLocationCmd
	cmd.SetVal(val)
	cmd.SetErr(err)
	return &cmd
}

// NewGeoPosCmdResult returns a GeoPosCmd initialised with val and err for testing.
func NewGeoPosCmdResult(val []*valkeycompat.GeoPos, err error) *valkeycompat.GeoPosCmd {
	var cmd valkeycompat.GeoPosCmd
	cmd.SetVal(val)
	cmd.SetErr(err)
	return &cmd
}

// NewCommandsInfoCmdResult returns a CommandsInfoCmd initialised with val and err for testing.
// func NewCommandsInfoCmdResult(val map[string]*valkeycompat.CommandInfo, err error) *valkeycompat.CommandsInfoCmd {
// 	var cmd valkeycompat.CommandsInfoCmd
// 	cmd.SetVal(val)
// 	cmd.SetErr(err)
// 	return &cmd
// }

// NewXMessageSliceCmdResult returns a XMessageSliceCmd initialised with val and err for testing.
func NewXMessageSliceCmdResult(val []valkeycompat.XMessage, err error) *valkeycompat.XMessageSliceCmd {
	var cmd valkeycompat.XMessageSliceCmd
	cmd.SetVal(val)
	cmd.SetErr(err)
	return &cmd
}

// NewXStreamSliceCmdResult returns a XStreamSliceCmd initialised with val and err for testing.
func NewXStreamSliceCmdResult(val []valkeycompat.XStream, err error) *valkeycompat.XStreamSliceCmd {
	var cmd valkeycompat.XStreamSliceCmd
	cmd.SetVal(val)
	cmd.SetErr(err)
	return &cmd
}

// NewXPendingResult returns a XPendingCmd initialised with val and err for testing.
func NewXPendingResult(val *valkeycompat.XPending, err error) *valkeycompat.XPendingCmd {
	var cmd valkeycompat.XPendingCmd
	cmd.SetVal(*val)
	cmd.SetErr(err)
	return &cmd
}
