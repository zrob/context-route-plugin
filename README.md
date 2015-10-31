# Context Route Plugin

This is a Cloud Foundry CLI plugin for manage routes with context paths.

## To Install
```
cf install-plugin https://github.com/zrob/context-route-plugin/blob/master/bin/OSX_binary?raw=true
```

## Commands

### create-context-route
`cf create-context-route SPACE DOMAIN HOST PATH`
##### Example
`cf create-context-route myspace bosh-lite.com apphost /v2`

### map-context-route
`cf map-context-route APP DOMAIN HOST PATH`
##### Example
`cf map-context-route myapp bosh-lite.com apphost /v2`

### unmap-context-route
`cf unmap-context-route APP DOMAIN HOST PATH`
##### Example
`cf unmap-context-route myapp bosh-lite.com apphost /v2`

### delete-context-route
`cf delete-context-route DOMAIN HOST PATH`
##### Example
`cf delete-context-route bosh-lite.com apphost /v2`

