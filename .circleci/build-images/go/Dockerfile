FROM circleci/golang:1.11.5
ENV PROTOC_VERSION 3.6.1

COPY GNUMakefile ./

RUN make bootstrap &&\
      rm GNUMakefile

RUN sudo apt-get update \
    && sudo apt-get -y upgrade \
    && sudo apt-get install -y unzip \
    && sudo rm -rf /var/lib/apt/lists/*

RUN wget -q -O /tmp/protoc.zip https://github.com/google/protobuf/releases/download/v${PROTOC_VERSION}/protoc-${PROTOC_VERSION}-linux-x86_64.zip \
    && unzip /tmp/protoc.zip -d /tmp/protoc3 \
    && sudo mv /tmp/protoc3/bin/* /usr/local/bin/ \
    && sudo mv /tmp/protoc3/include/* /usr/local/include/ \
    && sudo ln -s /usr/local/bin/protoc /usr/bin/protoc \
    && rm -rf /tmp/protoc*

