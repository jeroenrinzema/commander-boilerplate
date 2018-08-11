pipeline {
  agent any
  stages {
    stage('Clone repo') {
      agent any
      steps {
        sh '''ls -all;
pwd;'''
      }
    }
    stage('Build images') {
      steps {
        sh 'docker-compose -f docker/services/docker-compose.yml build query'
      }
    }
  }
}