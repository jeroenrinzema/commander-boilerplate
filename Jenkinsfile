pipeline {
  agent any
  stages {
    stage('Notification') {
      steps {
        slackSend(message: '${JOB_NAME} ${BUILD_DISPLAY_NAME} has started: ${BUILD_URL}', channel: 'server')
      }
    }
    stage('Build images') {
      steps {
        dir(path: 'docker/services/') {
          sh 'docker-compose build command'
        }

      }
    }
    stage('End Notification') {
      steps {
        slackSend(message: '${JOB_NAME} ${BUILD_DISPLAY_NAME} SUCCESS: ${BUILD_URL}', color: 'good', channel: 'server')
      }
    }
  }
}