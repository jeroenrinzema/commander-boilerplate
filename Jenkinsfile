pipeline {
  agent any
  stages {
    stage('Build images') {
      parallel {
        stage('Build images') {
          steps {
            dir(path: 'docker/services') {
              sh 'docker-compose build command'
            }

          }
        }
        stage('Notification') {
          steps {
            slackSend '$JOB_NAME $BUILD_DISPLAY_NAME has started: $BUILD_URL'
          }
        }
      }
    }
  }
}