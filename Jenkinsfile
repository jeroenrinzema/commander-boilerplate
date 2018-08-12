pipeline {
  agent {
    docker {
      image 'tmaier/docker-compose'
    }

  }
  stages {
    stage('Build images') {
      steps {
        dir(path: 'docker/services') {
          sh 'docker-compose build command'
        }

      }
    }
  }
}