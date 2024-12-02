package options

const ALL_PATTERN = "*"

type ActsOptions struct {
	Type  string
	State string
	Tag   string
	Key   string
	Ack   bool
}

type Message struct {
	Type  string `json:"type"`
	Pid   string `json:"pid"`
	Tid   string `json:"tid"`
	Key   string `json:"key"`
	State string `json:"state"`
	Name  string `json:"name"`

	Model   map[string]any `json:"model"`
	Inputs  map[string]any `json:"inputs"`
	Outputs map[string]any `json:"outputs"`
}

type Options func(*ActsOptions)
type Callback func(any)

func DefaultOptions() ActsOptions {
	return ActsOptions{Type: ALL_PATTERN, State: ALL_PATTERN, Tag: ALL_PATTERN, Key: ALL_PATTERN}
}

func WithType(typ string) Options {
	return func(v *ActsOptions) {
		v.Type = typ
	}
}

func WithState(state string) Options {
	return func(v *ActsOptions) {
		v.State = state
	}
}

func WithKey(key string) Options {
	return func(v *ActsOptions) {
		v.Key = key
	}
}

func WithTag(tag string) Options {
	return func(v *ActsOptions) {
		v.Tag = tag
	}
}

func WithAck(ack bool) Options {
	return func(v *ActsOptions) {
		v.Ack = ack
	}
}
