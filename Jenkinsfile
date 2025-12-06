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
    }
}
