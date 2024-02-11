# VAULT CLI

Buat read and write secret ke hashicorp vault

> Don't forget to edit config.json (copy from config.example.json)

## Usage:

#### Build App
```bash
go build
```
#### Format Command
```
./go-vault-cli [flags] [cmd] // you can also see this in help command
```

#### Help Command
```bash
./go-vault-cli -help // for help
```

#### write secret to Vault
```
./go-vault-cli -file=/path/to/secret -app=vault_path -branch=kv_name save
```

#### Read Secret from Vault
```
./go-vault-cli -app=vault_path -branch=kv_name read
```