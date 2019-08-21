echo $SSH_PRIV_KEY > priv_key.b64
base64 -d priv_key.b64 > priv.key
chmod 600 priv.key
ls -al
scp -i priv.key -P 4430 go-myubuntu abrar@plex.hnxlabs.com:~/bin/