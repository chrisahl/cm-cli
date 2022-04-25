// Copyright Contributors to the Open Cluster Management project
package hosted

import (
	genericclioptionscm "github.com/stolostron/cm-cli/pkg/genericclioptions"
	"k8s.io/cli-runtime/pkg/genericclioptions"
	"k8s.io/kubectl/pkg/cmd/get"
)

type Options struct {
	//CMFlags: The generic optiosn from the cm cli-runtime.
	CMFlags         *genericclioptionscm.CMFlags
	Hosting         string
	Hosted          string
	Timeout         int
	GetOptions      *get.GetOptions
	WithCredentials bool
}

func newOptions(cmFlags *genericclioptionscm.CMFlags, streams genericclioptions.IOStreams) *Options {
	return &Options{
		CMFlags:    cmFlags,
		GetOptions: get.NewGetOptions("cm", streams),
	}
}
