package logrlint

import "sort"

type loggerChecker struct {
	packageImport string
	funcs         stringSet
}

var loggerCheckersByName = loggerCheckerMap{
	"logr": {
		packageImport: "github.com/go-logr/logr",
		funcs: newStringSet(
			"(github.com/go-logr/logr.Logger).Error",
			"(github.com/go-logr/logr.Logger).Info",
			"(github.com/go-logr/logr.Logger).WithValues"),
	},
	"klog": {
		packageImport: "k8s.io/klog/v2",
		funcs: newStringSet(
			"k8s.io/klog/v2.InfoS",
			"k8s.io/klog/v2.InfoSDepth",
			"k8s.io/klog/v2.ErrorS",
			"(k8s.io/klog/v2.Verbose).InfoS",
			"(k8s.io/klog/v2.Verbose).InfoSDepth",
			"(k8s.io/klog/v2.Verbose).ErrorS",
		),
	},
}

type loggerCheckerMap map[string]loggerChecker

func (m loggerCheckerMap) Names() []string {
	names := make([]string, 0, len(m))
	for name := range m {
		names = append(names, name)
	}
	sort.Strings(names)
	return names
}