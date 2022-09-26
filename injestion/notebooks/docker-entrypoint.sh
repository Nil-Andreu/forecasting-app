#!/bin/bash

# Script to install Linux System Tools and Basic Python Components
apt-get update   # updates the package index cache
apt-get upgrade -y    # updates packages

# Installation of system tools
apt-get install -y bzip2 gcc git
apt-get install -y htop screen vim wget
apt-get upgrade -y bash
apt-get clean   # cleans up package index cache

# Installation of miniconda
wget https://repo.anaconda.com/miniconda/Miniconda3-latest-Linux-x86_64.sh -O Miniconda.sh   # download of miniconda
bash Miniconda.sh -b    # installation of miniconda
rm -rf Miniconda.sh
export PATH="/root/miniconda3/bin:$PATH"  # pre-pends the path

# Installation of python libraries
conda install -y pandas
conda install -y ipython
conda install -y jupyter seaborn statsmodels 
conda install -y pystan
conda install -y fbprophet
