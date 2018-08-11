pipeline {
  agent {
    docker {
      image 'tmaier/docker-compose'
    }
    
  }
  stages {
    stage('Build images') {
      agent any
      steps {
        dir(path: 'docker/services') {
          sh 'docker-compose build'
        }
        
      }
    }
  }
}