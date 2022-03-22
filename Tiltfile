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
docker_build(
    'jonashiltl/user-service', 
    '.', 
    dockerfile='services/user/Dockerfile', 
    live_update=[
        # Sync files from host to container
        sync('./services/user/internal', '/app/services/user/internal'),
        sync('./services/user/ent', '/app/services/user/ent'),
        sync('./services/user/main.go', '/app/services/user/main.go'),
        sync('./packages', '/app/packages'),
    ]
)
docker_build(
    'jonashiltl/party-service', 
    '.', 
    dockerfile='services/party/Dockerfile', 
    live_update=[
        # Sync files from host to container
        sync('./services/party/internal', '/app/services/party/internal'),
        sync('./services/party/main.go', '/app/services/party/main.go'),
        sync('./packages', '/app/packages'),
    ]
)
docker_build(
    'jonashiltl/story-service', 
    '.', 
    dockerfile='services/story/Dockerfile', 
    live_update=[
        # Sync files from host to container
        sync('./services/story/internal', '/app/services/story/internal'),
        sync('./services/story/main.go', '/app/services/story/main.go'),
        sync('./packages', '/app/packages'),
    ]
)
docker_build(
    'jonashiltl/comment-service', 
    '.', 
    dockerfile='services/comment/Dockerfile', 
    live_update=[
        # Sync files from host to container
        sync('./services/comment/internal', '/app/services/comment/internal'),
        sync('./services/comment/main.go', '/app/services/comment/main.go'),
        sync('./packages', '/app/packages'),
    ]
)
docker_build(
    'jonashiltl/notification-service', 
    '.', 
    dockerfile='services/notification/Dockerfile', 
    live_update=[
        # Sync files from host to container
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
    'k8s/services/comment.yaml', 
    'k8s/services/notification.yaml', 
    'k8s/services/party.yaml', 
    'k8s/services/story.yaml', 
    'k8s/services/user.yaml',
])
k8s_yaml([
    'k8s/deployments/comment.yaml', 
    'k8s/deployments/notification.yaml', 
    'k8s/deployments/party.yaml', 
    'k8s/deployments/story.yaml', 
    'k8s/deployments/user.yaml'
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
# k8s_resource('my-deployment',
#              # map one or more local ports to ports on your Pod
#              port_forwards=['5000:8080'],
#              # change whether the resource is started by default
#              auto_init=False,
#              # control whether the resource automatically updates
#              trigger_mode=TRIGGER_MODE_MANUAL
# )


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
