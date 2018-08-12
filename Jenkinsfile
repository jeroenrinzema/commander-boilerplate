pipeline {
  agent {
    docker {
      image 'docker'
    }

  }
  stages {
    stage('Build images') {
      steps {
        dir(path: 'command/') {
          sh 'docker build -t commander-command .'
        }

      }
    }
  }
}