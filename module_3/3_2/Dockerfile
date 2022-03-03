FROM ubuntu AS build
RUN mkdir /root/.ssh/
COPY id_rsa /root/.ssh/id_rsa

RUN \
  apt-get -y update && \
  apt-get -yqq install ssh

RUN \
  chmod 600 /root/.ssh/id_rsa && \
  chown -R root:root /root/.ssh && \
  touch /root/.ssh/known_hosts && \
  ssh-keyscan github.com >> /root/.ssh/known_hosts  


RUN \
  apt-get install -y git && \
  apt-get install -y golang-go

RUN \
  git clone "git@github.com:usomething/golang_study.git" && \
  go build -o /root/http_server golang_study/module_2/1_2/http_server.go
 

FROM scratch
COPY --from=build /root/http_server /bin/http_server
 
ENTRYPOINT ["/bin/http_server"]

