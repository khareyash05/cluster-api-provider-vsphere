/*
Copyright 2020 The Kubernetes Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package cmd

import (
	"os"

	"github.com/pkg/errors"
	"github.com/spf13/cobra"

	"sigs.k8s.io/cluster-api-provider-vsphere/packaging/flavorgen/flavors"
	"sigs.k8s.io/cluster-api-provider-vsphere/packaging/flavorgen/flavors/env"
	"sigs.k8s.io/cluster-api-provider-vsphere/packaging/flavorgen/flavors/util"
)

const flavorFlag = "flavor"

func RootCmd() *cobra.Command {
	rootCmd := &cobra.Command{
		Use:   "flavorgen",
		Short: "flavorgen generates clusterctl templates for Cluster API Provider vSphere",
		RunE: func(command *cobra.Command, args []string) error {
			return RunRoot(command)
		},
	}
	rootCmd.Flags().StringP(flavorFlag, "f", "", "Name of flavor to compile")
	return rootCmd
}

func Execute() {
	if err := RootCmd().Execute(); err != nil {
		os.Exit(1)
	}
}

func RunRoot(command *cobra.Command) error {
	flavor, err := command.Flags().GetString(flavorFlag)
	if err != nil {
		return errors.Wrapf(err, "error accessing flag %s for command %s", flavorFlag, command.Name())
	}
	switch flavor {
	case flavors.VIP:
		util.PrintObjects(flavors.MultiNodeTemplateWithKubeVIP())
	case flavors.ExternalLoadBalancer:
		util.PrintObjects(flavors.MultiNodeTemplateWithExternalLoadBalancer())
	case flavors.ClusterClass:
		util.PrintObjects(flavors.ClusterClassTemplateWithKubeVIP())
	case flavors.ClusterTopology:
		additionalReplacements := []util.Replacement{
			{
				Kind:      "Cluster",
				Name:      "${CLUSTER_NAME}",
				Value:     env.ControlPlaneMachineCountVar,
				FieldPath: []string{"spec", "topology", "controlPlane", "replicas"},
			},
		}
		util.Replacements = append(util.Replacements, additionalReplacements...)
		util.PrintObjects(flavors.ClusterTopologyTemplateKubeVIP())
	case flavors.Ignition:
		util.PrintObjects(flavors.MultiNodeTemplateWithKubeVIPIgnition())
	case flavors.NodeIPAM:
		util.PrintObjects(flavors.MultiNodeTemplateWithKubeVIPNodeIPAM())
	default:
		return errors.Errorf("invalid flavor")
	}
	return nil
}
