// Code generated by go-bindata.
// sources:
// pkg/hub/spokecluster/manifests/spokecluster-clusterrole.yaml
// pkg/hub/spokecluster/manifests/spokecluster-clusterrolebinding.yaml
// pkg/hub/spokecluster/manifests/spokecluster-namespace.yaml
// pkg/hub/spokecluster/manifests/spokecluster-work-role.yaml
// pkg/hub/spokecluster/manifests/spokecluster-work-rolebinding.yaml
// DO NOT EDIT!

package bindata

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
	"time"
)

type asset struct {
	bytes []byte
	info  os.FileInfo
}

type bindataFileInfo struct {
	name    string
	size    int64
	mode    os.FileMode
	modTime time.Time
}

func (fi bindataFileInfo) Name() string {
	return fi.name
}
func (fi bindataFileInfo) Size() int64 {
	return fi.size
}
func (fi bindataFileInfo) Mode() os.FileMode {
	return fi.mode
}
func (fi bindataFileInfo) ModTime() time.Time {
	return fi.modTime
}
func (fi bindataFileInfo) IsDir() bool {
	return false
}
func (fi bindataFileInfo) Sys() interface{} {
	return nil
}

var _pkgHubSpokeclusterManifestsSpokeclusterClusterroleYaml = []byte(`apiVersion: rbac.authorization.k8s.io/v1
kind: ClusterRole
metadata:
  name: system:open-cluster-management:spokecluster:{{ .SpokeClusterName }}
rules:
# Allow spoke agent to rotate its certificate
- apiGroups: ["certificates.k8s.io"]
  resources: ["certificatesigningrequests"]
  verbs: ["create", "get", "list", "watch"]
- apiGroups: ["register.open-cluster-management.io"]
  resources: ["spokeclusters/clientcertificates"]
  verbs: ["renew"]
# Allow spoke agent to get/list/update/watch its owner spoke cluster
- apiGroups: ["cluster.open-cluster-management.io"]
  resources: ["spokeclusters"]
  resourceNames: ["{{ .SpokeClusterName }}"]
  verbs: ["get", "list", "update", "watch"]
# Allow spoke agent to update the status of its owner spoke cluster
- apiGroups: ["cluster.open-cluster-management.io"]
  resources: ["spokeclusters/status"]
  resourceNames: ["{{ .SpokeClusterName }}"]
  verbs: ["patch", "update"]
`)

func pkgHubSpokeclusterManifestsSpokeclusterClusterroleYamlBytes() ([]byte, error) {
	return _pkgHubSpokeclusterManifestsSpokeclusterClusterroleYaml, nil
}

