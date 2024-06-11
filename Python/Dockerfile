FROM python:3.10 as build
WORKDIR /var/task
COPY ./ ./
RUN pip install -r requirement.txt --target ./
RUN echo "runner:x:2000:appuser" > group && echo "appuser:x:2000:2000:runner User:/var/task:/main" > passwd

# Final 
FROM gcr.io/distroless/python3-debian12
WORKDIR /var/task
COPY --from=build /var/task/group /etc/group
COPY --from=build /var/task ./
CMD ["/var/task/main.py"]

