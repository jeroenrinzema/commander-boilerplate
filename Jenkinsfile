pipeline {
  agent {
    docker {
      image 'docker/compose:1.22.0'
    }
    
  }
  stages {
    stage('Build images') {
      agent any
      steps {
        dir(path: 'docker/services') {
          sh 'docker-compose build query'
        }
        
      }
    }
  }
}