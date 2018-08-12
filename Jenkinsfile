pipeline {
  agent any
  stages {
    stage('Build images') {
      steps {
        dir(path: 'command/') {
          sh 'docker build -t commander_command .'
        }

      }
    }
  }
}