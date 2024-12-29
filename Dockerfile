FROM ubuntu:latest
LABEL authors="phinc"

ENTRYPOINT ["top", "-b"]