FROM golang:wheezy
MAINTAINER Kenneth Lee "kenneth.lee@gmail.com"

# environment
ENV DEBIAN_FRONTEND noninteractive
RUN echo "deb http://archive.ubuntu.com/ubuntu precise main universe" > /etc/apt/sources.list

####################################################################
# API 
####################################################################
ENV APP_HOME /go/src/github.com/kykl/goapi
ENV LOGS_HOME /var/log/goapi
RUN mkdir -p $LOGS_HOME

VOLUME $LOGS_HOME

COPY . $APP_HOME
WORKDIR $APP_HOME
ENV GOPATH $APP_HOME/Godeps/_workspace:$GOPATH

RUN go build

####################################################################
# Fluentd Agent
####################################################################
# update, curl, sudo
RUN apt-get update && apt-get -y --force-yes upgrade
RUN apt-get -y --force-yes install curl
RUN apt-get -y --force-yes install sudo
#
## fluentd
RUN curl -O http://packages.treasure-data.com/debian/RPM-GPG-KEY-td-agent && apt-key add RPM-GPG-KEY-td-agent && rm RPM-GPG-KEY-td-agent
RUN curl -L http://toolbelt.treasuredata.com/sh/install-ubuntu-precise-td-agent2.sh | sh
ADD conf/td-agent.conf /etc/td-agent/td-agent.conf
RUN curl -L https://raw.githubusercontent.com/fluent/fluentd/master/COPYING > /fluentd-license.txt
#
## fluent-plugin-bigquery
RUN /usr/sbin/td-agent-gem install fluent-plugin-bigquery --no-ri --no-rdoc -V
RUN curl -L https://raw.githubusercontent.com/kaizenplatform/fluent-plugin-bigquery/master/LICENSE.txt > fluent-plugin-bigquery-license.txt
#

# fluentd agent runs as td-agent
RUN chown td-agent:td-agent $APP_HOME/conf/devops-service-account.p12

EXPOSE 8080
ENTRYPOINT /etc/init.d/td-agent restart && $APP_HOME/goapi
