## zcryp

An open source implementation of the zcryp program used on the Olive O3HD, O4HD, and O6HD music player to decrypt firmware updates from Olive.  This tool is part of the Z-App framework which Olive used as the basis for their software.  This implementation is feature-complete but not bug-complete.

### Installing

[Install the go compiler for your operating system](http://golang.org/doc/install) and [configure your workspace environment](http://golang.org/doc/install#gopath).  Download zcryp using the following command.

```ShellSession
[dogshoes@oxcart ~]# go get github.com/dogshoes/zcryp
```

Finally, ensure that your $GOPATH/bin is in $PATH for ease of use.

```ShellSession
[dogshoes@oxcart ~]# export PATH=$PATH:$GOPATH/bin
```

### Obtaining the encryption key

The key zcryp needs can be found in /zapp/zbase/bin/Upgrade.sh on the olive player.  Activate the player's [SSH server through the manufacturer's backdoor](http://www.avsforum.com/forum/153-cd-players-dedicated-music-transports/1091695-official-olive-thread-opus-4-opus-6-melody-2-olive-2-olive-4-4hd-06hd-76.html#post22850481) and download the file via SCP.

```ShellSession
[dogshoes@oxcart ~]# scp root@myolive.local:/zapp/zbase/bin/Upgrade.sh .
[dogshoes@oxcart ~]# grep \$cryptCmd Upgrade.sh | awk '{ print $9 }' | uniq
```

### Use

Download and unpack the system update from Olive.  This container TAR archive is unencrypted.

```ShellSession
[dogshoes@oxcart ~]# mkdir scratch
[dogshoes@oxcart ~]# cd scratch
[dogshoes@oxcart scratch]# wget http://downloads.olive.us/SWUpgrade/olive4hd/432/SWUpgrade.tar
[dogshoes@oxcart scratch]# tar -xvf SWUpgrade.tar
```

You should end up with three files: SWUpgrade.xml which is unencrypted and contains information about the upgrade, FSSoftware.tar.gz which is encrypted and contains updates to be placed on the hard drive, and uImage.tar.gz which is encrypted and contains a new kernel.  These files can now be unencrypted and examined as needed.

```ShellSession
[dogshoes@oxcart scratch]# zcryp -i FSSoftware.tar.gz -o FSSoftware.unencrypted.tar.gz -m 1 -k encryptionkeyfromUpgradeSH
```

The same principles here can be used for ODBInstall.tar.gz and other encrypted resources consumed by the Olive players.

Since the method of encryption is a simple XOR, feeding an unencrypted file into zcryp will encrypt it.  This can be used to prepare modified updates for the Olive player.  However this should be done with extreme caution and no warranty is provided or implied.
