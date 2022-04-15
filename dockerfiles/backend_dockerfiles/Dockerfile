FROM prod-harbor-registry.cn-shanghai.cr.aliyuncs.com/prod/cmdb-backend:basic
LABEL Maintainer="lihan"
RUN rm -rf /app/cmdb_backend
WORKDIR /app
RUN echo "2022-03-02 16:25" > /app/init.txt 
RUN git clone git@code.corp.ecoplants.tech:cmdb_operation/cmdb_backend.git
RUN mkdir -p /app/cmdb_backend/backend/proto
RUN mkdir -p /app/cmdb_backend/build/
WORKDIR /app/cmdb_backend
RUN pip3 install supervisor
RUN rm -rf /usr/local/bin/supervisor*
RUN cp /usr/local/python3/bin/supervisor* /usr/local/bin/
RUN ln -s /usr/local/python3 /app/cmdb_backend/build/python
RUN ln -s /app/cmdb_backend/build/python/bin/python3 /app/cmdb_backend/build/python/bin/python
RUN pip3 install -r requirements.txt
RUN pip3 uninstall -y pycrypto
RUN pip3 uninstall -y pycryptodome
RUN pip3 install pycryptodome
RUN python3 setup.py gen_grpc -o /app/cmdb_backend/backend/proto/
CMD [ "bash", "/app/cmdb_backend/dev-support/start.sh" ]