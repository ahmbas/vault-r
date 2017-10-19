# vault-r

<h3>Usage</h3>

```
./vault-r --path=/secrets/production/mysql --json=/home/ubuntu/mysecrets.json --token=123-456-789 --host=http://vault.xxcd.org:8200 
```

<h3>JSON Format</h3>

```
{
	"MYSQL_HOST": "localhost",
	"MYSQL_PORT": "3306",
	"MYSQL_DATABASE": "default",
	"MYSQL_USER": "localhost",
	"MYSQL_PASSWORD": "password"
}
```


```
NAME:
   vaultr - Bulk write secrets to vault path

USAGE:
   vault-r [global options] command [command options] [arguments...]

VERSION:
   0.0.1

COMMANDS:
     help, h  Shows a list of commands or help for one command

GLOBAL OPTIONS:
   --path value, -p value   Vault path to add secrets ex: /secret/production/mysql
   --file value, -f value   JSON file to upload. Should be in following format {"<SECRET-NAME>":"<SECRET_VALUE>", ...}
   --host value, -H value   Vault host address. defaults http://127.0.0.1:8200
   --token value, -t value  Vault token with write access
   --help, -h               show help
   --version, -v            print the version
```
