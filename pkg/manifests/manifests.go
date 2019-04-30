package manifests

import (
	"bytes"
	"fmt"
	"io"
	"sort"
	"strings"
	yamlv2 "gopkg.in/yaml.v2"
	appsv1 "k8s.io/api/apps/v1"
	corev1 "k8s.io/api/core/v1"
	rbacv1 "k8s.io/api/rbac/v1"
	"k8s.io/apimachinery/pkg/util/yaml"
	"github.com/golang/glog"
	tunedv1 "github.com/openshift/cluster-node-tuning-operator/pkg/apis/tuned/v1"
	ntoconfig "github.com/openshift/cluster-node-tuning-operator/pkg/config"
)

const (
	nodeLabelsFile		= "/var/lib/tuned/ocp-node-labels.cfg"
	podLabelsFile		= "/var/lib/tuned/ocp-pod-labels.cfg"
	TunedServiceAccount	= "assets/tuned/01-service-account.yaml"
	TunedClusterRole	= "assets/tuned/02-cluster-role.yaml"
	TunedClusterRoleBinding	= "assets/tuned/03-cluster-role-binding.yaml"
	TunedConfigMapProfiles	= "assets/tuned/04-cm-tuned-profiles.yaml"
	TunedConfigMapRecommend	= "assets/tuned/05-cm-tuned-recommend.yaml"
	TunedDaemonSet		= "assets/tuned/06-ds-tuned.yaml"
	TunedCustomResource	= "assets/tuned/default-cr-tuned.yaml"
)

type tunedRecommend struct {
	Profile	string
	Data	string
}

func MustAssetReader(asset string) io.Reader {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return bytes.NewReader(MustAsset(asset))
}

type Factory struct{}

