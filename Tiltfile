# Welcome to Tilt!
#   To get you started as quickly as possible, we have created a
#   starter Tiltfile for you.
#
#   Uncomment, modify, and delete any commands as needed for your
#   project's configuration.

# Build Docker image
#   Tilt will automatically associate image builds with the resource(s)
#   that reference them (e.g. via Kubernetes or Docker Compose YAML).
#
#   More info: https://docs.tilt.dev/api.html#api.docker_build
#
load('ext://restart_process', 'docker_build_with_restart')

docker_build(
    'jonashiltl/user-service', 
    '.', 
    dockerfile='services/user/Dockerfile', 
    # entrypoint=['/app/services/user/user-service'],
    only=['./services/user', './packages', './go.mod', './go.sum', './tools.go'],
    live_update=[
        # Sync files from host to container
        sync('./services/user/internal', '/app/services/user/internal'),
        sync('./services/user/ent', '/app/services/user/ent'),
        sync('./services/user/main.go', '/app/services/user/main.go'),
        sync('./packages', '/app/packages'),
    ]
)
docker_build(
    'jonashiltl/user-service-sidecar', 
    '.', 
    dockerfile='services/user/sidecar.Dockerfile',
)

docker_build_with_restart(
    'jonashiltl/party-service', 
    '.', 
    dockerfile='services/party/Dockerfile', 
    entrypoint=['/app/services/party/party-service'],
    only=['./services/party', './packages', './go.mod', './go.sum', './tools.go'],
    live_update=[
        sync('./services/party/internal', '/app/services/party/internal'),
        sync('./services/party/main.go', '/app/services/party/main.go'),
        sync('./packages', '/app/packages'),
    ]
)
docker_build(
    'jonashiltl/party-service-sidecar', 
    '.', 
    dockerfile='services/party/sidecar.Dockerfile',
)

docker_build_with_restart(
    'jonashiltl/story-service', 
    '.', 
    dockerfile='services/story/Dockerfile', 
    entrypoint=['/app/services/story/story-service'],
    only=['./services/story', './packages', './go.mod', './go.sum', './tools.go'],
    live_update=[
        sync('./services/story/internal', '/app/services/story/internal'),
        sync('./services/story/main.go', '/app/services/story/main.go'),
        sync('./packages', '/app/packages'),
    ]
)
docker_build(
    'jonashiltl/story-service-sidecar', 
    '.', 
    dockerfile='services/story/sidecar.Dockerfile',
)

docker_build_with_restart(
    'jonashiltl/comment-service', 
    '.', 
    dockerfile='services/comment/Dockerfile', 
    entrypoint=['/app/services/comment/comment-service'],
    only=['./services/comment', './packages', './go.mod', './go.sum', './tools.go'],
    live_update=[
        sync('./services/comment/internal', '/app/services/comment/internal'),
        sync('./services/comment/main.go', '/app/services/comment/main.go'),
        sync('./packages', '/app/packages'),
    ]
)
docker_build(
    'jonashiltl/comment-service-sidecar', 
    '.', 
    dockerfile='services/comment/sidecar.Dockerfile',
)

docker_build_with_restart(
    'jonashiltl/notification-service', 
    '.', 
    dockerfile='services/notification/Dockerfile', 
    entrypoint=['/app/services/notification/notification-service'],
    only=['./services/notification', './packages', './go.mod', './go.sum', './tools.go'],
    live_update=[
        sync('./services/notification/internal', '/app/services/notification/internal'),
        sync('./services/notification/main.go', '/app/services/notification/main.go'),
        sync('./packages', '/app/packages'),
    ]
)

#load('ext://helm_remote', 'helm_remote')
#helm_remote(
#    "nats",
#    repo_name='nats',
#    repo_url='https://nats-io.github.io/k8s/helm/charts/'
#)
# Apply Kubernetes manifests
#   Tilt will build & push any necessary images, re-deploying your
#   resources as they change.
#
#   More info: https://docs.tilt.dev/api.html#api.k8s_yaml
#
k8s_yaml([
    'k8s/pods/comment.yaml', 
    'k8s/pods/party.yaml', 
    'k8s/pods/story.yaml', 
    'k8s/pods/user.yaml',
])
k8s_yaml([
    'k8s/services/comment.yaml', 
    'k8s/services/notification.yaml', 
    'k8s/services/party.yaml', 
    'k8s/services/story.yaml', 
    'k8s/services/user.yaml',
    'k8s/services/vespa.yaml',
])
k8s_yaml([
    'k8s/deployments/notification.yaml', 
])
k8s_yaml([
    'k8s/statefulsets/vespa.yaml', 
])


# Customize a Kubernetes resource
#   By default, Kubernetes resource names are automatically assigned
#   based on objects in the YAML manifests, e.g. Deployment name.
#
#   Tilt strives for sane defaults, so calling k8s_resource is
#   optional, and you only need to pass the arguments you want to
#   override.
#
#   More info: https://docs.tilt.dev/api.html#api.k8s_resource
#
k8s_resource('user', port_forwards=['8080:8080','8081:8081','8180:8180'])


# Run local commands
#   Local commands can be helpful for one-time tasks like installing
#   project prerequisites. They can also manage long-lived processes
#   for non-containerized services or dependencies.
#
#   More info: https://docs.tilt.dev/local_resource.html
#
# local_resource('install-helm',
#                cmd='which helm > /dev/null || brew install helm',
#                # `cmd_bat`, when present, is used instead of `cmd` on Windows.
#                cmd_bat=[
#                    'powershell.exe',
#                    '-Noninteractive',
#                    '-Command',
#                    '& {if (!(Get-Command helm -ErrorAction SilentlyContinue)) {scoop install helm}}'
#                ]
# )
