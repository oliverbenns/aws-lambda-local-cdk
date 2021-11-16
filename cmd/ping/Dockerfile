FROM lambda-builder as build
# build
ADD . .
RUN go build -o /main

# copy artifacts to a clean image
FROM public.ecr.aws/lambda/provided:al2
COPY --from=build /main /main
ENTRYPOINT [ "/main" ]  
