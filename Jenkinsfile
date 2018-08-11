pipeline {
  agent any
  stages {
    stage('Build images') {
      steps {
        dir(path: 'docker/services/') {
          sh 'docker-compose build'
        }

      }
    }
  }
}