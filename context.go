package asterisktoolkit

import "context"

type sessionContextArgsValue struct{}

// GetArguments can be run on the context passed to a SessionFunc
// in order to get the arguments the StasisApp was started with
// (if any).
//
// This can be useful for when an app is used in many places.
func GetArguments(ctx context.Context) []string {
	return ctx.Value(sessionContextArgsValue{}).([]string)
}
