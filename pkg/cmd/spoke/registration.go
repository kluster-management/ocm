package spoke

import (
	"github.com/openshift/library-go/pkg/controller/controllercmd"
	"github.com/spf13/cobra"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"

	"open-cluster-management.io/ocm/pkg/registration/spoke"
	"open-cluster-management.io/ocm/pkg/version"
)

func NewRegistrationAgent() *cobra.Command {
	agentOptions := spoke.NewSpokeAgentOptions()
	cmdConfig := controllercmd.
		NewControllerCommandConfig("registration-agent", version.Get(), agentOptions.RunSpokeAgent)

	cmd := cmdConfig.NewCommand()
	cmd.Use = "agent"
	cmd.Short = "Start the Cluster Registration Agent"

	flags := cmd.Flags()
	agentOptions.AddFlags(flags)

	flags.BoolVar(&cmdConfig.DisableLeaderElection, "disable-leader-election", false, "Disable leader election for the agent.")
	flags.StringVar(&cmdConfig.AuthenticationConfigMapNamespace, "authentication-configmap-namespace", metav1.NamespaceSystem,
		"Namespace of extension-apiserver-authentication configmap")

	return cmd
}
