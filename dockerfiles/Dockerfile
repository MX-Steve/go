FROM centos:7
RUN yum  -y install zlib-devel bzip2-devel gcc gcc-c++ make
RUN yum -y install openssl-devel ncurses-devel
RUN yum -y install sqlite-devel readline-devel
RUN yum -y install tk-devel gdbm-devel db4-devel libpcap-devel xz-devel
RUN yum -y install libffi-devel
RUN yum -y install wget mysql-devel gcc-devel python-devel
RUN mkdir /app
WORKDIR /app
RUN wget https://dev.mysql.com/get/mysql80-community-release-el7-3.noarch.rpm
RUN yum -y install mysql80-community-release-el7-3.noarch.rpmyum
RUN wget https://www.python.org/ftp/python/3.7.0/Python-3.7.0.tar.xz
RUN yum install mysql80-community-release-el7-3.noarch.rpm -y
RUN yum install mysql-community-client -y
RUN  tar -xf Python-3.7.0.tar.xz
WORKDIR /app/Python-3.7.0
RUN  mkdir /usr/local/python3
RUN   ./configure --prefix=/usr/local/python3/ --enable-ssl
RUN make && make install
RUN ln -s /usr/local/python3/bin/python3 /usr/local/bin/python3
RUN ln -s /usr/local/python3/bin/pip3 /usr/local/bin/pip3
RUN mkdir /root/.pip
COPY pip.conf /root/.pip/pip.conf
COPY requirements.txt /tmp/requirements.txt
RUN pip3 install --upgrade pip
RUN pip3 install -r /tmp/requirements.txt
