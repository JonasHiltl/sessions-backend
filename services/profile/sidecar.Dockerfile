FROM envoyproxy/envoy:v1.20-latest

COPY ./services/user/sidecar.envoy.yaml /etc/envoy/envoy.yaml
COPY ./services/user/descriptor.pb /tmp/envoy/descriptor.pb

EXPOSE 8080

CMD /usr/local/bin/envoy -c /etc/envoy/envoy.yaml