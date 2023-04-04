package main

import (
	"fmt"
	"os"
	"os/exec"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"gopkg.in/yaml.v2"
)

type Device struct {
	Name    string `yaml:"name"`
	Host    string `yaml:"host"`
	Port    int    `yaml:"port"`
	AuthKey string `yaml:"auth_key"`
}

func main() {
	rootCmd := &cobra.Command{
		Use:   "netinfo",
		Short: "Gather network information",
		Run: func(cmd *cobra.Command, args []string) {
			devices := viper.GetStringSlice("devices")

			for _, device := range devices {
				d := Device{}
				err := yaml.Unmarshal([]byte(device), &d)
				if err != nil {
					fmt.Printf("Error decoding YAML: %v", err)
					continue
				}

				fmt.Printf("Running commands on device: %s\n", d.Name)

				cmds := []string{"show ip route", "show interfaces"}

				for _, c := range cmds {
					fmt.Printf("Running command: %s\n", c)

					output, err := exec.Command("vtysh", "-c", c).Output()
					if err != nil {
						fmt.Printf("Error running command: %v", err)
						continue
					}

					fmt.Println(string(output))
				}
			}
		},
	}

	rootCmd.PersistentFlags().StringSlice("devices", []string{}, "List of devices in YAML format")
	viper.BindPFlags(rootCmd.PersistentFlags())

	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
