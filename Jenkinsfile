pipeline {
  agent {
    node {
      label 'master'
    }
  }
  parameters {
    string(name: 'distName', defaultValue: 'netd' , description: '镜像名')
    string(name: 'distNumber', defaultValue: "${BRANCH_NAME}", description: '镜像版本')
  }
  
  stages {
    stage('Build') {
      steps {
        script {
           docker.withRegistry('http://hub.sky-cloud.net', 'docker-image-builder') {
            docker.image('hub.sky-cloud.net/cicd/gobuilder:v2.0').inside {
              sh 'GOPROXY=https://goproxy.io CGO_ENABLE=0 GOOS=linux GOARCH=amd64 go build .'
            }
          }
        }
      }
    }
    stage('Build-dockerimage') {
      steps {
        script {
          def GIT_COMMITID = sh (script: 'git rev-parse --short HEAD ${GIT_COMMITID}', returnStdout: true).trim()
          def IMAGE_NAME = "${params.distName}:${params.distNumber}-jenkins-${GIT_COMMITID}-${BUILD_NUMBER}"
          sh "docker build -t hub.sky-cloud.net/nap2/${IMAGE_NAME} ."
          sh "docker push hub.sky-cloud.net/nap2/${IMAGE_NAME}"
        }
      }
    }
  }
}
