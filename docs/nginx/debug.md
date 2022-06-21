# how to debug nginx



- 修改编译命令，两种方式：

  - auto/cc/conf 文件，将 ngx_compile_opt="-c" 修改成 ngx_compile_opt="-c -g"
  - 在配置阶段执行 ./auto/configure --with-cc-opt="-O0" # 默认会加上 -g

- 在 nginx 执行配置生成 ./auto/configure --with-debug --prefix=

- 执行编译 make

- 添加 VSCode 配置 .vscode/lauch.json

  ```json
  {
      "version": "0.2.0",
      "configurations": [
          {
              "name": "(lldb) Launch",
              "type": "cppdbg",
              "request": "launch",
              "program": "${workspaceFolder}/objs/nginx",
              "args": ["-c", "conf/nginx.conf"],
              "stopAtEntry": true,
              "cwd": "${workspaceFolder}/objs",
              "environment": [],
              "externalConsole": false,
              "MIMode": "lldb"
          },
          { 
              "name": "(lldb) Attach",
              "type": "cppdbg",
              "request": "attach",
              "program": "${workspaceFolder}/objs/nginx",
              "processId": "{{nginx_pid}}", # ps -ef|grep nginx 填写worker进程PID
              "MIMode": "lldb"
          },
      ]
  }
  ```

- 修改 conf/nginx.conf 配置文件
- 在 VSCode 的 Run and Debug 模式下执行 Launch