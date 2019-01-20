# our image will be based from alpine image 

##########################################################
# syntax for image versions 
# FROM alpine         - will download latest version of alpine image from hub.docker.com (https://hub.docker.com/_/alpine/)
# FROM alpine:3.1     - will download alpine version 3.1 - of course if you need it

FROM alpine



EXPOSE 5000



ADD payment /bin/sh

##########################################################
# ENTRYPOINT - it can be used to set the primary command for the image.

ENTRYPOINT ["./bin/sh"]
