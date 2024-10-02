cd ~  # Reset the working directory to home
clear
echo "Setting up the environment..."
export secret="/root/executable-tutorial-dd2482/gosec/step1/secret.txt"
wget https://golang.org/dl/go1.21.0.linux-amd64.tar.gz
sudo tar -C /usr/local -xzf go1.21.0.linux-amd64.tar.gz
clear
echo "Configuring the environment..."
export PATH=$PATH:/usr/local/go/bin
export PATH=$PATH:~/go/bin
git clone https://github.com/lvainio/executable-tutorial-dd2482.git
cd executable-tutorial-dd2482/programs
echo DONE