FROM envoyproxy/envoy:v1.20-latest

COPY ./services/profile/sidecar.envoy.yaml /etc/envoy/envoy.yaml
COPY ./services/profile/descriptor.pb /tmp/envoy/descriptor.pb

EXPOSE 8080

CMD /usr/local/bin/envoy -c /etc/envoy/envoy.yaml