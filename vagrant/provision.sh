#!/bin/bash

# Update package lists
sudo apt-get update -y

# Your custom provisioning steps go here
# ...

# Optionally, you can echo a message to indicate successful provisioning
# ssh_pub_key="ssh-rsa AAAAB3NzaC1yc2EAAAADAQABAAACAQDIRD3yK4KxRVeudPRxQhc39NkcMUscjapNYE7VhFraP/r57ymXrRrovGLGWar+HZ6gIfBpGpUectKxZ4PoCwx/l0B6coTSFt0RGp+vDUiKB1JlKapZCNsee5GjzPVKxWCGcjQSAuFFYzMY6SUIJWRV4ZWtw4UDjKcZj0Komq37liBIflLjUVOFzNFkPxTsKZ4h2yKLJeVxyWH6ZoaMQ2do0Lo1sbjkleWHViOz4RCwuzze27U/FQ7ydq4uZ2Gz49LaF2/Sv33G1IYqJNGy/M0SeKLbJ8z2lysN++UT5tnQP8EZcz71BFl2q+x1v1kqFOAd7LkwFfX3UgcsCjc0Sdmc+cnl4z2pBz2Wl/32kE30MJE+VwkfObD2/0QnT5HhgGPaUkILyo5V8e7cmEk+o9GsLFK94deS4qtvGA/kFeID30L4NbvWA9Gjw2gAxJ2ye0Z8qAcim+NQCSAOU5P1PNYrj+0loB2aXFOEODh2QEFVQ/Pw0bQF/RNRTeE/xNNG0HWWbg/cAzbcDyexZpTGAiLhWcy7f2gaf9eo/2vf9VkMEVCarkbHX5mIsIphbgOjqP5wcRYFdgiTJ9PefGJZXsvJsqWqN86Rkp+c6UWfX4BgPQRd507cOhM4PBwgjiWTslQ1Zj5KqSEXBoHlv2PeFz6Yp9iqiBXJUCkTXOwl8wPHEw== your_email@example.com"
#
# echo "$ssh_pub_key" >> /home/vagrant/.ssh/authorized_keys
# echo "$ssh_pub_key" >> /root/.ssh/authorized_keys

echo "VM provisioned successfully."
