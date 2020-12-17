pipeline {
    agent any
    tools {
       go 'go1.15'
    }
    environment {
        GO114MODULE = 'on'
        CGO_ENABLED = 0 
        GOPATH = "${JENKINS_HOME}/jobs/${JOB_NAME}/builds/${BUILD_ID}"
    }
    stages {       
         
        stage('Pre Test') {
            steps {
                echo 'Change .env.example to .env'
                sh 'mv .env.example .env'
                echo 'Get Golang Version'
                sh 'go version'
                echo 'Dowloading dependencies'
                sh 'go mod download'
                echo 'Migrating Database'
            }
        }

        stage('Build') {
            steps {
                echo 'Compiling and building'
                sh 'go build'
            }
        }

        stage('Test') {
            steps {
                withEnv(["PATH+GO=${GOPATH}/bin"]){
                    echo 'Running Repository test'
                    sh 'go test ./api/Repository/Employe/...'
                }
            }
        }
        
    }

}
