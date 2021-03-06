package activesolution

import (
	"fmt"
	"os"
	"strings"

	"github.com/containerum/chkit/pkg/context"

	"github.com/containerum/chkit/pkg/model/solution"
	"github.com/containerum/chkit/pkg/porta"
	"github.com/containerum/chkit/pkg/util/activekit"
	"github.com/containerum/chkit/pkg/util/ferr"
	"github.com/containerum/chkit/pkg/util/namegen"
	"github.com/containerum/chkit/pkg/util/text"
	"github.com/containerum/kube-client/pkg/model"
	"github.com/sirupsen/logrus"
)

type WizardConfig struct {
	Solution   *solution.Solution
	Namespaces []string
	Templates  []string
	EditName   bool
}

func Wizard(ctx *context.Context, config WizardConfig) solution.Solution {
	var sol = func() solution.Solution {
		if config.Solution != nil {
			return *config.Solution
		}
		return solution.Solution{
			Name:   namegen.Aster() + "-" + namegen.Color(),
			Branch: "master",
		}
	}()
	userEnv := make(map[string]string)
	sol.Env = make(map[string]string)
	for k, v := range sol.Env {
		userEnv[k] = v
	}

	for exit := false; !exit; {
		var envItems activekit.MenuItems
		var i = 0
		for _, env := range sol.EnvironmentVars() {
			envItems = envItems.Append(&activekit.MenuItem{
				Label: fmt.Sprintf("Edit env      : %s", text.Crop(env.String(), 32)),
				Action: func(i int) func() error {
					envItem := env
					return func() error {
						envupd := envMenu(envItem.ToKube())
						delete(sol.Env, envItem.Name)
						delete(userEnv, envItem.Name)
						if envupd != nil {
							sol.Env[envItem.Name] = envupd.Value
							userEnv[envItem.Name] = envupd.Value
						} else {
							envItems.Delete(i)
						}
						return nil
					}
				}(i),
			})
			i++
		}
		var menu = activekit.MenuItems{
			func() *activekit.MenuItem {
				if config.EditName {
					return &activekit.MenuItem{
						Label: fmt.Sprintf("Edit name     : %s", activekit.OrString(sol.Name, "undefined, required")),
						Action: func() error {
							name := activekit.Promt("Type solution name (hit Enter to leave %s): ", activekit.OrString(sol.Name, "empty"))
							name = strings.TrimSpace(name)
							if name == "" {
								return nil
							}
							sol.Name = name
							return nil
						},
					}
				}
				return nil
			}(),
			{
				Label: fmt.Sprintf("Edit template : %s", activekit.OrString(sol.Template, "undefined, required")),
				Action: func() error {
					var menu activekit.MenuItems
					for _, templ := range config.Templates {
						menu = menu.Append(&activekit.MenuItem{
							Label: templ,
							Action: func(templ string) func() error {
								return func() error {
									sol.Template = templ
									if env, err := ctx.GetClient().GetSolutionsTemplatesEnvs(sol.Template, sol.Branch); err == nil {
										for k, v := range env.Env {
											if strings.HasPrefix(v, "{{") {
												env.Env[k] = ""
											}
										}
										sol.Env = env.Env
									}
									for k, v := range userEnv {
										sol.Env[k] = v
									}
									return nil
								}
							}(templ),
						})
					}
					(&activekit.Menu{
						Title: "Select template",
						Items: menu.Append(activekit.MenuItems{
							{
								Label: "Return to previous menu",
								Action: func() error {
									exit = true
									return nil
								},
							},
						}...),
					}).Run()
					return nil
				},
			},
			{
				Label: fmt.Sprintf("Edit branch   : %s", activekit.OrString(sol.Branch, "undefined, required")),
				Action: func() error {
					branch := activekit.Promt("Type branch branch (hit Enter to leave %s): ", activekit.OrString(sol.Branch, "empty"))
					branch = strings.TrimSpace(branch)
					if branch == "" {
						return nil
					}
					sol.Branch = branch
					if env, err := ctx.GetClient().GetSolutionsTemplatesEnvs(sol.Template, sol.Branch); err == nil {
						for k, v := range env.Env {
							if strings.HasPrefix(v, "{{") {
								env.Env[k] = ""
							}
						}
						sol.Env = env.Env
					}
					for k, v := range userEnv {
						sol.Env[k] = v
					}
					return nil
				},
			},
		}
		menu = menu.
			Append(envItems...).
			Append([]*activekit.MenuItem{
				{
					Label: "Add env",
					Action: func() error {
						env := envMenu(model.Env{})
						if env == nil {
							return nil
						}
						sol.Env[env.Name] = env.Value
						userEnv[env.Name] = env.Value
						menu.Append(&activekit.MenuItem{
							Label: fmt.Sprintf("Edit env      : %s", text.Crop(fmt.Sprintf("%s:%q", env.Name, env.Value), 32)),
							Action: func(i int) func() error {
								return func() error {
									env := envMenu(model.Env{
										Name:  env.Name,
										Value: env.Value,
									})
									delete(sol.Env, env.Name)
									delete(userEnv, env.Name)
									if env != nil {
										sol.Env[env.Name] = env.Value
										userEnv[env.Name] = env.Value
									} else {
										envItems.Delete(i)
									}
									return nil
								}
							}(menu.Len()),
						})
						return nil
					},
				},
				{
					Label: "Print to terminal",
					Action: func() error {
						data, err := sol.RenderYAML()
						if err != nil {
							logrus.WithError(err).Errorf("unable to render ingress to yaml")
							activekit.Attention(err.Error())
						}
						border := strings.Repeat("_", text.Width(data))
						fmt.Printf("%s\n%s\n%s\n", border, data, border)
						return nil
					},
				},
				{
					Label: "Export solution to file",
					Action: func() error {
						var fname = activekit.Promt("Type filename: ")
						fname = strings.TrimSpace(fname)
						if fname != "" {
							if err := (porta.Exporter{OutFile: fname}.Export(sol)); err != nil {
								ferr.Printf("unable to export solution:\n%v\n", err)
							}
						}
						return nil
					},
				},
				{
					Label: "Confirm",
					Action: func() error {
						if err := ValidateSolution(sol); err != nil {
							activekit.Attention(err.Error())
							return nil
						}
						exit = true
						return nil
					},
				},
				{
					Label: "Exit",
					Action: func() error {
						os.Exit(0)
						return nil
					},
				},
			}...)
		(&activekit.Menu{
			Items: menu,
		}).Run()
	}
	return sol
}
