pipeline {
  agent any
  stages {
    stage('Build images') {
      agent any
      steps {
        dir(path: '${WORKSPACE}/docker/services/') {
          sh '''pwd;
docker-compose build query;'''
        }
        
      }
    }
  }
}