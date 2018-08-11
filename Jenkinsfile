pipeline {
  agent any
  stages {
    stage('Build images') {
      agent any
      steps {
        dir(path: 'docker/services/') {
          sh '''pwd;
docker-compose build query;'''
        }
        
      }
    }
  }
}