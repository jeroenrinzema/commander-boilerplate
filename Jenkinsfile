pipeline {
  agent any
  stages {
    stage('Clone repo') {
      agent any
      steps {
        sh '''ls -all;
ls -all docker/services/docker-compose.yml;
pwd;
whoami;'''
      }
    }
    stage('Build images') {
      steps {
        sh 'docker-compose -f docker/services/docker-compose.yml build query'
      }
    }
  }
}