// Package rest provides primitives to interact with the openapi HTTP API.
//
// Code generated by github.com/deepmap/oapi-codegen DO NOT EDIT.
package rest

const (
	BearerAuthScopes = "BearerAuth.Scopes"
)

// AuthenticationRequest defines model for AuthenticationRequest.
type AuthenticationRequest struct {
	ApiKey string `json:"apiKey"`
	Target *struct {
		Anonymous  *bool                   `json:"anonymous,omitempty"`
		Attributes *map[string]interface{} `json:"attributes,omitempty"`
		Identifier string                  `json:"identifier"`
		Name       *string                 `json:"name,omitempty"`
	} `json:"target,omitempty"`
}

// AuthenticationResponse defines model for AuthenticationResponse.
type AuthenticationResponse struct {
	AuthToken string `json:"authToken"`
}

// Clause defines model for Clause.
type Clause struct {
	Attribute string   `json:"attribute"`
	Id        string   `json:"id"`
	Negate    bool     `json:"negate"`
	Op        string   `json:"op"`
	Values    []string `json:"values"`
}

// Distribution defines model for Distribution.
type Distribution struct {
	BucketBy   string              `json:"bucketBy"`
	Variations []WeightedVariation `json:"variations"`
}

// Error defines model for Error.
type Error struct {
	Code    string `json:"code"`
	Message string `json:"message"`
}

// Evaluation defines model for Evaluation.
type Evaluation struct {
	Flag       string  `json:"flag"`
	Identifier *string `json:"identifier,omitempty"`
	Kind       string  `json:"kind"`
	Value      string  `json:"value"`
}

// Evaluations defines model for Evaluations.
type Evaluations []Evaluation

// FeatureConfig defines model for FeatureConfig.
type FeatureConfig struct {
	DefaultServe         Serve           `json:"defaultServe"`
	Environment          string          `json:"environment"`
	Feature              string          `json:"feature"`
	Kind                 string          `json:"kind"`
	OffVariation         string          `json:"offVariation"`
	Prerequisites        *[]Prerequisite `json:"prerequisites,omitempty"`
	Project              string          `json:"project"`
	Rules                *[]ServingRule  `json:"rules,omitempty"`
	State                FeatureState    `json:"state"`
	VariationToTargetMap *[]VariationMap `json:"variationToTargetMap,omitempty"`
	Variations           []Variation     `json:"variations"`
	Version              *int64          `json:"version,omitempty"`
}

// FeatureState defines model for FeatureState.
type FeatureState string

// List of FeatureState
const (
	FeatureState_off FeatureState = "off"
	FeatureState_on  FeatureState = "on"
)

// Pagination defines model for Pagination.
type Pagination struct {
	ItemCount int  `json:"itemCount"`
	PageCount int  `json:"pageCount"`
	PageIndex int  `json:"pageIndex"`
	PageSize  int  `json:"pageSize"`
	Version   *int `json:"version,omitempty"`
}

// Prerequisite defines model for Prerequisite.
type Prerequisite struct {
	Feature    string   `json:"feature"`
	Variations []string `json:"variations"`
}

// Segment defines model for Segment.
type Segment struct {
	CreatedAt   *int64    `json:"createdAt,omitempty"`
	Environment *string   `json:"environment,omitempty"`
	Excluded    *[]Target `json:"excluded,omitempty"`

	// Unique identifier for the segment.
	Identifier string    `json:"identifier"`
	Included   *[]Target `json:"included,omitempty"`
	ModifiedAt *int64    `json:"modifiedAt,omitempty"`

	// Name of the segment.
	Name string `json:"name"`

	// An array of rules that can cause a user to be included in this segment.
	Rules   *[]Clause `json:"rules,omitempty"`
	Tags    *[]Tag    `json:"tags,omitempty"`
	Version *int64    `json:"version,omitempty"`
}

// Serve defines model for Serve.
type Serve struct {
	Distribution *Distribution `json:"distribution,omitempty"`
	Variation    *string       `json:"variation,omitempty"`
}

// ServingRule defines model for ServingRule.
type ServingRule struct {
	Clauses  []Clause `json:"clauses"`
	Priority int      `json:"priority"`
	RuleId   string   `json:"ruleId"`
	Serve    Serve    `json:"serve"`
}

// A name and value pair.
type Tag struct {
	Name  string  `json:"name"`
	Value *string `json:"value,omitempty"`
}

// Target defines model for Target.
type Target struct {
	Account     string                  `json:"account"`
	Anonymous   *bool                   `json:"anonymous,omitempty"`
	Attributes  *map[string]interface{} `json:"attributes,omitempty"`
	CreatedAt   *int64                  `json:"createdAt,omitempty"`
	Environment string                  `json:"environment"`
	Identifier  string                  `json:"identifier"`
	Name        string                  `json:"name"`
	Org         string                  `json:"org"`
	Project     string                  `json:"project"`
	Segments    *[]Segment              `json:"segments,omitempty"`
}

// TargetMap defines model for TargetMap.
type TargetMap struct {
	Identifier *string `json:"identifier,omitempty"`
	Name       string  `json:"name"`
}

// Variation defines model for Variation.
type Variation struct {
	Description *string `json:"description,omitempty"`
	Identifier  string  `json:"identifier"`
	Name        *string `json:"name,omitempty"`
	Value       string  `json:"value"`
}

// VariationMap defines model for VariationMap.
type VariationMap struct {
	TargetSegments *[]string    `json:"targetSegments,omitempty"`
	Targets        *[]TargetMap `json:"targets,omitempty"`
	Variation      string       `json:"variation"`
}

// WeightedVariation defines model for WeightedVariation.
type WeightedVariation struct {
	Variation string `json:"variation"`
	Weight    int    `json:"weight"`
}

// InternalServerError defines model for InternalServerError.
type InternalServerError Error

// NotFound defines model for NotFound.
type NotFound Error

// Unauthenticated defines model for Unauthenticated.
type Unauthenticated Error

// Unauthorized defines model for Unauthorized.
type Unauthorized Error

// AuthenticateJSONBody defines parameters for Authenticate.
type AuthenticateJSONBody AuthenticationRequest

// StreamParams defines parameters for Stream.
type StreamParams struct {
	APIKey string `json:"API-Key"`
}

// AuthenticateJSONRequestBody defines body for Authenticate for application/json ContentType.
type AuthenticateJSONRequestBody AuthenticateJSONBody
