pipeline {
  agent {
    docker {
      image 'docker/compose'
    }

  }
  stages {
    stage('Build images') {
      agent {
        docker {
          image 'docker/compose'
        }

      }
      steps {
        dir(path: 'docker/services/') {
          sh 'docker-compose build'
        }

      }
    }
  }
}