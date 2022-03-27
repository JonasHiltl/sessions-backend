## Deployment
```
kubectl port-forward vespa-0 19071 8080 &

vespa config set target local
vespa deploy --wait 300
```

## Doc Reference
[User Search with friend boosting](https://github.com/vespa-engine/sample-apps/tree/jobergum/friend-graph/friend-graph)