echo "writing private file"
echo $SSH_PRIV_KEY > ~/.ssh/priv_key.b64
base64 -d ~/.ssh/priv_key.b64 > ~/.ssh/travis.id_rsa
chmod 600 ~/.ssh/travis.id_rsa
echo "starting transfer"
rsync -Pav -e "ssh -i ~/.ssh/travis.id_rsa -p 4430" go-myubuntu abrar@plex.hnxlabs.com:~/bin/
echo "tranfer done"