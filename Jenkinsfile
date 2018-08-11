pipeline {
  agent any
  stages {
    stage('Build images') {
      agent any
      steps {
        dir(path: 'docker/services/') {
          sh 'docker-compose build query'
        }
        
      }
    }
  }
}