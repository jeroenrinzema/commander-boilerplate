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
        echo '$JOB_NAME'
      }
    }
  }
}