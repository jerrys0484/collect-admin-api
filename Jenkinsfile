pipeline {
    // Run on an agent where we want to use Go
    agent any

    // Ensure the desired Go version is installed for all stages,
    // using the name defined in the Global Tool Configuration
    tools { go 'go1.25.4' }

    stages {
        stage('Build') {
            steps {
                sh 'echo "开始构建..."'
                sh 'ls -la'
                sh 'go env -w GOPROXY=https://goproxy.cn,direct'
                sh 'go build -o collect-admin-api main.go'
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
                                    sourceFiles: 'collect-admin-api',
                                    // 远程目录（相对于系统配置中的“Remote Directory”）
                                    remoteDirectory: 'collect',
                                    // 传输完成后在远程执行的命令
                                    execCommand: '''
                                        # 这是一个在远程服务器上执行的脚本
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
