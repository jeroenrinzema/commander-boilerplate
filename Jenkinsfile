pipeline {
  agent any
  stages {
    stage('Build images') {
      steps {
        sh '''cd docker/services/;
docker-compose build;'''
      }
    }
  }
}