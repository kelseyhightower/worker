FROM scratch
MAINTAINER Kelsey Hightower <kelsey.hightower@gmail.com>
ADD worker /worker
ENTRYPOINT ["/worker"]
