package openfeature

// IClient defines the behaviour required of an openfeature client
type IClient interface {
	Metadata() ClientMetadata
	AddHooks(hooks ...Hook)
	GetBooleanValue(flag string, defaultValue bool, evalCtx EvaluationContext, options ...EvaluationOption) (bool, error)
	GetStringValue(flag string, defaultValue string, evalCtx EvaluationContext, options ...EvaluationOption) (string, error)
	GetNumberValue(flag string, defaultValue float64, evalCtx EvaluationContext, options ...EvaluationOption) (float64, error)
	GetObjectValue(flag string, defaultValue interface{}, evalCtx EvaluationContext, options ...EvaluationOption) (interface{}, error)
	GetBooleanValueDetails(flag string, defaultValue bool, evalCtx EvaluationContext, options ...EvaluationOption) (EvaluationDetails, error)
	GetStringValueDetails(flag string, defaultValue string, evalCtx EvaluationContext, options ...EvaluationOption) (EvaluationDetails, error)
	GetNumberValueDetails(flag string, defaultValue float64, evalCtx EvaluationContext, options ...EvaluationOption) (EvaluationDetails, error)
	GetObjectValueDetails(flag string, defaultValue interface{}, evalCtx EvaluationContext, options ...EvaluationOption) (EvaluationDetails, error)
}

// ClientMetadata provides a client's metadata
type ClientMetadata struct {
	name string
}

// Name returns the client's name
func (cm ClientMetadata) Name() string {
	return cm.name
}

// Client implements the behaviour required of an openfeature client
type Client struct {
	metadata ClientMetadata
	hooks []Hook
}

// GetClient returns a new Client. Name is a unique identifier for this client
func GetClient(name string) *Client {
	return &Client{
		metadata: ClientMetadata{name: name},
		hooks: []Hook{},
	}
}

// Metadata returns the client's metadata
func (c Client) Metadata() ClientMetadata {
	return c.metadata
}

// AddHooks appends to the client's collection of any previously added hooks
func (c *Client) AddHooks(hooks ...Hook) {
	c.hooks = append(c.hooks, hooks...)
}

// Type represents the type of a flg
type Type int64

const (
	Boolean Type = iota
	String
	Number
	Object
)

type EvaluationDetails struct {
	FlagKey  string
	FlagType Type
	ResolutionDetail
}

// GetBooleanValue return boolean evaluation for flag
func (c Client) GetBooleanValue(flag string, defaultValue bool, evalCtx EvaluationContext, options ...EvaluationOption) (bool, error) {
	resolution := api.provider.GetBooleanEvaluation(flag, defaultValue, evalCtx, options...)
	value := resolution.Value
	err := resolution.Error()
	if err != nil {
		value = defaultValue
	}
	return value, err
}

// GetStringValue return string evaluation for flag
func (c Client) GetStringValue(flag string, defaultValue string, evalCtx EvaluationContext, options ...EvaluationOption) (string, error) {
	resolution := api.provider.GetStringEvaluation(flag, defaultValue, evalCtx, options...)
	value := resolution.Value
	err := resolution.Error()
	if err != nil {
		value = defaultValue
	}
	return value, err
}

// GetNumberValue return number evaluation for flag
func (c Client) GetNumberValue(flag string, defaultValue float64, evalCtx EvaluationContext, options ...EvaluationOption) (float64, error) {
	resolution := api.provider.GetNumberEvaluation(flag, defaultValue, evalCtx, options...)
	value := resolution.Value
	err := resolution.Error()
	if err != nil {
		value = defaultValue
	}
	return value, err
}

// GetObjectValue return object evaluation for flag
func (c Client) GetObjectValue(flag string, defaultValue interface{}, evalCtx EvaluationContext, options ...EvaluationOption) (interface{}, error) {
	resolution := api.provider.GetObjectEvaluation(flag, defaultValue, evalCtx, options...)
	value := resolution.Value
	err := resolution.Error()
	if err != nil {
		value = defaultValue
	}
	return value, err
}

// GetBooleanValueDetails return boolean evaluation for flag
func (c Client) GetBooleanValueDetails(flag string, defaultValue bool, evalCtx EvaluationContext, options ...EvaluationOption) (EvaluationDetails, error) {
	resolution := api.provider.GetBooleanEvaluation(flag, defaultValue, evalCtx, options...)
	value := resolution.Value
	err := resolution.Error()
	if err != nil {
		value = defaultValue
	}
	return EvaluationDetails{
		FlagKey:          flag,
		FlagType:         Boolean,
		ResolutionDetail: ResolutionDetail{
			Value:     value,
			ErrorCode: resolution.ErrorCode,
			Reason:    resolution.Reason,
			Variant:   resolution.Variant,
		},
	}, err

}

// GetStringValueDetails return string evaluation for flag
func (c Client) GetStringValueDetails(flag string, defaultValue string, evalCtx EvaluationContext, options ...EvaluationOption) (EvaluationDetails, error) {
	resolution := api.provider.GetStringEvaluation(flag, defaultValue, evalCtx, options...)
	value := resolution.Value
	err := resolution.Error()
	if err != nil {
		value = defaultValue
	}
	return EvaluationDetails{
		FlagKey:          flag,
		FlagType:         String,
		ResolutionDetail: ResolutionDetail{
			Value:     value,
			ErrorCode: resolution.ErrorCode,
			Reason:    resolution.Reason,
			Variant:   resolution.Variant,
		},
	}, err
}

// GetNumberValueDetails return number evaluation for flag
func (c Client) GetNumberValueDetails(flag string, defaultValue float64, evalCtx EvaluationContext, options ...EvaluationOption) (EvaluationDetails, error) {
	resolution := api.provider.GetNumberEvaluation(flag, defaultValue, evalCtx, options...)
	value := resolution.Value
	err := resolution.Error()
	if err != nil {
		value = defaultValue
	}
	return EvaluationDetails{
		FlagKey:          flag,
		FlagType:         Number,
		ResolutionDetail: ResolutionDetail{
			Value:     value,
			ErrorCode: resolution.ErrorCode,
			Reason:    resolution.Reason,
			Variant:   resolution.Variant,
		},
	}, err
}

// GetObjectValueDetails return object evaluation for flag
func (c Client) GetObjectValueDetails(flag string, defaultValue interface{}, evalCtx EvaluationContext, options ...EvaluationOption) (EvaluationDetails, error) {
	resolution := api.provider.GetObjectEvaluation(flag, defaultValue, evalCtx, options...)
	value := resolution.Value
	err := resolution.Error()
	if err != nil {
		value = defaultValue
	}
	return EvaluationDetails{
		FlagKey:          flag,
		FlagType:         Object,
		ResolutionDetail: ResolutionDetail{
			Value:     value,
			ErrorCode: resolution.ErrorCode,
			Reason:    resolution.Reason,
			Variant:   resolution.Variant,
		},
	}, err
}