# Kong Gateway Admin API CLI - Kctl
``` 
 ______
< kctl >
 ------
        \   ^__^
         \  (oo)\_______
            (__)\       )\/\
                ||----w |
                ||     ||
```
 
### CLI for Kong Gateway Admin api management

功能：  
1. 批量获取
2. 批量创建
3. 批量更新
4. 批量删除

### Build  
```
GOOS=linux GOARCH=amd64 go build -o bin/kctl_amd64
GOOS=darwin GOARCH=arm64 go build -o bin/kctl_arm64
```