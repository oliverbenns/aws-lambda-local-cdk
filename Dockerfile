# This is the lambda Dockerfile that is needed in order to get the go mod files copied as you can't reference parent directories from a Dockerfile. Once built, it is used as a base image in the other locations. This will need to be rebuilt every time a dep changes...not good, improve this.

FROM public.ecr.aws/lambda/provided:al2 as build
# install compiler
RUN yum install -y golang
RUN go env -w GOPROXY=direct
# cache dependencies
ADD ./go.mod ./go.sum ./
RUN go mod download
