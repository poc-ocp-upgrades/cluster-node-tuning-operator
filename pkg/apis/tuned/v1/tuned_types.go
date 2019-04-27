package v1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

type TunedSpec struct {
	Profile		[]TunedProfile		`json:"profile"`
	Recommend	[]TunedRecommend	`json:"recommend"`
}
type TunedProfile struct {
	Name	*string	`json:"name"`
	Data	*string	`json:"data"`
}
type TunedRecommend struct {
	Profile		*string		`json:"profile"`
	Priority	*uint64		`json:"priority"`
	Match		[]TunedMatch	`json:"match,omitempty"`
}
type TunedMatch struct {
	Label	*string		`json:"label"`
	Value	*string		`json:"value,omitempty"`
	Type	*string		`json:"type,omitempty"`
	Match	[]TunedMatch	`json:"match,omitempty"`
}
type TunedStatus struct{}
type Tuned struct {
	metav1.TypeMeta		`json:",inline"`
	metav1.ObjectMeta	`json:"metadata,omitempty"`
	Spec			TunedSpec	`json:"spec,omitempty"`
	Status			TunedStatus	`json:"status,omitempty"`
}
type TunedList struct {
	metav1.TypeMeta	`json:",inline"`
	metav1.ListMeta	`json:"metadata,omitempty"`
	Items		[]Tuned	`json:"items"`
}

func init() {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	SchemeBuilder.Register(&Tuned{}, &TunedList{})
}
