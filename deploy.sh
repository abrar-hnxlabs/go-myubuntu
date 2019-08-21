echo $SSH_PRIV_KEY > priv_key.b64
base64 -d priv_key.b64 > priv.key
cat priv.key
scp -i priv.key -P 4430 go-myubuntu abrar@plex.hnxlabs.com:~/bin/