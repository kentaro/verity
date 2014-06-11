%w[
  git
].each do |p|
  package p do
    action :install
  end
end

file '/etc/profile.d/go.sh' do
  mode 0755
  content <<'EOS'
export PATH=/usr/local/go/bin:/home/vagrant/bin:$PATH
export GOPATH=/home/vagrant
EOS
  action :create
end

directory '/home/vagrant/src/github.com/kentaro' do
  owner 'vagrant'
  group 'vagrant'
  action :create
  recursive true
end

directory '/usr/local/go' do
  action :delete
  recursive true
end

execute 'install golang' do
  cwd '/usr/local/src'
  command <<EOS
wget https://storage.googleapis.com/golang/go1.2.2.linux-amd64.tar.gz &&
tar zxvf go1.2.2.linux-amd64.tar.gz                                   &&
mv go /usr/local/go
EOS
end
