if [[ $EUID -ne 0 ]] ; then
   echo -e  "Run tool as root.\n\tDetails : sudo $0  "
fi
apt install go -y

