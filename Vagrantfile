# -*- mode: ruby -*-
# vi: set ft=ruby :

Vagrant.configure("2") do |config|
  config.vm.box = "ubuntu/cosmic64"
  config.vm.hostname = "ccingo"
  config.vm.synced_folder ".", "/home/vagrant/ccingo"

  config.vm.provider "virtualbox" do |vb|
    vb.name = "ccingo"
    vb.memory = "1024"
    vb.cpus = 2
  end

  config.vm.provision "shell", inline: <<-SHELL
    # upgrade to newest packages
    apt-get update
    apt-get upgrade -y

    apt-get install -y gcc
  SHELL

  config.vm.provision "shell", privileged: false, inline: <<-SHELL
    sudo apt-get install -y make zsh
    sudo chsh -s /usr/bin/zsh vagrant
    curl -L https://raw.githubusercontent.com/bto/dotfiles/master/bin/installer.sh | bash
  SHELL
end
