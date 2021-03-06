package deplactive

import (
	"fmt"
	"path/filepath"

	"io"

	"github.com/containerum/chkit/pkg/model/container"
	"github.com/containerum/chkit/pkg/util/activekit"
	"github.com/containerum/kube-client/pkg/model"
)

func componentEditContainerConfigmaps(cont *container.Container, configmaps []string) activekit.MenuItems {
	var menuItems = make(activekit.MenuItems, 0, len(cont.ConfigMaps))
	for _, vol := range cont.ConfigMaps {
		menuItems = menuItems.Append(componentContainerConfigmap(&vol, configmaps))
	}
	return menuItems.Append(&activekit.MenuItem{
		Label: "Mount configmap",
		Action: func() error {
			var vol = &model.ContainerVolume{}
			if componentContainerConfigmap(vol, configmaps).Action() == nil {
				cont.ConfigMaps = append(cont.ConfigMaps, *vol)
			}
			return nil
		},
	})
}

func componentContainerConfigmap(oldConfigmap *model.ContainerVolume, configmaps []string) *activekit.MenuItem {
	var configmap = *oldConfigmap
	return &activekit.MenuItem{
		Label: fmt.Sprintf("Edit configmap mount %s", func() string {
			if configmap.MountPath != "" && configmap.Name != "" {
				return configmap.Name + " -> " + configmap.MountPath
			}
			if configmap.Name != "" {
				return configmap.Name
			}
			if configmap.MountPath != "" {
				return configmap.MountPath
			}
			return ""
		}()),
		Action: func() error {
			for exit := false; !exit; {
				_, err := (&activekit.Menu{
					Title: "Deployment -> Container -> Configmap",
					Items: activekit.MenuItems{
						{
							Label: fmt.Sprintf("Edit path %s",
								activekit.OrString(configmap.MountPath, "undefined, required")),
							Action: activekit.HandleString(
								fmt.Sprintf("Type mount path (hit Enter to leave %s): ",
									activekit.OrString(configmap.MountPath, "empty")),
								func(pathString string) bool {
									if !filepath.IsAbs(pathString) {
										fmt.Printf("Mount path must be absolute\n")
										return true
									}
									if pathString == "" {
										return true
									}
									configmap.MountPath = pathString
									return false
								}),
						},
						{
							Label: fmt.Sprintf("Set configmap : %s",
								activekit.OrString(configmap.Name, "undefined. required")),
							Action: func() error {
								_, err := (&activekit.Menu{
									Title: "Select configmap",
									Items: activekit.StringSelector(configmaps, func(configmapName string) error {
										configmap.Name = configmapName
										if configmap.MountPath == "" {
											configmap.MountPath = "/etc/" + configmap.Name
										}
										return nil
									}).Append(&activekit.MenuItem{
										Label: fmt.Sprintf("Return to previous menu, leave %s",
											activekit.OrString(configmap.Name, "empty")),
									}),
								}).Run()
								return err
							},
						},
						{
							Label: "Confirm",
							Action: func() error {
								if configmap.Name != "" && configmap.MountPath != "" {
									exit = true
									*oldConfigmap = configmap
									return nil
								}
								if configmap.Name == "" {
									fmt.Printf("Volume name must be non-empty!\n")
								}
								if configmap.MountPath == "" {
									fmt.Printf("Volume mountpath must be non-empty!\n")
								}
								exit = true
								return nil
							},
						},
						{
							Label: "Drop changes, return to previous menu",
							Action: func() error {
								exit = true
								return io.EOF
							},
						},
					},
				}).Run()
				if err != nil {
					return err
				}
			}
			return nil
		},
	}
}
