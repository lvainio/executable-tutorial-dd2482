echo "Setting up the environment..."
wget https://golang.org/dl/go1.21.0.linux-amd64.tar.gz
sudo tar -C /usr/local -xzf go1.21.0.linux-amd64.tar.gz
export PATH=$PATH:/usr/local/go/bin
export PATH=$PATH:~/go/bin
git clone https://github.com/lvainio/executable-tutorial-dd2482.git
cd executable-tutorial-dd2482/programs
echo DONE