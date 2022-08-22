FROM --platform=linux/amd64 golang:1.17.2-alpine

RUN \
    apk --no-cache add tzdata && \
	cp /usr/share/zoneinfo/Asia/Seoul /etc/localtime && \
	echo "Asia/Seoul" > /etc/timezone \
	apk del tzdata \

RUN \
  apk update && \
  apk add openrc --no-cache && \
  apk add openssh-server && \
  rc-update add sshd && \
  rc-status && \
  touch /run/openrc/softlevel \

RUN echo 'StrictHostKeyChecking no' >> /etc/ssh/sshd_config
RUN echo 'PasswordAuthentication yes' >> /etc/ssh/sshd_config

RUN adduser -h /home/extreme -s /bin/sh -D extreme
RUN echo -n 'extreme:password' | chpasswd

WORKDIR /app

COPY cli/go.* ./
RUN go mod download

COPY cli/*.go ./
COPY cli/cmd ./cmd

RUN go build -o efa
RUN mv efa /usr/bin

EXPOSE 22

CMD ["sh", "-c", "service sshd restart && sh"]