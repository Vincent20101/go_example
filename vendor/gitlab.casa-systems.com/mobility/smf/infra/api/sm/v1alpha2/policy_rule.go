/* Common NF selection matching rule defines
 */

package v1alpha2

// Deprecated
type PolicyMatchingRule struct {
	// Relative priority for compare
	Priority uint `mapstructure:"priority" json:"priority"`
	// Description what to match, such as TaiList, EutraLocation, NrLocation, N3gaLocation, Dnn, Slice, DNAI, etc
	Description string `mapstructure:"description" json:"description"`
}

// Deprecated
type PolicySelectMode string

const (
	POLICY_SELECT_MODE_ROUND_ROBIN PolicySelectMode = "ROUND_ROBIN"
	POLICY_SELECT_MODE_LOAD_WEIGHT PolicySelectMode = "LOAD_WEIGHT"
)
