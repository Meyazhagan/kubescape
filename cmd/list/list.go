package list

import (
	"fmt"
	"strings"

	"github.com/armosec/kubescape/cautils"
	"github.com/armosec/kubescape/cautils/logger"
	"github.com/armosec/kubescape/core/core"
	"github.com/armosec/kubescape/core/meta"
	v1 "github.com/armosec/kubescape/core/meta/datastructures/v1"
	"github.com/spf13/cobra"
)

var (
	listExample = `
  # List default supported frameworks names
  kubescape list frameworks
  
  # List all supported frameworks names
  kubescape list frameworks --account <account id>
	
  # List all supported controls names
  kubescape list controls

  # List all supported controls ids
  kubescape list controls --id 
  
  Control documentation:
  https://hub.armo.cloud/docs/controls
`
)

func GetListCmd(ks meta.IKubescape) *cobra.Command {
	var listPolicies = v1.ListPolicies{}

	listCmd := &cobra.Command{
		Use:     "list <policy> [flags]",
		Short:   "List frameworks/controls will list the supported frameworks and controls",
		Long:    ``,
		Example: listExample,
		Args: func(cmd *cobra.Command, args []string) error {
			supported := strings.Join(core.ListSupportActions(), ",")

			if len(args) < 1 {
				return fmt.Errorf("policy type requeued, supported: %s", supported)
			}
			if cautils.StringInSlice(core.ListSupportActions(), args[0]) == cautils.ValueNotFound {
				return fmt.Errorf("invalid parameter '%s'. Supported parameters: %s", args[0], supported)
			}
			return nil
		},
		RunE: func(cmd *cobra.Command, args []string) error {
			listPolicies.Target = args[0]

			if err := ks.List(&listPolicies); err != nil {
				logger.L().Fatal(err.Error())
			}
			return nil
		},
	}
	listCmd.PersistentFlags().StringVar(&listPolicies.Account, "account", "", "Armo portal account ID. Default will load account ID from configMap or config file")
	listCmd.PersistentFlags().StringVar(&listPolicies.Format, "format", "pretty-print", "output format. supported: 'pretty-printer'/'json'")
	listCmd.PersistentFlags().BoolVarP(&listPolicies.ListIDs, "id", "", false, "List control ID's instead of controls names")

	return listCmd
}