func NewFactory() *Factory {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return &Factory{}
}
func (f *Factory) TunedServiceAccount() (*corev1.ServiceAccount, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	sa, err := NewServiceAccount(MustAssetReader(TunedServiceAccount))
	if err != nil {
		return nil, err
	}
	return sa, nil
}
func (f *Factory) TunedClusterRole() (*rbacv1.ClusterRole, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	cr, err := NewClusterRole(MustAssetReader(TunedClusterRole))
	if err != nil {
		return nil, err
	}
	return cr, nil
}
func (f *Factory) TunedClusterRoleBinding() (*rbacv1.ClusterRoleBinding, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	crb, err := NewClusterRoleBinding(MustAssetReader(TunedClusterRoleBinding))
	if err != nil {
		return nil, err
	}
	return crb, nil
}
func (f *Factory) TunedConfigMapProfiles(tunedArray []tunedv1.Tuned) (*corev1.ConfigMap, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	cm, err := NewConfigMap(MustAssetReader(TunedConfigMapProfiles))
	if err != nil {
		return nil, err
	}
	m := map[string]string{}
	for _, tuned := range tunedArray {
		tunedConfigMapProfiles(&tuned, m)
	}
	tunedOcpProfiles, err := yamlv2.Marshal(&m)
	if err != nil {
		glog.Fatalf("error: %v", err)
	}
	cm.Data["tuned-profiles-data"] = string(tunedOcpProfiles)
	return cm, nil
}
func (f *Factory) TunedConfigMapRecommend(tunedArray []tunedv1.Tuned) (*corev1.ConfigMap, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	var (
		sb		strings.Builder
		aRecommendAll	[]tunedv1.TunedRecommend
	)
	cm, err := NewConfigMap(MustAssetReader(TunedConfigMapRecommend))
	if err != nil {
		return nil, err
	}
	for _, tuned := range tunedArray {
		if tuned.Spec.Recommend != nil {
			aRecommendAll = append(aRecommendAll, tuned.Spec.Recommend...)
		}
	}
	sort.Slice(aRecommendAll, func(i, j int) bool {
		if aRecommendAll[i].Priority != nil && aRecommendAll[j].Priority != nil {
			return *aRecommendAll[i].Priority < *aRecommendAll[j].Priority
		}
		return false
	})
	i := 0
	for _, r := range aRecommendAll {
		aRecommend := recommendWalk(&r)
		sb.WriteString(toRecommendConf(aRecommend, &i))
	}
	cm.Data["tuned-ocp-recommend"] = sb.String()
	return cm, nil
}
func (f *Factory) TunedDaemonSet() (*appsv1.DaemonSet, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	ds, err := NewDaemonSet(MustAssetReader(TunedDaemonSet))
	imageTuned := ntoconfig.NodeTunedImage()
	ds.Spec.Template.Spec.Containers[0].Image = imageTuned
	if err != nil {
		return nil, err
	}
	return ds, nil
}
func (f *Factory) TunedCustomResource() (*tunedv1.Tuned, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	cr, err := NewTuned(MustAssetReader(TunedCustomResource))
	if err != nil {
		return nil, err
	}
	return cr, nil
}
func NewServiceAccount(manifest io.Reader) (*corev1.ServiceAccount, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	sa := corev1.ServiceAccount{}
	if err := yaml.NewYAMLOrJSONDecoder(manifest, 100).Decode(&sa); err != nil {
		return nil, err
	}
	return &sa, nil
}
func NewClusterRole(manifest io.Reader) (*rbacv1.ClusterRole, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	cr := rbacv1.ClusterRole{}
	if err := yaml.NewYAMLOrJSONDecoder(manifest, 100).Decode(&cr); err != nil {
		return nil, err
	}
	return &cr, nil
}
func NewClusterRoleBinding(manifest io.Reader) (*rbacv1.ClusterRoleBinding, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	crb := rbacv1.ClusterRoleBinding{}
	if err := yaml.NewYAMLOrJSONDecoder(manifest, 100).Decode(&crb); err != nil {
		return nil, err
	}
	return &crb, nil
}
func NewConfigMap(manifest io.Reader) (*corev1.ConfigMap, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	cm := corev1.ConfigMap{}
	if err := yaml.NewYAMLOrJSONDecoder(manifest, 100).Decode(&cm); err != nil {
		return nil, err
	}
	return &cm, nil
}
func NewDaemonSet(manifest io.Reader) (*appsv1.DaemonSet, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	ds := appsv1.DaemonSet{}
	if err := yaml.NewYAMLOrJSONDecoder(manifest, 100).Decode(&ds); err != nil {
		return nil, err
	}
	return &ds, nil
}
func NewTuned(manifest io.Reader) (*tunedv1.Tuned, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	o := tunedv1.Tuned{}
	if err := yaml.NewYAMLOrJSONDecoder(manifest, 100).Decode(&o); err != nil {
		return nil, err
	}
	return &o, nil
}
func toRecommendLine(match *tunedv1.TunedMatch) string {
	_logClusterCodePath()
	defer _logClusterCodePath()
	var (
		sb		strings.Builder
		labelsFile	string
	)
	if match.Type != nil {
		if *match.Type == "pod" {
			labelsFile = podLabelsFile
		} else {
			glog.Errorf("Label type: %s unknown, using \"node\".", *match.Type)
		}
	} else {
		labelsFile = nodeLabelsFile
	}
	sb.WriteString(labelsFile)
	sb.WriteString("=.*")
	if match.Label != nil {
		sb.WriteString("\\b")
		sb.WriteString(*match.Label)
		sb.WriteString("=")
		if match.Value != nil {
			sb.WriteString(*match.Value)
			sb.WriteString("\\n")
		}
	} else {
	}
	return sb.String()
}
func toRecommendConf(recommend []tunedRecommend, i *int) string {
	_logClusterCodePath()
	defer _logClusterCodePath()
	var sb strings.Builder
	for _, r := range recommend {
		fmt.Fprintf(&sb, "[%s,%d]\n", r.Profile, *i)
		sb.WriteString(r.Data)
		sb.WriteString("\n\n")
		*i++
	}
	return sb.String()
}
func matchWalk(match *tunedv1.TunedMatch, p tunedRecommend) []tunedRecommend {
	_logClusterCodePath()
	defer _logClusterCodePath()
	var (
		sb		strings.Builder
		aRecommend	[]tunedRecommend
	)
	if match.Label != nil {
		for _, m := range match.Match {
			sb.WriteString(p.Data)
			if len(p.Data) > 0 {
				sb.WriteString("\n")
			}
			sb.WriteString(toRecommendLine(&m))
			if len(m.Match) == 0 {
				aRecommend = append(aRecommend, tunedRecommend{Profile: p.Profile, Data: sb.String()})
			} else {
				aRecommend = append(aRecommend, matchWalk(&m, tunedRecommend{Profile: p.Profile, Data: sb.String()})...)
			}
			sb.Reset()
		}
	}
	return aRecommend
}
func recommendWalk(r *tunedv1.TunedRecommend) []tunedRecommend {
	_logClusterCodePath()
	defer _logClusterCodePath()
	var aRecommend []tunedRecommend
	if r.Profile != nil {
		if len(r.Match) == 0 {
			sRecommend := toRecommendLine(&tunedv1.TunedMatch{})
			aRecommend = append(aRecommend, tunedRecommend{Profile: *r.Profile, Data: sRecommend})
		}
		for _, m := range r.Match {
			sRecommend := toRecommendLine(&m)
			if len(m.Match) == 0 {
				aRecommend = append(aRecommend, tunedRecommend{Profile: *r.Profile, Data: sRecommend})
			} else {
				aRecommend = append(aRecommend, matchWalk(&m, tunedRecommend{Profile: *r.Profile, Data: sRecommend})...)
			}
		}
	} else {
	}
	return aRecommend
}
func tunedConfigMapProfiles(tuned *tunedv1.Tuned, m map[string]string) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	if tuned.Spec.Profile != nil {
		for _, v := range tuned.Spec.Profile {
			if v.Name != nil && v.Data != nil {
				if _, found := m[*v.Name]; found {
					glog.Warningf("WARNING: Duplicate profile %s", *v.Name)
				}
				m[*v.Name] = *v.Data
			}
		}
	}
}
