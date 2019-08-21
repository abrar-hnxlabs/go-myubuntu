echo $SSH_PRIV_KEY > priv_key.b64
base64 -d priv_key.b64 > priv.key
scp -y -i priv.key -P 4430 go-myubuntu abrar@plex.hnxlabs.com:~/bin/