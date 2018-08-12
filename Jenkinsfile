pipeline {
  agent any
  stages {
    stage('Build images') {
      steps {
        dir(path: 'docker/services/') {
          sh 'docker-compose build command'
        }

      }
    }
    stage('Notify Success') {
      steps {
        slackSend(message: '$JOB_NAME $BUILD_DISPLAY_NAME SUCCESS', color: 'good', botUser: true, channel: 'server')
      }
    }
  }
}