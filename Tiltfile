# USER MINKUBE TUNNEL TO ACCESS LOADBALANCER SERVICE FROM LOCALHOST
# minikube tunnel


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
    only=[
        './services/user', 
        './packages', 
        './go.mod', 
        './go.sum', 
    ],
    live_update=[
        # Sync files from host to container
        sync('./services/user', '/app/services/user'),
        sync('./packages', '/app/packages'),
    ]
)

docker_build(
    'jonashiltl/party-service', 
    '.', 
    dockerfile='services/party/Dockerfile', 
    only=[
        './services/party', 
        './packages', 
        './go.mod', 
        './go.sum', 
    ],
    live_update=[
        sync('./services/party', '/app/services/party'),
        sync('./packages', '/app/packages'),
    ]
)

docker_build(
    'jonashiltl/relation-service', 
    '.', 
    dockerfile='services/relation/Dockerfile', 
    only=[
        './services/relation', 
        './packages', 
        './go.mod', 
        './go.sum', 
    ],
    live_update=[
        sync('./services/relation', '/app/services/relation'),
        sync('./packages', '/app/packages'),
    ]
)

docker_build(
    'jonashiltl/story-service', 
    '.', 
    dockerfile='services/story/Dockerfile', 
    only=[
        './services/story', 
        './packages', 
        './go.mod', 
        './go.sum', 
    ],
    live_update=[
        sync('./services/story', '/app/services/story'),
        sync('./packages', '/app/packages'),
    ]
)

docker_build(
    'jonashiltl/comment-service', 
    '.', 
    dockerfile='services/comment/Dockerfile', 
    only=[
        './services/comment', 
        './packages', 
        './go.mod', 
        './go.sum', 
    ],
    live_update=[
        sync('./services/comment', '/app/services/comment'),
        sync('./packages', '/app/packages'),
    ]
)

docker_build(
    'jonashiltl/notification-service', 
    '.', 
    dockerfile='services/notification/Dockerfile', 
    only=[
        './services/notification', 
        './packages', 
        './go.mod', 
        './go.sum', 
    ],
    live_update=[
        sync('./services/notification', '/app/services/notification'),
        sync('./packages', '/app/packages'),
    ]
)

docker_build(
    'jonashiltl/aggregator', 
    '.', 
    dockerfile='services/aggregator/Dockerfile', 
    only=[
        './services/aggregator', 
        './packages', 
        './go.mod', 
        './go.sum', 
    ],
    live_update=[
        sync('./services/aggregator', '/app/services/aggregator'),
        sync('./packages', '/app/packages'),
    ]
)

docker_build(
    'jonashiltl/scylla-sync', 
    '.', 
    dockerfile='services/scylla-sync/Dockerfile', 
    only=[
        './services/scylla-sync', 
        './packages', 
        './go.mod', 
        './go.sum', 
    ],
    live_update=[
        sync('./services/scylla-sync', '/app/services/scylla-sync'),
        sync('./services/scylla-sync/main.go', '/app/services/scylla-sync/main.go'),
        sync('./packages', '/app/packages'),
    ]
)

load('ext://helm_remote', 'helm_remote')
helm_remote(
    "nats",
    repo_name='nats',
    repo_url='https://nats-io.github.io/k8s/helm/charts/',
    set=['nats.jetstream.enabled=true'],
)
# Apply Kubernetes manifests
#   Tilt will build & push any necessary images, re-deploying your
#   resources as they change.
#
#   More info: https://docs.tilt.dev/api.html#api.k8s_yaml
#
k8s_yaml([
    'k8s/deployments/comment.yaml',
    'k8s/deployments/notification.yaml',
    'k8s/deployments/party.yaml',
    'k8s/deployments/user.yaml',
    'k8s/deployments/story.yaml',
    'k8s/deployments/relation.yaml',
    'k8s/deployments/aggregator.yaml',
    'k8s/deployments/scylla-sync.yaml',
])
k8s_yaml([
    'k8s/services/comment.yaml',
    'k8s/services/notification.yaml',
    'k8s/services/party.yaml',
    'k8s/services/story.yaml',
    'k8s/services/user.yaml',
    'k8s/services/vespa.yaml',
    'k8s/services/aggregator.yaml',
    'k8s/services/relation.yaml',
#    'k8s/services/scylla.yaml',
#    'k8s/services/mongo.yaml',
])
k8s_yaml([
    'k8s/statefulsets/vespa.yaml', 
])
#k8s_yaml([
#    'k8s/endpoints/mongo.yaml',
#    'k8s/endpoints/scylla.yaml',
#])


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
# k8s_resource('profile', port_forwards=['8080:8080','8081:8081','8180:8180'])


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
