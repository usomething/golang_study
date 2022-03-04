FROM ubuntu AS build

RUN \
  apt-get -y update && \
  export DEBIAN_FRONTEND=noninteractive && \
  apt-get install -y wget && \
  apt-get install -y golang-go && \
  apt-get install -y unzip

RUN \
  wget https://codeload.github.com/usomething/golang_study/zip/refs/heads/main --no-check-certificate && \
  unzip main && \
  go build -o /root/http_server golang_study-main/module_2/1_2/http_server.go


FROM ubuntu AS run_go
WORKDIR /root
COPY --from=build /root/http_server /root/http_server

ENTRYPOINT ["/root/http_server"]