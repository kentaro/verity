# -*- mode: ruby -*-
# vi: set ft=ruby :

# Vagrantfile API/syntax version. Don't touch unless you know what you're doing!
VAGRANTFILE_API_VERSION = "2"

Vagrant.configure(VAGRANTFILE_API_VERSION) do |config|
  config.vm.box = "CentOS-6.5"
  config.vm.box_url = "https://developer.nrel.gov/downloads/vagrant-boxes/CentOS-6.5-x86_64-v20140504.box"

  config.vm.provision "chef_solo" do |chef|
    chef.add_recipe "verity"
  end

  config.vm.synced_folder ".", "/home/vagrant/src/github.com/kentaro/verity"
end
