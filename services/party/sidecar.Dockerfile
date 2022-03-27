FROM envoyproxy/envoy:v1.20-latest

COPY ./services/party/sidecar.envoy.yaml /etc/envoy/envoy.yaml
COPY ./services/party/descriptor.pb /tmp/envoy/descriptor.pb

EXPOSE 8080

CMD /usr/local/bin/envoy -c /etc/envoy/envoy.yaml