FROM artifacts.iflytek.com/docker-private/aiaas/sqjian:ubuntu20.04
COPY objs /gts/objs
COPY contrib/configure/development /gts/objs
COPY contrib/configure/common /gts/objs
ARG WORKDIR=/gts/objs
ENV LD_LIBRARY_PATH ${WORKDIR}
WORKDIR ${WORKDIR}