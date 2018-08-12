pipeline {
  agent {
    docker {
      image 'tmaier/docker-compose:18.06'
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