func pkgHubSpokeclusterManifestsSpokeclusterClusterroleYaml() (*asset, error) {
	bytes, err := pkgHubSpokeclusterManifestsSpokeclusterClusterroleYamlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "pkg/hub/spokecluster/manifests/spokecluster-clusterrole.yaml", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _pkgHubSpokeclusterManifestsSpokeclusterClusterrolebindingYaml = []byte(`apiVersion: rbac.authorization.k8s.io/v1
kind: ClusterRoleBinding
metadata:
  name: system:open-cluster-management:spokecluster:{{ .SpokeClusterName }}
roleRef:
  apiGroup: rbac.authorization.k8s.io
  kind: ClusterRole
  name: system:open-cluster-management:spokecluster:{{ .SpokeClusterName }}
subjects:
- kind: Group
  apiGroup: rbac.authorization.k8s.io
  name: system:open-cluster-management:{{ .SpokeClusterName }}
`)

func pkgHubSpokeclusterManifestsSpokeclusterClusterrolebindingYamlBytes() ([]byte, error) {
	return _pkgHubSpokeclusterManifestsSpokeclusterClusterrolebindingYaml, nil
}

func pkgHubSpokeclusterManifestsSpokeclusterClusterrolebindingYaml() (*asset, error) {
	bytes, err := pkgHubSpokeclusterManifestsSpokeclusterClusterrolebindingYamlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "pkg/hub/spokecluster/manifests/spokecluster-clusterrolebinding.yaml", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _pkgHubSpokeclusterManifestsSpokeclusterNamespaceYaml = []byte(`apiVersion: v1
kind: Namespace
metadata:
  name: {{ .SpokeClusterName }}
`)

func pkgHubSpokeclusterManifestsSpokeclusterNamespaceYamlBytes() ([]byte, error) {
	return _pkgHubSpokeclusterManifestsSpokeclusterNamespaceYaml, nil
}

func pkgHubSpokeclusterManifestsSpokeclusterNamespaceYaml() (*asset, error) {
	bytes, err := pkgHubSpokeclusterManifestsSpokeclusterNamespaceYamlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "pkg/hub/spokecluster/manifests/spokecluster-namespace.yaml", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _pkgHubSpokeclusterManifestsSpokeclusterWorkRoleYaml = []byte(`apiVersion: rbac.authorization.k8s.io/v1
kind: Role
metadata:
  name: {{ .SpokeClusterName }}:spoke-work
  namespace: {{ .SpokeClusterName }}
rules:
# Allow spoke work agent to send event to hub
- apiGroups: ["", "events.k8s.io"]
  resources: ["events"]
  verbs: ["create", "patch", "update"]
# Allow spoke work agent to get/list/watch/update manifestworks
- apiGroups: ["work.open-cluster-management.io"]
  resources: ["manifestworks"]
  verbs: ["get", "list", "watch", "update"]
# Allow spoke work agent to update the status of manifestwork
- apiGroups: ["work.open-cluster-management.io"]
  resources: ["manifestworks/status"]
  verbs: ["patch", "update"]
`)

func pkgHubSpokeclusterManifestsSpokeclusterWorkRoleYamlBytes() ([]byte, error) {
	return _pkgHubSpokeclusterManifestsSpokeclusterWorkRoleYaml, nil
}

func pkgHubSpokeclusterManifestsSpokeclusterWorkRoleYaml() (*asset, error) {
	bytes, err := pkgHubSpokeclusterManifestsSpokeclusterWorkRoleYamlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "pkg/hub/spokecluster/manifests/spokecluster-work-role.yaml", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _pkgHubSpokeclusterManifestsSpokeclusterWorkRolebindingYaml = []byte(`apiVersion: rbac.authorization.k8s.io/v1
kind: RoleBinding
metadata:
  name: {{ .SpokeClusterName }}:spoke-work
  namespace: {{ .SpokeClusterName }}
roleRef:
  apiGroup: rbac.authorization.k8s.io
  kind: Role
  name: {{ .SpokeClusterName }}:spoke-work
subjects:
  # Bind the role with spoke agent user group, the role will be as a common role for all spoke agents
  # TODO: we will consider bind a specific role for each spoke agent by spoke agent name
  - kind: Group
    apiGroup: rbac.authorization.k8s.io
    name: system:open-cluster-management:{{ .SpokeClusterName }}
`)

func pkgHubSpokeclusterManifestsSpokeclusterWorkRolebindingYamlBytes() ([]byte, error) {
	return _pkgHubSpokeclusterManifestsSpokeclusterWorkRolebindingYaml, nil
}

func pkgHubSpokeclusterManifestsSpokeclusterWorkRolebindingYaml() (*asset, error) {
	bytes, err := pkgHubSpokeclusterManifestsSpokeclusterWorkRolebindingYamlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "pkg/hub/spokecluster/manifests/spokecluster-work-rolebinding.yaml", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

// Asset loads and returns the asset for the given name.
// It returns an error if the asset could not be found or
// could not be loaded.
func Asset(name string) ([]byte, error) {
	cannonicalName := strings.Replace(name, "\\", "/", -1)
	if f, ok := _bindata[cannonicalName]; ok {
		a, err := f()
		if err != nil {
			return nil, fmt.Errorf("Asset %s can't read by error: %v", name, err)
		}
		return a.bytes, nil
	}
	return nil, fmt.Errorf("Asset %s not found", name)
}

// MustAsset is like Asset but panics when Asset would return an error.
// It simplifies safe initialization of global variables.
func MustAsset(name string) []byte {
	a, err := Asset(name)
	if err != nil {
		panic("asset: Asset(" + name + "): " + err.Error())
	}

	return a
}

// AssetInfo loads and returns the asset info for the given name.
// It returns an error if the asset could not be found or
// could not be loaded.
func AssetInfo(name string) (os.FileInfo, error) {
	cannonicalName := strings.Replace(name, "\\", "/", -1)
	if f, ok := _bindata[cannonicalName]; ok {
		a, err := f()
		if err != nil {
			return nil, fmt.Errorf("AssetInfo %s can't read by error: %v", name, err)
		}
		return a.info, nil
	}
	return nil, fmt.Errorf("AssetInfo %s not found", name)
}

// AssetNames returns the names of the assets.
func AssetNames() []string {
	names := make([]string, 0, len(_bindata))
	for name := range _bindata {
		names = append(names, name)
	}
	return names
}

// _bindata is a table, holding each asset generator, mapped to its name.
var _bindata = map[string]func() (*asset, error){
	"pkg/hub/spokecluster/manifests/spokecluster-clusterrole.yaml":        pkgHubSpokeclusterManifestsSpokeclusterClusterroleYaml,
	"pkg/hub/spokecluster/manifests/spokecluster-clusterrolebinding.yaml": pkgHubSpokeclusterManifestsSpokeclusterClusterrolebindingYaml,
	"pkg/hub/spokecluster/manifests/spokecluster-namespace.yaml":          pkgHubSpokeclusterManifestsSpokeclusterNamespaceYaml,
	"pkg/hub/spokecluster/manifests/spokecluster-work-role.yaml":          pkgHubSpokeclusterManifestsSpokeclusterWorkRoleYaml,
	"pkg/hub/spokecluster/manifests/spokecluster-work-rolebinding.yaml":   pkgHubSpokeclusterManifestsSpokeclusterWorkRolebindingYaml,
}

// AssetDir returns the file names below a certain
// directory embedded in the file by go-bindata.
// For example if you run go-bindata on data/... and data contains the
// following hierarchy:
//     data/
//       foo.txt
//       img/
//         a.png
//         b.png
// then AssetDir("data") would return []string{"foo.txt", "img"}
// AssetDir("data/img") would return []string{"a.png", "b.png"}
// AssetDir("foo.txt") and AssetDir("notexist") would return an error
// AssetDir("") will return []string{"data"}.
func AssetDir(name string) ([]string, error) {
	node := _bintree
	if len(name) != 0 {
		cannonicalName := strings.Replace(name, "\\", "/", -1)
		pathList := strings.Split(cannonicalName, "/")
		for _, p := range pathList {
			node = node.Children[p]
			if node == nil {
				return nil, fmt.Errorf("Asset %s not found", name)
			}
		}
	}
	if node.Func != nil {
		return nil, fmt.Errorf("Asset %s not found", name)
	}
	rv := make([]string, 0, len(node.Children))
	for childName := range node.Children {
		rv = append(rv, childName)
	}
	return rv, nil
}

type bintree struct {
	Func     func() (*asset, error)
	Children map[string]*bintree
}

var _bintree = &bintree{nil, map[string]*bintree{
	"pkg": {nil, map[string]*bintree{
		"hub": {nil, map[string]*bintree{
			"spokecluster": {nil, map[string]*bintree{
				"manifests": {nil, map[string]*bintree{
					"spokecluster-clusterrole.yaml":        {pkgHubSpokeclusterManifestsSpokeclusterClusterroleYaml, map[string]*bintree{}},
					"spokecluster-clusterrolebinding.yaml": {pkgHubSpokeclusterManifestsSpokeclusterClusterrolebindingYaml, map[string]*bintree{}},
					"spokecluster-namespace.yaml":          {pkgHubSpokeclusterManifestsSpokeclusterNamespaceYaml, map[string]*bintree{}},
					"spokecluster-work-role.yaml":          {pkgHubSpokeclusterManifestsSpokeclusterWorkRoleYaml, map[string]*bintree{}},
					"spokecluster-work-rolebinding.yaml":   {pkgHubSpokeclusterManifestsSpokeclusterWorkRolebindingYaml, map[string]*bintree{}},
				}},
			}},
		}},
	}},
}}

// RestoreAsset restores an asset under the given directory
func RestoreAsset(dir, name string) error {
	data, err := Asset(name)
	if err != nil {
		return err
	}
	info, err := AssetInfo(name)
	if err != nil {
		return err
	}
	err = os.MkdirAll(_filePath(dir, filepath.Dir(name)), os.FileMode(0755))
	if err != nil {
		return err
	}
	err = ioutil.WriteFile(_filePath(dir, name), data, info.Mode())
	if err != nil {
		return err
	}
	err = os.Chtimes(_filePath(dir, name), info.ModTime(), info.ModTime())
	if err != nil {
		return err
	}
	return nil
}

// RestoreAssets restores an asset under the given directory recursively
func RestoreAssets(dir, name string) error {
	children, err := AssetDir(name)
	// File
	if err != nil {
		return RestoreAsset(dir, name)
	}
	// Dir
	for _, child := range children {
		err = RestoreAssets(dir, filepath.Join(name, child))
		if err != nil {
			return err
		}
	}
	return nil
}

func _filePath(dir, name string) string {
	cannonicalName := strings.Replace(name, "\\", "/", -1)
	return filepath.Join(append([]string{dir}, strings.Split(cannonicalName, "/")...)...)
}
