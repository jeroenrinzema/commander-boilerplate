pipeline {
  agent any
  stages {
    stage('Build images') {
      steps {
        sh '''TAG = sh (
  script: \'git tag --points-at HEAD\',
  returnStdout: true
).trim()

TAG=$TAG;
cd docker/services/;
docker-compose build;'''
        }
      }
    }
  }