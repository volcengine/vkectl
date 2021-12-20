package main

import (
	"fmt"
	"net/url"
	"os"
	"reflect"

	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"

	"github.com/volcengine/vkectl/pkg/client/app"
	"github.com/volcengine/vkectl/pkg/client/resource"
	"github.com/volcengine/vkectl/pkg/client/security"
	"github.com/volcengine/vkectl/pkg/version"
)

func main() {
	var rootCmd = &cobra.Command{
		Use:   "vkectl [OPTIONS] MODULE ACTION",
		Short: "Volcengine VKE Command Line Interface Version " + version.Get().Version,
		Example: `export AK=YOUR AK
export SK=YOUR SK
export HOST=YOUR HOST
export REGION=YOUR REGION
vkectl resource GetCluster`,
	}

	data := rootCmd.PersistentFlags().StringP("data", "d", "", "json data of action")
	verbose := rootCmd.PersistentFlags().BoolP("verbose", "v", false, "verbose output")

	ak := os.Getenv("AK")
	sk := os.Getenv("SK")
	host := os.Getenv("HOST")
	region := os.Getenv("REGION")

	addModuleCmd := func(module string, client interface{}) {
		cmd := &cobra.Command{
			Use:     module + " ACTION",
			Short:   "VKE " + module + " module actions",
			Example: "vkectl -d '{}' " + module + " ACTION",
		}
		val := reflect.ValueOf(client)
		tpe := reflect.TypeOf(client)
		for i := 0; i < val.NumMethod(); i++ {
			mval := val.Method(i)
			mtpe := tpe.Method(i)
			actionCmd := &cobra.Command{
				Use: mtpe.Name,
				PreRunE: func(cmd *cobra.Command, args []string) error {
					if ak == "" {
						return errors.New("AK is not set, can set by \"export AK=YOUR AK\"")
					}
					if sk == "" {
						return errors.New("SK is not set, can set by \"export SK=YOUR SK\"")
					}
					if host == "" {
						return errors.New("HOST is not set, can set by \"export AK=YOUR HOST\"")
					}
					if region == "" {
						return errors.New("REGION is not set, can set by \"export REGION=YOUR REGION\"")
					}
					return nil
				},
				RunE: func(cmd *cobra.Command, args []string) error {
					if *verbose {
						logrus.SetLevel(logrus.DebugLevel)
					} else {
						logrus.SetLevel(logrus.PanicLevel)
					}
					ret := mval.Call([]reflect.Value{reflect.ValueOf(*data), reflect.ValueOf(url.Values{})})
					if err := ret[2].Interface().(error); err != nil {
						fmt.Printf("call api: %v", err)
					} else {
						fmt.Println(ret[0].Interface())
					}
					return nil
				},
			}
			cmd.AddCommand(actionCmd)
		}
		rootCmd.AddCommand(cmd)
	}

	addModuleCmd("app", app.NewAPIClient(ak, sk, host, "vke", region))
	addModuleCmd("resource", resource.NewAPIClient(ak, sk, host, "vke", region))
	addModuleCmd("security", security.NewAPIClient(ak, sk, host, "vke", region))

	rootCmd.AddCommand(
		&cobra.Command{
			Use:   "version",
			Short: "Print the version number of vkectl",
			Run: func(cmd *cobra.Command, args []string) {
				fmt.Println(version.Get().Version)
			},
		},
	)
	if err := rootCmd.Execute(); err != nil {
		fmt.Printf("excute error: %v", err)
	}
}
