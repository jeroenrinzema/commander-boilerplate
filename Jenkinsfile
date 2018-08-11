pipeline {
  agent any
  stages {
    stage('Build images') {
      agent any
      steps {
        sh 'docker-compose -f docker/services/docker-compose.yml build query'
      }
    }
  }
}