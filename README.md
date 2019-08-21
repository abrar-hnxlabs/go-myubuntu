# go-myubuntu
My ubuntu utils written in golang

# Travis Deploys

## Rsync remote deploy without password.
### Generating and adding keys to server
Create a key pair using the below command. Use a custom filename.
```
ssh-keygen
```
Once the key pair is created copy the id to remote server
```
ssh-copy-id -i travis.id_rsa abrar@plex.hnxlabs.com
```
Note: use private key for identity file

### Encoding
Encode the private key to base64
```
base64 travis.id_rsa > travis.id_rsa.b64
```
Copy the env variable to CI system.
Decode the base64 env variable to create the private key file again.
```
echo $PRIV_KEY > ~/.ssh/travis.id_rsa.b64
base64 -d ~/.ssh/travis.id_rsa.b64 > ~/.ssh/travis.id_rsa
```
Ssh to remote system using the identity file just decoded from env.

```
ssh -i ~/.ssh/travis.id_rsa abrar@plex.hnxlabs.com
```


