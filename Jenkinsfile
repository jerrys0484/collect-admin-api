pipeline {
    // Run on an agent where we want to use Go
    agent any

    // Ensure the desired Go version is installed for all stages,
    // using the name defined in the Global Tool Configuration
    tools { go 'go1.25.4' }

    stages {
        stage('Load Config') {
            steps {
                configFileProvider(
                    [configFile(fileId: 'a71f7ca2-1232-4c47-b7cf-4a9e1ab4bb4a', targetLocation: 'bin/collect-admin-api.yaml')]
                ) {
                    sh '''
                        echo "配置文件已生成："
                        cat bin/collect-admin-api.yaml
                    '''
                }
            }
        }
        stage('Build') {
            steps {
                sh 'echo "开始构建..."'
                sh 'go env -w GOPROXY=https://goproxy.cn,direct'
                sh 'go build -o bin/collect-admin-api main.go'
                sh 'cp -rf resource bin/resource'
                sh 'echo "构建完成"'
                sh 'ls -la bin'
            }
        }
        stage('Deploy via SSH') {
            steps {
                sshPublisher(
                    publishers: [
                        sshPublisherDesc(
                            configName: 'Developer', // 在系统设置中配置的名称
                            transfers: [
                                sshTransfer(
                                    sourceFiles: 'bin/**',
                                    // 远程目录（相对于系统配置中的“Remote Directory”）
                                    remoteDirectory: 'collect',
                                    // 传输完成后在远程执行的命令
                                    execCommand: '''
                                        # 执行重启服务
                                        sh bin/restart.sh
                                        echo "部署完成于 $(date)"
                                    ''',
                                    // 可选：是否在传输前清空远程目录
                                    cleanRemote: false
                                )
                            ],
                            // 可选：全局执行命令（在所有传输完成后执行一次）
                            execCommand: 'echo 所有文件传输完毕',
                            // 可选：是否使用sudo
                            useSftpForExec: false
                        )
                    ]
                )
            }
        }
    }
}
