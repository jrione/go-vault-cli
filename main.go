package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"reflect"
	"time"

	"github.com/hashicorp/vault-client-go"
	cfg "github.com/jrione/go-vault-cli/config"
)

type Vault struct {
	Ctx    context.Context
	Client *vault.Client
}

func (v *Vault) Save() {
	_, err := v.Client.Write(v.Ctx, "/"+os.Getenv("APP_NAME")+"/data/"+os.Getenv("BRANCH"), map[string]any{
		"data": cfg.GetExternalEnv(),
	})
	if err != nil {
		log.Fatal(err)
		return
	}
	fmt.Println("Secret Successfully Saved!")
}

func (v Vault) Read() {
	res, err := v.Client.Read(v.Ctx, "/"+os.Getenv("APP_NAME")+"/data/"+os.Getenv("BRANCH"))
	if err != nil {
		log.Fatal(err)
		return
	}
	data, ok := res.Data["data"].(map[string]interface{})
	if !ok {
		log.Fatal("No Map Interface")
	}

	for k, v := range data {
		fmt.Printf("%s=%s\n", k, v)
	}
}

func main() {
	cfg.RunFlag()
	cfg.LoadConfig()
	cmd := cfg.RunArgs()

	vaultAddr := fmt.Sprintf("%v://%v:%v", cfg.Env.Vault.Protocol, cfg.Env.Vault.Url, cfg.Env.Vault.Port)
	client, err := vault.New(
		vault.WithAddress(vaultAddr),
		vault.WithRequestTimeout(30*time.Second),
	)
	if err != nil {
		log.Fatal(err)
	}
	client.SetToken(cfg.Env.Vault.Token)

	var v Vault
	v.Ctx = context.Background()
	v.Client = client
	methodName := cfg.ToMethod(cmd)

	defer reflect.ValueOf(&v).MethodByName(methodName).Call([]reflect.Value{})
}
