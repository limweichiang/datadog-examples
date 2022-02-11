#!/bin/sh

sudo apt remove --purge -y datadog-agent
sudo rm /etc/apt/sources.list.d/datadog.list
sudo rm /etc/apt/sources.list.d/ansible_datadog_7.list 
sudo apt-key del "A292 3DFF 56ED A6E7 6E55  E492 D3A8 0E30 382E 94DE"
sudo apt-key del "D75C EA17 048B 9ACB F186  794B 3263 7D44 F14F 620E"
rm ddagent-install.log
