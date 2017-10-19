package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"

	cleanhttp "github.com/hashicorp/go-cleanhttp"
	vault "github.com/hashicorp/vault/api"
	cli "github.com/urfave/cli"
)

func main() {

	var secretPath string
	var jsonPath string
	var vaultHost string
	var vaultToken string

	app := cli.NewApp()
	app.Name = "vaultr"
	app.Version = "0.0.1"
	app.Usage = "Bulk write secrets to vault path"
	app.Flags = []cli.Flag{
		cli.StringFlag{
			Name:        "path, p",
			Value:       "",
			Usage:       "Vault path to add secrets ex: /secret/production/mysql",
			Destination: &secretPath,
		},
		cli.StringFlag{
			Name:        "file, f",
			Value:       "/home/ahmedbashir/webapps/menaops_infra/puppet-roles-olx-dbzops/secrets.json",
			Usage:       "JSON file to upload. Should be in following format {\"<SECRET-NAME>\":\"<SECRET_VALUE>\", ...}",
			Destination: &jsonPath,
		},
		cli.StringFlag{
			Name:        "host, H",
			Value:       "",
			Usage:       "Vault host address. defaults http://127.0.0.1:8200",
			Destination: &vaultHost,
		},
		cli.StringFlag{
			Name:        "token, t",
			Value:       "",
			Usage:       "Vault token with write access",
			Destination: &vaultToken,
		},
	}

	app.Action = func(c *cli.Context) error {

		//Read file
		raw, err := ioutil.ReadFile(jsonPath)
		if err != nil {
			fmt.Println(err.Error())
			os.Exit(1)
		}
		var secrets = map[string]string{}

		err = json.Unmarshal(raw, &secrets)
		if err != nil {
			fmt.Println(err.Error())
			os.Exit(1)
		}
		//Create client
		vaultClient, err := vault.NewClient(&vault.Config{Address: vaultHost, HttpClient: cleanhttp.DefaultClient()})
		if err != nil {
			fmt.Printf("Could not connect to vault at %v", vaultHost)
			os.Exit(1)
		}
		vaultClient.SetToken(vaultToken)
		ch := make(chan string)

		for k, v := range secrets {

			go func(v string, k string) {
				secret := make(map[string]interface{})
				secret["value"] = v
				_, err = vaultClient.Logical().Write(k, secret)
				if err != nil {
					ch <- fmt.Sprintf("Could not add %v :%v", k, err.Error())
				}
				ch <- fmt.Sprintf("Added %v to %v", k, secretPath)
			}(v, k)
		}

		for i := 0; i < len(secrets); i++ {
			fmt.Println(<-ch)
		}

		return nil
	}

	app.Run(os.Args)
}
