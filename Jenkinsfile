pipeline {
  agent any
  stages {
    stage('Build command images') {
      agent any
      steps {
        dir(path: 'command/') {
          sh 'docker build . -t commander-command'
        }

      }
    }
  }
}