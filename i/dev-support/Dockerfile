FROM prod-harbor-registry.cn-shanghai.cr.aliyuncs.com/prod/itsm-backend:basic
LABEL Maintainer="lihan"
WORKDIR /app
RUN echo "2022-03-17 15:00" > /app/init.txt 
RUN git clone git@code.corp.ecoplants.tech:devops/itsm_backend.git
RUN mkdir -p /app/itsm_backend/build/
WORKDIR /app/itsm_backend
RUN ln -s /usr/local/python3 /app/itsm_backend/build/python
RUN ln -s /app/itsm_backend/build/python/bin/python3 /app/itsm_backend/build/python/bin/python
RUN pip3 install -r requirements.txt
CMD [ "bash", "/app/itsm_backend/dev-support/start.sh" ]