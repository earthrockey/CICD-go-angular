pipeline {
    tools {
        go 'go-1.12.1'
    }
    stages {
        stage('Test') {
            steps {
                sh 'go run main.go'
            }
        }
    }
}