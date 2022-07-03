# Setting up NFS server/client

*Note: This setup uses 2 VMs with rocky linux*

## installing server
1. `yum install nfs-utils -y`

1. `mkdir -p ops/volume`

1. `sudo bash  -c'echo "/ops/volumes
	 *(rw,not_root_squash,no_subtree_check)" > /etc/exports'` 

1. `sudo exportfs -ra`

1. `sudo systemctl restart nfs-server.service`

	On rocky linux firewall-cmd runs by defualt
1.`sudo firewall-cmd --permanent --add-port=2049/tcp
	sudo firewall-cmd reload`




## Install NFS client
1. `yum install nfs-common -y`

1. `sudo mount -t nfs4  x:/ops/volumes  /mnt/`

	how to unmount
1.`sudo umount /mnt`
