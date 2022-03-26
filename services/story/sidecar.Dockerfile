FROM envoyproxy/envoy:v1.20-latest

COPY ./services/story/sidecar.envoy.yaml.yaml /etc/envoy/envoy.yaml
COPY ./services/story/descriptor.pb /tmp/envoy/descriptor.pb

EXPOSE 5000

CMD /usr/local/bin/envoy -c /etc/envoy/envoy.yaml