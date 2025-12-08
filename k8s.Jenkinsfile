pipeline {
    // Run on an agent where we want to use Go
    agent any

    // Ensure the desired Go version is installed for all stages,
    // using the name defined in the Global Tool Configuration
    tools { go 'go1.25.4' }

    stages {
        stage('Deploy to K8s') {
            steps {
                withKubeConfig([
                    credentialsId: 'k8s-kubeconfig',
                    serverUrl: 'http://192.168.0.102:8001'
                ]) {
                    sh 'kubectl get pods'
                    sh 'kubectl apply -f k8s/deployment.yaml'
                }
            }
        }
    }
}
