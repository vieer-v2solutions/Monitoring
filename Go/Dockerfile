FROM golang:1.22 as build
WORKDIR /var/task
COPY ./ ./
RUN go mod init app && go get ./ && go vet -v && go test -v && CGO_ENABLED=0 go build -o main
RUN echo "runner:x:2000:appuser" > group && echo "appuser:x:2000:2000:runner User:/var/task:/main" > passwd

## Final Stage
FROM gcr.io/distroless/static-debian12
WORKDIR /var/task
COPY --from=build /var/task/passwd /etc/passwd
COPY --from=build /var/task/group /etc/group
COPY --from=build /var/task/home.html ./
USER appuser
COPY --from=build --chown=appuser:appuser /var/task/main ./
CMD ["/var/task/main"]
EXPOSE 8080
