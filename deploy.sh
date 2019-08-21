echo $SSH_PRIV_KEY > priv_key.b64
base64 -d priv_key.b64 > travis.id_rsa
chmod 600 travis.id_rsa
ls -al
rsync -Pav -e "ssh -i travis.id_rsa -p 4430" go-myubuntu abrar@plex.hnxlabs.com:~/bin/